package ies

import "rrc/utils"

// MAC-ParametersCommon-lch-ToSCellRestriction ::= ENUMERATED
type MacParameterscommonLchToscellrestriction struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonLchToscellrestrictionEnumeratedNothing = iota
	MacParameterscommonLchToscellrestrictionEnumeratedSupported
)
