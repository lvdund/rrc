package ies

// NPDSCH-MultiTB-Config-NB-r16 ::= SEQUENCE
type NpdschMultitbConfigNbR16 struct {
	MultitbConfigR16   NpdschMultitbConfigNbR16MultitbConfigR16
	HarqAckbundlingR16 *bool
}
