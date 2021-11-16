package cgroups

import (
	"github.com/sirupsen/logrus"
	"mydocker/cgroups/subsystems"
	"strconv"
)

const (
	resourceTypeMemory = "Memory"
	resourceTypeCpu    = "cpu"
	resourceTypeCpuSet = "cpuSet"
)

type Manager struct {
	// cgroup在hierarchy中的路径 相当于创建的cgroup目录相对于root cgroup目录的路径
	CustomCGroupName string
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

// 将进程pid加入到这个cgroup中
func (c *Manager) Apply(pid int) error {
	err := c.Subsystem.Apply(c.CustomCGroupName, pid)
	if err != nil {
		logrus.Errorf("memory cgroup apply failed, pid: %d", pid)
		return err
	}

	return nil
}

// 设置cgroup资源限制
func (c *Manager) Set(memoryLimitByte int) error {
	if err := c.Subsystem.Set(c.CustomCGroupName, strconv.Itoa(memoryLimitByte)); err != nil {
		logrus.Errorf("memory cgroup apply failed, pid: %d", memoryLimitByte)
		return err
	}

	return nil
}

//释放 cgroup
func (c *Manager) Destroy() error {
	if err := c.Subsystem.Remove(c.CustomCGroupName); err != nil {
		logrus.Errorf("memory cgroup apply failed, customCGroupName: %s", c.CustomCGroupName)
		return err
	}

	return nil
}
