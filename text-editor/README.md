# Text Editor

A simple, lightweight, multi-platform text editor. Open, edit, and save any
plain-text file through a small single-window interface.

Built with [Go](https://go.dev) and [Fyne](https://fyne.io) — chosen because
Go compiles to native binaries for Windows, Linux, and macOS, and Fyne is a
lightweight, simple native-feeling GUI toolkit.

## Features

- New / Open / Save / Save As for any plain-text file (no file-extension restrictions)
- Undo / Redo
- Unsaved-changes indicator in the title bar, with a confirm-discard prompt on
  New / Open / Quit while there are unsaved changes
- Keyboard shortcuts (Ctrl on Linux/Windows, Cmd on macOS): New, Open, Save,
  Save As (+Shift), Undo, Redo, Quit

Not included in this initial version — see [Future ideas](#future-ideas).

## Running from source

```sh
cd text-editor
go run .
```

## Building

### Linux (native)

Requires GL/X11 development headers to compile Fyne's OpenGL bindings:

```sh
sudo apt-get install gcc libgl1-mesa-dev xorg-dev libxkbcommon-dev pkg-config
cd text-editor
CGO_ENABLED=1 go build -o build/text-editor-linux .
```

### Windows and macOS

Fyne requires CGO to bind to the platform's OpenGL driver, so a plain
`GOOS=windows go build` (or `GOOS=darwin`) from a Linux machine will fail at
link time without a matching C cross-toolchain (mingw-w64 for Windows, a
macOS SDK for Darwin). There are two practical ways to get real Windows and
macOS binaries:

1. **GitHub Actions build matrix (recommended)** — see
   `.github/workflows/build.yml` at the repo root. It builds natively on
   `ubuntu-latest`, `windows-latest`, and `macos-latest`, so each binary is
   compiled on its own OS with no cross-toolchain needed, and uploads all
   three as workflow artifacts.
2. **[`fyne-cross`](https://github.com/fyne-io/fyne-cross)** — a Docker-based
   cross-compiler that bundles mingw-w64 for Windows. macOS targets require
   supplying your own macOS SDK (extracted from Xcode via
   `fyne-cross darwin-sdk-extractor`), since Apple's license doesn't allow
   redistributing a prebuilt image containing it.
   ```sh
   go install github.com/fyne-io/fyne-cross@latest
   cd text-editor
   fyne-cross windows
   fyne-cross linux
   fyne-cross darwin   # requires a supplied macOS SDK
   ```

## Testing

```sh
go vet ./...
xvfb-run -a go test ./...
```

GUI rendering itself can't be verified headlessly (Fyne needs a real GL
context, which a virtual framebuffer like Xvfb doesn't provide). Before
calling a change done, manually verify on a real desktop session:

- New / Open / Edit / Save / Save As
- Undo / Redo
- Quit with unsaved changes — confirms the discard prompt appears
- Opening files with arbitrary/no extensions (`.txt`, `.md`, `.log`, no
  extension at all) to confirm nothing is filtered out
- Cmd-based shortcuts specifically need checking on a real macOS machine

## Future ideas

Not part of this initial version, but reasonable next steps: syntax
highlighting, line numbers, find/replace, multiple tabs/documents, themes,
and persisted settings.
