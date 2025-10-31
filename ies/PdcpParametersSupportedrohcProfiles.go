package ies

import "rrc/utils"

// PDCP-Parameters-supportedROHC-Profiles ::= SEQUENCE
type PdcpParametersSupportedrohcProfiles struct {
	Profile0x0000 utils.BOOLEAN
	Profile0x0001 utils.BOOLEAN
	Profile0x0002 utils.BOOLEAN
	Profile0x0003 utils.BOOLEAN
	Profile0x0004 utils.BOOLEAN
	Profile0x0006 utils.BOOLEAN
	Profile0x0101 utils.BOOLEAN
	Profile0x0102 utils.BOOLEAN
	Profile0x0103 utils.BOOLEAN
	Profile0x0104 utils.BOOLEAN
}
