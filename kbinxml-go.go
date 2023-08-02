package main

import "fmt"
import "github.com/beevik/etree"

const (
	UTF8 = "UTF-8"

	SIGNATURE = 0xA0

	SIG_COMPRESED    = 0x42
	SIG_UNCOMPRESSED = 0x45
)

type KBinXML struct {
	xmlDoc etree.Document

	encoding   string
	compressed bool

	dataSize    int
	dataMemSize int
	memSize     int
}

func (KBinXML) test() string {

	return "aa"
}

func (self KBinXML) fromText(input []byte) {

	self.xmlDoc = *etree.NewDocument()
	self.xmlDoc.ReadFromBytes(input)

	self.encoding = UTF8
	self.compressed = true
	self.dataSize = 0

}

func (self KBinXML) toText() string {
	ret, err := self.xmlDoc.WriteToString()

	if err != nil {
		fmt.Println("err")
	}

	return ret
}

func (KBinXML) isBinaryXML(input []byte) bool {
	if len(input) < 2 {
		return false
	}

	// TODO: also check if input[0] against sig compressed/uncompressed
	if input[0] != SIGNATURE {
		return false
	}

	return true
}

// TODO: Learn more about the purpose of this function
func (self KBinXML) getDataMemSize() {

	// dataLen := 0

	// self.xmlDoc.Attr

	// for t, e := range self.xmlDoc.Root().ChildElements() {

	// }

}

func (KBinXML) dataGrabAuto() int {
	return 0
}

func (KBinXML) dataAppendAuto() int {
	return 0
}

func (KBinXML) dataGrabString() int {
	return 0
}

func (KBinXML) dataAppendString() int {
	return 0
}

func (KBinXML) dataGrabAligned() int {
	return 0
}

func (KBinXML) dataAppendAligned() int {
	return 0
}

func (KBinXML) addNamespace() int {
	return 0
}

func (KBinXML) nodeToBinary() int {
	return 0
}

func (KBinXML) toBinary() int {
	return 0
}

func (KBinXML) fromBinary() int {
	return 0
}

func main() {

	fmt.Println("test")

}
