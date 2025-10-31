package ies

import "rrc/utils"

// SidelinkParametersNR-r16-p0-OLPC-Sidelink-r17 ::= ENUMERATED
type SidelinkparametersnrR16P0OlpcSidelinkR17 struct {
	Value utils.ENUMERATED
}

const (
	SidelinkparametersnrR16P0OlpcSidelinkR17EnumeratedNothing = iota
	SidelinkparametersnrR16P0OlpcSidelinkR17EnumeratedSupported
)
