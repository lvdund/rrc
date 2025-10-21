package ies

import "rrc/utils"

// LAA-Parameters-v1530 ::= SEQUENCE
type LaaParametersV1530 struct {
	AulR15           *utils.ENUMERATED
	LaaPuschMode1R15 *utils.ENUMERATED
	LaaPuschMode2R15 *utils.ENUMERATED
	LaaPuschMode3R15 *utils.ENUMERATED
}
