package ies

import "rrc/utils"

// PDCP-Config-NB-r13-headerCompression-r13-rohc-profiles-r13 ::= SEQUENCE
type PdcpConfigNbR13HeadercompressionR13RohcProfilesR13 struct {
	Profile0x0002 utils.BOOLEAN
	Profile0x0003 utils.BOOLEAN
	Profile0x0004 utils.BOOLEAN
	Profile0x0006 utils.BOOLEAN
	Profile0x0102 utils.BOOLEAN
	Profile0x0103 utils.BOOLEAN
	Profile0x0104 utils.BOOLEAN
}
