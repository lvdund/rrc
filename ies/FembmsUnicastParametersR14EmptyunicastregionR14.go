package ies

import "rrc/utils"

// FeMBMS-Unicast-Parameters-r14-emptyUnicastRegion-r14 ::= ENUMERATED
type FembmsUnicastParametersR14EmptyunicastregionR14 struct {
	Value utils.ENUMERATED
}

const (
	FembmsUnicastParametersR14EmptyunicastregionR14EnumeratedNothing = iota
	FembmsUnicastParametersR14EmptyunicastregionR14EnumeratedSupported
)
