package proto

import (
	"fmt"
)

// GRPCClient is an implementation of KV that talks over RPC.
type GRPCClient struct{ client RoboSdkClient }

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl RoboSdkApi
}

func (m *GRPCServer) mustEmbedUnimplementedRoboSdkServer() {
	fmt.Println("mustEmbedUnimplementedKVServer")
}
