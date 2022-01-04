package constants

import (
	"jvm/src/instructions/base"
	"jvm/src/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// do nothing
}
