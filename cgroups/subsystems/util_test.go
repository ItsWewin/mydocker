package subsystems

import "testing"

func TestGetcgroupRootPath(t *testing.T) {
	cpath, err := GetcgroupRootPath("memory", "my-test-cgroup-name")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("memory cpath: %s", cpath)

	cpath, err = GetcgroupRootPath("cpu", "my-test-cgroup-name")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("cpu cpath: %s", cpath)

	cpath, err = GetcgroupRootPath("cpuset", "my-test-cgroup-name")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("cpuset cpath: %s", cpath)
}