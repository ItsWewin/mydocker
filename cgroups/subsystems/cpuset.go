package subsystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type CpuSetLimit struct {
}

func (m *CpuSetLimit) SourceType() string {
	return "cpuset"
}

func (m *CpuSetLimit) Set(customCgroupName string, res *ResourceConfig) error {
	subCGroupPath, err := GetcgroupRootPath(m.SourceType(), customCgroupName)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join(subCGroupPath, "cpuset.cpus"), []byte(res.CpuSet), 0644); err != nil {
		return fmt.Errorf("set cpuset fail %v", err)
	}

	return nil
}

func (m *CpuSetLimit) Apply(customCgroupName string, pid int) error {
	subCGroupPath, err := GetcgroupRootPath(m.SourceType(), customCgroupName)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join(subCGroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
		return fmt.Errorf("set cpuset fail %v", err)
	}

	return nil
}

func (m *CpuSetLimit) Remove(customCgroupName string) error {
	subCGroupPath, err := GetcgroupRootPath(m.SourceType(), customCgroupName)
	if err != nil {
		return fmt.Errorf("remove cpuset failed: %v", err)
	}

	err = os.RemoveAll(subCGroupPath)
	if err != nil {
		return fmt.Errorf("remove cpuset fail: %v", err)
	}

	return nil
}
