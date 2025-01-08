package main

import (
	"context"
	"rm-exporter/backend"
	"slices"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx         context.Context
	rm_reader   backend.RmReader
	tablet_addr string
	selection   backend.FileSelection

	rm_export   backend.RmExport
	export_from int
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.export_from = -1
}

func (a *App) ReadTabletDocs(tablet_addr string) error {
	a.tablet_addr = tablet_addr

	err := a.rm_reader.Read(tablet_addr)
	if err != nil {
		return err
	}

	a.selection = backend.NewFileSelection(a.rm_reader.GetChildren())
	return nil
}

func (a *App) GetTabletFolder(id backend.DocId) []backend.DocInfo {
	return a.rm_reader.GetFolder(id)
}

func (a *App) GetTabletFolderSelection(id backend.DocId) []backend.SelectionInfo {
	return a.selection.GetFolderSelection(id)
}

func (a *App) IsIpValid(s string) bool {
	return backend.IsIpValid(s)
}

func (a *App) SetExportOptions(export backend.RmExport) {
	a.rm_export = export
}

func (a *App) Export() {
	files := a.GetCheckedFiles()
	if a.export_from == -1 {
		a.rm_export.WrappingFolderName = "rM Export (" + time.Now().Format(time.DateTime) + ")"
	}
	a.export_from = a.rm_export.ExportMultiple(a.ctx, a.tablet_addr, files, a.export_from)
}

func (a *App) OnItemSelect(id backend.DocId, selection bool) {
	a.selection.Select(id, selection)
}

/* Includes path for every checked file */
func (a *App) GetCheckedFiles() []backend.DocInfo {
	ids := a.selection.GetCheckedFiles()
	files := a.rm_reader.GetElementsByIds(ids)
	paths := a.rm_reader.GetPaths(ids)
	for i := 0; i < len(files); i += 1 {
		files[i].Path = &paths[i]
	}
	slices.SortFunc(files, func(i, j backend.DocInfo) int {
		if *i.Path < *j.Path {
			return -1
		}
		if *i.Path == *j.Path {
			return 0
		}
		return 1
	})
	return files
}

func (a *App) GetCheckedFilesCount() int {
	return a.selection.GetCheckedFilesCount()
}

func (a *App) GetPaths(ids []backend.DocId) []string {
	return a.rm_reader.GetPaths(ids)
}

func (a *App) DirectoryDialog() string {
	dir, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		return ""
	}
	return dir
}
