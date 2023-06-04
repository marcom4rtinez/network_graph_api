package jagw_connector

import (
	"context"
	"log"

	"github.com/jalapeno-api-gateway/jagw-go/jagw"
)

func GetAllLinks() *jagw.LsNodeEdgeResponse {
	client := getJagwRequestClient()
	defer closeJagwRequestClient()
	request := &jagw.TopologyRequest{}
	response, err := client.GetLsNodeEdges(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling GetLsLinks on request service: %s", err)
	}
	return response
}
