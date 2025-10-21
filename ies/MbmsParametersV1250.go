package ies

import "rrc/utils"

// MBMS-Parameters-v1250 ::= SEQUENCE
type MbmsParametersV1250 struct {
	MbmsAsyncdcR12 *utils.ENUMERATED
}
