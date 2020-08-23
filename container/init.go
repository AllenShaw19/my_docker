package container

import (
	"os"
	"syscall"

	log "github.com/sirupsen/logrus"
)

// RunContainerInitProcess Function
func RunContainerInitProcess(command string, args []string) error {
	log.Infof("command %s", command)

	// MS_NOEXEC 在本文件系统中不允许运行其他程序
	// MS_NOSUID 在本系统运行时，不允许set-user-id或者set-group-id
	// MS_NODEV 默认参数
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV

	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	argv := []string{command}

	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		log.Errorf(err.Error())
	}
	return nil
}
