package gossip

import (
    "bytemurmur.com/mmbc/protos"
    "golang.org/x/net/context"
    "google.golang.org/grpc/peer"
    "log"
)

type communicate struct {
}

func (comm *communicate) GossipStream(serv *protos.Gossip_GossipStreamServer) error {

    return nil
}

func (comm *communicate) Ping(ctx *context.Context, msg *protos.Empty) (*protos.Empty, error) {
    var remoteAddress string
    pr, ok := peer.FromContext(ctx)
    if ok {
        if addr := pr.Addr; addr != nil {
            remoteAddress = addr.String()
        }
    }

    log.Println(remoteAddress)

    return &protos.Empty{}, nil
}
