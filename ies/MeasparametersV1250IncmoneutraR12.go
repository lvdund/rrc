package ies

import "rrc/utils"

// MeasParameters-v1250-incMonEUTRA-r12 ::= ENUMERATED
type MeasparametersV1250IncmoneutraR12 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1250IncmoneutraR12EnumeratedNothing = iota
	MeasparametersV1250IncmoneutraR12EnumeratedSupported
)
