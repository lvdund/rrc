package ies

import "rrc/utils"

// LWA-Parameters-v1440 ::= SEQUENCE
type LwaParametersV1440 struct {
	LwaRlcUmR14 *utils.ENUMERATED
}
