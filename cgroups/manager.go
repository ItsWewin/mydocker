package cgroups

import (
	"github.com/sirupsen/logrus"
	"mydocker/cgroups/subsystems"
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
	Subsystems []subsystems.Subsystem
}

func NewCgroupManager(customCGroupName string) *Manager {
	m := &Manager{
		CustomCGroupName: customCGroupName,
	}

	m.Subsystems = []subsystems.Subsystem{
		&subsystems.MemoryLimit{},
		//&subsystems.CpuLimit{},
		&subsystems.CpuSetLimit{},
	}

	return m
}

// 将进程pid加入到这个cgroup中
func (c *Manager) Apply(pid int) error {
	for _, subSystem := range c.Subsystems {
		err := subSystem.Apply(c.CustomCGroupName, pid)
		if err != nil {
			logrus.Errorf("cgroup apply failed, err: %s, pid: %d", err, pid)
			return err
		}
	}

	return nil
}

// 设置cgroup资源限制
func (c *Manager) Set(res *subsystems.ResourceConfig) error {
	for _, system := range c.Subsystems {
		if err := system.Set(c.CustomCGroupName, res); err != nil {
			logrus.Errorf("cgroup set failed, system: %s, err: %s, res: %#v", system, err, res)
			return err
		}
	}

	return nil
}

//释放 cgroup
func (c *Manager) Destroy() error {
	for _, system := range c.Subsystems {
		if err := system.Remove(c.CustomCGroupName); err != nil {
			logrus.Errorf("cgroup destroy failed, err: %s, customCGroupName: %s", err, c.CustomCGroupName)
			return err
		}
	}

	return nil
}