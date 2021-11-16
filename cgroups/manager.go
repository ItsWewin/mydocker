package cgroups

import (
	"mydocker/cgroups/subsystems"
)

const (
	resourceTypeMemory = "Memory"
	resourceTypeCpu = "cpu"
	resourceTypeCpuSet = "cpuSet"
)

type Manager struct {
	// cgroup在hierarchy中的路径 相当于创建的cgroup目录相对于root cgroup目录的路径
	CustomCGroupName     string
	// 资源配置
	Subsystem subsystems.Subsystem
}

func NewCgroupManager(customCGroupName string, resourceType string) *Manager {
	m := &Manager{
		CustomCGroupName: customCGroupName,
	}

	var subsystem subsystems.Subsystem
	switch resourceType {
	case resourceTypeMemory:
		subsystem = &subsystems.MemoryLimit{}
	case resourceTypeCpu:
		subsystem = &subsystems.CpuLimit{}
	case resourceTypeCpuSet:
		subsystem = &subsystems.CpuSetLimit{}
	default:
		return nil
	}
	m.Subsystem = subsystem

	return m
}

//// 将进程pid加入到这个cgroup中
//func (c *Manager) Apply(pid int) error {
//	switch c.ResourceType {
//	case "resourceTypeMemory":
//		subsystems.MemorySystem
//	}
//}
//
//// 设置cgroup资源限制
//func (c *Manager) Set(res *subsystems.ResourceConfig) error {
//	for _, subSysIns := range(subsystems.SubsystemsIns) {
//		subSysIns.Set(c.Path, res)
//	}
//	return nil
//}
//
////释放cgroup
//func (c *Manager) Destroy() error {
//	for _, subSysIns := range(subsystems.SubsystemsIns) {
//		if err := subSysIns.Remove(c.Path); err != nil {
//			logrus.Warnf("remove cgroup fail %v", err)
//		}
//	}
//	return nil
//}
//
