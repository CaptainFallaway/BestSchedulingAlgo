package internal

type Cpu struct {
	IScheduler
	IDiagram

	workTime float64
}

func NewCpu(scheduler IScheduler, diagram IDiagram, optWorkTime ...float64) *Cpu {
	var workTime float64 = 100

	if len(optWorkTime) > 0 {
		workTime = optWorkTime[0]
	}

	return &Cpu{
		IScheduler: scheduler,
		IDiagram:   diagram,
		workTime:   workTime,
	}
}

func (c *Cpu) Work() {
	c.SubtractTime(c.workTime)
	c.Update(c.GetWorkTimes(), c.GetProcessNames())
}
