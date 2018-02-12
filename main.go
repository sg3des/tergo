package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"

	"github.com/mattn/go-gtk/gtk"
)

func init() {
	log.SetFlags(log.Lshortfile)

	defaultConf()
	readConf()

	log.Printf("%#v", &Conf)
}

func main() {
	gtk.Init(nil)

	w := NewWindow("tergo", 400, 300)
	w.SetBinds(Conf.Binds)
	w.NewTab()
	w.ShowAll()

	gtk.Main()
}

//Configuration

//Conf contains global configuration
var Conf struct {
	TabCloseButton bool
	TabHeight      int

	Font string

	Background  string
	Foreground  string
	CursorColor string
	Palette     []string

	Binds map[string]string
}

func defaultConf() {
	Conf.TabCloseButton = true
	Conf.TabHeight = 16
	Conf.Font = "Liberation Mono 7.5"
	Conf.Background = "#222"
	Conf.Foreground = "#bbb"
	Conf.CursorColor = "#e8e8e8"
	Conf.Palette = []string{"#2e3436", "#cc0000", "#4e9a06", "#c4a000", "#3465a4", "#75507b", "#06989a", "#b1b1b1", "#555753", "#ef2929", "#8ae234", "#fce94f", "#729fcf", "#ad7fa8", "#34e2e2", "#acacac"}

	Conf.Binds = map[string]string{
		"NewTab":   "ctrl+t",
		"CloseTab": "ctrl+w",

		"Copy":  "shift+ctrl+c",
		"Paste": "shift+ctrl+v",

		"Quit": "ctrl+q",
	}
}

func readConf() {
	data, ok := lookupConf()
	if !ok {
		return
	}
	log.Println(string(data))

	err := yaml.Unmarshal(data, &Conf)
	log.Println(err)
	if err != nil {
		fmt.Println("failed read config by reason:", err)
	}
}

func lookupConf() ([]byte, bool) {
	filenames := []string{
		"testdata/tergo.conf",
		"tergo.conf",
		filepath.Join(os.Getenv("XDG_CONFIG_HOME"), "tergo/tergo.conf"),
		filepath.Join(os.Getenv("HOME"), ".config/tergo/tergo.conf"),
	}

	for _, filename := range filenames {
		log.Println(filename)
		data, err := ioutil.ReadFile(filename)
		if err == nil {
			return data, true
		}
	}

	return nil, false
}
