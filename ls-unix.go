// +build !windows

package main

import (
	"fmt"
	"os"
	"os/user"
	"syscall"
)

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
