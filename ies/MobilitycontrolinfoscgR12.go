package ies

import "rrc/utils"

// MobilityControlInfoSCG-r12 ::= SEQUENCE
// Extensible
type MobilitycontrolinfoscgR12 struct {
	T307R12                  utils.ENUMERATED
	UeIdentityscgR12         *CRnti
	RachConfigdedicatedR12   *RachConfigdedicated
	CipheringalgorithmscgR12 *CipheringalgorithmR12
}
