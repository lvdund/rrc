package ies

// CSI-RS-TriggeredList-r12 ::= SEQUENCE OF MeasCSI-RS-Id-r12
// SIZE (1..maxCSI-RS-Meas-r12)
type CsiRsTriggeredlistR12 struct {
	Value []MeascsiRsIdR12 `lb:1,ub:maxCSIRsMeasR12`
}
