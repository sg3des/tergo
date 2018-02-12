package main

import (
	"fmt"
	"os"
	"reflect"
	"unsafe"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

type Window struct {
	window   *gtk.Window
	notebook *gtk.Notebook

	tabs []NotebookTab

	keyEvents chan *gdk.EventKey
	binds     map[Key]reflect.Value
}

func NewWindow(title string, width, height int) *Window {
	var w = new(Window)

	w.window = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	w.window.SetTitle(title)
	w.window.SetDefaultSize(width, height)
	w.window.Connect("destroy", gtk.MainQuit)
	w.window.Connect("key-press-event", w.KeysEvent)

	w.notebook = gtk.NewNotebook()
	w.window.Add(w.notebook)

	w.keyEvents = make(chan *gdk.EventKey)
	go w.KeysHandler()

	w.binds = make(map[Key]reflect.Value)

	return w
}

func (w *Window) ShowAll() {
	w.window.ShowAll()
}

type NotebookTab interface {
	GetChilds() (*gtk.Widget, gtk.IWidget)
	Close()
}

func (w *Window) NewTab() {
	t := NewTab("")
	w.tabs = append(w.tabs, t)

	child, label := t.GetChilds()

	n := w.notebook.AppendPage(child, label)

	w.notebook.ChildSet(child, "tab-expand", true)
	w.notebook.SetReorderable(child, true)
	w.notebook.ShowAll()

	w.notebook.SetCurrentPage(n)
	// w.notebook.SetFocusChild(child)

	// w.notebook.ShowAll()
	child.GrabFocus()
}

func (w *Window) CloseTab() {
	n := w.notebook.GetCurrentPage()
	t := w.tabs[n]

	child, _ := t.GetChilds()
	w.notebook.RemovePage(child, n)

	w.tabs = append(w.tabs[:n], w.tabs[n+1:]...)

	// t.Close()

	if len(w.tabs) == 0 {
		w.Quit()
	}
}

func (w *Window) Quit() {
	// gtk.MainQuit()
	os.Exit(0)
}

//
// KEYS
//

func (w *Window) SetBinds(binds map[string]string) {
	for method, key := range binds {
		k, ok := ParseKey(key)
		if !ok {
			fmt.Printf("failed parse key %s for method %s\n", key, method)
		}

		f, ok := w.lookupMethod(method)
		if !ok {
			fmt.Printf("method %s unexists\n", method)
		}

		w.binds[k] = f
	}
}

func (w *Window) lookupMethod(s string) (reflect.Value, bool) {
	v := reflect.ValueOf(w)
	method := v.MethodByName(s)
	if !method.IsValid() {
		return reflect.Value{}, false
	}
	return method, true
}

func (w *Window) KeysEvent(ctx *glib.CallbackContext) {
	arg := ctx.Args(0)
	key := *(**gdk.EventKey)(unsafe.Pointer(&arg))

	w.keyEvents <- key
}

func (w *Window) KeysHandler() {
	for {
		key := <-w.keyEvents
		k := NewKey(key)

		f, ok := w.binds[k]
		if !ok {
			continue
		}

		f.Call([]reflect.Value{})
	}
}
