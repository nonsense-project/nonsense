package protowire

import (
	"github.com/nonsense-project/nonsense/v2/app/appmessage"
	"github.com/pkg/errors"
)

func (x *NonsensedMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NonsensedMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *NonsensedMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
