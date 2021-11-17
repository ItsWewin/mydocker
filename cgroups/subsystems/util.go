package subsystems

import (
	"bufio"
	"os"
	"path"
	"strings"
)

// rootPath + resourceType + customCgrpup +
// /sys/fs/cgroup/memory/mydocker-cgroup/memory.limit_in_bytes
// /sys/fs/cgroup/memory/mydocker-cgroup/tasks

func FindCgroupMountpoint(subsystem string) string {
	f, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return ""
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		fields := strings.Split(txt, " ")
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == subsystem {
				return fields[4]
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return ""
	}

	return ""
}


func GetcgroupRootPath(resourceType string, customCgrpupName string) (string, error) {
	cgroupRootPath := FindCgroupMountpoint(resourceType)
	cpath := path.Join(cgroupRootPath, customCgrpupName)

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
