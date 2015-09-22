package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
)

func silenceLog() {
	log.SetOutput(ioutil.Discard)
}

func TestFireSpikesSynapses(t *testing.T) {
	silenceLog()
	na := NewNeuron("A")
	nb := NewNeuron("B")
	nc := NewNeuron("C")

	na.AddSynapse(&nb, 10)
	na.AddSynapse(&nc, 5)

	go na.Fire()

	assert.Equal(t, 10, <-nb.dendrite)
	assert.Equal(t, 5, <-nc.dendrite)
}

func TestFireZerosPotential(t *testing.T) {
	silenceLog()
	na := NewNeuron("A")
	na.potential = 10
	na.Fire()
	assert.Zero(t, na.potential)
}

func TestHasReachedThreshold(t *testing.T) {
	silenceLog()
	na := NewNeuron("A")
	na.potential = 0
	assert.False(t, na.HasReachedThreshold())

	na.potential = ACTION_POTENTIAL_THRESHOLD + 1
	assert.True(t, na.HasReachedThreshold())
}

func TestListen(t *testing.T) {
	silenceLog()
	na := NewNeuron("A")
	nb := NewNeuron("B")

	na.AddSynapse(&nb, 11)

	go na.Listen()

	assert.Zero(t, nb.potential)

	na.dendrite <- ACTION_POTENTIAL_THRESHOLD
	assert.Zero(t, nb.potential)

	na.dendrite <- 1

	assert.Equal(t, 11, <-nb.dendrite)
}
