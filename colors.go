package main

import (
	"fmt"
	"strconv"
	"strings"
)

// see the 256-color codes
// http://i.stack.imgur.com/UQVe5.png

type Color int8

const (
	Black Color   = 0
	Red           = 1
	Green         = 2
	Yellow        = 3
	Blue          = 4
	Magenta       = 5
	Cyan          = 6
	White         = 7
	BrightBlack   = 60
	BrightRed     = 61
	BrightGreen   = 62
	BrightYellow  = 63
	BrightBlue    = 64
	BrightMagenta = 65
	BrightCyan    = 66
	BrightWhite   = 67
)

func NamedFg(color Color) string {
	colored := []string{"\x1b[", strconv.Itoa(int(color) + 30), "m"}
	return strings.Join(colored, "")
}

func NamedBg(color Color) string {
	colored := []string{"\x1b[", strconv.Itoa(int(color) + 40), "m"}
	return strings.Join(colored, "")
}

// Fg wraps an 8-bit foreground color code in the ANSI escape sequence
func Fg(code int) string {
	colored := []string{"\x1b[38;5;", strconv.Itoa(code), "m"}
	return strings.Join(colored, "")
}

// Bg wraps an 8-bit background color code in the ANSI escape sequence
func Bg(code int) string {
	colored := []string{"\x1b[48;5;", strconv.Itoa(code), "m"}
	return strings.Join(colored, "")
}

// Rgb2code converts RGB values (up to 5) to an 8-bit color code
func Rgb2code(r int, g int, b int) int {
	code := 36*r + 6*g + b + 16
	if code < 16 || 231 < code {
		panic(fmt.Errorf("Invalid RGB values (%d, %d, %d)", r, g, b))
	}
	return code
}

// Gray2code converts a scalar of "grayness" to an 8-bit color code
func Gray2code(lightness int) int {
	code := lightness + 232
	if code < 232 || 255 < code {
		panic(fmt.Errorf("Invalid lightness value (%d) for gray", lightness))
	}
	return code
}

// FgRGB converts RGB values (up to 5) to an ANSI-escaped foreground 8-bit color code
func FgRGB(r int, g int, b int) string {
	return Fg(Rgb2code(r, g, b))
}

// BgRGB converts RGB values (up to 5) to an ANSI-escaped background 8-bit color code
func BgRGB(r int, g int, b int) string {
	return Bg(Rgb2code(r, g, b))
}

// FgGray converts a scalar of "grayness" to an ANSI-escaped foreground 8-bit color code
func FgGray(lightness int) string {
	return Fg(Gray2code(lightness))
}

// BgGray converts a scalar of "grayness" to an ANSI-escaped background 8-bit color code
func BgGray(lightness int) string {
	return Bg(Gray2code(lightness))
}

const (
	// Reset undoes ANSI color codes
	Reset = "\x1b[0m"
	Bold = "\x1b[1m"
)

