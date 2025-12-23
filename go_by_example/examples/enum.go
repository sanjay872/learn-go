package examples

import "fmt"

type serverstate int // type of enum is int

// Possible States of ServerState
const (
	StateIdl serverstate = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[serverstate]string{
	StateIdl:       "idl",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (s serverstate) String() string {
	return stateName[s]
}

func transition(s serverstate) serverstate {
	switch s {
	case StateIdl:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdl
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("Unknow state!"))
	}
}

func Enum() {
	s1 := transition(StateIdl)

	fmt.Println(s1)

	s2 := transition(s1)

	fmt.Println(s2)
}
