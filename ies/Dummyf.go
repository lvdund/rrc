package ies

import "rrc/utils"

// DummyF ::= SEQUENCE
type Dummyf struct {
	MaxnumberperiodiccsiReportperbwp       utils.INTEGER `lb:0,ub:4`
	MaxnumberaperiodiccsiReportperbwp      utils.INTEGER `lb:0,ub:4`
	MaxnumbersemipersistentcsiReportperbwp utils.INTEGER `lb:0,ub:4`
	SimultaneouscsiReportsallcc            utils.INTEGER `lb:0,ub:32`
}
