package main

import (
	"context"
	"fmt"

	"github.com/nonsense-project/nonsense/v2/cmd/nonsensewallet/daemon/client"
	"github.com/nonsense-project/nonsense/v2/cmd/nonsensewallet/daemon/pb"
	"github.com/nonsense-project/nonsense/v2/cmd/nonsensewallet/utils"
)

func balance(conf *balanceConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()
	response, err := daemonClient.GetBalance(ctx, &pb.GetBalanceRequest{})
	if err != nil {
		return err
	}

	pendingSuffix := ""
	if response.Pending > 0 {
		pendingSuffix = " (pending)"
	}
	if conf.Verbose {
		pendingSuffix = ""
		println("Address                                                                       Available             Pending")
		println("-----------------------------------------------------------------------------------------------------------")
		for _, addressBalance := range response.AddressBalances {
			fmt.Printf("%s %s %s\n", addressBalance.Address, utils.FormatKls(addressBalance.Available), utils.FormatKls(addressBalance.Pending))
		}
		println("-----------------------------------------------------------------------------------------------------------")
		print("                                                 ")
	}
	fmt.Printf("Total balance, NNN %s %s%s\n", utils.FormatKls(response.Available), utils.FormatKls(response.Pending), pendingSuffix)

	return nil
}
