package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type SynapseRecord struct {
	source      string
	destination string
	weight      int
}

func CSVToRecords(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error on open", err)
	}

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("Error on CSV parse", err)
	}
	return records
}

func CSVToNeurons(filename string) map[string]*Neuron {
	records := CSVToRecords(filename)
	neurons := make(map[string]*Neuron)

	for _, record := range records {
		neuron := NewNeuron(record[0])
		neurons[record[0]] = &neuron
	}

	return neurons
}

func CSVToSynapseRecords(filename string) []SynapseRecord {
	records := CSVToRecords(filename)
	synapseRecords := make([]SynapseRecord, len(records))

	for i, record := range records {
		weight, err := strconv.Atoi(record[2])
		if err != nil {
			log.Fatal("Couldn't parse weight", err)
		}
		synapseRecords[i] = SynapseRecord{record[0], record[1], weight}
	}

	return synapseRecords
}
