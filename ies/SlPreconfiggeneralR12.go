package ies

import "rrc/utils"

// SL-PreconfigGeneral-r12 ::= SEQUENCE
// Extensible
type SlPreconfiggeneralR12 struct {
	RohcProfilesR12               interface{}
	CarrierfreqR12                ArfcnValueeutraR9
	MaxtxpowerR12                 PMax
	AdditionalspectrumemissionR12 Additionalspectrumemission
	SlBandwidthR12                utils.ENUMERATED
	TddConfigslR12                TddConfigslR12
	ReservedR12                   utils.BITSTRING
}
