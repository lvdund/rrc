package ies

// FreqPriorityNR ::= SEQUENCE
type Freqprioritynr struct {
	Carrierfreq                ArfcnValuenr
	Cellreselectionpriority    Cellreselectionpriority
	Cellreselectionsubpriority *Cellreselectionsubpriority
}
