package ies

import "rrc/utils"

// ServingCellConfig-moreThanOneNackOnlyMode-r17 ::= ENUMERATED
type ServingcellconfigMorethanonenackonlymodeR17 struct {
	Value utils.ENUMERATED
}

const (
	ServingcellconfigMorethanonenackonlymodeR17EnumeratedNothing = iota
	ServingcellconfigMorethanonenackonlymodeR17EnumeratedMode2
)
