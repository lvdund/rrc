package ies

import "rrc/utils"

// CSI-ReportFramework ::= SEQUENCE
type CsiReportframework struct {
	MaxnumberperiodiccsiPerbwpForcsiReport        utils.INTEGER `lb:0,ub:4`
	MaxnumberaperiodiccsiPerbwpForcsiReport       utils.INTEGER `lb:0,ub:4`
	MaxnumbersemipersistentcsiPerbwpForcsiReport  utils.INTEGER `lb:0,ub:4`
	MaxnumberperiodiccsiPerbwpForbeamreport       utils.INTEGER `lb:0,ub:4`
	MaxnumberaperiodiccsiPerbwpForbeamreport      utils.INTEGER `lb:0,ub:4`
	MaxnumberaperiodiccsiTriggeringstatepercc     CsiReportframeworkMaxnumberaperiodiccsiTriggeringstatepercc
	MaxnumbersemipersistentcsiPerbwpForbeamreport utils.INTEGER `lb:0,ub:4`
	SimultaneouscsiReportspercc                   utils.INTEGER `lb:0,ub:8`
}
