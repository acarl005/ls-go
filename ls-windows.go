// +build windows

package main

import (
	"os"
	"path/filepath"
	"syscall"
	"unsafe"
)

var (
	libadvapi32                    = syscall.NewLazyDLL("advapi32.dll")
	procGetFileSecurity            = libadvapi32.NewProc("GetFileSecurityW")
	procGetSecurityDescriptorOwner = libadvapi32.NewProc("GetSecurityDescriptorOwner")
)

func getOwnerAndGroup(parentDir string, fileInfo *os.FileInfo) (string, string) {
	path := filepath.Join(parentDir, (*fileInfo).Name())

	var needed uint32
	procGetFileSecurity.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(path))),
		0x00000001, /* OWNER_SECURITY_INFORMATION */
		0,
		0,
		uintptr(unsafe.Pointer(&needed)))
	buf := make([]byte, needed)
	r1, _, err := procGetFileSecurity.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(path))),
		0x00000001, /* OWNER_SECURITY_INFORMATION */
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(needed),
		uintptr(unsafe.Pointer(&needed)))
	if r1 == 0 && err != nil {
		return "", ""
	}
	var ownerDefaulted uint32
	var sid *syscall.SID
	r1, _, err = procGetSecurityDescriptorOwner.Call(
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(unsafe.Pointer(&sid)),
		uintptr(unsafe.Pointer(&ownerDefaulted)))
	if r1 == 0 && err != nil {
		return "", ""
	}
	uid, gid, _, err := sid.LookupAccount("")
	if r1 == 0 && err != nil {
		return "", ""
	}
	return uid, gid
}

func deviceNumbers(absPath string) string {
	return ""
}
