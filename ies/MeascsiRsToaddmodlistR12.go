package ies

// MeasCSI-RS-ToAddModList-r12 ::= SEQUENCE OF MeasCSI-RS-Config-r12
// SIZE (1..maxCSI-RS-Meas-r12)
type MeascsiRsToaddmodlistR12 struct {
	Value []MeascsiRsConfigR12 `lb:1,ub:maxCSIRsMeasR12`
}
