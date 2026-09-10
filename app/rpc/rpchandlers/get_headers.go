package rpchandlers

import (
	"github.com/nonsense-project/nonsense/v2/app/appmessage"
	"github.com/nonsense-project/nonsense/v2/app/rpc/rpccontext"
	"github.com/nonsense-project/nonsense/v2/infrastructure/network/netadapter/router"
)

// HandleGetHeaders handles the respectively named RPC command
func HandleGetHeaders(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetHeadersResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
