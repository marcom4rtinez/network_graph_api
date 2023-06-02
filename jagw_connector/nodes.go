package jagw_connector

import (
	"context"
	"log"

	"github.com/jalapeno-api-gateway/jagw-go/jagw"
)

func GetAllNodes() *jagw.LsNodeResponse {
	client := getJagwRequestClient()
	defer closeJagwRequestClient()
	request := &jagw.TopologyRequest{}

	response, err := client.GetLsNodes(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling GetLsNodes on request service: %s", err)
	}

	return response
}
