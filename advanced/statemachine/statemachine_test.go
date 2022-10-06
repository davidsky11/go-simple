package statemachine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMachineStructure(t *testing.T) {
	machine := Machine{
		ID:      "machine-1",
		Initial: "on",
		States: StateMap{
			"on": MachineState{
				On: TransitionMap{
					"TOGGLE": MachineTransition{
						To: "off",
					},
				},
			},
			"off": MachineState{
				On: TransitionMap{
					"TOGGLE": MachineTransition{
						To: "on",
					},
				},
			},
		},
	}
	output := machine.Transition("TOGGLE")
	assert.Equal(t, output, "off", "Transition should occur on toggle.")
	output = machine.Transition("TOGGLE")
	assert.Equal(t, output, "on", "Transition should occur on toggle.")
}

func TestMachineCondition(t *testing.T) {
	machine := Machine{
		ID:      "machine-1",
		Initial: "on",
		States: StateMap{
			"on": MachineState{
				On: TransitionMap{
					"TOGGLE": MachineTransition{
						To: "off",
						Cond: func(curr, next string) bool {
							return curr == ""
						},
					},
				},
			},
			"off": MachineState{
				On: TransitionMap{
					"TOGGLE": MachineTransition{
						To: "on",
					},
				},
			},
		},
	}
	output := machine.Transition("TOGGLE")
	assert.Equal(t, output, "on", "Transition should not occur on toggle. due to condition")
}

func TestMultipleActions(t *testing.T) {
	times := 0
	actions := []func(string, string){
		func(c, n string) { times++ },
		func(c, n string) { times++ },
		func(c, n string) { times++ },
	}
	machine := Machine{
		ID:      "machine-1",
		Initial: "on",
		States: StateMap{
			"on": MachineState{
				On: TransitionMap{
					"TOGGLE": MachineTransition{
						To:      "off",
						Actions: actions,
					},
				},
			},
			"off": MachineState{
				On: TransitionMap{
					"TOGGLE": MachineTransition{
						To: "on",
					},
				},
			},
		},
	}
	machine.Transition("TOGGLE")
	assert.Equal(t, times, len(actions), "actions are not called")
}

func TestStateSubscribers(t *testing.T) {
	times := 0
	machine := Machine{
		ID:      "machine-1",
		Initial: "on",
		Subscribers: []func(string, string){
			func(c, n string) { times++ },
		},
		States: StateMap{
			"on": MachineState{
				On: TransitionMap{
					"TOGGLE": MachineTransition{
						To: "off",
					},
				},
			},
			"off": MachineState{
				On: TransitionMap{
					"TOGGLE": MachineTransition{
						To: "on",
					},
				},
			},
		},
	}
	for i := 0; i < 10; i++ {
		machine.Transition("TOGGLE")
	}
	assert.Equal(t, times, 10, "subscribers are not called")
}
