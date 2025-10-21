package ies

import "rrc/utils"

// FeMBMS-Unicast-Parameters-r14 ::= SEQUENCE
type FembmsUnicastParametersR14 struct {
	UnicastFembmsmixedscellR14 *utils.ENUMERATED
	EmptyunicastregionR14      *utils.ENUMERATED
}
