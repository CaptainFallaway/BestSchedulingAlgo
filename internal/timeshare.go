package internal

// Linked list node
type llnode struct {
	Process
	next *llnode
}

func subtractTime(time uint16, node llnode) {
	node.ExecTime -= time

	if node.ExecTime <= 0 {

	}
}

type TimeShare struct {
	processCount uint16
	head         *llnode
}

func (ts *TimeShare) SubtractTime(time uint16) {

}

func (ts *TimeShare) AddProcesses(processes []Process) {
	// TODO
}
