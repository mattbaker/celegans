package main

func main() {

	neurons := CSVToNeurons("./neurons.small.csv")
	synapseRecords := CSVToSynapseRecords("./connections.small.csv")

	for _, synapseRecord := range synapseRecords {
		source := neurons[synapseRecord.source]
		destination := neurons[synapseRecord.destination]
		source.AddSynapse(destination, synapseRecord.weight)
	}

	for _, neuron := range neurons {
		go neuron.Listen()
	}

	neurons["A"].Fire()

	for {
	}
}
