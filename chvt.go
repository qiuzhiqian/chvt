package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

const (
	MIN_NR_CONSOLES      uint = 1      /* must be at least 1 */
	MAX_NR_CONSOLES      uint = 63     /* serial lines start at 64 */
	MAX_NR_USER_CONSOLES uint = 63     /* must be root to allocate above this */
	VT_OPENQRY           uint = 0x5600 /* find available vt */
	VT_GETMODE           uint = 0x5601 /* get mode of active vt */
	VT_SETMODE           uint = 0x5602 /* set mode of active vt */
	VT_AUTO              uint = 0x00   /* auto vt switching */
	VT_PROCESS           uint = 0x01   /* process controls switching */
	VT_ACKACQ            uint = 0x02   /* acknowledge switch */
	VT_GETSTATE          uint = 0x5603 /* get global vt state info */
	VT_SENDSIG           uint = 0x5604 /* signal to send to bitmask of vts */
	VT_RELDISP           uint = 0x5605 /* release display */
	VT_ACTIVATE          uint = 0x5606 /* make vt active */
	VT_WAITACTIVE        uint = 0x5607 /* wait for vt active */
	VT_DISALLOCATE       uint = 0x5608 /* free memory associated to vt */
	VT_RESIZE            uint = 0x5609 /* set kernel's idea of screensize */
	VT_RESIZEX           uint = 0x560A /* set kernel's idea of screensize + more */
	VT_LOCKSWITCH        uint = 0x560B /* disallow vt switching */
	VT_UNLOCKSWITCH      uint = 0x560C /* allow vt switching */
	VT_GETHIFONTMASK     uint = 0x560D /* return hi font mask */
	VT_EVENT_SWITCH      uint = 0x0001 /* Console switch */
	VT_EVENT_BLANK       uint = 0x0002 /* Screen blank */
	VT_EVENT_UNBLANK     uint = 0x0004 /* Screen unblank */
	VT_EVENT_RESIZE      uint = 0x0008 /* Resize display */
	VT_MAX_EVENT         uint = 0x000F
	VT_WAITEVENT         uint = 0x560E /* Wait for an event */
	VT_SETACTIVATE       uint = 0x560F /* Activate and set the mode of a console */
)

const (
	KDGKBTYPE uint = 0x4B33 /* get keyboard type */
	KB_84     byte = 0x01
	KB_101    byte = 0x02 /* this is what we always answer */
	KB_OTHER  byte = 0x03
)

// Chvt switch current tty to ttyN
func Chvt(num uint) error {
	fd := getfd("")
	if fd == nil {
		return fmt.Errorf("nil fd")
	}

	err := unix.IoctlSetInt(int(fd.Fd()), VT_ACTIVATE, int(num))
	if err != nil {
		return fmt.Errorf("VT_ACTIVATE error:%s", err.Error())
	}
	err = unix.IoctlSetInt(int(fd.Fd()), VT_WAITACTIVE, int(num))
	if err != nil {
		return fmt.Errorf("VT_WAITACTIVE error:%s", err.Error())
	}
	return nil
}

// IsConsole can determine whether it is a console
func IsConsole(file *os.File) bool {
	ret, err := unix.IoctlGetInt(int(file.Fd()), KDGKBTYPE)
	if err == nil && ((byte(ret) == KB_101) || (byte(ret) == KB_84)) {
		return true
	}
	return false
}

// OpenConsole will open console by path name with O_RDWR or O_WRONLY or O_RDONLY
func OpenConsole(name string) *os.File {
	file, err := os.OpenFile(name, os.O_RDWR, 0)
	if os.IsPermission(err) {
		file, err = os.OpenFile(name, os.O_WRONLY, 0)
	}

	if os.IsPermission(err) {
		file, err = os.OpenFile(name, os.O_RDONLY, 0)
	}

	if os.IsPermission(err) {
		return nil
	}

	if !IsConsole(file) {
		file.Close()
		return nil
	}
	return file
}

func getfd(name string) *os.File {
	fileList := []string{
		name,
		"/dev/tty",
		"/dev/tty0",
		"/dev/vc/0",
		"/dev/console",
	}

	for _, ifile := range fileList {
		if ifile == "" {
			continue
		}
		file := OpenConsole(ifile)
		if file != nil {
			return file
		}

		if name == ifile {
			return nil
		}
	}

	stdFdList := []*os.File{
		os.Stdin,
		os.Stdout,
		os.Stderr,
	}
	for _, fd := range stdFdList {
		if IsConsole(fd) {
			return fd
		}
	}

	return nil
}
