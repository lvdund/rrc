package ies

// CSI-RS-ConfigNZP-EMIMO-v1430 ::= SEQUENCE
type CsiRsConfignzpEmimoV1430 struct {
	NzpResourceconfiglistextR14 []NzpResourceconfigR13 `lb:0,ub:4`
	CdmtypeV1430                *CsiRsConfignzpEmimoV1430CdmtypeV1430
}
