package main

import (
	"context"
	"fmt"
	"log"

	// import osmosis gamm types, presumably this will work for other modules as well
	"github.com/osmosis-labs/osmosis/x/gamm/types"
	// importing the general purpose Cosmos blockchain client
	"github.com/tendermint/starport/starport/pkg/cosmosclient"
)

func main() {

	// create an instance of cosmosclient NOTE: NodeAddress should not be hardcoded in production
	cosmos, err := cosmosclient.New(context.Background(), cosmosclient.WithNodeAddress("http://66.228.33.64:1317"))
	if err != nil {
		log.Fatal(err)
	}

	// instantiate a query client
	queryClient := types.NewQueryClient(cosmos.Context)

	// query the osmosis blockchain
	// store all posts in queryResp
	queryResp, err := queryClient.Pool(context.Background(), &types.QueryPoolRequest{PoolId: uint64(1),})
	if err != nil {
		log.Fatal(err)
	}

	// print response from querying all the posts
	//fmt.Print("\n\nAll posts:\n\n")
	fmt.Println(queryResp)
}
