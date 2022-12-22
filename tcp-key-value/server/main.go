package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/thanhfphan/tcp-key-value/common"
)

type KVStore struct {
	mu   sync.Mutex
	data map[string]string
}

func (kv *KVStore) Get(args *common.GetArgs, reply *common.GetReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	reply.Value = kv.data[args.Key]
	fmt.Printf("get(%s) = %s\n", args.Key, reply.Value)

	return nil
}

func (kv *KVStore) Put(args *common.PutArgs, reply *common.PutReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	kv.data[args.Key] = args.Value
	fmt.Printf("put(%s) = %s\n", args.Key, args.Value)

	return nil
}

func main() {
	kv := &KVStore{}
	kv.data = map[string]string{}

	server := rpc.NewServer()
	server.Register(kv)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", common.TCPPort))
	if err != nil {
		panic(err)
	}

	go func() {
		defer listener.Close()
		for {
			conn, err := listener.Accept()
			if err != nil {
				break
			}

			go server.ServeConn(conn)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()
	<-done
	fmt.Println("exited")
}
