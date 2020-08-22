package main

import (
	"fmt"
	"net"
	"sync"
	"sync/atomic"
)

func main() {
	lis, err := net.Listen("tcp", ":1789")
	if err != nil {
		return
	}

	lis = LimitListener(lis, 1)
}

func LimitListener(l net.Listener, n int) net.Listener {
	return &limitListener{l, 0, make(chan struct{}, n)}
}

type limitListener struct {
	net.Listener
	closed int32
	limit  chan struct{}
}

func (l *limitListener) acquire() { l.limit <- struct{}{} }

func (l *limitListener) release() { <-l.limit }

func (l *limitListener) shutdown() bool { return atomic.LoadInt32(&l.closed) != 0 }

func (l *limitListener) Accept() (net.Conn, error) {
	if l.shutdown() {
		return nil, fmt.Errorf("accept tcp %s: use of closed network connection", l.Listener.Addr().String())
	}
	l.acquire()
	c, err := l.Listener.Accept()
	if err != nil {
		l.release()
		return nil, err
	}
	return &limitListenerConn{Conn: c, release: l.release}, nil
}

func (l *limitListener) Close() error {
	if l.shutdown() {
		return nil
	}
	atomic.StoreInt32(&l.closed, 1)
	close(l.limit)
	return l.Listener.Close()
}

type limitListenerConn struct {
	net.Conn
	releaseOnce sync.Once
	release     func()
}

func (l *limitListenerConn) Close() error {
	err := l.Conn.Close()
	l.releaseOnce.Do(l.release)
	return err
}
