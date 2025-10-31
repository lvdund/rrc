package ies

// RRCConnectionRelease-NB-criticalExtensions-c1 ::= CHOICE
const (
	RrcconnectionreleaseNbCriticalextensionsC1ChoiceNothing = iota
	RrcconnectionreleaseNbCriticalextensionsC1ChoiceRrcconnectionreleaseR13
	RrcconnectionreleaseNbCriticalextensionsC1ChoiceSpare1
)

type RrcconnectionreleaseNbCriticalextensionsC1 struct {
	Choice                  uint64
	RrcconnectionreleaseR13 *RrcconnectionreleaseNbR13
	Spare1                  *struct{}
}
