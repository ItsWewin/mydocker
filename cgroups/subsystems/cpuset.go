package subsystems

type CpuSetLimit struct {
}

func (m *CpuSetLimit) SourceType() string {
	return "cpuset"
}

func (m *CpuSetLimit) Set(path string, resourceLimit string) error {
	return nil
}

func (m *CpuSetLimit) Apply(path string, pid int) error {
	return nil
}

func (m *CpuSetLimit) Remove(path string) error {
	return nil
}
