package main
import (
	"bytes"
	"fmt"
	"buftools"
	"encoding/binary"
)

func main() {


	b := bytes.NewBuffer(nil)

	buftools.WriteUint16(b, uint16(259))
	buftools.WriteUint16(b, uint16(65535))
	buftools.WriteString(b, "你好123\n你哈")

	r := bytes.NewReader(b.Bytes())

	fmt.Printf("read:%v", buftools.ReadUint16(r))
	fmt.Printf("read:%v", buftools.ReadUint16(r))
	fmt.Printf("read str:%v", buftools.ReadString(r))


	b2 := bytes.NewBuffer(nil)

	err := binary.Write(b2, binary.LittleEndian, []byte("abc123456"))
	if err!=nil{
		fmt.Println("err:",err.Error())
	}
	fmt.Printf("\nresult:%v\n",b2.Bytes())
	fmt.Printf("\nresult:%v\n",b2.String())

	b3 := bytes.NewBuffer(nil)
	testval:=uint64(0xff01)
	binary.Write(b3, binary.LittleEndian, testval)
	fmt.Printf("\nresult b3:%v\n",b3.Bytes())



}

