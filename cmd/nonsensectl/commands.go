package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/nonsense-project/nonsense/v2/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.NonsensedMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.NonsensedMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.NonsensedMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.NonsensedMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.NonsensedMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.NonsensedMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.NonsensedMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.NonsensedMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.NonsensedMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.NonsensedMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.NonsensedMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.NonsensedMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.NonsensedMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.NonsensedMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.NonsensedMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.NonsensedMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.NonsensedMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.NonsensedMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.NonsensedMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.NonsensedMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.NonsensedMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.NonsensedMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.NonsensedMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.NonsensedMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.NonsensedMessage_BanRequest{}),
	reflect.TypeOf(protowire.NonsensedMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
