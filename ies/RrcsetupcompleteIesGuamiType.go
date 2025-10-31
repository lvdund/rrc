package ies

import "rrc/utils"

// RRCSetupComplete-IEs-guami-Type ::= ENUMERATED
type RrcsetupcompleteIesGuamiType struct {
	Value utils.ENUMERATED
}

const (
	RrcsetupcompleteIesGuamiTypeEnumeratedNothing = iota
	RrcsetupcompleteIesGuamiTypeEnumeratedNative
	RrcsetupcompleteIesGuamiTypeEnumeratedMapped
)
