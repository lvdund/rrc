package ies

import "rrc/utils"

// SupportedBandEUTRA-v1320 ::= SEQUENCE
type SupportedbandeutraV1320 struct {
	IntrafreqCeNeedforgapsR13 *utils.ENUMERATED
	UePowerclassNR13          *utils.ENUMERATED
}
