package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index")
}
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.stringIndex)
}
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUft8Info)
	return utf8Info.str
}

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_INTEGER:
		return &ConstantIntegerInfo{}
	case CONSTANT_FLOAT:
		return &ConstantFloatInfo{}
	case CONSTANT_LONG:
		return &ConstantLongInfo{}
	case CONSTANT_DOUBLE:
		return &ConstantDoubleInfo{}
	case CONSTANT_UTF8:
		return &ConstantUft8Info{}
	case CONSTANT_STRING:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_CLASS:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_FIELDREF:
		return &ConstantFieldRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_METHODREF:
		return &ConstantMethodRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_INTERFACEMETHODREF:
		return &ConstantInterfaceRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_NAMEANDTYPE:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_METHODTYPE:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_METHODHANDLE:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_INVOKEDYNAMIC:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag")
	}
}
