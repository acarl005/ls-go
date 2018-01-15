package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	set "gopkg.in/fatih/set.v0"
)

func Fg(code int) string {
	colored := []string{"\x1b[38;5;", strconv.Itoa(code), "m"}
	return strings.Join(colored, "")
}

func Bg(code int) string {
	colored := []string{"\x1b[48;5;", strconv.Itoa(code), "m"}
	return strings.Join(colored, "")
}

func Rgb2code(r int, g int, b int) int {
	code := 36*r + 6*g + b + 16
	if code < 16 || 231 < code {
		panic(errors.New(fmt.Sprintf("Invalid RGB values (%i, %i, %i)", r, g, b)))
	}
	return code
}

func Gray2code(lightness int) int {
	code := lightness + 232
	if code < 232 || 255 < code {
		panic(errors.New(fmt.Sprintf("Invalid lightness value (%i) for gray", lightness)))
	}
	return code
}

func FgRGB(r int, g int, b int) string {
	return Fg(Rgb2code(r, g, b))
}

func BgRGB(r int, g int, b int) string {
	return Bg(Rgb2code(r, g, b))
}

func FgGray(lightness int) string {
	return Fg(Gray2code(lightness))
}

func BgGray(lightness int) string {
	return Bg(Gray2code(lightness))
}

const (
	Reset = "\x1b[0m"
	Bold  = "\x1b[1m"
)

