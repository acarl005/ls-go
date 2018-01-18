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

var (
	MediaTypes    = set.New("png", "gif", "jpg", "jpeg", "ico", "svg", "webp", "tiff", "pxm", "mp3", "m4a", "wav", "webm", "avi", "wmv", "mkv", "mp4", "flv", "mov")
	CompressTypes = set.New("tar", "gz", "tgz", "zip", "rar")
	FileColor     = map[string][2]string{
		"js":        [2]string{FgRGB(4, 4, 0), FgRGB(2, 2, 0)},
		"json":      [2]string{FgRGB(4, 4, 0), FgRGB(2, 2, 0)},
		"jsx":       [2]string{Fg(87), Fg(73)},
		"coffee":    [2]string{Fg(136), Fg(94)},
		"cson":      [2]string{Fg(136), Fg(94)},
		"java":      [2]string{Fg(136), Fg(94)},
		"plist":     [2]string{FgRGB(4, 0, 4), FgRGB(2, 0, 2)},
		"sh":        [2]string{FgRGB(4, 0, 4), FgRGB(2, 0, 2)},
		"bash":      [2]string{FgRGB(4, 0, 4), FgRGB(2, 0, 2)},
		"zsh":       [2]string{FgRGB(4, 0, 4), FgRGB(2, 0, 2)},
		"py":        [2]string{FgRGB(0, 3, 0), FgRGB(0, 1, 0)},
		"ipynb":     [2]string{FgRGB(0, 3, 0), FgRGB(0, 1, 0)},
		"pyc":       [2]string{FgGray(8), FgGray(5)},
		"rb":        [2]string{FgRGB(5, 1, 0), FgRGB(3, 1, 1)},
		"go":        [2]string{Fg(121), Fg(109)},
		"scala":     [2]string{Fg(124), Fg(52)},
		"ts":        [2]string{Fg(33), Fg(20)},
		"c":         [2]string{Fg(33), Fg(20)},
		"cpp":       [2]string{Fg(33), Fg(20)},
		"h":         [2]string{Fg(33), Fg(20)},
		"cs":        [2]string{Fg(33), Fg(20)},
		"rs":        [2]string{Fg(208), Fg(94)},
		"hs":        [2]string{Fg(99), Fg(57)},
		"m":         [2]string{Fg(33), Fg(20)},
		"php":       [2]string{Fg(30), Fg(22)},
		"sql":       [2]string{Fg(193), Fg(148)},
		"tsql":      [2]string{Fg(193), Fg(148)},
		"psql":      [2]string{Fg(193), Fg(148)},
		"plsql":     [2]string{Fg(193), Fg(148)},
		"plpgsql":   [2]string{Fg(193), Fg(148)},
		"swift":     [2]string{Fg(223), Fg(215)},
		"log":       [2]string{FgGray(8), FgGray(5)},
		"lock":      [2]string{FgGray(8), FgGray(5)},
		"md":        [2]string{Fg(87), Fg(73)},
		"markdown":  [2]string{Fg(87), Fg(73)},
		"html":      [2]string{Fg(87), Fg(73)},
		"xml":       [2]string{Fg(87), Fg(73)},
		"css":       [2]string{Fg(219), Fg(207)},
		"scss":      [2]string{Fg(219), Fg(207)},
		"_compress": [2]string{FgRGB(5, 0, 0), FgRGB(3, 0, 0)},
		"_media":    [2]string{Fg(141), Fg(99)},
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
			"name": BgRGB(0, 0, 2) + FgGray(23),
			"ext":  FgRGB(2, 2, 5),
		},
		".dir": map[string]string{
			"name": BgRGB(0, 0, 1) + FgGray(23),
			"ext":  FgRGB(2, 2, 5),
		},
		"folderHeader": map[string]string{
			"arrow":      FgRGB(3, 2, 0),
			"main":       BgGray(2) + FgRGB(3, 2, 0),
			"slash":      FgGray(5),
			"lastFolder": Bold + FgRGB(5, 5, 0),
			"error":      BgRGB(5, 0, 0) + FgRGB(5, 5, 0),
		},
		"link": map[string]string{
			"arrow":  FgRGB(1, 0, 1),
			"path":   FgRGB(4, 0, 4),
			"broken": BgRGB(5, 0, 0) + FgRGB(5, 5, 0),
		},
		"stats": map[string]string{
			"text":   BgGray(2) + FgGray(15),
			"number": Fg(24),
			"ms":     Fg(39),
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
