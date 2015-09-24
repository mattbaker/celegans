# celegans

An implementation of the celegans connectome in Go.

Based on @interintel's [work](http://interintelligence.org/archCE.htm) ([Python code](https://github.com/Connectome/GoPiGo))

One go routine is kicked off per Neuron, which listens indefinitely for incoming "signals" from other Neurons.

Neurons are modeled as structs. Synapses are modeled as Go channels between each Neuron's go routine.

Neurons have a potential value that starts at 0. Potential is increased by the signals (ints) coming in to the Neuron via synapses. If the potential exceeds a threshold, the Neuron "fires" — causing signals to be sent from the firing Neuron to other Neurons.

This is a work in progress.
