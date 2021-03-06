package db

import (
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/cloud"
	"golang.org/x/net/context"
	"google.golang.org/cloud/datastore"
	"golang.org/x/oauth2"

	"github.com/ShareSound/RPC-Server/shared"
)

var datastoreCtx context.Context
var defaultTokenSource oauth2.TokenSource = nil
func initDatastore(){

	datastoreCtx = context.Background()

	var err error
	defaultTokenSource, err = google.DefaultTokenSource(datastoreCtx,
		datastore.ScopeDatastore,
	)
	if err != nil || defaultTokenSource == nil{
		log.Fatalf("Error getting google default token source")
		return
	}
}

type DataStoreClient struct{
	Client *datastore.Client
	Ctx	context.Context
}

func GetNewDataStoreClient() (*DataStoreClient, error){
	client, err := datastore.NewClient(datastoreCtx, shared.PROJECT_ID, cloud.WithTokenSource(defaultTokenSource))
	return &DataStoreClient{
		Client: client,
		Ctx: datastoreCtx,
	}, err
}
func GetDatastoreContext() context.Context { return datastoreCtx }

func (this *DataStoreClient) NewKey(kind, name string, id int64, parent *datastore.Key) *datastore.Key {
	//Wrapper
	return datastore.NewKey(this.Ctx, kind, name, id, parent)
}
func (this *DataStoreClient) Run(query *datastore.Query) *datastore.Iterator {
	return this.Client.Run(this.Ctx, query)
}
func (this *DataStoreClient) RunInTransaction(f func(tx *datastore.Transaction) error, opts ...datastore.TransactionOption) (*datastore.Commit, error){
	return this.Client.RunInTransaction(this.Ctx, f, opts...)
}
