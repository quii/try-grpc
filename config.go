package try_grpc

import "fmt"

var (
	Port          = ":5000"
	FridgeAddress = fmt.Sprintf("localhost%s", Port)
)