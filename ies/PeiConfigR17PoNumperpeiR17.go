package ies

import "rrc/utils"

// PEI-Config-r17-po-NumPerPEI-r17 ::= ENUMERATED
type PeiConfigR17PoNumperpeiR17 struct {
	Value utils.ENUMERATED
}

const (
	PeiConfigR17PoNumperpeiR17EnumeratedNothing = iota
	PeiConfigR17PoNumperpeiR17EnumeratedPo1
	PeiConfigR17PoNumperpeiR17EnumeratedPo2
	PeiConfigR17PoNumperpeiR17EnumeratedPo4
	PeiConfigR17PoNumperpeiR17EnumeratedPo8
)
