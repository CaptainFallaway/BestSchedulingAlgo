package internal

type Cpu struct {
	IScheduler

	workTime uint16
}

func NewCpu(scheduler IScheduler, optWorkTime ...uint16) *Cpu {
	var workTime uint16 = 100

	if len(optWorkTime) > 0 {
		workTime = optWorkTime[0]
	}

	return &Cpu{
		IScheduler: scheduler,
		workTime:   workTime,
	}
}

func (c *Cpu) Work() {
	c.SubtractTime(c.workTime)
}
