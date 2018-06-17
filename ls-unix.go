// +build !windows

package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"syscall"
)

func isUnknownUser(err error) bool {
	if strings.Contains(err.Error(), "unknown userid") {
		return true
	}
	return false
}

func isUnknownGroup(err error) bool {
	if strings.Contains(err.Error(), "unknown group") {
		return true
	}
	return false
}

func getOwnerAndGroup(fileInfo *os.FileInfo) (string, string) {
	stat_t := (*fileInfo).Sys().(*syscall.Stat_t)
	uid := fmt.Sprint(stat_t.Uid)
	gid := fmt.Sprint(stat_t.Gid)
	owner, err := user.LookupId(uid)
	if err != nil {
		if isUnknownUser(err) {
			owner = new(user.User)
			owner.Username = uid
		} else {
			check(err)
		}
	}
	group, err := user.LookupGroupId(gid)
	if err != nil {
		if isUnknownGroup(err) {
			group = new(user.Group)
			group.Name = gid
		} else {
			check(err)
		}
	}
	return owner.Username, group.Name
}
