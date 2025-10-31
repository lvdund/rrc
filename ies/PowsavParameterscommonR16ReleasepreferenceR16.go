package ies

import "rrc/utils"

// PowSav-ParametersCommon-r16-releasePreference-r16 ::= ENUMERATED
type PowsavParameterscommonR16ReleasepreferenceR16 struct {
	Value utils.ENUMERATED
}

const (
	PowsavParameterscommonR16ReleasepreferenceR16EnumeratedNothing = iota
	PowsavParameterscommonR16ReleasepreferenceR16EnumeratedSupported
)
