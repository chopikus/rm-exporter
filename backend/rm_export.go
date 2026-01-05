package backend

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type RmExportOptions struct {
	Pdf               bool
	Rmdoc             bool
	Location          string // path to the folder to export
	PreserveDirStruct bool
}

type RmExport struct {
	Options RmExportOptions

	items       []DocInfo
	export_from int // index of the first item to be exported

	tablet_addr        string
	wrappingFolderName string
	client             http.Client
	ctx                context.Context
	paths              Paths
}

func InitExport(ctx context.Context, options RmExportOptions, items []DocInfo, tablet_addr string) RmExport {
	client := http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
		},
		Timeout: 5 * time.Minute,
	}

	t := time.Now().Format(time.DateTime)
	folderName := "rM Export (" + t + ")"

	return RmExport{
		Options:            options,
		items:              items,
		export_from:        0,
		tablet_addr:        tablet_addr,
		wrappingFolderName: folderName,
		client:             client,
		ctx:                ctx,
		paths:              initPaths(),
	}
}

/*
Exports all items passed in Init() method.
Calls the callbacks when:
* item started downloading;
* item download has finished;
* item download has failed.

Supports retries.
In case the last export succeeded on all items, it starts the export again from the first item;
otherwise, the export starts from the first failed item.
*/
func (r *RmExport) Export(started, finished func(item DocInfo), failed func(item DocInfo, err error)) {
	formats := []string{}
	if r.Options.Rmdoc {
		formats = append(formats, "rmdoc")
	}
	if r.Options.Pdf {
		formats = append(formats, "pdf")
	}

	runtime.LogInfof(r.ctx, "[%v] Export formats: %v", time.Now().UTC(), formats)
	runtime.LogInfof(r.ctx, "[%v] In export location, using a wrapper folder with a name: %v", time.Now().UTC(), r.wrappingFolderName)

	for i := r.export_from; i < len(r.items); i++ {
		item := r.items[i]
		started(item)

		for _, format := range formats {
			err := r.exportOne(item, format)
			if err != nil {
				r.export_from = i
				failed(item, err)
				return
			}
		}

		finished(item)
	}
}

func (r *RmExport) lookupDir(id DocId) error {
	runtime.LogInfof(r.ctx, "[%v] looking up dir, id=%v", time.Now().UTC(), id)

	url := "http://" + r.tablet_addr + "/documents/" + id

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &bytes.Buffer{})
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", "Mozilla/5.0 (U; Linux x86_64; en-US) Gecko/20100101 Firefox/133.0")
	_, err = r.client.Do(req)

	if err != nil {
		return err
	}

	return nil
}

func (r *RmExport) exportOne(item DocInfo, format string) error {
	time.Sleep(150 * time.Millisecond)
	err := r.lookupDir(item.ParentId)
	if err != nil {
		return err
	}

	time.Sleep(150 * time.Millisecond)
	err = r.download(item, format)
	if err != nil {
		return err
	}

	return nil
}

func (r *RmExport) download(item DocInfo, format string) error {
	if item.IsFolder {
		return nil
	}
	runtime.LogInfof(r.ctx, "[%v] downloading an item, id=%v", time.Now().UTC(), item.Id)

	out, err := r.createFile(r.wrappingFolderName, item, format)
	if err != nil {
		return err
	}
	defer out.Close()

	url := "http://" + r.tablet_addr + "/download/" + item.Id + "/" + format

	req, err := http.NewRequest(http.MethodGet, url, &bytes.Buffer{})
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", "Mozilla/5.0 (U; Linux x86_64; en-US) Gecko/20100101 Firefox/133.0")
	resp, err := r.client.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("tablet returned HTTP code %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}

func (r *RmExport) createFile(folderName string, item DocInfo, format string) (*os.File, error) {
	itemPath := item.TabletPath
	if !r.Options.PreserveDirStruct {
		/* If the directory structure is preserved,
		 * create a slice only with the last element.
		 * This flattens the directory structure. */
		itemPath = itemPath[len(itemPath)-1:]
	}

	path, err := r.paths.getFilePathUnique(r.Options.Location, folderName, itemPath, format)
	if err != nil {
		return nil, fmt.Errorf("failed to find a path, id=%v, (%v)", item.Id, err.Error())
	}

	runtime.LogDebugf(r.ctx, "[%v] exporting to path %v, id=%v", time.Now().UTC(), path, item.Id)

	path = filepath.FromSlash(path)
	dir, _ := filepath.Split(path)

	/* Permission 0755: The owner can read, write, execute.
	   Everyone else can read and execute but not modify the file.*/
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, err
	}

	out, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return out, nil
}
