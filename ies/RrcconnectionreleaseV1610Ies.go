package ies

import "rrc/utils"

// RRCConnectionRelease-v1610-IEs ::= SEQUENCE
type RrcconnectionreleaseV1610Ies struct {
	FulliRntiR16             *IRntiR15
	ShortiRntiR16            *ShortiRntiR15
	PurConfigR16             *Setuprelease
	RrcInactiveconfigV1610   *RrcInactiveconfigV1610
	ReleaseidlemeasconfigR16 *utils.ENUMERATED
	AltfreqprioritiesR16     *utils.ENUMERATED
	T323R16                  *utils.ENUMERATED
	Noncriticalextension     *RrcconnectionreleaseV1650Ies
}
