package ies

import "rrc/utils"

// CellChangeOrder ::= SEQUENCE
// Extensible
type Cellchangeorder struct {
	T304          utils.ENUMERATED
	TargetratType *interface{}
}
