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

var HWKeyCodes = map[uint16]string{
	9:  "ESC",
	10: "1",
	11: "2",
	12: "3",
	13: "4",
	14: "5",
	15: "6",
	16: "7",
	17: "8",
	18: "9",
	19: "0",
	20: "-",
	21: "=",
	22: "BACKSPACE",
	23: "TAB",
	24: "Q",
	25: "W",
	26: "E",
	27: "R",
	28: "T",
	29: "Y",
	30: "U",
	31: "I",
	32: "O",
	33: "P",
	34: "[",
	35: "]",
	36: "ENTER",
	37: "LEFTCTRL",
	38: "A",
	39: "S",
	40: "D",
	41: "F",
	42: "G",
	43: "H",
	44: "J",
	45: "K",
	46: "L",
	47: ";",
	48: "'",
	49: "`",
	50: "LEFTSHIFT",
	51: "\\",
	52: "Z",
	53: "X",
	54: "C",
	55: "V",
	56: "B",
	57: "N",
	58: "M",
	59: ",",
	60: ".",
	61: "/",
	62: "RIGHTSHIFT",
	63: "*",
	64: "LEFTALT",
	65: "SPACE",
	66: "CAPSLOCK",
	67: "F1",
	68: "F2",
	69: "F3",
	70: "F4",
	71: "F5",
	72: "F6",
	73: "F7",
	74: "F8",
	75: "F9",
	76: "F10",
	77: "NUMLOCK",
	78: "SCRLOCK",
	79: "NUM7",
	80: "NUM8",
	81: "NUM9",
	82: "NUM-",
	83: "NUM4",
	84: "NUM5",
	85: "NUM6",
	86: "NUM+",
	87: "NUM1",
	88: "NUM2",
	89: "NUM3",
	90: "NUM0",
	91: "NUM.",
	94: "LEFTALT",
	95: "F11",
	96: "F12",

	104: "NUMENTER",
	105: "RIGHTCTRL",
	106: "NUM/",
	107: "PRTSCN",
	108: "RIGHTALT",

	110: "HOME",
	111: "TOP",
	112: "PAGEUP",
	113: "LEFT",
	114: "RIGHT",
	115: "END",
	116: "DOWN",
	117: "PAGEDOWN",
	118: "INSERT",
	119: "DELETE",

	121: "MUTE",
	122: "VOL-",
	123: "VOL+",
	127: "PAUSE",
	133: "SUPER",
	135: "MENU",
	148: "CALC",
	172: "PLAY/PAUSE",
}

func getKey(i uint16) string {
	if c, ok := HWKeyCodes[i]; ok {
		return c
	}
	return ""
}

func getMod(m uint32) (ctrl, alt, shift, super bool) {
	ctrl = m&Ctrl != 0
	alt = m&Alt != 0
	shift = m&Shift != 0
	super = m&Super != 0

	return
}

type Key struct {
	Key   string
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
	ss := strings.Split(key, " + ")
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
			k.Key = s
		}
	}

	if k.Key != "" {
		ok = true
	}

	return
}
