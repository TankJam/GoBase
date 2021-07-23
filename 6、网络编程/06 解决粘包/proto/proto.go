package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// Encode 消息封装
func Encode(msg string) ([]byte, error){
	// 读取消息的长度，转换为int32类型(占4个字节)
	var length = int32(len(msg))

	var pkg = new(bytes.Buffer)

	// 写入消息头
	// 大端和小端
	// 按照小端的顺序 把 length 写到pkg
	err := binary.Write(pkg, binary.LittleEndian, length)  // 封装数据长度
	if err != nil{
		return nil, err
	}

	// 写入需要传输的数据
	err = binary.Write(pkg, binary.LittleEndian, []byte(msg))
	if err != nil{
		return nil, err
	}
	return pkg.Bytes(), nil
}


// Decode 消息拆解
func Decode(reader *bufio.Reader) (string, error){
	// 获取需要读取客户端传递数据的长度
	lengthByte, _ := reader.Peek(4)  // 读取前4个字节的数据
	lengthBuff := bytes.NewBuffer(lengthByte)  // 读取数据的对象
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)  // 将读取数据长度给length
	if err != nil{
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数
	if int32(reader.Buffered()) < length + 4 {
		return "", err
	}

	// 读取真正的消息数据
	pack := make([]byte, int(4 + length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}

	return string(pack[4:]), nil
}