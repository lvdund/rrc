package ies

import "rrc/utils"

// BandNR-jointReleaseConfiguredGrantType2-r16 ::= ENUMERATED
type BandnrJointreleaseconfiguredgranttype2R16 struct {
	Value utils.ENUMERATED
}

const (
	BandnrJointreleaseconfiguredgranttype2R16EnumeratedNothing = iota
	BandnrJointreleaseconfiguredgranttype2R16EnumeratedSupported
)
