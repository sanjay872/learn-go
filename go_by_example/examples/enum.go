package examples

type serverstate int

// Possible values of ServerState

const (
	StateIdl serverstate = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[serverstate]string{
	StateIdl:       "",
	StateConnected: "",
	StateError:     "",
	StateRetrying:  "",
}

func Enum() {

}
