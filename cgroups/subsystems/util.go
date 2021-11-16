package subsystems

import (
	"os"
	"path"
)

// rootPath + resourceType + customCgrpup +
// /sys/fs/cgroup/memory/mydocker-cgroup/memory.limit_in_bytes
// /sys/fs/cgroup/memory/mydocker-cgroup/tasks

const cgroupRootPath = `/sys/fs/cgroup`
func GetcgroupRootPath(resourceType string, customCgrpupName string) (string, error) {
	 cpath := path.Join(cgroupRootPath, resourceType, customCgrpupName)

	 if _, err := os.Stat(cpath); err != nil {
	 	if os.IsNotExist(err) {
	 		err := os.Mkdir(cpath, 0755)
	 		if err != nil {
	 			return "", err
			}
			return cpath, nil
		} else {
			return "", err
		}
	 }

	 return cpath, nil
}
