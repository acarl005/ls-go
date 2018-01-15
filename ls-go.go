package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/willf/pad"
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

type DisplayItem struct {
	display  string
	info     os.FileInfo
	basename string
	ext      string
	link     string
}

// configure the arguments
var (
	True       = true
	sizeUnits  = []string{" B", "kB", "MB", "GB", "TB"}
	dateFormat = "02.Jan'06"
	timeFormat = "15:04:05"
	args       = arguments{
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

func argsPostParse() {
	if *args.long {
		args.bytes = &True
		args.mdate = &True
		args.owner = &True
		args.perms = &True
	}
}

func main() {
	// auto-generate help text for the command with -h
	kingpin.CommandLine.HelpFlag.Short('h')

	// parse the arguments and populate the struct
	kingpin.Parse()
	argsPostParse()

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

	if *args.recurse {
		for _, item := range items {
			if item.IsDir() {
				listDir(path + item.Name() + "/")
			}
		}
	}
}

func listFiles(parentDir string, items *[]os.FileInfo) {
	//prettyPath := prettifyPath(parentDir)
	absPath, err := filepath.Abs(parentDir)
	check(err)
	files := []DisplayItem{}
	dirs := []DisplayItem{}

	longestOwnerName := 0
	longestGroupName := 0
	if *args.owner {
		for _, fileInfo := range *items {
			owner, group := getOwnerAndGroup(&fileInfo)
			longestOwnerName = max(longestOwnerName, len(owner))
			longestGroupName = max(longestGroupName, len(group))
		}
	}

	for _, fileInfo := range *items {
		displayItem := DisplayItem{info: fileInfo}
		filePath := path.Join(absPath, fileInfo.Name())

		var basename, ext string
		// if this is a dotfile (hidden file)
		if fileInfo.Name()[0] == '.' {
			// we can skip everything with this file if we aren't using the `all` option
			if !*args.all {
				continue
			}
			ext = fileInfo.Name()[1:]
			basename = ""
		} else {
			ext = filepath.Ext(fileInfo.Name())
			basename = fileInfo.Name()[:len(fileInfo.Name())-len(ext)]
			if ext != "" {
				ext = ext[1:]
			}
		}
		displayItem.ext = ext
		displayItem.basename = basename

		if fileInfo.IsDir() {
			dirs = append(dirs, displayItem)
		} else {
			files = append(files, displayItem)
		}

		owner, group := getOwnerAndGroup(&fileInfo)
		ownerColor, groupColor := getOwnerAndGroupColors(owner, group)

		if *args.perms {
			displayItem.display += permString(fileInfo, ownerColor, groupColor)
		}

		if *args.owner {
			ownerInfo := []string{Reset, ownerColor + owner, groupColor + group, Reset}
			displayItem.display += strings.Join(ownerInfo, " ")
		}

		if *args.bytes {
			displayItem.display += sizeString(fileInfo.Size())
		}

		if *args.mdate {
			displayItem.display += timeString(fileInfo.ModTime())
		}

		if fileInfo.IsDir() {
			displayItem.display += dirString(&displayItem)
		} else {
			displayItem.display += fileString(&displayItem)
		}

		if *args.links && fileInfo.Mode()&os.ModeSymlink != 0 {
			link, err := os.Readlink(filePath)
			fmt.Println(link)
			check(err)
			displayItem.link = link
		}
		fmt.Println(displayItem.display)
		//fmt.Println(fileInfo.Name())
	}
}

/*
linkString = (file) ->
    reset + colors['_link']['arrow'] + " ► " + colors['_link'][(file in stats.brokenLinks) and 'broken' or 'path'] + fs.readlinkSync(file)

nameString = (name, ext, stat) ->
    key = if mediaTypes.has(ext) then "_media" else ext
    name = " " + colors[colors[key]? and key or '_default'][0] + name + reset
    if ((stat.mode >> 6) & 0b001) && args.icons
        ' ' + BW(1) + fg(0, 5, 0) + '>_' + reset + name
    else
        name

extString = (ext) ->
    key = if mediaTypes.has(ext) then "_media" else ext
    colors[colors[key]? and key or '_default'][1] + '.' + ext + reset

dirString = (name, ext) ->
    c = name and '_dir' or '_.dir'
    name = '📂 ' + name if args.icons
    colors[c][0] + (name and (" " + name) or "") + (if ext then colors[c][1] + '.' + colors[c][2] + ext else "") + " "
*/

func fileString(item *DisplayItem) string {
	ext := strings.ToLower(item.ext)
	var colors [2]string
	if MediaTypes.Has(ext) {
		colors = FileColor["_media"]
	} else if CompressTypes.Has(ext) {
		colors = FileColor["_compress"]
	} else if val, hasExt := FileColor[ext]; hasExt {
		colors = val
	} else {
		colors = FileColor["_default"]
	}
	ext = item.ext
	if ext != "" {
		ext = "." + ext
	}
	execIcon := ""
	executable := item.info.Mode() & 0111
	if executable != 0 && *args.icons {
		execIcon = BgGray(1) + FgRGB(0, 5, 0) + ">_" + Reset + " "
	}
	displayStrings := []string{execIcon, colors[0], item.basename, colors[1], ext, Reset}
	return strings.Join(displayStrings, "")
}

func dirString(item *DisplayItem) string {
	colors := ConfigColor["dir"]
	if item.basename == "" {
		colors = ConfigColor[".dir"]
	}
	displayStrings := []string{colors["name"], " "}
	if *args.icons {
		displayStrings = append(displayStrings, "📂 ")
	}
	ext := item.ext
	if ext != "" {
		ext = "." + ext
	}
	displayStrings = append(displayStrings, item.basename, colors["ext"], ext, " ", Reset)
	return strings.Join(displayStrings, "")
}

func permString(info os.FileInfo, ownerColor string, groupColor string) string {
	defaultColor := PermsColor["other"]["_default"]
	modeString := info.Mode().String()
	coloredStrings := []string{defaultColor, modeString[0:1]}
	coloredStrings = append(coloredStrings, ownerColor+modeString[1:4])
	coloredStrings = append(coloredStrings, groupColor+modeString[4:7])
	coloredStrings = append(coloredStrings, defaultColor+modeString[7:], Reset)
	return strings.Join(coloredStrings, " ")
}

func sizeString(size int64) string {
	for i, unit := range sizeUnits {
		base := int64(math.Pow(10, float64(i*3)))
		if size < base*1000 {
			return SizeColor[unit] + pad.Left(fmt.Sprintf("%d", size/base), 4, " ") + unit + Reset
		}
	}
	return strconv.Itoa(int(size))
}

func timeString(modtime time.Time) string {
	dateStr := modtime.Format(dateFormat)
	timeStr := modtime.Format(timeFormat)
	hour, err := strconv.Atoi(timeStr[0:2])
	check(err)
	// generate a color based on the hour of the day. darkest around midnight and whitest around noon
	timeColor := 14 - int(8*math.Cos(math.Pi*float64(hour)/12))
	colored := []string{FgGray(22), dateStr, FgGray(timeColor) + timeStr, Reset}
	return strings.Join(colored, " ")
}

// when we list out any subdirectories, print those paths conspicuously above the contents
// this helps with visual separation
func printFolderHeader(path string) {
	colors := ConfigColor["folderHeader"]
	headerString := colors["arrow"] + "►" + colors["main"] + " "
	prettyPath := prettifyPath(path)

	if prettyPath == "/" {
		headerString += "/"
	} else {
		folders := strings.Split(prettyPath, "/")
		coloredFolders := make([]string, 0, len(folders))
		for i, folder := range folders {
			if i == len(folders)-1 { // different color for the last folder in the path
				coloredFolders = append(coloredFolders, colors["lastFolder"]+folder)
			} else {
				coloredFolders = append(coloredFolders, colors["main"]+folder)
			}
		}
		headerString += strings.Join(coloredFolders, colors["slash"]+"/")
	}

	fmt.Println(headerString + " " + Reset)
}

func prettifyPath(path string) string {
	prettyPath, err := filepath.Abs(path)
	check(err)
	pwd := os.Getenv("PWD")
	home := os.Getenv("HOME")

	if strings.HasPrefix(prettyPath, pwd) {
		prettyPath = "." + prettyPath[len(pwd):]
	} else if strings.HasPrefix(prettyPath, home) {
		prettyPath = "~" + prettyPath[len(home):]
	}
	return prettyPath
}

func getOwnerAndGroup(fileInfo *os.FileInfo) (string, string) {
	stat_t := (*fileInfo).Sys().(*syscall.Stat_t)
	uid := fmt.Sprint(stat_t.Uid)
	gid := fmt.Sprint(stat_t.Gid)
	owner, err := user.LookupId(uid)
	check(err)
	group, err := user.LookupGroupId(gid)
	check(err)
	return owner.Username, group.Name
}

func getOwnerAndGroupColors(owner string, group string) (string, string) {
	if owner == os.Getenv("USER") {
		owner = "_self"
	}
	ownerColor := PermsColor["user"][owner]
	if ownerColor == "" {
		ownerColor = PermsColor["user"]["_default"]
	}
	groupColor := PermsColor["group"][group]
	if groupColor == "" {
		groupColor = PermsColor["group"]["_default"]
	}
	return ownerColor, groupColor
}

// Go doesn't provide a `Max` function for ints like it does for floats (wtf?)
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func check(err error) {
	// TODO: handle no permission error, broken links, invalid path args
	if err != nil {
		panic(err)
	}
}
