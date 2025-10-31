package ies

import "rrc/utils"

// PDCP-Config-NB-r13-headerCompression-r13-rohc ::= SEQUENCE
// Extensible
type PdcpConfigNbR13HeadercompressionR13Rohc struct {
	MaxcidR13   utils.INTEGER `lb:0,ub:16383`
	ProfilesR13 PdcpConfigNbR13HeadercompressionR13RohcProfilesR13
}
