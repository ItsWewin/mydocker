package subsystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type MemoryLimit struct {

}

var MemorySystem = MemoryLimit{}

func (m *MemoryLimit) SourceType() string {
	return "memory"
}

func (m *MemoryLimit) Set(customCgroupName string, resourceLimit string) error {
	subCGroupPath, err := GetcgroupRootPath(m.SourceType(), customCgroupName)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join(subCGroupPath, "memory.limit_in_bytes"), []byte(resourceLimit), 0644); err != nil {
		return fmt.Errorf("set cgroup memory fail %v", err)
	}

	return nil
}

func (m *MemoryLimit) Apply(customCgroupName string, pid int) error {
	subCGroupPath, err := GetcgroupRootPath(m.SourceType(), customCgroupName)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join(subCGroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
		return fmt.Errorf("set cgroup memory fail %v", err)
	}

	return nil

}

func (m *MemoryLimit) Remove(customCgroupName string) error {
	subCGroupPath, err := GetcgroupRootPath(m.SourceType(), customCgroupName)
	if err != nil {
		return err
	}

	if err := os.Remove(subCGroupPath); err != nil {
		return fmt.Errorf("remove cgroup memory failed: %v", err)
	}

	return nil
}
