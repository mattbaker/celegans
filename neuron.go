package main

import (
	"log"
	"time"
)

var ACTION_POTENTIAL_THRESHOLD int = 30
var DELAY_BETWEEN_FIRINGS time.Duration = 10
var SIGNAL_BUFFER_SIZE = 2048

type Neuron struct {
	tag       string
	dendrite  chan int  // Inbound signals
	synapses  []Synapse // Outbound signals
	potential int
}

func (n *Neuron) Fire() {
	log.Println(n.tag, "Fired")
	n.potential = 0
	for _, synapse := range n.synapses {
		/* Non-blocking send, we drop signals
		 * if the system is overwhelmed
		 */
		select {
		case synapse.channel <- synapse.weight:
		default:
			log.Println("Signal from", n.tag, "dropped")
		}
		time.Sleep(DELAY_BETWEEN_FIRINGS * time.Millisecond)
	}
}

func (n *Neuron) HasReachedThreshold() bool {
	return n.potential > ACTION_POTENTIAL_THRESHOLD
}

func (n *Neuron) Listen() {
	for actionPotential := range n.dendrite {
		n.potential += actionPotential
		if n.HasReachedThreshold() {
			n.Fire()
		}
	}
}

func (n *Neuron) AddSynapse(destination *Neuron, weight int) {
	n.synapses = append(n.synapses, Synapse{destination.dendrite, weight})
}

func NewNeuron(tag string) Neuron {
	return Neuron{tag, make(chan int, SIGNAL_BUFFER_SIZE), []Synapse{}, 0}
}
