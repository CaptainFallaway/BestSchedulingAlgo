package internal

type IDiagram interface {
	Update(processes []float64)
}

type IScheduler interface {
	SubtractTime(time float64)
	AddProcesses(processes []Process)
	GetWorkTimes() []float64
}

type Process struct {
	Name     string
	ExecTime float64
	Prio     uint16
}
