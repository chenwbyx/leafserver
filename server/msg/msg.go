package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)

//var Processor network.Processor
var Processor = protobuf.NewProcessor()

func init() {
	Processor.Register(&UserLogin{})
	Processor.Register(&UserRegister{})
	Processor.Register(&UserST{})
}
