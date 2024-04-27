package internal

type Fifo struct {
	Arr []Process
}

func (f *Fifo) SubtractTime(time float64) {
	if len(f.Arr) == 0 {
		return
	}

	f.Arr[0].ExecTime -= time

	if f.Arr[0].ExecTime <= 0 {
		f.Arr = f.Arr[1:]
	}
}

func (f *Fifo) AddProcesses(processes []Process) {
	f.Arr = append(f.Arr, processes...)
}

func (f *Fifo) GetWorkTimes() []float64 {
	temp := make([]float64, 0, len(f.Arr))

	for _, item := range f.Arr {
		temp = append(temp, item.ExecTime)
	}

	return temp
}

func (f *Fifo) GetProcessNames() []string {
	temp := make([]string, 0, len(f.Arr))

	for _, item := range f.Arr {
		temp = append(temp, item.Name)
	}

	return temp
}
