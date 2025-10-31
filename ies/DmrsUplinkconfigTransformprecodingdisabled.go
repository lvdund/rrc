package ies

import "rrc/utils"

// DMRS-UplinkConfig-transformPrecodingDisabled ::= SEQUENCE
// Extensible
type DmrsUplinkconfigTransformprecodingdisabled struct {
	Scramblingid0 *utils.INTEGER `lb:0,ub:65535`
	Scramblingid1 *utils.INTEGER `lb:0,ub:65535`
	DmrsUplinkR16 *DmrsUplinkconfigTransformprecodingdisabledDmrsUplinkR16
}
