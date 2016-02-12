package main

import(
	"./liftDriver"
	"./types"
	. "fmt"
	. "time"

)

func main() {
	//section 1 initialize
	var init_success bool =	liftDriver.LiftDriver_Initialize()
	Println(init_success)

	//section 2 resolver master-slave hierarki

	//section 3 master program

	/*

	//section 4 Slave code

	UddrAddr masterIP = 
	UddrConn masterConn = 

	buttonChannel := make(chan types.ButtonOrder)
	floorChannel := make(chan int)

	go liftDriver.LiftDriver_DetectButtonEvent(buttonChannel)
	go
	for{
		select:
		case buttonOrder := <- buttonChannel{
			//make msg corresponding to protocol
			network.Network_sendMessage(message, masterConnection)
		}
		case floor := <- floorChannel{

		}

	}

*/

}