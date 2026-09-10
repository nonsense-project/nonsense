package rpchandlers

import (
	"github.com/nonsense-project/nonsense/v2/app/appmessage"
	"github.com/nonsense-project/nonsense/v2/app/rpc/rpccontext"
	"github.com/nonsense-project/nonsense/v2/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
