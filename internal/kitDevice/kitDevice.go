/*
 * kitDevice.go Device-difine
 * by:Leftyer
 * dt:2025-11-01
 */
package kitDevice

import (
	"net"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	gopsnet "github.com/shirou/gopsutil/net"
)

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

/* Device-get BootTime-difine by:Leftyer dt:2025-11-01 */
func GetBootTime() time.Time {
	if t, err := host.BootTime(); err == nil {
		return time.Unix(int64(t), 0)
	}
	return time.Time{}
}

/* Device-get BIOS-UUID-difine by:Leftyer dt:2025-11-01 */
func GetBIOSUUID() string {
	info := GetDeviceInfo()
	if info == nil {
		return ""
	}
	return info.HostID
}

/* Device-get CPU-Name-difine by:Leftyer dt:2025-11-01 */
func GetCPUName() string {
	if cpus, err := cpu.Info(); err == nil && len(cpus) > 0 {
		return cpus[0].ModelName
	}
	return ""
}

/* Device-get CPU-Cores-difine by:Leftyer dt:2025-11-01 */
func GetCPUCores() int {
	if cnt, err := cpu.Counts(true); err == nil {
		return cnt
	}
	return 0
}

/* Device-get Memory-GB-difine by:Leftyer dt:2025-11-01 */
func GetMemoryGB() float64 {
	if vm, err := mem.VirtualMemory(); err == nil {
		return float64(vm.Total) / 1024 / 1024 / 1024
	}
	return 0
}

/* Device-get Disk-TotalGB-difine by:Leftyer dt:2025-11-01 */
func GetDiskTotalGB() float64 {
	if d, err := disk.Usage("/"); err == nil {
		return float64(d.Total) / 1024 / 1024 / 1024
	}
	return 0
}

/* Device-get IP-List-difine by:Leftyer dt:2025-11-01 */
func GetIPList() []string {
	var ips []string
	if ifs, err := gopsnet.Interfaces(); err == nil {
		for _, iface := range ifs {
			for _, addr := range iface.Addrs {
				ip, _, err := net.ParseCIDR(addr.Addr)
				if err == nil && !ip.IsLoopback() && ip.To4() != nil {
					ips = append(ips, ip.String())
				}
			}
		}
	}
	return ips
}

/* Device-get MAC-List-difine by:Leftyer dt:2025-11-01 */
func GetMACList() []string {
	var macs []string
	if ifs, err := gopsnet.Interfaces(); err == nil {
		for _, iface := range ifs {
			if iface.HardwareAddr != "" {
				macs = append(macs, iface.HardwareAddr)
			}
		}
	}
	return macs
}

/* Device-get Default-IP-difine by:Leftyer dt:2025-11-01 */
func GetDefaultIP() string {
	if conn, err := net.Dial("udp", "8.8.8.8:80"); err == nil {
		defer conn.Close()
		if addr := conn.LocalAddr().(*net.UDPAddr); addr != nil {
			return addr.IP.String()
		}
	}
	return ""
}
