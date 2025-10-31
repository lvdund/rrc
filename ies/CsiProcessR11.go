package ies

import "rrc/utils"

// CSI-Process-r11 ::= SEQUENCE
// Extensible
type CsiProcessR11 struct {
	CsiProcessidR11            CsiProcessidR11
	CsiRsConfignzpidR11        CsiRsConfignzpidR11
	CsiImConfigidR11           CsiImConfigidR11
	PCAndcbsrlistR11           PCAndcbsrPairR13a
	CqiReportbothprocR11       *CqiReportbothprocR11
	CqiReportperiodicprocidR11 *utils.INTEGER `lb:0,ub:maxCQIProcextR11`
	CqiReportaperiodicprocR11  *CqiReportaperiodicprocR11
}
