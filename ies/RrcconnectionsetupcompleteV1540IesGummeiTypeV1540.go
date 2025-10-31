package ies

import "rrc/utils"

// RRCConnectionSetupComplete-v1540-IEs-gummei-Type-v1540 ::= ENUMERATED
type RrcconnectionsetupcompleteV1540IesGummeiTypeV1540 struct {
	Value utils.ENUMERATED
}

const (
	RrcconnectionsetupcompleteV1540IesGummeiTypeV1540EnumeratedNothing = iota
	RrcconnectionsetupcompleteV1540IesGummeiTypeV1540EnumeratedMappedfrom5g_V1540
)
