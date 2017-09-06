package protox

import "github.com/golang/protobuf/proto"

type Message interface {
	Validate() error
	Reset()
	String() string
	ProtoMessage()
}

func Unmarshal(buf []byte, msg Message) error {
	if err := proto.Unmarshal(buf, msg); err != nil {
		return err
	}
	return msg.Validate()
}

func Marshal(msg Message) ([]byte, error) {
	if err := msg.Validate(); err != nil {
		return nil, err
	}
	return proto.Marshal(msg)
}
