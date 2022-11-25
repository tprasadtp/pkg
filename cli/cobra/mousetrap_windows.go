//go:build windows

package cobra

import (
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"
)

var (
	kernel                       = syscall.MustLoadDLL("kernel32.dll")
	procCreateToolhelp32Snapshot = kernel.MustFindProc("CreateToolhelp32Snapshot")
	procProcess32First           = kernel.MustFindProc("Process32FirstW")
	procProcess32Next            = kernel.MustFindProc("Process32NextW")
)

// defined by the Win32 API
const th32cs_snapprocess uintptr = 0x2

type processEntry32 struct {
	size              uint32
	cntUsage          uint32
	processID         uint32
	defaultHeapID     int
	moduleID          uint32
	cntThreads        uint32
	parentProcessID   uint32
	priorityClassBase int32
	flags             uint32
	exeFile           [syscall.MAX_PATH]uint16
}

func getProcessEntry(pid int) (pe *processEntry32, err error) {
	snapshot, _, err := procCreateToolhelp32Snapshot.Call(th32cs_snapprocess, uintptr(0))
	if snapshot == uintptr(syscall.InvalidHandle) {
		err = fmt.Errorf("CreateToolhelp32Snapshot: %v", err)
		return
	}
	defer syscall.CloseHandle(syscall.Handle(snapshot))

	var processEntry processEntry32
	processEntry.size = uint32(unsafe.Sizeof(processEntry))
	ok, _, err := procProcess32First.Call(snapshot, uintptr(unsafe.Pointer(&processEntry)))
	if ok == 0 {
		err = fmt.Errorf("Process32First: %v", err)
		return
	}

	for {
		if processEntry.processID == uint32(pid) {
			pe = &processEntry
			return
		}

		ok, _, err = procProcess32Next.Call(snapshot, uintptr(unsafe.Pointer(&processEntry)))
		if ok == 0 {
			err = fmt.Errorf("Process32Next: %v", err)
			return
		}
	}
}

func preExecHook(cmd *Command) {
	if MousetrapHelpText != "" {
		ppid := os.Getppid()
		pe, err := getProcessEntry(ppid)
		if err != nil {
			return false
		}
		// check if parent process is explorer.exe
		name := syscall.UTF16ToString(pe.exeFile[:])
		if name == "explorer.exe" {
			cmd.Print(MousetrapHelpText)
			if MousetrapDisplayDuration > 0 {
				time.Sleep(MousetrapDisplayDuration)
			} else {
				cmd.Println("Please press Return/Enter to continue...")
				fmt.Scanln()
			}
			os.Exit(1)
		}
	}
}
