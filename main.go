package main

func main() {
	neurons := CSVToNeurons("./neurons.csv")
	synapseRecords := CSVToSynapseRecords("./connections.csv")

	// Connect neurons and synapses
	for _, synapseRecord := range synapseRecords {
		source := neurons[synapseRecord.source]
		destination := neurons[synapseRecord.destination]
		source.AddSynapse(destination, synapseRecord.weight)
	}

	// Kick off one go routine per neuron
	for _, neuron := range neurons {
		go neuron.Listen()
	}

	/* Tim's original work kicked off the system
	 * with 50 stimulations of the "food sense neurons"
	 */
	foodNeurons := []*Neuron{
		neurons["ADFR"],
		neurons["ASGR"],
		neurons["ASIR"],
		neurons["ASJR"],
		neurons["AWCR"],
		neurons["AWAR"],
		neurons["ADFL"],
		neurons["ASGL"],
		neurons["ASIL"],
		neurons["ASJL"],
		neurons["AWCL"],
		neurons["AWAL"],
	}

	for i := 0; i < 50; i++ {
		for _, neuron := range foodNeurons {
			neuron.dendrite <- ACTION_POTENTIAL_THRESHOLD + 1
		}
	}

	// Block indefinitely
	ch := make(chan bool)
	<-ch
}
