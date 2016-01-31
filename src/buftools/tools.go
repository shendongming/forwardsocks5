package buftools
import (
	"encoding/binary"
	"fmt"
	"io"
)


func WriteString(writer io.Writer, val string) {
	buf := []byte(val)
	WriteUint16(writer, uint16(len(buf)))
	writer.Write(buf)
}


func ReadString(reader io.Reader) string {
	string_len := ReadUint16(reader)

	buf := make([]byte, string_len)
	n, err := reader.Read(buf)
	if err != nil {
		fmt.Println("errror:", err.Error())
	}
	fmt.Println("read n:", n)
	return string(buf)
}



func ReadUint16(reader io.Reader) uint16 {
	var buf16 [2]byte;
	a := buf16[:]
	fmt.Printf("a1:%v\n", a)
	reader.Read(a)
	val := binary.LittleEndian.Uint16(a)
	fmt.Printf("reader :%v\n", val)
	return val

}

func WriteUint16(writer io.Writer, val uint16) {
	var buf16 [2]byte;
	a := buf16[:]
	fmt.Printf("a1:%v\n", a)
	binary.LittleEndian.PutUint16(a, val)
	fmt.Printf("a2:%v\n", a)
	writer.Write(a)

}