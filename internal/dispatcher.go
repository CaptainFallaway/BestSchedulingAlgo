package internal

type Dispatcher struct {
	Cpus []*Cpu // TODO: Make an interface for the cpu objects
}

func NewDispatcher(cpus ...*Cpu) *Dispatcher {
	return &Dispatcher{
		Cpus: cpus,
	}
}

// Add the processes to the cpus schedulers
func (d *Dispatcher) Push(processes []Process) {
	for _, cpu := range d.Cpus {
		cpu.AddProcesses(processes)
	}
}

func (d *Dispatcher) Work() {
	for _, cpu := range d.Cpus {
		go cpu.Work()
	}
}
