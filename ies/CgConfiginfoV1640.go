package ies

// CG-ConfigInfo-v1640-IEs ::= SEQUENCE
type CgConfiginfoV1640 struct {
	ServcellinfolistmcgNrR16    *ServcellinfolistmcgNrR16
	ServcellinfolistmcgEutraR16 *ServcellinfolistmcgEutraR16
	Noncriticalextension        *CgConfiginfoV1700
}
