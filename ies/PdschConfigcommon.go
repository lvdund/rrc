package ies

import "rrc/utils"

// PDSCH-ConfigCommon ::= SEQUENCE
type PdschConfigcommon struct {
	Referencesignalpower utils.INTEGER
	PB                   utils.INTEGER
}
