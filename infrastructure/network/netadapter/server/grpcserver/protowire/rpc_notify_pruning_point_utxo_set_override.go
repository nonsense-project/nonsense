package protowire

import (
	"github.com/nonsense-project/nonsense/v2/app/appmessage"
	"github.com/pkg/errors"
)

func (x *NonsensedMessage_NotifyPruningPointUTXOSetOverrideRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NonsensedMessage_NotifyPruningPointUTXOSetOverrideRequest is nil")
	}
	return &appmessage.NotifyPruningPointUTXOSetOverrideRequestMessage{}, nil
}

func (x *NonsensedMessage_NotifyPruningPointUTXOSetOverrideRequest) fromAppMessage(_ *appmessage.NotifyPruningPointUTXOSetOverrideRequestMessage) error {
	x.NotifyPruningPointUTXOSetOverrideRequest = &NotifyPruningPointUTXOSetOverrideRequestMessage{}
	return nil
}

func (x *NonsensedMessage_NotifyPruningPointUTXOSetOverrideResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NonsensedMessage_NotifyPruningPointUTXOSetOverrideResponse is nil")
	}
	return x.NotifyPruningPointUTXOSetOverrideResponse.toAppMessage()
}

func (x *NonsensedMessage_NotifyPruningPointUTXOSetOverrideResponse) fromAppMessage(message *appmessage.NotifyPruningPointUTXOSetOverrideResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.NotifyPruningPointUTXOSetOverrideResponse = &NotifyPruningPointUTXOSetOverrideResponseMessage{
		Error: err,
	}
	return nil
}

func (x *NotifyPruningPointUTXOSetOverrideResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NotifyPruningPointUTXOSetOverrideResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}
	return &appmessage.NotifyPruningPointUTXOSetOverrideResponseMessage{
		Error: rpcErr,
	}, nil
}

func (x *NonsensedMessage_PruningPointUTXOSetOverrideNotification) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NonsensedMessage_PruningPointUTXOSetOverrideNotification is nil")
	}
	return &appmessage.PruningPointUTXOSetOverrideNotificationMessage{}, nil
}

func (x *NonsensedMessage_PruningPointUTXOSetOverrideNotification) fromAppMessage(_ *appmessage.PruningPointUTXOSetOverrideNotificationMessage) error {
	x.PruningPointUTXOSetOverrideNotification = &PruningPointUTXOSetOverrideNotificationMessage{}
	return nil
}

func (x *NonsensedMessage_StopNotifyingPruningPointUTXOSetOverrideRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NonsensedMessage_StopNotifyingPruningPointUTXOSetOverrideRequest is nil")
	}
	return &appmessage.StopNotifyingPruningPointUTXOSetOverrideRequestMessage{}, nil
}

func (x *NonsensedMessage_StopNotifyingPruningPointUTXOSetOverrideRequest) fromAppMessage(_ *appmessage.StopNotifyingPruningPointUTXOSetOverrideRequestMessage) error {
	x.StopNotifyingPruningPointUTXOSetOverrideRequest = &StopNotifyingPruningPointUTXOSetOverrideRequestMessage{}
	return nil
}

func (x *NonsensedMessage_StopNotifyingPruningPointUTXOSetOverrideResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NonsensedMessage_StopNotifyingPruningPointUTXOSetOverrideResponse is nil")
	}
	return x.StopNotifyingPruningPointUTXOSetOverrideResponse.toAppMessage()
}

func (x *NonsensedMessage_StopNotifyingPruningPointUTXOSetOverrideResponse) fromAppMessage(
	message *appmessage.StopNotifyingPruningPointUTXOSetOverrideResponseMessage) error {

	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.StopNotifyingPruningPointUTXOSetOverrideResponse = &StopNotifyingPruningPointUTXOSetOverrideResponseMessage{
		Error: err,
	}
	return nil
}

func (x *StopNotifyingPruningPointUTXOSetOverrideResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "StopNotifyingPruningPointUTXOSetOverrideResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}
	return &appmessage.StopNotifyingPruningPointUTXOSetOverrideResponseMessage{
		Error: rpcErr,
	}, nil
}
