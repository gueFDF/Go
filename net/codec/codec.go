package codec

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

// Ecode 将消息进行编码
func Encode(message string) ([]byte, error) {
	//和获取消息长度
	length := int32(len(message))

	pkg := new(bytes.Buffer)

	//在头部写入消息长度
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err

	}

	//写入实际消息
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil

}

// Decode 将消息解码
func Decode(reader *bufio.Reader) (string, error) {
	//读取长度,前四个字节
	lengthByte,_:=reader.Peek(4)
	lengthBuff:=bytes.NewBuffer(lengthByte)

	var length int32

	//将长度转化为int32
	err:=binary.Read(lengthBuff,binary.LittleEndian,&length)
	if err !=nil{
		return "",nil
	}
	//判断是否有足够的可读数据
	if int32(reader.Buffered())<length+4 {
		fmt.Println("可读数据不够")
		return "" ,err 
	}

	//读取真正的消息数据

	pack:=make([]byte,int(4+length))

	_,err=reader.Read(pack)
	if err!=nil {
		return "",err
	}

	return string(pack[4:]),nil


}
