package internal

type Dispatcher struct {
	Cpus []Cpu // TODO: Make an interface for the cpu objects
}

// Add the processes to the cpus schedulers
func (d *Dispatcher) Push(processes []Process) {
	for _, cpu := range d.Cpus {
		cpu.AddProcesses(processes)
	}
}
