package ies

import "rrc/utils"

// MAC-Parameters-r12 ::= SEQUENCE
type MacParametersR12 struct {
	LogicalchannelsrProhibittimerR12 *utils.ENUMERATED
	LongdrxCommandR12                *utils.ENUMERATED
}
