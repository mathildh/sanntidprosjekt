package types

type MotorDirection int
type ButtonType int
const (
	MotorDirection_DOWN MotorDirection = iota - 1
	MotorDirection_STOP 
	MotorDirection_UP 
	)

const (
	ButtonType_UP ButtonType = iota
	ButtonType_DOWN 
	ButtonType_INTERNAL 
	)

type ButtonOrder struct{
	Button ButtonType
	Floor int
}
