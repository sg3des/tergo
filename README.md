# tergo

Lightweight terminal emulator on gtk2

## Config

Configuration in yaml format:

```yaml
TabCloseButton: false
TabHeight: 13

Font: "Liberation Mono 7.5"

Background: "#222"
Foreground: "#bbb"
CursorColor: "#e8e8e8"
Palette: ["#2e3436", "#cc0000", "#4e9a06", "#c4a000", "#3465a4", "#75507b", "#06989a", "#b1b1b1", "#555753", "#ef2929", "#8ae234", "#fce94f", "#729fcf", "#ad7fa8", "#34e2e2", "#acacac"]

Binds:
  NewTab: ctrl + t
  CloseTab: ctrl + w
  Quit: ctrl + q

  NextTab: "ctrl + pagedown"
  PrevTab: "ctrl + pageup"

  Copy: shift + ctrl + c
  Paste: shift + ctrl + v
```
