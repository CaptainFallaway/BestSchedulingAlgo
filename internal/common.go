package internal

type IDiagram interface {
	Update(processes []Process)
}

type IScheduler interface {
	SubtractTime(time uint16)
	AddProcesses(processes []Process)
}

type Process struct {
	Name     string
	ExecTime uint16
	Prio     uint8
}
