package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davecgh/go-spew/spew"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

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
	sortSize *bool
	sortTime *bool
	sortKind *bool
	stats    *bool
	icons    *bool
	recurse  *bool
	find     *string
}

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
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()
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

	if len(files) > 0 {
		pwd, err := os.Getwd()
		check(err)
		listFiles(pwd, &files)
	}

	for _, dir := range dirs {
		listDir(dir)
	}
}

func listDir(path string) {
	items, err := ioutil.ReadDir(path)
	check(err)

	if len(items) > 0 {
		listFiles(path, &items)
	}
}

func listFiles(parentDir string, items *[]os.FileInfo) {
	fmt.Println(parentDir)
	for _, fileInfo := range *items {
		//fmt.Println(fileInfo)
		//fmt.Printf("%+v\n", fileInfo)
		spew.Dump(fileInfo)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
