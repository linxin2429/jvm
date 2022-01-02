package classfile

type ConstantClassInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

func (self *ConstantClassInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
