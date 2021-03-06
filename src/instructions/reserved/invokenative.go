package reserved

import (
	"jvm/src/instructions/base"
	"jvm/src/native"
	_ "jvm/src/native/java/io"
	_ "jvm/src/native/java/lang"
	_ "jvm/src/native/java/security"
	_ "jvm/src/native/java/util/concurrent/atomic"
	_ "jvm/src/native/sun/io"
	_ "jvm/src/native/sun/misc"
	_ "jvm/src/native/sun/reflect"
	"jvm/src/rtda"
)

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()
	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}
	nativeMethod(frame)
}
