# rm-exporter

https://github.com/user-attachments/assets/0996b7c3-85a7-45ea-90d7-cfabd989c5dc

As you might know, reMarkable supports exporting notes locally through the USB connection.

Unfortunately, the default local export has a few flaws:
  * Large notes (10MB+) often can't be exported, **the UI doesn't wait for long enough for a note to download**;
  * Downloading a folder is not possible; only notes one-by-one.

This tool aims to fix those problems.

## Features
* Supports exporting as many folders & notes as you want;
* Can download both .pdf and .rmdoc;
* Retries the download **from the last failed note**;
* Waits for large notes long enough (the timeout can be specified);
* Doesn't require reMarkable account or internet connection;
* Works with out of the box reMarkable software;
* Has a nice GUI.

## Usage
Releases for Windows/MacOS/Linux are available on the 'Releases' tab of the repository.

The tool is built with [wailsv2](https://github.com/wailsapp/wails). The UI is implemented in Typescript/Svelte, file downloading itself is done in Golang.

### Supported rM software version
Around 3.10+, around that version the local server requests got updated.

Tested on Version 3.22.0.64 on reMarkable 2.

### Steps before running the `rm-exporter`
* Enable USB connection in the Storage settings. Without the permission the app can't find the tablet;
* For long exports with large number of files, turn off Sleep Mode in the Battery settings. For some reason the local export doesn't prevent the tablet from going to sleep.

### Building steps
1. Install [wails v2](https://wails.io/docs/gettingstarted/installation).
2. Clone the project
3. `wails build`
