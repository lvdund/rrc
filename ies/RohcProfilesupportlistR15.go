package ies

import "rrc/utils"

// ROHC-ProfileSupportList-r15 ::= SEQUENCE
type RohcProfilesupportlistR15 struct {
	Profile0x0001R15 utils.BOOLEAN
	Profile0x0002R15 utils.BOOLEAN
	Profile0x0003R15 utils.BOOLEAN
	Profile0x0004R15 utils.BOOLEAN
	Profile0x0006R15 utils.BOOLEAN
	Profile0x0101R15 utils.BOOLEAN
	Profile0x0102R15 utils.BOOLEAN
	Profile0x0103R15 utils.BOOLEAN
	Profile0x0104R15 utils.BOOLEAN
}