/*
    'coffee':   [ bold+fgc(136),  fgc(130) ]
    'js':       [ bold+Fg(4,4,0), Fg(2,2,0) ]
    'json':     [ bold+Fg(4,4,0), Fg(2,2,0) ]
    'cson':     [ bold+fgc(136),  fgc(130) ]
    'jsx':      [ bold+fgc(14),   fgc(6) ]
    'plist':    [ bold+Fg(4,0,4), Fg(2,0,2) ]
    'sh':       [ bold+Fg(4,0,4), Fg(2,0,2) ]
    'bash':     [ bold+Fg(4,0,4), Fg(2,0,2) ]
    'zsh':      [ bold+Fg(4,0,4), Fg(2,0,2) ]
    'cpp':      [ bold+Fg(4,0,4), Fg(2,0,2) ]
    'h':        [ bold+Fg(4,0,4), Fg(2,0,2) ]
    'py':       [ bold+Fg(0,3,0), Fg(0,1,0) ]
    'pyc':      [      fw(8),     fw(5) ]
    'rb':       [ bold+Fg(5,1,0), Fg(3,0,0) ]
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
    'tar':      [      Fg(5,0,0), Fg(3,0,0) ]
    'gz':       [      Fg(5,0,0), Fg(3,0,0) ]
    'tgz':      [      Fg(5,0,0), Fg(3,0,0) ]
    'zip':      [      Fg(5,0,0), Fg(3,0,0) ]
    'rar':      [      Fg(5,0,0), Fg(3,0,0) ]
    #
    '_default': [      fw(23),    fw(12) ]
    '_dir':     [ bold+BG(0,0,2)+fw(23), Fg(1,1,5), Fg(2,2,5) ]
    '_.dir':    [ bold+BG(0,0,1)+fw(23), bold+BG(0,0,1)+Fg(1,1,5), bold+BG(0,0,1)+Fg(2,2,5) ]
    '_link':    { 'arrow': Fg(1,0,1), 'path': Fg(4,0,4), 'broken': BG(5,0,0)+Fg(5,5,0) }
    '_arrow':     fw(1)
    '_header':  [ bold+BW(2)+Fg(3,2,0),  fw(4), bold+BW(2)+Fg(5,5,0) ]
    '_media':   [      fgc(141),  fgc(54) ]
    #
    '_size':    { b: fgc(20), kB: fgc(33), MB: fgc(81), GB: fgc(123) }
    '_users':   { root:  Fg(5,0,2), default: Fg(0,3,3) }
    '_groups':  { wheel: Fg(3,0,0), staff: Fg(0,2,0), admin: Fg(2,2,0), default: Fg(2,0,2) }
    '_error':   [ bold+BG(5,0,0)+Fg(5,5,0), bold+BG(5,0,0)+Fg(5,5,5) ]

mediaTypes = new Set ['png', 'gif', 'jpg', 'jpeg', 'ico', 'svg', 'webp', 'tiff', 'pxm', 'mp3', 'm4a', 'wav', 'webm', 'avi', 'wmv', 'mkv', 'mp4', 'flv', 'mov']
*/
var (
	MediaTypes    = set.New("png", "gif", "jpg", "jpeg", "ico", "svg", "webp", "tiff", "pxm", "mp3", "m4a", "wav", "webm", "avi", "wmv", "mkv", "mp4", "flv", "mov")
	CompressTypes = set.New("tar", "gz", "tgz", "zip", "rar")
	FileColor     = map[string][2]string{
		"coffee":    [2]string{Bold + Fg(136), Fg(130)},
		"js":        [2]string{Bold + FgRGB(4, 4, 0), FgRGB(2, 2, 0)},
		"json":      [2]string{Bold + FgRGB(4, 4, 0), FgRGB(2, 2, 0)},
		"cson":      [2]string{Bold + Fg(136), Fg(130)},
		"jsx":       [2]string{Bold + Fg(14), Fg(6)},
		"plist":     [2]string{Bold + FgRGB(4, 0, 4), FgRGB(2, 0, 2)},
		"sh":        [2]string{Bold + FgRGB(4, 0, 4), FgRGB(2, 0, 2)},
		"bash":      [2]string{Bold + FgRGB(4, 0, 4), FgRGB(2, 0, 2)},
		"zsh":       [2]string{Bold + FgRGB(4, 0, 4), FgRGB(2, 0, 2)},
		"py":        [2]string{Bold + FgRGB(0, 3, 0), FgRGB(0, 1, 0)},
		"pyc":       [2]string{FgGray(8), FgGray(5)},
		"rb":        [2]string{Bold + FgRGB(5, 1, 0), FgRGB(3, 0, 0)},
		"java":      [2]string{Bold + Fg(136), Fg(130)},
		"go":        [2]string{Bold + Fg(121), Fg(115)},
		"scala":     [2]string{Bold + Fg(124), Fg(88)},
		"c":         [2]string{Bold + Fg(99), Fg(97)},
		"cpp":       [2]string{Bold + Fg(99), Fg(97)},
		"h":         [2]string{Bold + Fg(99), Fg(97)},
		"php":       [2]string{Bold + Fg(30), Fg(22)},
		"log":       [2]string{FgGray(8), FgGray(5)},
		"swp":       [2]string{FgGray(8), FgGray(5)},
		"md":        [2]string{Fg(87), Fg(73)},
		"markdown":  [2]string{Fg(87), Fg(73)},
		"html":      [2]string{Fg(87), Fg(73)},
		"xml":       [2]string{Fg(87), Fg(73)},
		"css":       [2]string{Fg(219), Fg(207)},
		"scss":      [2]string{Fg(219), Fg(207)},
		"_compress": [2]string{FgRGB(5, 0, 0), FgRGB(3, 0, 0)},
		"_media":    [2]string{Fg(141), Fg(54)},
		"_default":  [2]string{FgGray(23), FgGray(12)},
	}
	SizeColor = map[string]string{
		" B": Fg(27),
		"kB": Fg(33),
		"MB": Fg(81),
		"GB": Fg(123),
		"TB": Fg(159),
	}
	ConfigColor = map[string]map[string]string{
		"dir": map[string]string{
			"name": Bold + BgRGB(0, 0, 2) + FgGray(23),
			"ext":  FgRGB(2, 2, 5),
		},
		".dir": map[string]string{
			"name": Bold + BgRGB(0, 0, 1) + FgGray(23),
			"ext":  FgRGB(2, 2, 5),
		},
		"folderHeader": map[string]string{
			"arrow":      FgRGB(3, 2, 0),
			"main":       BgGray(2) + FgRGB(3, 2, 0),
			"slash":      FgGray(5),
			"lastFolder": Bold + FgRGB(5, 5, 0),
		},
	}
	PermsColor = map[string]map[string]string{
		"user": map[string]string{
			"root":     FgRGB(5, 0, 2),
			"daemon":   FgRGB(4, 2, 1),
			"_self":    FgRGB(0, 4, 0),
			"_default": FgRGB(0, 3, 3),
		},
		"group": map[string]string{
			"wheel":    FgRGB(3, 0, 0),
			"staff":    FgRGB(0, 2, 0),
			"admin":    FgRGB(2, 2, 0),
			"_default": FgRGB(2, 0, 2),
		},
		"other": map[string]string{
			"_default": BgGray(2) + FgGray(15),
		},
	}
)
