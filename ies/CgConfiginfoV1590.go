package ies

// CG-ConfigInfo-v1590-IEs ::= SEQUENCE
type CgConfiginfoV1590 struct {
	ServfrequenciesmnNr  *[]ArfcnValuenr `lb:1,ub:maxNrofServingCells1`
	Noncriticalextension *CgConfiginfoV1610
}
