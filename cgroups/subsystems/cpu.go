package subsystems

import (
	"fmt"
	"io/ioutil"
	"path"
)

type CpuLimit struct {
}

var CpuSystem = CpuLimit{}

func (m *CpuLimit) SourceType() string {
	return "cpu"
}

func (m *CpuLimit) Set(customCgroupName string, resourceLimit string) error {
	subCGroupPath, err := GetcgroupRootPath(m.SourceType(), customCgroupName)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join(subCGroupPath, "cpu"), []byte(resourceLimit), 0644); err != nil {
		return fmt.Errorf("set cgroup memory fail %v", err)
	}

	return nil
}

func (m *CpuLimit) Apply(path string, pid int) error {
	return nil
}

func (m *CpuLimit) Remove(path string) error {
	return nil
}
