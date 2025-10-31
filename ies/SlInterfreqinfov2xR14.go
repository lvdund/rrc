package ies

// SL-InterFreqInfoV2X-r14 ::= SEQUENCE
// Extensible
type SlInterfreqinfov2xR14 struct {
	PlmnIdentitylistR14   *PlmnIdentitylist
	V2xCommcarrierfreqR14 ArfcnValueeutraR9
	SlMaxtxpowerR14       *PMax
	SlBandwidthR14        *SlInterfreqinfov2xR14SlBandwidthR14
	V2xSchedulingpoolR14  *SlCommresourcepoolv2xR14
	V2xUeConfiglistR14    *SlV2xUeConfiglistR14
}
