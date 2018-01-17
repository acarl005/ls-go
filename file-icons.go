package main

import (
	"strings"
)

func getIconForFile(name, ext string) string {
	key := strings.ToLower(ext)
	alias, hasAlias := aliases[key]
	if hasAlias {
		key = alias
	}
	icon := icons["file"]
	betterIcon, hasBetterIcon := icons[key]
	if hasBetterIcon {
		icon = betterIcon
	}
	bestIcon, hasBestIcon := icons[name+"."+ext]
	if hasBestIcon {
		icon = bestIcon
	}
	return icon
}

func getIconForFolder(name string) string {
	icon := folders["folder"]
	betterIcon, hasBetterIcon := folders[name]
	if hasBetterIcon {
		icon = betterIcon
	}
	return icon
}

var icons = map[string]string{
	"ai":           "\ue7b4",
	"android":      "\ue70e",
	"apple":        "\uf179",
	"audio":        "\uf001",
	"avro":         "\ue60b",
	"c":            "\ue61e",
	"clj":          "\ue768",
	"coffee":       "\uf0f4",
	"conf":         "\ue615",
	"cpp":          "\ue61d",
	"css":          "\ue749",
	"d":            "\ue7af",
	"dart":         "\ue798",
	"db":           "\uf1c0",
	"diff":         "\uf440",
	"doc":          "\uf1c2",
	"ebook":        "\ue28b",
	"env":          "\uf462",
	"epub":         "\ue28a",
	"erl":          "\ue7b1",
	"file":         "\uf15b",
	"font":         "\uf031",
	"gform":        "\uf298",
	"git":          "\uf1d3",
	"go":           "\ue626",
	"gruntfile.js": "\ue74c",
	"hs":           "\ue777",
	"html":         "\uf13b",
	"image":        "\uf1c5",
	"iml":          "\ue7b5",
	"java":         "\ue204",
	"js":           "\ue74e",
	"json":         "\ue60b",
	"jsx":          "\ue7ba",
	"less":         "\ue758",
	"log":          "\uf18d",
	"lua":          "\ue620",
	"md":           "\uf48a",
	"mustache":     "\ue60f",
	"npmignore":    "\ue71e",
	"pdf":          "\uf1c1",
	"php":          "\ue73d",
	"pl":           "\ue769",
	"ppt":          "\uf1c4",
	"psd":          "\ue7b8",
	"py":           "\ue606",
	"r":            "\uf25d",
	"rb":           "\ue21e",
	"rdb":          "\ue76d",
	"rss":          "\uf09e",
	"rubydoc":      "\ue73b",
	"sass":         "\ue603",
	"scala":        "\ue737",
	"shell":        "\uf489",
	"sqlite3":      "\ue7c4",
	"styl":         "\ue600",
	"tex":          "\ue600",
	"ts":           "\ue628",
	"twig":         "\ue61c",
	"txt":          "\uf15c",
	"video":        "\uf03d",
	"vim":          "\ue62b",
	"windows":      "\uf17a",
	"xls":          "\uf1c3",
	"xml":          "\ue619",
	"yarn.lock":    "\ue718",
	"yml":          "\uf481",
	"zip":          "\uf410",
}

var aliases = map[string]string{
	"apk":              "android",
	"gradle":           "android",
	"ds_store":         "apple",
	"localized":        "apple",
	"mp3":              "audio",
	"ogg":              "audio",
	"m4a":              "audio",
	"wav":              "audio",
	"flac":             "audio",
	"editorconfig":     "conf",
	"scss":             "css",
	"docx":             "doc",
	"gdoc":             "doc",
	"mobi":             "ebook",
	"eot":              "font",
	"otf":              "font",
	"ttf":              "font",
	"woff":             "font",
	"woff2":            "font",
	"gitconfig":        "git",
	"gitignore":        "git",
	"gitignore_global": "git",
	"lhs":              "hs",
	"bmp":              "image",
	"gif":              "image",
	"ico":              "image",
	"jpeg":             "image",
	"jpg":              "image",
	"png":              "image",
	"svg":              "image",
	"webp":             "image",
	"tiff":             "image",
	"pxm":              "image",
	"jar":              "java",
	"properties":       "json",
	"tsx":              "jsx",
	"license":          "md",
	"markdown":         "md",
	"mkd":              "md",
	"rdoc":             "md",
	"readme":           "md",
	"gslides":          "ppt",
	"pptx":             "ppt",
	"pyc":              "py",
	"rdata":            "r",
	"rds":              "r",
	"gemfile":          "rb",
	"gemspec":          "rb",
	"guardfile":        "rb",
	"lock":             "rb",
	"procfile":         "rb",
	"rakefile":         "rb",
	"rspec":            "rb",
	"rspec_parallel":   "rb",
	"rspec_status":     "rb",
	"ru":               "rb",
	"erb":              "rubydoc",
	"slim":             "rubydoc",
	"bash":             "shell",
	"bash_history":     "shell",
	"bash_profile":     "shell",
	"bashrc":           "shell",
	"fish":             "shell",
	"sh":               "shell",
	"zsh":              "shell",
	"zsh-theme":        "shell",
	"zshrc":            "shell",
	"stylus":           "styl",
	"cls":              "tex",
	"avi":              "video",
	"mkv":              "video",
	"mp4":              "video",
	"ogv":              "video",
	"webm":             "video",
	"mov":              "video",
	"flv":              "video",
	"bat":              "windows",
	"exe":              "windows",
	"ini":              "windows",
	"csv":              "xls",
	"gsheet":           "xls",
	"xlsx":             "xls",
	"xul":              "xml",
	"yaml":             "yml",
	"gz":               "zip",
	"rar":              "zip",
	"tar":              "zip",
}

var folders = map[string]string{
	".atom":        "\ue764",
	".git":         "\uf1d3",
	".github":      "\uf408",
	".rvm":         "\ue21e",
	".Trash":       "\uf1f8",
	".vscode":      "\ue70c",
	"config":       "\ue5fc",
	"folder":       "\uf115",
	"hidden":       "\uf023",
	"lib":          "\uf121",
	"node_modules": "\ue718",
}
