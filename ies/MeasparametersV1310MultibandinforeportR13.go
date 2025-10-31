package ies

import "rrc/utils"

// MeasParameters-v1310-multiBandInfoReport-r13 ::= ENUMERATED
type MeasparametersV1310MultibandinforeportR13 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1310MultibandinforeportR13EnumeratedNothing = iota
	MeasparametersV1310MultibandinforeportR13EnumeratedSupported
)
