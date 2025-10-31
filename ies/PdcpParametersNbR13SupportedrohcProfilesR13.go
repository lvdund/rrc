package ies

import "rrc/utils"

// PDCP-Parameters-NB-r13-supportedROHC-Profiles-r13 ::= SEQUENCE
type PdcpParametersNbR13SupportedrohcProfilesR13 struct {
	Profile0x0002 utils.BOOLEAN
	Profile0x0003 utils.BOOLEAN
	Profile0x0004 utils.BOOLEAN
	Profile0x0006 utils.BOOLEAN
	Profile0x0102 utils.BOOLEAN
	Profile0x0103 utils.BOOLEAN
	Profile0x0104 utils.BOOLEAN
}
