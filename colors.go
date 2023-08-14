package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// see the 256-color codes
// http://i.stack.imgur.com/UQVe5.png

type Color int8

const (
	Black         Color = 0
	Red                 = 1
	Green               = 2
	Yellow              = 3
	Blue                = 4
	Magenta             = 5
	Cyan                = 6
	White               = 7
	BrightBlack         = 60
	BrightRed           = 61
	BrightGreen         = 62
	BrightYellow        = 63
	BrightBlue          = 64
	BrightMagenta       = 65
	BrightCyan          = 66
	BrightWhite         = 67
)

var lightThemeMap = map[int]int{
	67: 0,
	7:  60,
	60: 7,
	0:  67,
}

func applyTheme(index int) int {
	if *args.light {
		newIndex, hasNewIndex := lightThemeMap[index]
		if hasNewIndex {
			return newIndex
		} else if index >= 60 {
			return index - 60
		}
		return index + 60
	}
	return index
}

func NamedFg(color Color) string {
	index := applyTheme(int(color))
	colored := []string{"\x1b[", strconv.Itoa(index + 30), "m"}
	return strings.Join(colored, "")
}

func NamedBg(color Color) string {
	index := applyTheme(int(color))
	colored := []string{"\x1b[", strconv.Itoa(index + 40), "m"}
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

func avg(a float64, b float64) int {
	m := (a + b) / 2
	return int(math.Ceil(m))
}

// Rgb2code converts RGB values (up to 5) to an 8-bit color code
func Rgb2code(r int, g int, b int, theme bool) int {
	if theme && *args.light {
		r_, g_, b_ := float64(r), float64(g), float64(b)
		r = 5 - avg(g_, b_)
		g = 5 - avg(r_, b_)
		b = 5 - avg(r_, g_)
	}
	code := 36*r + 6*g + b + 16
	if code < 16 || 231 < code {
		panic(fmt.Errorf("Invalid RGB values (%d, %d, %d)", r, g, b))
	}
	return code
}

// Gray2code converts a scalar of "grayness" to an 8-bit color code
func Gray2code(lightness int) int {
	if *args.light {
		lightness = 23 - lightness
	}
	code := lightness + 232
	if code < 232 || 255 < code {
		panic(fmt.Errorf("Invalid lightness value (%d) for gray", lightness))
	}
	return code
}

// FgRGB converts RGB values (up to 5) to an ANSI-escaped foreground 8-bit color code
func FgRGB(r int, g int, b int) string {
	return Fg(Rgb2code(r, g, b, false))
}

// Theme-sensitive version
func FgRGBT(r int, g int, b int) string {
	return Fg(Rgb2code(r, g, b, true))
}

// BgRGB converts RGB values (up to 5) to an ANSI-escaped background 8-bit color code
func BgRGB(r int, g int, b int) string {
	return Bg(Rgb2code(r, g, b, false))
}

// Theme-sensitive version
func BgRGBT(r int, g int, b int) string {
	return Bg(Rgb2code(r, g, b, true))
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
	Bold  = "\x1b[1m"
)

var (
	// FileColor is a mapping of filetypes to colors
	FileColor map[string]FileColorConfig

	// SizeColor has the color mappings for ranges of file sizes
	SizeColor map[string]string

	// ConfigColor holds mappings for other various colors
	ConfigColor map[string]map[string]string

	// PermsColor holds color mappings for users and groups
	PermsColor map[string]map[string]string
)

var (
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
		"jade":     "js",
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
		"glp":      "media",
		"gltf":     "media",
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
		"sixel":    "media",
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
)

type FileColorConfig struct {
	light       string
	dark        string
	themeSwitch bool
}

func generateColors() {
	FileColor = map[string]FileColorConfig{
		"as":      {FgRGB(5, 0, 0), FgRGB(3, 0, 0), true},
		"asm":     {FgRGBT(5, 4, 3), FgRGBT(5, 3, 1), false},
		"apk":     {FgRGB(1, 5, 0), FgRGB(1, 3, 0), true},
		"bf":      {FgRGBT(5, 4, 3), FgRGBT(5, 3, 1), false},
		"bzl":     {FgRGB(1, 5, 0), FgRGB(1, 3, 0), true},
		"c":       {FgRGB(0, 3, 5), FgRGB(0, 1, 5), true},
		"clj":     {FgRGB(5, 1, 2), FgRGB(4, 0, 2), true},
		"coffee":  {FgRGB(3, 2, 0), FgRGB(2, 1, 0), true},
		"cr":      {FgRGB(1, 5, 0), FgRGB(1, 3, 0), true},
		"cson":    {FgRGB(3, 2, 0), FgRGB(2, 1, 0), true},
		"css":     {FgRGB(5, 3, 5), FgRGB(5, 1, 5), true},
		"dart":    {FgRGB(0, 4, 3), FgRGB(0, 2, 3), true},
		"diff":    {FgRGB(0, 5, 0), FgRGB(5, 0, 0), true},
		"elm":     {FgRGB(0, 5, 5), FgRGB(0, 3, 5), true},
		"erl":     {FgRGB(4, 0, 1), FgRGB(2, 0, 1), true},
		"ex":      {FgRGB(2, 1, 5), FgRGB(1, 0, 5), true},
		"f":       {FgRGB(5, 2, 0), FgRGB(2, 1, 0), true},
		"fs":      {FgRGB(0, 4, 5), FgRGB(0, 2, 4), true},
		"gb":      {FgRGB(0, 4, 3), FgRGB(0, 2, 1), true},
		"go":      {FgRGB(2, 5, 3), FgRGB(2, 3, 3), true},
		"graphql": {FgRGB(5, 3, 5), FgRGB(5, 1, 5), true},
		"groovy":  {FgRGBT(5, 4, 3), FgRGBT(5, 3, 1), false},
		"gv":      {FgRGB(3, 2, 5), FgRGB(2, 1, 5), true},
		"hs":      {FgRGB(2, 1, 5), FgRGB(1, 0, 5), true},
		"html":    {FgRGB(1, 5, 5), FgRGB(1, 3, 3), true},
		"hx":      {FgRGB(5, 2, 0), FgRGB(2, 1, 0), true},
		"ino":     {FgRGB(0, 4, 3), FgRGB(0, 2, 1), true},
		"java":    {FgRGB(3, 2, 0), FgRGB(2, 1, 0), true},
		"jl":      {FgRGB(3, 2, 5), FgRGB(2, 1, 5), true},
		"js":      {FgRGB(4, 4, 0), FgRGB(2, 2, 0), true},
		"jsx":     {FgRGB(1, 5, 5), FgRGB(1, 3, 3), true},
		"kt":      {FgRGB(2, 1, 5), FgRGB(1, 0, 5), true},
		"lock":    {FgGray(11), FgGray(7), false},
		"log":     {FgGray(11), FgGray(7), false},
		"lua":     {FgRGB(0, 3, 5), FgRGB(0, 1, 5), true},
		"m":       {FgRGBT(5, 4, 3), FgRGBT(5, 3, 1), false},
		"md":      {FgRGB(1, 5, 5), FgRGB(1, 3, 3), true},
		"ml":      {FgRGB(5, 2, 0), FgRGB(2, 1, 0), true},
		"nim":     {FgRGB(4, 4, 0), FgRGB(2, 2, 0), true},
		"php":     {FgRGB(0, 3, 4), FgRGB(0, 2, 2), true},
		"pl":      {FgRGB(2, 1, 5), FgRGB(1, 0, 5), true},
		"py":      {FgRGB(0, 4, 0), FgRGB(0, 2, 0), true},
		"r":       {FgRGB(0, 5, 5), FgRGB(0, 3, 5), true},
		"rb":      {FgRGB(5, 1, 0), FgRGB(3, 1, 0), true},
		"rs":      {FgRGB(4, 3, 0), FgRGB(2, 2, 0), true},
		"scala":   {FgRGB(5, 0, 0), FgRGB(3, 0, 0), true},
		"sh":      {FgRGB(4, 0, 4), FgRGB(2, 0, 2), true},
		"sol":     {FgRGB(0, 3, 5), FgRGB(0, 1, 5), true},
		"sql":     {FgRGB(4, 5, 3), FgRGB(3, 4, 0), true},
		"svelte":  {FgRGB(5, 2, 0), FgRGB(5, 0, 0), true},
		"swift":   {FgRGBT(5, 4, 3), FgRGBT(5, 3, 1), false},
		"vim":     {FgRGB(0, 4, 0), FgRGB(0, 2, 0), true},
		"vue":     {FgRGB(0, 4, 3), FgRGB(0, 2, 1), true},
		"xml":     {FgRGB(1, 5, 5), FgRGB(1, 3, 3), true},
		"zig":     {FgRGB(4, 4, 0), FgRGB(2, 2, 0), true},

		"compiled": {FgGray(11), FgGray(7), false},
		"compress": {FgRGB(5, 0, 0), FgRGB(3, 0, 0), true},
		"document": {FgRGB(5, 0, 0), FgRGB(3, 0, 0), true},
		"media":    {FgRGB(3, 2, 5), FgRGB(2, 1, 5), true},
		"_default": {FgGray(20), FgGray(11), false},
	}

	SizeColor = map[string]string{
		"B": FgRGBT(0, 1, 5),
		"K": FgRGBT(0, 2, 5),
		"M": FgRGBT(1, 4, 5),
		"G": FgRGBT(2, 5, 5),
		"T": FgRGBT(3, 5, 5),
	}

	ConfigColor = map[string]map[string]string{
		"dir": {
			"name": Bold + BgRGBT(0, 0, 2) + FgGray(23),
		},
		".dir": {
			"name": Bold + BgRGBT(0, 0, 1) + FgGray(18),
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
			"nameDir":  Bold + BgRGBT(0, 0, 2) + NamedFg(BrightCyan),
			"arrow":    NamedFg(BrightGreen),
			"arrowDir": NamedFg(BrightCyan),
			"broken":   NamedFg(BrightRed),
		},
		"device": {
			"name": Bold + BgGray(3) + NamedFg(BrightYellow), // /dev
		},
		"socket": {
			"name": Bold + BgRGBT(1, 0, 1) + FgGray(23),
		},
		"pipe": {
			"name": Bold + BgRGBT(2, 1, 0) + FgGray(23),
		},
		"stats": {
			"text":   BgGray(2) + FgGray(15),
			"number": FgRGBT(0, 2, 3),
			"ms":     FgRGBT(0, 3, 5),
		},
	}

	PermsColor = map[string]map[string]string{
		"user": {
			"root":     FgRGB(5, 0, 2),
			"daemon":   FgRGB(4, 2, 1),
			"_self":    FgRGB(0, 5, 0),
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
}
