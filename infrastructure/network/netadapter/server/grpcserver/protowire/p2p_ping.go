package protowire

import (
	"github.com/nonsense-project/nonsense/v2/app/appmessage"
	"github.com/pkg/errors"
)

func (x *NonsensedMessage_Ping) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NonsensedMessage_Ping is nil")
	}
	return x.Ping.toAppMessage()
}

func (x *PingMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PingMessage is nil")
	}
	return &appmessage.MsgPing{
		Nonce: x.Nonce,
	}, nil
}

func (x *NonsensedMessage_Ping) fromAppMessage(msgPing *appmessage.MsgPing) error {
	x.Ping = &PingMessage{
		Nonce: msgPing.Nonce,
	}
	return nil
}
