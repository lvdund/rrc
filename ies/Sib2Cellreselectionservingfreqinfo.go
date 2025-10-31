package ies

// SIB2-cellReselectionServingFreqInfo ::= SEQUENCE
// Extensible
type Sib2Cellreselectionservingfreqinfo struct {
	SNonintrasearchp           *Reselectionthreshold
	SNonintrasearchq           *Reselectionthresholdq
	Threshservinglowp          Reselectionthreshold
	Threshservinglowq          *Reselectionthresholdq
	Cellreselectionpriority    Cellreselectionpriority
	Cellreselectionsubpriority *Cellreselectionsubpriority
}
