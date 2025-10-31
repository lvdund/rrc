package ies

import "rrc/utils"

// MAC-Parameters-NB-v1530-sr-SPS-BSR-r15 ::= ENUMERATED
type MacParametersNbV1530SrSpsBsrR15 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersNbV1530SrSpsBsrR15EnumeratedNothing = iota
	MacParametersNbV1530SrSpsBsrR15EnumeratedSupported
)
