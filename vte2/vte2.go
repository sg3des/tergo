// Copyright (C) 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vte2

// #include "vte.go.h"
// #cgo pkg-config: vte
import "C"

import (
	"unsafe"

	"github.com/mattn/go-gtk/gtk"
	//"github.com/mattn/go-gtk/pango"
)

func (t Terminal) VteToGtk() *gtk.Widget {
	return gtk.WidgetFromNative(unsafe.Pointer(t.Widget))
}

type Terminal struct {
	Widget *C.GtkWidget
}

func NewTerminal() Terminal {
	return Terminal{C.vte_terminal_new()}
}

func (t Terminal) Fork(args []string, wd string) int {
	wdptr := C.CString(wd)
	defer C.free(unsafe.Pointer(wdptr))

	argv := C.make_strings(C.int(len(args) + 1))
	defer C.free(unsafe.Pointer(argv))
	for i, arg := range args {
		ptr := C.CString(arg)
		defer C.free(unsafe.Pointer(ptr))
		C.set_string(argv, C.int(i), ptr)
	}
	C.set_string(argv, C.int(len(args)), nil)

	var pid C.GPid

	C.vte_terminal_fork_command_full(C.toVTerminal(t.Widget), C.VTE_PTY_DEFAULT,
		wdptr, argv, nil, C.G_SPAWN_SEARCH_PATH, nil, nil, &pid, nil)

	return int(pid)
}

func (t Terminal) GetIconTitle() string {
	return C.GoString(C.vte_terminal_get_icon_title(C.toVTerminal(t.Widget)))
}

func (t Terminal) GetWindowTitle() string {
	return C.GoString(C.vte_terminal_get_window_title(C.toVTerminal(t.Widget)))
}

func (t Terminal) SetFontFromString(font string) {
	C.vte_terminal_set_font_from_string(C.toVTerminal(t.Widget), C.CString(font))
}

func (t Terminal) SetFont(fd *C.PangoFontDescription) {
	C.vte_terminal_set_font(C.toVTerminal(t.Widget), fd)
}

func (t Terminal) GetFont() *C.PangoFontDescription {
	return C.vte_terminal_get_font(C.toVTerminal(t.Widget))
}

// func (t Terminal) SetFontSize() {
// 	origin := &pango.FontDescription{(*C.PangoFontDescription)(t.GetFont())}
// 	newfont := origin.Copy()
// 	newfont.SetSize(+1)
// 	t.SetFont(newfont.GFontDescription)
// }

func createColor(pattern string) C.GdkColor {
	var color C.GdkColor
	ptr := C.CString(pattern)
	defer C.free(unsafe.Pointer(ptr))
	C.gdk_color_parse(C.toGstr(ptr), &color)
	return color
}

func (t Terminal) SetColors(foreground string, background string, palette []string) {
	fColor := createColor(foreground)
	bColor := createColor(background)
	pColors := new([16]C.GdkColor)
	for i := 0; i < len(pColors); i++ {
		C.gdk_color_parse((*C.gchar)(C.CString(palette[i])), &pColors[i])
	}
	C.vte_terminal_set_colors(C.toVTerminal(t.Widget), &fColor, &bColor, (*C.GdkColor)(unsafe.Pointer(pColors)), 16)
}

func (t Terminal) SetColorCursor(pattern string) {
	color := createColor(pattern)
	C.vte_terminal_set_color_cursor(C.toVTerminal(t.Widget), &color)
}

func (t Terminal) SetColorForeground(pattern string) {
	color := createColor(pattern)
	C.vte_terminal_set_color_foreground(C.toVTerminal(t.Widget), &color)
}

func (t Terminal) SetColorBackground(pattern string) {
	color := createColor(pattern)
	C.vte_terminal_set_color_background(C.toVTerminal(t.Widget), &color)
}

// func (t Terminal) GetCurrentDirectory() string {
// 	return C.GoString(C.vte_terminal_get_current_directory_uri(C.toVTerminal(t.Widget)))
// }

func (t Terminal) Copy() {
	// C.vte_terminal_copy_clipboard(C.toVTerminal(t.Widget))
	C.vte_terminal_copy_primary(C.toVTerminal(t.Widget))
}

func (t Terminal) Paste() {
	C.vte_terminal_paste_primary(C.toVTerminal(t.Widget))
}

func (t Terminal) SetScrollbackLines(n int) {
	C.vte_terminal_set_scrollback_lines(C.toVTerminal(t.Widget), C.glong(n))
}

func (t Terminal) SetWordChars(chars string) {
	ptr := C.CString(chars)
	defer C.free(unsafe.Pointer(ptr))

	C.vte_terminal_set_word_chars(C.toVTerminal(t.Widget), ptr)
}
