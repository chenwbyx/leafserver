/*
 * ///////////////////////////////////////////////////////////////////////
 * Auther: chenWB
 * Date: 2019-10-28
 * ///////////////////////////////////////////////////////////////////////
 * 在 Leaf 中，默认的 Protobuf Processor 将一个完整的 Protobuf 消息定义为如下格式：
 * -------------------------
 * | id | protobuf message |
 * -------------------------
 * 其中 id 为 2 个字节。如果你选择使用 TCP 协议时，在网络中传输的消息格式如下：
 * -------------------------------
 * | len | id | protobuf message |
 * -------------------------------
 * 如果你选择使用 WebSocket 协议时，发送的消息格式如下：
 * -------------------------
 * | id | protobuf message |
 * -------------------------
 * 其中 len 默认为两个字节，len 和 id 默认使用网络字节序。
 * //TODO:本测试使用TCP测试
 * 最终生成的字节流为：
 * -------------------------------------------------------
 * | len(2字节) | id(2字节) | protobuf message(len(data)) |
 * -------------------------------------------------------
 * 其中len = len(id) + len(data)
 */
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/chenwbyx/leafserver/server/msg"
	"github.com/golang/protobuf/proto"
	"net"
)

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}

	info := &msg.UserRegister{
		LoginName: "wssssenbdddddoqw",
		LoginPW: "123cx456",
	}
	data, err  := proto.Marshal(info)

	// len + data
	m := make([]byte, 4+len(data))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data) + 2))

	copy(m[2:], IntToBytes(1)[2:])
	copy(m[4:], data)
	fmt.Println(m)
	// 发送消息
	conn.Write(m)

}