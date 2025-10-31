package ies

import "rrc/utils"

// RRCConnectionSetupComplete-v1020-IEs-rn-SubframeConfigReq-r10 ::= ENUMERATED
type RrcconnectionsetupcompleteV1020IesRnSubframeconfigreqR10 struct {
	Value utils.ENUMERATED
}

const (
	RrcconnectionsetupcompleteV1020IesRnSubframeconfigreqR10EnumeratedNothing = iota
	RrcconnectionsetupcompleteV1020IesRnSubframeconfigreqR10EnumeratedRequired
	RrcconnectionsetupcompleteV1020IesRnSubframeconfigreqR10EnumeratedNotrequired
)
