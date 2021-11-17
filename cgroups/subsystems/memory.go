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

func (m *MemoryLimit) SourceType() string {
	return "memory"
}

func (m *MemoryLimit) Set(customCgroupName string, res *ResourceConfig) error {
	subCGroupPath, err := GetcgroupRootPath(m.SourceType(), customCgroupName)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join(subCGroupPath, "memory.limit_in_bytes"), []byte(res.MemoryLimit), 0644); err != nil {
		return fmt.Errorf("set memory fail %v", err)
	}

	return nil
}

func (m *MemoryLimit) Apply(customCgroupName string, pid int) error {
	subCGroupPath, err := GetcgroupRootPath(m.SourceType(), customCgroupName)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join(subCGroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
		return fmt.Errorf("applay memory fail %v", err)
	}

	return nil

}

func (m *MemoryLimit) Remove(customCgroupName string) error {
	subCGroupPath, err := GetcgroupRootPath(m.SourceType(), customCgroupName)
	if err != nil {
		return err
	}

	if err := os.RemoveAll(subCGroupPath); err != nil {
		return fmt.Errorf("remove memory failed: %v", err)
	}

	return nil
}
