package ies

import "rrc/utils"

// DMRS-UplinkConfig-transformPrecodingDisabled-dmrs-Uplink-r16 ::= ENUMERATED
type DmrsUplinkconfigTransformprecodingdisabledDmrsUplinkR16 struct {
	Value utils.ENUMERATED
}

const (
	DmrsUplinkconfigTransformprecodingdisabledDmrsUplinkR16EnumeratedNothing = iota
	DmrsUplinkconfigTransformprecodingdisabledDmrsUplinkR16EnumeratedEnabled
)