var (
	// FileColor is a mapping of filetypes to colors
	FileColor = map[string][2]string{
		"as":      {Fg(196), Fg(124)},
		"asm":     {Fg(223), Fg(215)},
		"bf":      {Fg(223), Fg(215)},
		"c":       {Fg(39), Fg(27)},
		"clj":     {Fg(204), Fg(162)},
		"coffee":  {Fg(136), Fg(94)},
		"cr":      {Fg(82), Fg(70)},
		"cson":    {Fg(136), Fg(94)},
		"css":     {Fg(219), Fg(207)},
		"dart":    {Fg(43), Fg(31)},
		"diff":    {Fg(10), Fg(9)},
		"elm":     {Fg(51), Fg(39)},
		"erl":     {Fg(161), Fg(89)},
		"ex":      {Fg(99), Fg(57)},
		"f":       {Fg(208), Fg(94)},
		"fs":      {Fg(45), Fg(32)},
		"gb":      {Fg(43), Fg(29)},
		"go":      {Fg(121), Fg(109)},
		"graphql": {Fg(219), Fg(207)},
		"groovy":  {Fg(223), Fg(215)},
		"gv":      {Fg(141), Fg(99)},
		"hs":      {Fg(99), Fg(57)},
		"html":    {Fg(87), Fg(73)},
		"ino":     {Fg(43), Fg(29)},
		"java":    {Fg(136), Fg(94)},
		"jl":      {Fg(141), Fg(99)},
		"js":      {FgRGB(4, 4, 0), FgRGB(2, 2, 0)},
		"jsx":     {Fg(87), Fg(73)},
		"lock":    {FgGray(8), FgGray(5)},
		"log":     {FgGray(8), FgGray(5)},
		"lua":     {Fg(39), Fg(27)},
		"m":       {Fg(208), Fg(196)},
		"md":      {Fg(87), Fg(73)},
		"ml":      {Fg(208), Fg(94)},
		"php":     {Fg(30), Fg(22)},
		"pl":      {Fg(99), Fg(57)},
		"py":      {Fg(34), Fg(28)},
		"r":       {Fg(51), Fg(39)},
		"rb":      {FgRGB(5, 1, 0), FgRGB(3, 1, 1)},
		"rs":      {Fg(208), Fg(94)},
		"scala":   {Fg(196), Fg(124)},
		"sh":      {FgRGB(4, 0, 4), FgRGB(2, 0, 2)},
		"sol":     {Fg(39), Fg(27)},
		"sql":     {Fg(193), Fg(148)},
		"svelte":  {Fg(208), Fg(196)},
		"swift":   {Fg(223), Fg(215)},
		"vim":     {Fg(34), Fg(28)},
		"vue":     {Fg(43), Fg(29)},
		"xml":     {Fg(87), Fg(73)},

		"compiled": {FgGray(8), FgGray(5)},
		"compress": {FgRGB(5, 0, 0), FgRGB(3, 0, 0)},
		"document": {FgRGB(5, 0, 0), FgRGB(3, 0, 0)},
		"media":    {Fg(141), Fg(99)},
		"_default": {NamedFg(White), NamedFg(BrightBlack)},
	}
	// FileAliases converts alternative extensions to their canonical mapping in FileColor
	FileAliases = map[string]string{
		"s":        "asm",
		"b":        "bf",
		"c++":      "c",
		"cc":       "c",
		"cpp":      "c",
		"cs":       "c",
		"cxx":      "c",
		"d":        "c",
		"h":        "c",
		"h++":      "c",
		"hh":       "c",
		"hpp":      "c",
		"hxx":      "c",
		"pxd":      "c",
		"cljc":     "clj",
		"cljs":     "clj",
		"class":    "compiled",
		"elc":      "compiled",
		"hi":       "compiled",
		"o":        "compiled",
		"pyc":      "compiled",
		"7z":       "compress",
		"Z":        "compress",
		"bz2":      "compress",
		"deb":      "compress",
		"dmg":      "compress",
		"dpkg":     "compress",
		"gz":       "compress",
		"iso":      "compress",
		"jar":      "compress",
		"lzma":     "compress",
		"par":      "compress",
		"rar":      "compress",
		"rpm":      "compress",
		"tar":      "compress",
		"tc":       "compress",
		"tgz":      "compress",
		"txz":      "compress",
		"whl":      "compress",
		"xz":       "compress",
		"z":        "compress",
		"zip":      "compress",
		"less":     "css",
		"sass":     "css",
		"scss":     "css",
		"styl":     "css",
		"djvu":     "document",
		"doc":      "document",
		"docx":     "document",
		"dvi":      "document",
		"eml":      "document",
		"fotd":     "document",
		"odp":      "document",
		"odt":      "document",
		"pdf":      "document",
		"ppt":      "document",
		"pptx":     "document",
		"rtf":      "document",
		"xls":      "document",
		"xlsx":     "document",
		"f03":      "f",
		"f77":      "f",
		"f90":      "f",
		"f95":      "f",
		"for":      "f",
		"fpp":      "f",
		"ftn":      "f",
		"fsi":      "fs",
		"fsscript": "fs",
		"fsx":      "fs",
		"dna":      "gb",
		"gsh":      "groovy",
		"gvy":      "groovy",
		"gy":       "groovy",
		"htm":      "html",
		"xhtml":    "html",
		"bson":     "js",
		"json":     "js",
		"mjs":      "js",
		"ts":       "js",
		"cjs":      "js",
		"tsx":      "jsx",
		"cjsx":     "jsx",
		"mat":      "m",
		"markdown": "md",
		"mkd":      "md",
		"rst":      "md",
		"sml":      "ml",
		"mli":      "ml",
		"aac":      "media",
		"alac":     "media",
		"audio":    "media",
		"avi":      "media",
		"bmp":      "media",
		"cbr":      "media",
		"cbz":      "media",
		"eps":      "media",
		"flac":     "media",
		"flv":      "media",
		"gif":      "media",
		"ico":      "media",
		"image":    "media",
		"jpeg":     "media",
		"jpg":      "media",
		"m2v":      "media",
		"m4a":      "media",
		"mka":      "media",
		"mkv":      "media",
		"mov":      "media",
		"mp3":      "media",
		"mp4":      "media",
		"mpeg":     "media",
		"mpg":      "media",
		"nef":      "media",
		"ogg":      "media",
		"ogm":      "media",
		"ogv":      "media",
		"opus":     "media",
		"orf":      "media",
		"pbm":      "media",
		"pgm":      "media",
		"png":      "media",
		"pnm":      "media",
		"ppm":      "media",
		"pxm":      "media",
		"stl":      "media",
		"svg":      "media",
		"tif":      "media",
		"tiff":     "media",
		"video":    "media",
		"vob":      "media",
		"wav":      "media",
		"webm":     "media",
		"webp":     "media",
		"wma":      "media",
		"wmv":      "media",
		"xpm":      "media",
		"php3":     "php",
		"php4":     "php",
		"php5":     "php",
		"phpt":     "php",
		"phtml":    "php",
		"ipynb":    "py",
		"pickle":   "py",
		"pkl":      "py",
		"pyx":      "py",
		"bash":     "sh",
		"csh":      "sh",
		"fish":     "sh",
		"ksh":      "sh",
		"zsh":      "sh",
		"plpgsql":  "sql",
		"plsql":    "sql",
		"psql":     "sql",
		"tsql":     "sql",
		"vimrc":    "vim",
	}
	// SizeColor has the color mappings for ranges of file sizes
	SizeColor = map[string]string{
		"B": Fg(27),
		"K": Fg(33),
		"M": Fg(81),
		"G": Fg(123),
		"T": Fg(159),
	}
	// ConfigColor holds mappings for other various colors
	ConfigColor = map[string]map[string]string{
		"dir": {
			"name": Bold + BgRGB(0, 0, 2) + FgGray(23),
		},
		".dir": {
			"name": Bold + BgRGB(0, 0, 1) + FgGray(18),
		},
		"folderHeader": {
			"arrow":      NamedFg(Yellow),
			"main":       BgGray(2) + NamedFg(Yellow),
			"slash":      FgGray(5),
			"lastFolder": Bold + NamedFg(BrightYellow),
			"error":      NamedFg(BrightYellow) + NamedBg(Red),
		},
		"link": {
			"name":     NamedFg(BrightGreen),
			"nameDir":  Bold + BgRGB(0, 0, 2) + NamedFg(BrightCyan),
			"arrow":    NamedFg(BrightGreen),
			"arrowDir": NamedFg(BrightCyan),
			"broken":   NamedFg(BrightRed),
		},
		"device": {
			"name": Bold + BgGray(3) + NamedFg(BrightYellow), // /dev
		},
		"socket": {
			"name": Bold + Bg(53) + Fg(15),
		},
		"pipe": {
			"name": Bold + Bg(94) + Fg(15),
		},
		"stats": {
			"text":   BgGray(2) + FgGray(15),
			"number": Fg(24),
			"ms":     Fg(39),
		},
	}
	// PermsColor holds color mappings for users and groups
	PermsColor = map[string]map[string]string{
		"user": {
			"root":     FgRGB(5, 0, 2),
			"daemon":   FgRGB(4, 2, 1),
			"_self":    FgRGB(0, 4, 0),
			"_default": FgRGB(0, 3, 3),
		},
		"group": {
			"wheel":    FgRGB(3, 0, 0),
			"staff":    FgRGB(0, 2, 0),
			"admin":    FgRGB(2, 2, 0),
			"_default": FgRGB(2, 0, 2),
		},
		"other": {
			"_default": FgGray(15),
		},
	}
)
