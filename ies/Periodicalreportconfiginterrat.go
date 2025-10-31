package ies

import "rrc/utils"

// PeriodicalReportConfigInterRAT ::= SEQUENCE
// Extensible
type Periodicalreportconfiginterrat struct {
	Reportinterval               Reportinterval
	Reportamount                 PeriodicalreportconfiginterratReportamount
	Reportquantity               Measreportquantity
	Maxreportcells               utils.INTEGER                          `lb:0,ub:maxCellReport`
	ReportquantityutraFddR16     *MeasreportquantityutraFddR16          `ext`
	IncludecommonlocationinfoR16 *utils.BOOLEAN                         `ext`
	IncludebtMeasR16             *utils.Setuprelease[BtNamelistR16]     `ext`
	IncludewlanMeasR16           *utils.Setuprelease[WlanNamelistR16]   `ext`
	IncludesensorMeasR16         *utils.Setuprelease[SensorNamelistR16] `ext`
	ReportquantityrelayR17       *SlMeasreportquantityR16               `ext`
}
