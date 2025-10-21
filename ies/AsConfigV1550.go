package ies

import "rrc/utils"

// AS-Config-v1550 ::= SEQUENCE
type AsConfigV1550 struct {
	TdmPatternconfigR15 *interface{}
	PMaxeutraR15        *PMax
}
