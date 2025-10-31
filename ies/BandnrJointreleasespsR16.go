package ies

import "rrc/utils"

// BandNR-jointReleaseSPS-r16 ::= ENUMERATED
type BandnrJointreleasespsR16 struct {
	Value utils.ENUMERATED
}

const (
	BandnrJointreleasespsR16EnumeratedNothing = iota
	BandnrJointreleasespsR16EnumeratedSupported
)
