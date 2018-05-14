package gossip

import (
	"bytemurmur.com/mmbc/protos"
	"google.golang.org/grpc"
	"sync"
)

type Handler func (*protos.Envelope) error;

type connection struct{
	handler      Handler
	outBuf       chan *protos.Envelope
	lock 		 *sync.RWMutex
	clientStream *protos.Gossip_GossipStreamClient
	serverStream *protos.Gossip_GossipStreamServer
}

func NewConnection() *connection {
	return &connection{
		outBuf:  make(chan *protos.Envelope, 10),
		lock:	 &sync.RWMutex{},
	}
}

func (conn *connection)serverConnection() error {

	return nil
}

func (conn *connection)readFromStream() {
	stream := conn.getStream()

	if stream == nil {

	}

	for {
		if envelope, err := stream.Recv(); err != nil {
			conn.handler(envelope)
		}
	}
}

func (conn *connection)writeStream() {
	stream := conn.getStream()

	if stream == nil {

	}

	for {
		envelop := <- conn.outBuf
		if err := stream.Send(envelop); err != nil {

		}
	}
}

type Stream interface {
	Send(envelope *protos.Envelope) error
	Recv() (*protos.Envelope, error)
	grpc.Stream
}

func (conn *connection)getStream() Stream {

	if conn.clientStream != nil {
		return conn.clientStream
	}

	if conn.serverStream != nil {
		return conn.serverStream
	}

	return nil
}
