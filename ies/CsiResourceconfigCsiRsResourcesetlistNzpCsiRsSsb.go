package ies

// CSI-ResourceConfig-csi-RS-ResourceSetList-nzp-CSI-RS-SSB ::= SEQUENCE
type CsiResourceconfigCsiRsResourcesetlistNzpCsiRsSsb struct {
	NzpCsiRsResourcesetlist *[]NzpCsiRsResourcesetid `lb:1,ub:maxNrofNZPCsiRsResourcesetsperconfig`
	CsiSsbResourcesetlist   *[]CsiSsbResourcesetid   `lb:1,ub:maxNrofCSISsbResourcesetsperconfig`
}
