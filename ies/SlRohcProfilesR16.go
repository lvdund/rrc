package ies

import "rrc/utils"

// SL-RoHC-Profiles-r16 ::= SEQUENCE
type SlRohcProfilesR16 struct {
	Profile0x0001R16 utils.BOOLEAN
	Profile0x0002R16 utils.BOOLEAN
	Profile0x0003R16 utils.BOOLEAN
	Profile0x0004R16 utils.BOOLEAN
	Profile0x0006R16 utils.BOOLEAN
	Profile0x0101R16 utils.BOOLEAN
	Profile0x0102R16 utils.BOOLEAN
	Profile0x0103R16 utils.BOOLEAN
	Profile0x0104R16 utils.BOOLEAN
}
