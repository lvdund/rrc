package ies

import "rrc/utils"

// CSI-ReportFrameworkExt-r16 ::= SEQUENCE
type CsiReportframeworkextR16 struct {
	MaxnumberaperiodiccsiPerbwpForcsiReportextR16 utils.INTEGER `lb:0,ub:8`
}
