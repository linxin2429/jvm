package references

import (
	"jvm/src/instructions/base"
	"jvm/src/rtda"
)

type INVOKE_SPECIAL struct{ base.Index16Instruction }

// hack!
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}
