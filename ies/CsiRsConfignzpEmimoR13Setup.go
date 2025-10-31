package ies

// CSI-RS-ConfigNZP-EMIMO-r13-setup ::= SEQUENCE
type CsiRsConfignzpEmimoR13Setup struct {
	NzpResourceconfiglistR13 []NzpResourceconfigR13 `lb:1,ub:2`
	CdmtypeR13               *CsiRsConfignzpEmimoR13SetupCdmtypeR13
}
