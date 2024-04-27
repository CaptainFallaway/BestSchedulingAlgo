package internal

type TimeShare struct {
	Arr []Process
}

func (ts *TimeShare) SubtractTime(time float64) {
	subTime := time / float64(len(ts.Arr))

	auxArr := make([]Process, 0, len(ts.Arr))
	for _, proc := range ts.Arr {
		proc.ExecTime -= subTime
		if proc.ExecTime > 0 {
			auxArr = append(auxArr, proc)
		}
	}

	ts.Arr = auxArr
}

func (ts *TimeShare) AddProcesses(processes []Process) {
	ts.Arr = append(ts.Arr, processes...)
}

func (ts *TimeShare) GetWorkTimes() []float64 {
	temp := make([]float64, 0, len(ts.Arr))

	for _, item := range ts.Arr {
		temp = append(temp, item.ExecTime)
	}

	return temp
}

func (ts *TimeShare) GetProcessNames() []string {
	temp := make([]string, 0, len(ts.Arr))

	for _, item := range ts.Arr {
		temp = append(temp, item.Name)
	}

	return temp
}
