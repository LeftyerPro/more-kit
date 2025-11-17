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

/*---------------  Comp Area ---------------*/

/* Comp-Image by:Leftyer dt:2025-11-01 */
func CompImage(input, output string, fileType int) error {
	return kitComp.CompImg(kitComp.CompConfig{InputPath: input, OutPutPath: output, FileType: fileType})
}

/*---------------  Device Area ---------------*/

/* Device-GetInfo by:Leftyer dt:2025-11-01 */
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

/* Device-GetBootTime by:Leftyer dt:2025-11-01 */
func DeviceGetBootTime() string {
	return kitDevice.GetBootTime().Format("2006-01-02 15:04:05")
}

/* Device-GetBIOSUUID by:Leftyer dt:2025-11-01 */
func DeviceGetBIOSUUID() string {
	return kitDevice.GetBIOSUUID()
}

/* Device-GetCPUName by:Leftyer dt:2025-11-01 */
func DeviceGetCPUName() string {
	return kitDevice.GetCPUName()
}

/* Device-GetCPUCores by:Leftyer dt:2025-11-01 */
func DeviceGetCPUCores() int {
	return kitDevice.GetCPUCores()
}

/* Device-GetMemoryGB by:Leftyer dt:2025-11-01 */
func DeviceGetMemoryGB() float64 {
	return kitDevice.GetMemoryGB()
}

/* Device-GetDiskTotalGB by:Leftyer dt:2025-11-01 */
func DeviceGetDiskTotalGB() float64 {
	return kitDevice.GetDiskTotalGB()
}

/* Device-GetIPList by:Leftyer dt:2025-11-01 */
func DeviceGetIPList() []string {
	return kitDevice.GetIPList()
}

/* Device-GetMACList by:Leftyer dt:2025-11-01 */
func DeviceGetMACList() []string {
	return kitDevice.GetMACList()
}

/* Device-GetDefaultIP by:Leftyer dt:2025-11-01 */
func DeviceGetDefaultIP() string {
	return kitDevice.GetDefaultIP()
}

/*---------------  File Area ---------------*/

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

/* File-Read by:Leftyer dt:2025-11-01 */
func FileRead(filePath string) (string, error) {
	return kitFile.ReadFile(filePath)
}

/* File-Write by:Leftyer dt:2025-11-01 */
func FileWrite(filePath string, content string) error {
	return kitFile.WriteFile(filePath, content)
}

/* Json-Read by:Leftyer dt:2025-11-01 */
func JsonRead[T any](path string) (T, error) {
	return kitFile.ReadJSON[T](path)
}

/* Json-Write by:Leftyer dt:2025-11-01 */
func JsonWrite[T any](path string, obj T) error {
	return kitFile.WriteJSON(path, obj)
}
