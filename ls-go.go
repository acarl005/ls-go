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
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/willf/pad"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type DisplayItem struct {
	display  string
	info     os.FileInfo
	basename string
	ext      string
	link     *LinkInfo
}

type LinkInfo struct {
	path string
	info os.FileInfo
}

// configure the arguments
var (
	True             = true
	sizeUnits        = []string{" B", "kB", "MB", "GB", "TB"}
	dateFormat       = "02.Jan'06"
	timeFormat       = "15:04:05"
	start      int64 = 0
)

func main() {
	start = time.Now().UnixNano()
	// auto-generate help text for the command with -h
	kingpin.CommandLine.HelpFlag.Short('h')

	// parse the arguments and populate the struct
	kingpin.Parse()
	argsPostParse()

	// separate the directories from the regular files
	dirs := []string{}
	files := []os.FileInfo{}
	for _, pathStr := range *args.paths {
		fileStat, err := os.Stat(pathStr)
		check(err)
		if fileStat.IsDir() {
			dirs = append(dirs, pathStr)
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
	for i, dir := range dirs {
		// print a blank line between directories, but not before the first one
		if i > 0 {
			fmt.Println("")
		}
		listDir(dir)
	}
}

func listDir(pathStr string) {
	items, err := ioutil.ReadDir(pathStr)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") || strings.Contains(err.Error(), "permission denied") {
			fmt.Println(ConfigColor["folderHeader"]["error"] + "► " + prettifyPath(pathStr) + Reset)
			fmt.Println(err.Error())
			return
		} else {
			check(err)
		}
	}

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
		printFolderHeader(pathStr)
	}

	if len(items) > 0 {
		listFiles(pathStr, &items)
	}

	if *args.recurse {
		for _, item := range items {
			if item.IsDir() {
				fmt.Println("") // put a blank line between directories
				listDir(path.Join(pathStr, item.Name()))
			}
		}
	}
}

func listFiles(parentDir string, items *[]os.FileInfo) {
	absPath, err := filepath.Abs(parentDir)
	check(err)
	files := []*DisplayItem{}
	dirs := []*DisplayItem{}

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
			dirs = append(dirs, &displayItem)
		} else {
			files = append(files, &displayItem)
		}

		if fileInfo.Mode()&os.ModeSymlink != 0 {
			getLinkInfo(&displayItem, absPath)
		}

		owner, group := getOwnerAndGroup(&fileInfo)
		ownerColor, groupColor := getOwnerAndGroupColors(owner, group)

		if *args.perms {
			displayItem.display += permString(fileInfo, ownerColor, groupColor)
		}

		if *args.owner {
			ownerInfo := []string{Reset + ownerColor + owner, groupColor + group, Reset}
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
			displayItem.display += linkString(&displayItem, absPath)
		}
	}

	if !*args.files {
		if *args.sortTime {
			sort.Sort(ByTime(dirs))
		}

		for _, dir := range dirs {
			fmt.Println(dir.display)
		}
	}

	if !*args.dirs {
		if *args.sortSize {
			sort.Sort(BySize(files))
		}

		if *args.sortTime {
			sort.Sort(ByTime(files))
		}

		if *args.sortKind {
			sort.Sort(ByKind(files))
		}

		for _, file := range files {
			fmt.Println(file.display)
		}
	}

	if *args.stats {
		printStats(len(files), len(dirs))
	}
}

func getLinkInfo(item *DisplayItem, absPath string) {
	fullPath := path.Join(absPath, item.info.Name())
	linkPath, err1 := os.Readlink(fullPath)
	check(err1)
	linkInfo, err2 := os.Stat(linkPath)
	if *args.linkRel {
		linkRel, _ := filepath.Rel(absPath, linkPath)
		if linkRel != "" && len(linkRel) <= len(linkPath) {
			linkPath = linkRel
		}
	}
	link := LinkInfo{
		path: linkPath,
	}
	item.link = &link
	if linkInfo != nil {
		link.info = linkInfo
	} else if !strings.Contains(err2.Error(), "no such file or directory") {
		check(err2)
	}
}

func linkString(item *DisplayItem, absPath string) string {
	colors := ConfigColor["link"]
	displayStrings := []string{colors["arrow"], "►"}
	if item.link.info == nil {
		displayStrings = append(displayStrings, colors["broken"]+item.link.path+Reset)
	} else {
		displayStrings = append(displayStrings, colors["path"]+item.link.path+Reset)
	}
	return strings.Join(displayStrings, " ")
}

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
	executable := isExecutableScript(item)
	if executable && *args.icons {
		execIcon = BgGray(1) + FgRGB(0, 5, 0) + ">_" + Reset + " "
	}
	displayStrings := []string{execIcon, colors[0], item.basename, colors[1], ext, Reset}
	return strings.Join(displayStrings, "")
}

func isExecutableScript(item *DisplayItem) bool {
	executable := false
	if item.link != nil {
		if item.link.info != nil {
			executable = (item.link.info.Mode()&0111) != 0 && !item.link.info.IsDir()
		}
	} else if item.info.Mode()&0111 != 0 {
		executable = true
	}
	return executable
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
	coloredStrings = append(coloredStrings, defaultColor+modeString[7:], Reset, Reset)
	return strings.Join(coloredStrings, " ")
}

func sizeString(size int64) string {
	for i, unit := range sizeUnits {
		base := int64(math.Pow(10, float64(i*3)))
		if size < base*1000 {
			return SizeColor[unit] + pad.Left(fmt.Sprintf("%d", size/base), 4, " ") + unit + " " + Reset
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
	colored := []string{FgGray(22) + dateStr, FgGray(timeColor) + timeStr, Reset}
	return strings.Join(colored, " ")
}

// when we list out any subdirectories, print those paths conspicuously above the contents
// this helps with visual separation
func printFolderHeader(pathStr string) {
	colors := ConfigColor["folderHeader"]
	headerString := colors["arrow"] + "►" + colors["main"] + " "
	prettyPath := prettifyPath(pathStr)

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

func prettifyPath(pathStr string) string {
	prettyPath, err := filepath.Abs(pathStr)
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

func printStats(numFiles, numDirs int) {
	colors := ConfigColor["stats"]
	end := time.Now().UnixNano()
	microSeconds := (end - start) / int64(time.Microsecond)
	milliSeconds := float64(microSeconds) / 1000
	statStrings := []string{
		colors["text"],
		colors["number"] + strconv.Itoa(numDirs),
		colors["text"] + "dirs",
		colors["number"] + strconv.Itoa(numFiles),
		colors["text"] + "files",
		colors["ms"] + fmt.Sprintf("%.2f", milliSeconds),
		colors["text"] + "ms",
		Reset,
	}
	fmt.Println(strings.Join(statStrings, " "))
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
