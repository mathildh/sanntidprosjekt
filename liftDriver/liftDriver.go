package liftDriver

/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
#include "elev.h"
*/
import "C"

import (
	"../types"
	"time"
)

const BUTTON_RATE = time.Millisecond * 0.1
const FLOOR_RATE = time.Millisecond

const TOTAL_FLOORS = 4 //TENK PÅ DETTE???????????????????????????????????????????????????????????
const TOTAL_BUTTON_TYPES = 3
//should we create a button_type enum in go??
//type buttonType int

func LiftDriver_Initialize() bool {
	return int(C.elev_init()) != 0
}

func LiftDriver_setMotorDirection(direction types.MotorDirection) {
	C.elev_set_motor_direction(C.elev_motor_direction_t(C.int(direction)))
}
func LiftDriver_SetButtonLamp(button types.ButtonType, floor int, onOrOff int) {
	C.elev_set_button_lamp(C.elev_button_type_t(C.int(button)), C.int(floor), C.int(onOrOff))
}
func LiftDriver_SetFloorIndicator(floor int) {
	C.elev_set_floor_indicator(C.int(floor))
}

func LiftDriver_SetDoorLamp(onOrOff int) {
	C.elev_set_door_open_lamp(C.int(onOrOff))
}

func LiftDriver_GetButtonSignal(button types.ButtonType, floor int) bool {
	return int(C.elev_get_button_signal(C.elev_button_type_t(C.int(button)), C.int(floor))) != 0
}

func LiftDriver_GetFloor() int {
	return int(C.elev_get_floor_sensor_signal())
}



func LiftDriver_DetectButtonEvent(buttonChannel chan types.ButtonOrder){
	var previous[TOTAL_FLOORS][TOTAL_BUTTON_TYPES] bool 
	for{
		time.Sleep(BUTTON_RATE)
		for floor:=0; floor < TOTAL_FLOORS; floor++{
			for button := ButtonType_DOWN; button < ButtonType_INTERNAL; button++{
				var ifButtonPressed bool = LiftDriver_GetButtonSignal(button,floor)

				if  ifButtonPressed &&  !previous[button][floor] {
					buttonOrder := make(types.ButtonOrder)
					buttonOrder.ButtonType = button
					buttonOrder.Floor = floor
					buttonChannel <- buttonOrder
					previous[floor][button] = 1
				} else if ifButtonPressed != previous[button][floor]{
					previous[floor][button] = 0
				}
			}
		}
	}
}

func LiftDriver_DetectFloorEvent(floorChannel chan int){
	var previousFloor int = -1
	for{
		time.Sleep(FLOOR_RATE)
		var currentFloor int = LiftDriver_GetFloor()
		if currentFloor != previousFloor && currentFloor != -1{
			floorChannel <- currentFloor
			previousFloor = currentFloor
		}
	}
}
