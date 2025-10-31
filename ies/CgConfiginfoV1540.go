package ies

// CG-ConfigInfo-v1540-IEs ::= SEQUENCE
type CgConfiginfoV1540 struct {
	PhInfomcg            *PhTypelistmcg
	Measresultreportcgi  *CgConfiginfoV1540IesMeasresultreportcgi
	Noncriticalextension *CgConfiginfoV1560
}
