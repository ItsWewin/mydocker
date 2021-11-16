package subsystems

type CpuLimit struct {

}

var CpuSystem = CpuLimit{}

func (m *CpuLimit) SourceType() string {
	return "cpu"
}

func (m *CpuLimit) Set(path string, resourceLimit string) error {

}

func (m *CpuLimit) Apply(path string, pid int) error {

}

func (m *CpuLimit) Remove(path string) error {

}
