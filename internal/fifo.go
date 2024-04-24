package internal

type Fifo struct {
	Arr []Process
}

func (f *Fifo) SubtractTime(time uint16) {
	if len(f.Arr) == 0 {
		return
	}

	item := f.Arr[0]
	item.ExecTime -= time

	if item.ExecTime <= 0 {
		f.Arr = f.Arr[1:]
	}
}

func (f *Fifo) AddProcesses(processes []Process) {
	for _, process := range processes {
		f.Arr = append(f.Arr, process)
	}
}
