package main

import(
	"github.com/ShareSound/RPC-Server/account/handlers"
	"github.com/ShareSound/RPC-Server/rpc/account"
	"github.com/mshockwave/thrift-go/thrift"
	"github.com/ShareSound/RPC-Server/shared"
)

func main() {

	account_handler := handlers.NewAccountServiceHandler()
	account_processor := account.NewAccountServiceProcessor(account_handler)

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()
	transport, err := thrift.NewTServerSocket(":4444")
	if(err != nil ){
		panic(err)
	}

	server := thrift.NewTSimpleServer4(account_processor, transport, transportFactory, protocolFactory)
	shared.LogD.Println("Starting server on :4444...")
	server.Serve()
}
