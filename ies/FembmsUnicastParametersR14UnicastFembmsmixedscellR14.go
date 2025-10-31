package ies

import "rrc/utils"

// FeMBMS-Unicast-Parameters-r14-unicast-fembmsMixedSCell-r14 ::= ENUMERATED
type FembmsUnicastParametersR14UnicastFembmsmixedscellR14 struct {
	Value utils.ENUMERATED
}

const (
	FembmsUnicastParametersR14UnicastFembmsmixedscellR14EnumeratedNothing = iota
	FembmsUnicastParametersR14UnicastFembmsmixedscellR14EnumeratedSupported
)
