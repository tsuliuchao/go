package protocol

const (
	RPC_ADDITION  = "Calculator.Addition"
	RPC_SUBTRACTION = "Calculator.Subtraction"
	RPC_MULTIPLICATION = "Calculator.Multiplication"
	RPC_DIVISION = "Calculator.Division"
)

type Param struct {
	A int32
	B int32
}