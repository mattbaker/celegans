package main

import (
	"log"
	"time"
)

var ACTION_POTENTIAL_THRESHOLD int = 15

type Neuron struct {
	tag       string
	dendrite  chan int  //inbound
	synapses  []Synapse //outbound
	potential int
}

func (n *Neuron) Fire() {
	log.Println(n.tag, "Fired")
	n.potential = 0
	for _, synapse := range n.synapses {
		synapse.channel <- synapse.weight
		time.Sleep(1 * time.Millisecond)
	}
}

func (n *Neuron) HasReachedThreshold() bool {
	return n.potential > ACTION_POTENTIAL_THRESHOLD
}

func (n *Neuron) Listen() {
	for actionPotential := range n.dendrite {
		log.Println("Received on", n.tag)
		n.potential += actionPotential
		if n.HasReachedThreshold() {
			log.Println("Firing on", n.tag)
			n.Fire()
		}
	}
}

func (n *Neuron) AddSynapse(destination *Neuron, weight int) {
	n.synapses = append(n.synapses, Synapse{destination.dendrite, weight})
}

func NewNeuron(tag string) Neuron {
	return Neuron{tag, make(chan int), []Synapse{}, 0}
}
