package subsystems

type Subsystem interface {
	SourceType() string
	Set(path string, resourceLimit string) error
	Apply(path string, pid int) error
	Remove(path string) error
}