package subsystems

type CpuSetLimit struct {

}

var CpuSetSystem = CpuSetLimit{}

func (m *CpuSetLimit) SourceType() string {
	return "cpuset"
}

func (m *CpuSetLimit) Set(path string, resourceLimit string) error {

}

func (m *CpuSetLimit) Apply(path string, pid int) error {

}

func (m *CpuSetLimit) Remove(path string) error {

}
