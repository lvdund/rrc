package ies

// CG-Config-v1590-IEs ::= SEQUENCE
type CgConfigV1590 struct {
	ScellfrequenciessnNr    *[]ArfcnValuenr    `lb:1,ub:maxNrofServingCells1`
	ScellfrequenciessnEutra *[]ArfcnValueeutra `lb:1,ub:maxNrofServingCells1`
	Noncriticalextension    *CgConfigV1610
}
