package protowire

import (
	"github.com/nonsense-project/nonsense/v2/app/appmessage"
	"github.com/pkg/errors"
)

func (x *NonsensedMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NonsensedMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *NonsensedMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
