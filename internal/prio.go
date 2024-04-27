package internal

type Prio struct {
	Arr []Process
}

func sum(arr []Process) float64 {
	var acc uint16
	for _, item := range arr {
		acc += item.Prio
	}
	return float64(acc)
}

func (p *Prio) SubtractTime(time float64) {
	execTimes := make([]float64, 0, len(p.Arr))
	prioSum := sum(p.Arr)

	for _, proc := range p.Arr {
		execTimes = append(execTimes, (float64(proc.Prio)*time)/prioSum)
	}

	for i := range p.Arr {
		p.Arr[i].ExecTime -= execTimes[i]
	}

	newArr := make([]Process, 0)
	for _, proc := range p.Arr {
		if proc.ExecTime > 0 {
			newArr = append(newArr, proc)
		}
	}

	p.Arr = newArr
}

func (p *Prio) AddProcesses(processes []Process) {
	p.Arr = append(p.Arr, processes...)
}

func (p *Prio) GetWorkTimes() []float64 {
	temp := make([]float64, 0, len(p.Arr))

	for _, item := range p.Arr {
		temp = append(temp, item.ExecTime)
	}

	return temp
}

func (p *Prio) GetProcessNames() []string {
	temp := make([]string, 0, len(p.Arr))

	for _, item := range p.Arr {
		temp = append(temp, item.Name)
	}

	return temp
}
