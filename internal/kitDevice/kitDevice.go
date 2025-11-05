/*
 * kitDevice.go Device-difine
 * by:Leftyer
 * dt:2025-11-01
 */
package kitDevice

import "github.com/shirou/gopsutil/host"

/* Device-config-difine by:Leftyer dt:2025-11-01 */
type DeviceConfig struct{}

/* Device-get Info-difine by:Leftyer dt:2025-11-01 */
func GetDeviceInfo() *host.InfoStat {
	if hostInfo, err := host.Info(); err == nil && hostInfo.HostID != "" {
		return hostInfo
	}
	return nil
}

/* Device-get Id-difine by:Leftyer dt:2025-11-01 */
func GetDeviceId() string {
	info := GetDeviceInfo()
	if info == nil {
		return ""
	}
	return info.HostID
}

/* Device-get Name-difine by:Leftyer dt:2025-11-01 */
func GetDeviceName() string {
	info := GetDeviceInfo()
	if info == nil {
		return ""
	}
	return info.Hostname
}
