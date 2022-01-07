package rtda

import "jvm/src/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
