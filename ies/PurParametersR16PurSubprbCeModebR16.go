package ies

import "rrc/utils"

// PUR-Parameters-r16-pur-SubPRB-CE-ModeB-r16 ::= ENUMERATED
type PurParametersR16PurSubprbCeModebR16 struct {
	Value utils.ENUMERATED
}

const (
	PurParametersR16PurSubprbCeModebR16EnumeratedNothing = iota
	PurParametersR16PurSubprbCeModebR16EnumeratedSupported
)
