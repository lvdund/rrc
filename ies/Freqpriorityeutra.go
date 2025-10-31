package ies

// FreqPriorityEUTRA ::= SEQUENCE
type Freqpriorityeutra struct {
	Carrierfreq             ArfcnValueeutra
	Cellreselectionpriority Cellreselectionpriority
}
