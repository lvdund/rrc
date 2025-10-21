package ies

import "rrc/utils"

// MeasResults ::= SEQUENCE
// Extensible
type Measresults struct {
	Measid               Measid
	Measresultpcell      interface{}
	Measresultneighcells *interface{}
}
