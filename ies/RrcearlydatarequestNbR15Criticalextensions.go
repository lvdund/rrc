package ies

// RRCEarlyDataRequest-NB-r15-criticalExtensions ::= CHOICE
const (
	RrcearlydatarequestNbR15CriticalextensionsChoiceNothing = iota
	RrcearlydatarequestNbR15CriticalextensionsChoiceRrcearlydatarequestR15
	RrcearlydatarequestNbR15CriticalextensionsChoiceLater
)

type RrcearlydatarequestNbR15Criticalextensions struct {
	Choice                 uint64
	RrcearlydatarequestR15 *RrcearlydatarequestNbR15
	Later                  *RrcearlydatarequestNbR15CriticalextensionsLater
}
