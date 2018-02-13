package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-gtk/pango"
	"github.com/sg3des/tergo/vte2"
)

type Tab struct {
	label *gtk.Label
	term  vte2.Terminal
	wdgt  *gtk.Widget

	pid int
}

func NewTab(wd string) *Tab {
	var t = new(Tab)
	t.term = vte2.NewTerminal()
	t.wdgt = t.term.VteToGtk()

	t.term.SetFontFromString(Conf.Font)
	t.term.SetColors(Conf.Foreground, Conf.Background, Conf.Palette)
	t.term.SetColorCursor(Conf.CursorColor)
	t.term.SetScrollbackLines(Conf.TermLines)
	t.term.SetWordChars(Conf.WordChars)

	t.pid = t.term.Fork([]string{os.Getenv("SHELL")}, wd)

	t.label = gtk.NewLabel(t.term.GetWindowTitle())
	t.label.SetEllipsize(pango.ELLIPSIZE_MIDDLE)
	t.wdgt.Connect("window-title-changed", t.changeTabLabel)

	return t
}

func (t *Tab) GetChilds() (*gtk.Widget, gtk.IWidget) {
	return t.wdgt, t.label
}

func (t *Tab) Copy() {
	t.term.Copy()
}

func (t *Tab) Paste() {
	t.term.Paste()
}

func (t *Tab) GetCurrentWD() string {
	wd, err := os.Readlink(fmt.Sprintf("/proc/%d/cwd", t.pid))
	if err != nil {
		return ""
	}

	return wd
}

func (t *Tab) changeTabLabel() {
	t.label.SetText(t.term.GetWindowTitle())
}
