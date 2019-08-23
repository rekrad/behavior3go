package actions

import (
	b3 "github.com/rekrad/behavior3go"
	. "github.com/rekrad/behavior3go/core"
)

type Runner struct {
	Action
}

func (this *Runner) OnTick(tick *Tick) b3.Status {
	return b3.RUNNING
}
