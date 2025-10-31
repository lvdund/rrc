package ies

// MobilityControlInfoV2X-r14 ::= SEQUENCE
type Mobilitycontrolinfov2xR14 struct {
	V2xCommtxpoolexceptionalR14 *SlCommresourcepoolv2xR14
	V2xCommrxpoolR14            *SlCommrxpoollistv2xR14
	V2xCommsyncconfigR14        *SlSyncconfiglistv2xR14
	CbrMobilitytxconfiglistR14  *SlCbrCommontxconfiglistR14
}
