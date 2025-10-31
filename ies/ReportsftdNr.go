package ies

import "rrc/utils"

// ReportSFTD-NR ::= SEQUENCE
// Extensible
type ReportsftdNr struct {
	ReportsftdMeas            utils.BOOLEAN
	Reportrsrp                utils.BOOLEAN
	ReportsftdNeighmeas       *utils.BOOLEAN `ext`
	DrxSftdNeighmeas          *utils.BOOLEAN `ext`
	Cellsforwhichtoreportsftd *[]Physcellid  `lb:1,ub:maxCellSFTD,ext`
}
