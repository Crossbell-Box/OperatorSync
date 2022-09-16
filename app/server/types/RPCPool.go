// Modified from https://dgraph.io/blog/post/rpc-vs-grpc/

package types

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"net"
	"net/rpc"
	"strings"
	"time"
)

type Pool struct {
	clients chan *rpc.Client
	Addr    string
}

func NewPool(addr string, dialNow bool) *Pool {
	p := new(Pool)
	p.Addr = addr
	p.clients = make(chan *rpc.Client, consts.RPCSETTINGS_MAXCAP)
	if dialNow {
		client, err := p.dialNew()
		if err != nil {
			return nil
		}
		p.clients <- client
	}
	return p
}

func (p *Pool) Call(serviceMethod string, args interface{}, reply interface{}) error {
	client, err := p.get()
	if err != nil {
		return err
	}
	if err = client.Call(serviceMethod, args, reply); err != nil {
		return err
	}

	select {
	case p.clients <- client:
		return nil
	default:
		return client.Close()
	}
}

func (p *Pool) dialNew() (*rpc.Client, error) {
	d := &net.Dialer{
		Timeout: consts.RPCSETTINGS_DIAL_TIMEOUT,
	}
	var (
		nConn net.Conn
		err   error
	)
	// Retry
	for i := 0; i < consts.RPCSETTINGS_DIAL_MAXTRY; i++ {
		nConn, err = d.Dial("tcp", p.Addr)
		if err == nil {
			break
		}
		if !strings.Contains(err.Error(), "refused") {
			break
		}
		time.Sleep(consts.RPCSETTINGS_DIAL_RETRY_INTERVAL)
	}
	if err != nil {
		return nil, err
	}

	return rpc.NewClient(nConn), nil
}

func (p *Pool) get() (*rpc.Client, error) {
	select {
	case client := <-p.clients:
		return client, nil
	default:
		return p.dialNew()
	}
}

func (p *Pool) Close() error {
	// We're not doing a clean exit here. A clean exit here would require
	// synchronization, which seems unnecessary for now. But, we should
	// add one if required later.
	return nil
}
