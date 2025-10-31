package ies

import "rrc/utils"

// MeasParameters-v1250-incMonUTRA-r12 ::= ENUMERATED
type MeasparametersV1250IncmonutraR12 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1250IncmonutraR12EnumeratedNothing = iota
	MeasparametersV1250IncmonutraR12EnumeratedSupported
)
