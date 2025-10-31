package ies

import "rrc/utils"

// SPUCCH-Config-v1550-setup-twoAntennaPortActivatedSPUCCH-Format3-v1550 ::= SEQUENCE
type SpucchConfigV1550SetupTwoantennaportactivatedspucchFormat3V1550 struct {
	N3spucchAnListV1550 []utils.INTEGER `lb:1,ub:4`
}
