package subsystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type CpuLimit struct {
}

var CpuSystem = CpuLimit{}

func (m *CpuLimit) SourceType() string {
	return "cpu"
}

func (m *CpuLimit) Set(customCgroupName string, res *ResourceConfig) error {
	subCGroupPath, err := GetcgroupRootPath(m.SourceType(), customCgroupName)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join(subCGroupPath, "cpu.shares"), []byte(res.CpuShare), 0644); err != nil {
		return fmt.Errorf("set cpu fail %v", err)
	}

	return nil
}

func (m *CpuLimit) Apply(customCgroupName string, pid int) error {
	subCGroupPath, err := GetcgroupRootPath(m.SourceType(), customCgroupName)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join(subCGroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
		return fmt.Errorf("apply cpu fail %v", err)
	}

	return nil
}

func (m *CpuLimit) Remove(customCgroupName string) error {
	subCGroupPath, err := GetcgroupRootPath(m.SourceType(), customCgroupName)
	if err != nil {
		return err
	}

	err = os.RemoveAll(subCGroupPath)
	if err != nil {
		return fmt.Errorf("remove cpu share fail: %v", err)
	}

	return nil
}
