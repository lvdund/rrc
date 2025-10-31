package ies

import "rrc/utils"

// PDSCH-Config-pucch-sSCellDynDCI-1-2-r17 ::= ENUMERATED
type PdschConfigPucchSscelldyndci12R17 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigPucchSscelldyndci12R17EnumeratedNothing = iota
	PdschConfigPucchSscelldyndci12R17EnumeratedEnabled
)
