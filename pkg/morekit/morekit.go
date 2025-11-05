/*
 * morekit.go Kit-difine
 * by:Leftyer
 * dt:2025-11-01
 */
package morekit

import (
	"github.com/LeftyerPro/more-kit/internal/kitComp"
	"github.com/LeftyerPro/more-kit/internal/kitDevice"
	"github.com/LeftyerPro/more-kit/internal/kitFile"
	"github.com/shirou/gopsutil/host"
)

/* Comp-Image by:Leftyer dt:2025-11-01 */
func CompImage(input, output string, fileType int) error {
	c := &kitComp.CompConfig{
		InputPath:  input,
		OutPutPath: output,
		FileType:   fileType,
	}
	return c.CompImg()
}

/* Device-GetId by:Leftyer dt:2025-11-01 */
func DeviceGetInfo() *host.InfoStat {
	return kitDevice.GetDeviceInfo()
}

/* Device-GetId by:Leftyer dt:2025-11-01 */
func DeviceGetId() string {
	return kitDevice.GetDeviceId()
}

/* Device-GetName by:Leftyer dt:2025-11-01 */
func DeviceGetName() string {
	return kitDevice.GetDeviceName()
}

/* Folder-IsExist by:Leftyer dt:2025-11-01 */
func FolderIsExist(path string, isClear int) bool {
	return kitFile.IsFolderExist(path, isClear)
}

/* Folder-Cpoy by:Leftyer dt:2025-11-01 */
func FolderCpoy(src, dst string) error {
	return kitFile.CopyFolder(src, dst)
}

/* File-IsExist by:Leftyer dt:2025-11-01 */
func FileIsExist(path string) bool {
	return kitFile.IsFileExist(path)
}

/* File-Cpoy by:Leftyer dt:2025-11-01 */
func FileCpoy(src, dst string) error {
	return kitFile.CopyFile(src, dst)
}
