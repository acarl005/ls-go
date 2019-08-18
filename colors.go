package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// see the color codes
// http://i.stack.imgur.com/UQVe5.png

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
	FileColor = map[string][2]string{
		"c":       [2]string{Fg(33), Fg(20)},
		"clj":     [2]string{Fg(204), Fg(162)},
		"coffee":  [2]string{Fg(136), Fg(94)},
		"cr":      [2]string{Fg(82), Fg(70)},
		"cs":      [2]string{Fg(33), Fg(20)},
		"cson":    [2]string{Fg(136), Fg(94)},
		"css":     [2]string{Fg(219), Fg(207)},
		"dart":    [2]string{Fg(43), Fg(31)},
		"diff":    [2]string{Fg(10), Fg(9)},
		"elm":     [2]string{Fg(51), Fg(39)},
		"erl":     [2]string{Fg(161), Fg(89)},
		"ex":      [2]string{Fg(99), Fg(57)},
		"graphql": [2]string{Fg(219), Fg(207)},
		"groovy":  [2]string{Fg(223), Fg(215)},
		"go":      [2]string{Fg(121), Fg(109)},
		"hs":      [2]string{Fg(99), Fg(57)},
		"html":    [2]string{Fg(87), Fg(73)},
		"java":    [2]string{Fg(136), Fg(94)},
		"js":      [2]string{FgRGB(4, 4, 0), FgRGB(2, 2, 0)},
		"mjs":     [2]string{FgRGB(4, 4, 0), FgRGB(2, 2, 0)},
		"json":    [2]string{FgRGB(4, 4, 0), FgRGB(2, 2, 0)},
		"jsx":     [2]string{Fg(87), Fg(73)},
		"lock":    [2]string{FgGray(8), FgGray(5)},
		"log":     [2]string{FgGray(8), FgGray(5)},
		"lua":     [2]string{Fg(33), Fg(20)},
		"m":       [2]string{Fg(33), Fg(20)},
		"md":      [2]string{Fg(87), Fg(73)},
		"php":     [2]string{Fg(30), Fg(22)},
		"plist":   [2]string{FgRGB(4, 0, 4), FgRGB(2, 0, 2)},
		"py":      [2]string{FgRGB(0, 3, 0), FgRGB(0, 1, 0)},
		"r":       [2]string{Fg(51), Fg(39)},
		"rb":      [2]string{FgRGB(5, 1, 0), FgRGB(3, 1, 1)},
		"rs":      [2]string{Fg(208), Fg(94)},
		"scala":   [2]string{Fg(124), Fg(52)},
		"sh":      [2]string{FgRGB(4, 0, 4), FgRGB(2, 0, 2)},
		"sol":     [2]string{Fg(33), Fg(20)},
		"sql":     [2]string{Fg(193), Fg(148)},
		"swift":   [2]string{Fg(223), Fg(215)},
		"vue":     [2]string{Fg(43), Fg(29)},
		"xml":     [2]string{Fg(87), Fg(73)},

		"compiled": [2]string{FgGray(8), FgGray(5)},
		"compress": [2]string{FgRGB(5, 0, 0), FgRGB(3, 0, 0)},
		"document": [2]string{FgRGB(5, 0, 0), FgRGB(3, 0, 0)},
		"media":    [2]string{Fg(141), Fg(99)},
		"_default": [2]string{FgGray(23), FgGray(12)},
	}
	FileAliases = map[string]string{
		"d":        "c",
		"cpp":      "c",
		"cxx":      "c",
		"c++":      "c",
		"cc":       "c",
		"h":        "c",
		"hpp":      "c",
		"hxx":      "c",
		"h++":      "c",
		"hh":       "c",
		"ts":       "js",
		"tsx":      "jsx",
		"bash":     "sh",
		"zsh":      "sh",
		"fish":     "sh",
		"ksh":      "sh",
		"csh":      "sh",
		"ipynb":    "py",
		"pickle":   "py",
		"pkl":      "py",
		"tsql":     "sql",
		"psql":     "sql",
		"plsql":    "sql",
		"plpgsql":  "sql",
		"less":     "css",
		"sass":     "css",
		"scss":     "css",
		"styl":     "css",
		"gvy":      "groovy",
		"gy":       "groovy",
		"gsh":      "groovy",
		"markdown": "md",
		"mkd":      "md",
		"png":      "media",
		"gif":      "media",
		"jpg":      "media",
		"jpeg":     "media",
		"ico":      "media",
		"svg":      "media",
		"webp":     "media",
		"bmp":      "media",
		"ppm":      "media",
		"pgm":      "media",
		"pbm":      "media",
		"pnm":      "media",
		"stl":      "media",
		"eps":      "media",
		"cbr":      "media",
		"cbz":      "media",
		"xpm":      "media",
		"orf":      "media",
		"nef":      "media",
		"tiff":     "media",
		"pxm":      "media",
		"mp3":      "media",
		"m4a":      "media",
		"wav":      "media",
		"flac":     "media",
		"alac":     "media",
		"aac":      "media",
		"ogg":      "media",
		"wma":      "media",
		"mka":      "media",
		"opus":     "media",
		"webm":     "media",
		"avi":      "media",
		"wmv":      "media",
		"mkv":      "media",
		"mp4":      "media",
		"flv":      "media",
		"mov":      "media",
		"m2v":      "media",
		"mpeg":     "media",
		"mpg":      "media",
		"ogm":      "media",
		"ogv":      "media",
		"vob":      "media",
		"video":    "media",
		"audio":    "media",
		"image":    "media",
		"Z":        "compress",
		"z":        "compress",
		"bz2":      "compress",
		"7z":       "compress",
		"iso":      "compress",
		"dmg":      "compress",
		"tc":       "compress",
		"par":      "compress",
		"xz":       "compress",
		"txz":      "compress",
		"lzma":     "compress",
		"deb":      "compress",
		"rpm":      "compress",
		"dpkg":     "compress",
		"jar":      "compress",
		"tar":      "compress",
		"gz":       "compress",
		"tgz":      "compress",
		"zip":      "compress",
		"rar":      "compress",
		"whl":      "compress",
		"pdf":      "document",
		"djvu":     "document",
		"doc":      "document",
		"docx":     "document",
		"dvi":      "document",
		"eml":      "document",
		"fotd":     "document",
		"odp":      "document",
		"odt":      "document",
		"ppt":      "document",
		"pptx":     "document",
		"rtf":      "document",
		"xls":      "document",
		"xlsx":     "document",
		"pyc":      "compiled",
		"class":    "compiled",
		"elc":      "compiled",
		"o":        "compiled",
		"hi":       "compiled",
	}
	SizeColor = map[string]string{
		"B": Fg(27),
		"K": Fg(33),
		"M": Fg(81),
		"G": Fg(123),
		"T": Fg(159),
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
			"error":      BgRGB(5, 0, 0) + FgRGB(5, 5, 0),
		},
		"link": map[string]string{
			"name":   Bold + FgRGB(0, 5, 0),
			"arrow":  FgRGB(1, 0, 1),
			"path":   FgRGB(4, 0, 4),
			"broken": BgRGB(5, 0, 0) + FgRGB(5, 5, 0),
		},
		"device": map[string]string{
			"name": Bold + BgGray(3) + Fg(220),
		},
		"socket": map[string]string{
			"name": Bold + Bg(53) + Fg(15),
		},
		"pipe": map[string]string{
			"name": Bold + Bg(94) + Fg(15),
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
