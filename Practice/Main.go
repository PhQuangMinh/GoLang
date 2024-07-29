package main

import "Practice/controller/consumerreceive"

func main() {
	//consumercreate.CreateNewCall("QueueSolve")
	consumerreceive.ReceiveCall("QueueSolve")
}
