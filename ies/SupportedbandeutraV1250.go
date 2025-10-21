package ies

import "rrc/utils"

// SupportedBandEUTRA-v1250 ::= SEQUENCE
type SupportedbandeutraV1250 struct {
	Dl256qamR12 *utils.ENUMERATED
	Ul64qamR12  *utils.ENUMERATED
}
