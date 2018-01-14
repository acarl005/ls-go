package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// declare the struct that holds all the arguments
type arguments struct {
	paths    *[]string
	all      *bool
	bytes    *bool
	mdate    *bool
	owner    *bool
	perms    *bool
	long     *bool
	dirs     *bool
	files    *bool
	links    *bool
	sortSize *bool
	sortTime *bool
	sortKind *bool
	stats    *bool
	icons    *bool
	recurse  *bool
	find     *string
}

// configure the arguments
var (
	args = arguments{
		kingpin.Arg("paths", "the files(s) and/or folder(s) to display").Default(".").Strings(),
		kingpin.Flag("all", "show hidden files").Short('a').Bool(),
		kingpin.Flag("bytes", "include size").Short('b').Bool(),
		kingpin.Flag("mdate", "include modification date").Short('m').Bool(),
		kingpin.Flag("owner", "include owner and group").Short('o').Bool(),
		kingpin.Flag("perms", "include permissions for owner, group, and other").Short('p').Bool(),
		kingpin.Flag("long", "include size, date, owner, and permissions").Short('l').Bool(),
		kingpin.Flag("dirs", "only show directories").Short('d').Bool(),
		kingpin.Flag("files", "only show files").Short('f').Bool(),
		kingpin.Flag("links", "show paths for symlinks").Short('L').Bool(),
		kingpin.Flag("size", "sort items by size").Short('s').Bool(),
		kingpin.Flag("time", "sort items by time").Short('t').Bool(),
		kingpin.Flag("kind", "sort items by extension").Short('k').Bool(),
		kingpin.Flag("stats", "show statistics").Short('S').Bool(),
		kingpin.Flag("icons", "show folder icon before dirs").Short('i').Bool(),
		kingpin.Flag("recurse", "traverse all dirs recursively").Short('r').Bool(),
		kingpin.Flag("find", "filter items with a regexp").Short('F').String(),
	}
)

func main() {
	// auto-generate help text for the command with -h
	kingpin.CommandLine.HelpFlag.Short('h')

	// parse the arguments and populate the struct
	kingpin.Parse()

	// separate the directories from the regular files
	dirs := []string{}
	files := []os.FileInfo{}
	for _, path := range *args.paths {
		fileStat, err := os.Stat(path)
		check(err)
		if fileStat.IsDir() {
			dirs = append(dirs, path)
		} else {
			files = append(files, fileStat)
		}
	}

	// list files first
	if len(files) > 0 {
		pwd := os.Getenv("PWD")
		listFiles(pwd, &files)
	}

	// then list the contents of each directory
	for _, dir := range dirs {
		listDir(dir)
	}
}

func listDir(path string) {
	items, err := ioutil.ReadDir(path)
	check(err)

	// filter by the regexp if one was passed
	if len(*args.find) > 0 {
		filteredItems := []os.FileInfo{}
		for _, fileInfo := range items {
			re, err := regexp.Compile(*args.find)
			check(err)
			if re.MatchString(fileInfo.Name()) {
				filteredItems = append(filteredItems, fileInfo)
			}
		}
		items = filteredItems
	}

	if len(*args.find) > 0 && len(items) == 0 {
	} else if len(*args.paths) == 1 && (*args.paths)[0] == "." && !*args.recurse {
	} else {
		printFolderHeader(path)
	}

	if len(items) > 0 {
		listFiles(path, &items)
	}
}

func listFiles(parentDir string, items *[]os.FileInfo) {
	fmt.Println(parentDir)
	for _, fileInfo := range *items {
		fmt.Println(fileInfo.Name())
	}
}

/*
   s = colors['_arrow'] + "►" + colors['_header'][0] + " "
   ps = path.resolve(ps) if ps[0] != '~'
   if _s.startsWith(ps, process.env.PWD)
       ps = "./" + ps.substr(process.env.PWD.length)
   else if _s.startsWith(p, process.env.HOME)
       ps = "~" + p.substr(process.env.HOME.length)

   if ps == '/'
       s += '/'
   else
       sp = ps.split('/')
       s += colors['_header'][0] + sp.shift()
       while sp.length
           pn = sp.shift()
           if pn
               s += colors['_header'][1] + '/'
               s += colors['_header'][sp.length == 0 and 2 or 0] + pn
   log reset
   log s + " " + reset
   log reset
*/
func printFolderHeader(path string) {
	headerString := ConfigColor["folderHeader"]["arrow"] + "►" + ConfigColor["header"]["main"] + " "
	prettyPath, err := filepath.Abs(path)
	check(err)
	pwd := os.Getenv("PWD")
	home := os.Getenv("HOME")

	if strings.HasPrefix(prettyPath, pwd) {
		prettyPath = "." + prettyPath[len(pwd):]
	} else if strings.HasPrefix(prettyPath, home) {
		prettyPath = "~" + prettyPath[len(home):]
	}

	if prettyPath == "/" {
		headerString += "/"
	} else {
		folders := strings.Split(prettyPath, "/")
		coloredFolders := make([]string, 0, len(folders))
		for i, folder := range folders {
			if i == len(folders)-1 {
				coloredFolders = append(coloredFolders, ConfigColor["folderHeader"]["lastFolder"]+folder)
			} else {
				coloredFolders = append(coloredFolders, ConfigColor["folderHeader"]["main"]+folder)
			}
		}
		headerString += strings.Join(coloredFolders, ConfigColor["folderHeader"]["slash"]+"/")
	}

	fmt.Println(headerString + " " + Reset)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
