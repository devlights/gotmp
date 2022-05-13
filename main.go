package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		sigCtx, sigCxl   = signal.NotifyContext(mainCtx, os.Interrupt)
	)
	defer mainCxl()
	defer sigCxl()

	var (
		errCh = make(chan error)
	)
	go runServer(sigCtx, errCh)

	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		go runClient(sigCtx, errCh)
	}
LOOP:
	for {
		select {
		case err := <-errCh:
			fmt.Fprintln(os.Stderr, err)
			mainCxl()
		case <-sigCtx.Done():
			fmt.Println("SIGNAL RECV")
			break LOOP
		}
	}

	fmt.Println("MAIN DONE")
}

func runClient(pCtx context.Context, errCh chan<- error) {

}

func runServer(pCtx context.Context, errCh chan<- error) {
	server, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8888,
	})
	if err != nil {
		errCh <- err
		return
	}
	defer server.Close()

	fmt.Println("Server listening on localhost:8888")
LOOP:
	for {
		select {
		case <-pCtx.Done():
			fmt.Println("stop server")
			break LOOP
		default:
			server.SetDeadline(time.Now().Add(1 * time.Second))

			conn, err := server.Accept()
			if err != nil {
				break LOOP
			}
			fmt.Println(conn)
			conn.Close()
		}
	}

	errCh <- errors.New("SERVER SHUTDOWN")
}
