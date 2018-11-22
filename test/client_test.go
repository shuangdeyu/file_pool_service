package test

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"log"
	"testing"
)

type Out struct {
	OutData interface{}
	OutMsg  string
}

type Reply struct {
	Out *Out `msg:"Out"`
}

func TestClient(t *testing.T) {
	//addr2 := flag.String("addr", "106.14.113.45:6667", "server address")
	addr2 := flag.String("addr", "localhost:6000", "server address")
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr2, "")
	xclient := client.NewXClient("FilePoolService", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := map[string]interface{}{
		"UserId": 1,
	}

	reply := &Reply{}
	err := xclient.Call(context.Background(), "GetUserInfo", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Println(reply.Out.OutData, reply.Out.OutMsg)
}
