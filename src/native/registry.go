package native

import "jvm/src/rtda"

type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}
