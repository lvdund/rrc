package ies

// MobilityControlInfoSCG-r12 ::= SEQUENCE
// Extensible
type MobilitycontrolinfoscgR12 struct {
	T307R12                  MobilitycontrolinfoscgR12T307R12
	UeIdentityscgR12         *CRnti
	RachConfigdedicatedR12   *RachConfigdedicated
	CipheringalgorithmscgR12 *CipheringalgorithmR12
}
