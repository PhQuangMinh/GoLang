package callcontroller

import (
	"PhoneCall/controller/callcontroller/apilist"
	"PhoneCall/controller/callcontroller/consumer"
	"flag"
)

func CallController() {
	newCall := flag.Bool("new", false, "Create New Call")
	receiveCall := flag.Bool("rec", false, "Receive Call")
	updateResult := flag.Bool("up", false, "Update Result")
	getList := flag.Bool("gel", false, "Get List")
	getItem := flag.Bool("gei", false, "Get Item")

	flag.Parse()
	if *newCall {
		consumer.CreateNewCall("QueueSolve")
	} else if *receiveCall {
		consumer.ReceiveCall("QueueSolve")
	} else if *updateResult {
		consumer.UpdateResult("QueueResult")
	} else if *getList {
		apilist.GetList()
	} else if *getItem {
		apilist.GetItem()
	}
}
