package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func fg(code int) string {
	colored := []string{"\x1b[38;5;", strconv.Itoa(code), "m"}
	return strings.Join(colored, "")
}

func bg(code int) string {
	colored := []string{"\x1b[48;5;", strconv.Itoa(code), "m"}
	return strings.Join(colored, "")
}

func rgb2code(r int, g int, b int) int {
	code := 36*r + 6*g + b + 16
	if code < 16 || 231 < code {
		panic(errors.New(fmt.Sprintf("Invalid RGB values (%i, %i, %i)", r, g, b)))
	}
	return code
}

func gray2code(lightness int) int {
	code := lightness + 232
	if code < 232 || 255 < code {
		panic(errors.New(fmt.Sprintf("Invalid lightness value (%i) for gray", lightness)))
	}
	return code
}

func fgRGB(r int, g int, b int) string {
	return fg(rgb2code(r, g, b))
}

func bgRGB(r int, g int, b int) string {
	return bg(rgb2code(r, g, b))
}

func fgGray(lightness int) string {
	return fg(gray2code(lightness))
}

func bgGray(lightness int) string {
	return bg(gray2code(lightness))
}

var (
	Reset = "\x1b[0m"
	Bold  = "\x1b[1m"
)

/*
    'coffee':   [ bold+fgc(136),  fgc(130) ]
    'js':       [ bold+fg(4,4,0), fg(2,2,0) ]
    'json':     [ bold+fg(4,4,0), fg(2,2,0) ]
    'cson':     [ bold+fgc(136),  fgc(130) ]
    'jsx':      [ bold+fgc(14),   fgc(6) ]
    'plist':    [ bold+fg(4,0,4), fg(2,0,2) ]
    'sh':       [ bold+fg(4,0,4), fg(2,0,2) ]
    'bash':     [ bold+fg(4,0,4), fg(2,0,2) ]
    'zsh':      [ bold+fg(4,0,4), fg(2,0,2) ]
    'cpp':      [ bold+fg(4,0,4), fg(2,0,2) ]
    'h':        [ bold+fg(4,0,4), fg(2,0,2) ]
    'py':       [ bold+fg(0,3,0), fg(0,1,0) ]
    'pyc':      [      fw(8),     fw(5) ]
    'rb':       [ bold+fg(5,1,0), fg(3,0,0) ]
    'java':     [ bold+fgc(136),  fgc(130) ]
    'go':       [ bold+fgc(123),  fgc(117) ]
    'scala':    [ bold+fgc(124),  fgc(88) ]
    'c':        [ bold+fgc(99),   fgc(97) ]
    'cpp':      [ bold+fgc(99),   fgc(97) ]
    'php':      [ bold+fgc(30),   fgc(22) ]
    'log':      [      fw(8),     fw(5) ]
    'swp':      [      fw(8),     fw(5) ]
    'md':       [      fgc(87),   fgc(73) ]
    'markdown': [      fgc(87),   fgc(73) ]
    'html':     [      fgc(87),   fgc(73) ]
    'css':      [      fgc(219),  fgc(207) ]
    'scss':     [      fgc(219),  fgc(207) ]
    'tar':      [      fg(5,0,0), fg(3,0,0) ]
    'gz':       [      fg(5,0,0), fg(3,0,0) ]
    'tgz':      [      fg(5,0,0), fg(3,0,0) ]
    'zip':      [      fg(5,0,0), fg(3,0,0) ]
    'rar':      [      fg(5,0,0), fg(3,0,0) ]
    #
    '_default': [      fw(23),    fw(12) ]
    '_dir':     [ bold+BG(0,0,2)+fw(23), fg(1,1,5), fg(2,2,5) ]
    '_.dir':    [ bold+BG(0,0,1)+fw(23), bold+BG(0,0,1)+fg(1,1,5), bold+BG(0,0,1)+fg(2,2,5) ]
    '_link':    { 'arrow': fg(1,0,1), 'path': fg(4,0,4), 'broken': BG(5,0,0)+fg(5,5,0) }
    '_arrow':     fw(1)
    '_header':  [ bold+BW(2)+fg(3,2,0),  fw(4), bold+BW(2)+fg(5,5,0) ]
    '_media':   [      fgc(141),  fgc(54) ]
    #
    '_size':    { b: fgc(20), kB: fgc(33), MB: fgc(81), GB: fgc(123) }
    '_users':   { root:  fg(5,0,2), default: fg(0,3,3) }
    '_groups':  { wheel: fg(3,0,0), staff: fg(0,2,0), admin: fg(2,2,0), default: fg(2,0,2) }
    '_error':   [ bold+BG(5,0,0)+fg(5,5,0), bold+BG(5,0,0)+fg(5,5,5) ]

mediaTypes = new Set ['png', 'gif', 'jpg', 'jpeg', 'ico', 'svg', 'webp', 'tiff', 'pxm', 'mp3', 'm4a', 'wav', 'webm', 'avi', 'wmv', 'mkv', 'mp4', 'flv', 'mov']
*/
var (
	ConfigColor = map[string]map[string]string{
		"dir": map[string]string{
			"name": Bold + bgRGB(0, 0, 2) + fgGray(23),
			"ext":  fgRGB(2, 2, 5),
		},
		".dir": map[string]string{
			"name": Bold + bgRGB(0, 0, 1) + fgGray(23),
			"ext":  fgRGB(2, 2, 5),
		},
		"folderHeader": map[string]string{
			"arrow":      fg(58),
			"main":       bgGray(2) + fgRGB(3, 2, 0),
			"slash":      fgGray(4),
			"lastFolder": Bold + fgRGB(5, 5, 0),
		},
	}
)
