package main

import (
	"fmt"
	"net/rpc"

	"github.com/thanhfphan/tcp-key-value/common"
)

func connect() (*rpc.Client, error) {
	client, err := rpc.Dial("tcp", fmt.Sprintf(":%d", common.TCPPort))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func get(key string) (string, error) {
	client, err := connect()
	if err != nil {
		return "", err
	}
	defer client.Close()

	args := common.GetArgs{
		Key: key,
	}
	reply := common.GetReply{}
	err = client.Call("KVStore.Get", &args, &reply)
	if err != nil {
		return "", err
	}

	return reply.Value, nil
}

func put(key, value string) error {
	client, err := connect()
	if err != nil {
		return err
	}
	defer client.Close()

	args := common.PutArgs{
		Key:   key,
		Value: value,
	}
	reply := common.PutReply{}

	return client.Call("KVStore.Put", &args, &reply)
}

func main() {
	err := put("name", "vegeta")
	if err != nil {
		panic(err)
	}

	value, err := get("name")
	if err != nil {
		panic(err)
	}

	fmt.Printf("get(name) -> %s\n", value)
}
