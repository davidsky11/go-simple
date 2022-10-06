package statemachine

import (
	"errors"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParallelToggle(t *testing.T) {
	times := 0
	machineOne := Machine{
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
	machineTwo := Machine{
		ID:      "machine-2",
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

	parallel := ParallelMachine{
		Machines: Machines{
			"machine-1": &machineOne,
			"machine-2": &machineTwo,
		},
		Subscribers: ParallelSubscribers{
			func(c, n ParallelState) { times++ },
			func(c, n ParallelState) { times++ },
			func(c, n ParallelState) { times++ },
		},
	}

	next, err := parallel.Transition("machine-1.TOGGLE")
	assert.Equal(t, ParallelState{"machine-1": "off", "machine-2": "on"}, next, "Transition should occur on toggle.")
	assert.Equal(t, nil, err, "Error should not occur in correct transition")
	assert.Equal(t, 3, times, "Subscribers should be called on transition")

	next, err = parallel.Transition("machine-one")
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("event format doesn't match"), err)
		assert.Equal(t, next, ParallelState{"machine-1": "off", "machine-2": "on"}, "Transition should not occur on error.")
	} else {
		t.Error("error should occur when key format doesn't match")
	}

	next, err = parallel.Transition("machine-one.TOGGLE")
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("machine key doesn't match"), err)
		assert.Equal(t, next, ParallelState{"machine-1": "off", "machine-2": "on"}, "Transition should not occur on error.")
	} else {
		t.Error("error should occur when machine key doesn't exist")
	}

	assert.Equal(t, 3, times, "Subscribers should not be called on error on transition")
}
