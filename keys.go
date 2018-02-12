package main

import (
	"strings"

	"github.com/mattn/go-gtk/gdk"
)

type Mod int

const (
	None  Mod = 0
	Shift     = 1
	Ctrl      = 4
	Alt       = 8
	Super     = 0x4000040
)

var KeyCodes = map[uint16]byte{
	24: 'Q',
	25: 'W',
	26: 'E',
	27: 'R',
	28: 'T',
	29: 'Y',
	30: 'U',
	31: 'I',
	32: 'O',
	33: 'P',
	38: 'A',
	39: 'S',
	40: 'D',
	41: 'F',
	42: 'G',
	43: 'H',
	44: 'J',
	45: 'K',
	46: 'L',
	52: 'Z',
	53: 'X',
	54: 'C',
	55: 'V',
	56: 'B',
	57: 'N',
	58: 'M',
	59: ',',
	60: '.',
	61: '/',
	47: ';',
	48: '\'',
	34: '[',
	35: ']',
	49: '`',
	10: '1',
	11: '2',
	12: '3',
	13: '4',
	14: '5',
	15: '6',
	16: '7',
	17: '8',
	18: '9',
	19: '0',
	20: '-',
	21: '=',
	51: '\\',
}

func getKey(i uint16) byte {
	if c, ok := KeyCodes[i]; ok {
		return c
	}
	return 0
}

func getMod(m uint32) (ctrl, alt, shift, super bool) {
	ctrl = m&Ctrl != 0
	alt = m&Alt != 0
	shift = m&Shift != 0
	super = m&Super != 0

	return
}

type Key struct {
	Char  byte
	Ctrl  bool
	Alt   bool
	Shift bool
	SUper bool
}

func NewKey(ke *gdk.EventKey) Key {
	c := getKey(ke.HardwareKeycode)
	ctrl, alt, shift, super := getMod(uint32(ke.State))
	return Key{c, ctrl, alt, shift, super}
}

func ParseKey(key string) (k Key, ok bool) {
	ss := strings.Split(key, "+")
	for _, s := range ss {
		s = strings.TrimSpace(s)
		if len(s) == 0 {
			continue
		}

		s = strings.ToUpper(s)

		switch s {
		case "CTRL":
			k.Ctrl = true
		case "ALT":
			k.Alt = true
		case "SHIFT":
			k.Shift = true
		case "SUPER":
			k.SUper = true
		default:
			k.Char = []byte(s)[0]
		}
	}

	if k.Char != 0 {
		ok = true
	}

	return
}
