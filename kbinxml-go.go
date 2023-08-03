package main

import "fmt"
import "github.com/beevik/etree"

import "os"

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

	binFile, err := os.ReadFile("testcases_out.kbin")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(binFile[0])

}

const (
	BUCKET_NONE = 0
	BUCKET_BYTE
	BUCKET_SHORT
	BUCKET_INT
)

// ID, bytes, bucket, XML name

type BinType struct {
	XMLNames   []string
	bucketType int

	count int
}

var binTypeMap = map[int]BinType{
	0x01: {[]string{"void"}, BUCKET_NONE, 1},
	0x02: {[]string{"s8"}, BUCKET_BYTE, 1},
	0x03: {[]string{"u8"}, BUCKET_BYTE, 1},
	0x04: {[]string{"s16"}, BUCKET_SHORT, 1},
	0x05: {[]string{"u8"}, BUCKET_SHORT, 1},
	0x06: {[]string{"s32"}, BUCKET_INT, 1},
	0x07: {[]string{"u32"}, BUCKET_INT, 1},
	0x08: {[]string{"s64"}, BUCKET_BYTE, 1},
	0x09: {[]string{"u64"}, BUCKET_BYTE, 1},

	0x0a: {[]string{"bin", "binary"}, BUCKET_INT, 1},
	0x0b: {[]string{"str", "string"}, BUCKET_INT, 1},
	0x0c: {[]string{"ip4"}, BUCKET_INT, 1},
	0x0d: {[]string{"time"}, BUCKET_INT, 1},

	0x0e: {[]string{"float", "f"}, BUCKET_INT, 1},
	0x0f: {[]string{"double", "d"}, BUCKET_INT, 1},

	0x10: {[]string{"2s8"}, BUCKET_INT, 2},
	0x11: {[]string{"2u8"}, BUCKET_INT, 2},
	0x12: {[]string{"2s16"}, BUCKET_INT, 2},
	0x13: {[]string{"2u16"}, BUCKET_INT, 2},
	0x14: {[]string{"2u32"}, BUCKET_INT, 2},
	0x15: {[]string{"2u32"}, BUCKET_INT, 2},
	0x16: {[]string{"2u64", "vs64"}, BUCKET_INT, 2},
	0x17: {[]string{"2u64", "vu64"}, BUCKET_INT, 2},
	0x18: {[]string{"2f"}, BUCKET_BYTE, 2},
	0x19: {[]string{"2d", "vd"}, BUCKET_INT, 2},

	0x1a: {[]string{"3s8"}, BUCKET_INT, 3},
	0x1b: {[]string{"3u8"}, BUCKET_INT, 3},
	0x1c: {[]string{"3s16"}, BUCKET_INT, 3},
	0x1d: {[]string{"3u16"}, BUCKET_INT, 3},
	0x1e: {[]string{"3s32"}, BUCKET_INT, 3},
	0x1f: {[]string{"3u32"}, BUCKET_INT, 3},
	0x20: {[]string{"3s64"}, BUCKET_INT, 3},
	0x21: {[]string{"3u64"}, BUCKET_INT, 3},
	0x22: {[]string{"3f"}, BUCKET_INT, 3},
	0x23: {[]string{"3d"}, BUCKET_INT, 3},

	0x24: {[]string{"4s8"}, BUCKET_INT, 4},
	0x25: {[]string{"4u8"}, BUCKET_INT, 4},
	0x26: {[]string{"4s16"}, BUCKET_INT, 4},
	0x27: {[]string{"4u16"}, BUCKET_INT, 4},
	0x28: {[]string{"4s32"}, BUCKET_INT, 4},
	0x29: {[]string{"4u32"}, BUCKET_INT, 4},
	0x2a: {[]string{"4s64"}, BUCKET_INT, 4},
	0x2b: {[]string{"4u64"}, BUCKET_INT, 4},
	0x2c: {[]string{"4f"}, BUCKET_INT, 4},
	0x2d: {[]string{"4d"}, BUCKET_INT, 4},

	0x2e: {[]string{"attr"}, BUCKET_INT, 1},
	0x2f: {[]string{"array"}, BUCKET_INT, 0},

	0x30: {[]string{"vs8"}, BUCKET_INT, 16},
	0x31: {[]string{"vu8"}, BUCKET_INT, 16},
	0x32: {[]string{"vs16"}, BUCKET_INT, 8},
	0x33: {[]string{"vu16"}, BUCKET_INT, 8},

	0x34: {[]string{"bool", "b"}, BUCKET_BYTE, 1},

	0x35: {[]string{"2b"}, BUCKET_SHORT, 2},
	0x36: {[]string{"3b"}, BUCKET_INT, 3},
	0x37: {[]string{"4b"}, BUCKET_INT, 4},
	0x38: {[]string{"vb"}, BUCKET_INT, 16},
}
