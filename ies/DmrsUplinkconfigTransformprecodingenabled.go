package ies

import "rrc/utils"

// DMRS-UplinkConfig-transformPrecodingEnabled ::= SEQUENCE
// Extensible
type DmrsUplinkconfigTransformprecodingenabled struct {
	NpuschIdentity                  *utils.INTEGER `lb:0,ub:1007`
	Sequencegrouphopping            *DmrsUplinkconfigTransformprecodingenabledSequencegrouphopping
	Sequencehopping                 *DmrsUplinkconfigTransformprecodingenabledSequencehopping
	DmrsUplinktransformprecodingR16 *utils.Setuprelease[DmrsUplinktransformprecodingR16]
}
