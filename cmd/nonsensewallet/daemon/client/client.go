package client

import (
	"context"
	"time"

	"github.com/nonsense-project/nonsense/v2/cmd/nonsensewallet/daemon/server"

	"github.com/pkg/errors"

	"github.com/nonsense-project/nonsense/v2/cmd/nonsensewallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the nonsensewalletd server, and returns the client instance
func Connect(address string) (pb.NonsensewalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("nonsensewallet daemon is not running, start it with `nonsensewallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewNonsensewalletdClient(conn), func() {
		conn.Close()
	}, nil
}
