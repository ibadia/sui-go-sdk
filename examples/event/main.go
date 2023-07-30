package main

import (
	"context"
	"fmt"

	"github.com/yasir7ca/sui-go-sdk/constant"
	"github.com/yasir7ca/sui-go-sdk/models"
	"github.com/yasir7ca/sui-go-sdk/sui"
	"github.com/yasir7ca/sui-go-sdk/utils"
)

var ctx = context.Background()
var cli = sui.NewSuiClient(constant.BvTestnetEndpoint)

func main() {
	SuiGetEvents()
	SuiXQueryEvents()
}

func SuiGetEvents() {
	rsp, err := cli.SuiGetEvents(ctx, models.SuiGetEventsRequest{
		Digest: "HATq5p7MNynkBL5bLsdVqL3K38PxWHbqs7vndGiz5qrA",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiXQueryEvents() {
	rsp, err := cli.SuiXQueryEvents(ctx, models.SuiXQueryEventsRequest{
		SuiEventFilter: models.EventFilterByMoveEventType{
			MoveEventType: "0x3::validator::StakingRequestEvent",
		},
		Limit:           5,
		DescendingOrder: false,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}
