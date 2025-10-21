package ies

import "rrc/utils"

// MeasRSS-DedicatedConfig-r16 ::= SEQUENCE
type MeasrssDedicatedconfigR16 struct {
	RssConfigcarrierinfoR16 *RssConfigcarrierinfoR16
	CellstoaddmodlistV1610  *CellstoaddmodlistV1610
}
