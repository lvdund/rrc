package ies

import "rrc/utils"

// PeriodicalReportConfig ::= SEQUENCE
// Extensible
type Periodicalreportconfig struct {
	Rstype                       NrRsType
	Reportinterval               Reportinterval
	Reportamount                 PeriodicalreportconfigReportamount
	Reportquantitycell           Measreportquantity
	Maxreportcells               utils.INTEGER `lb:0,ub:maxCellReport`
	ReportquantityrsIndexes      *Measreportquantity
	MaxnrofrsIndexestoreport     *utils.INTEGER `lb:0,ub:maxNrofIndexesToReport`
	Includebeammeasurements      utils.BOOLEAN
	Useallowedcelllist           utils.BOOLEAN
	MeasrssiReportconfigR16      *MeasrssiReportconfigR16                     `ext`
	IncludecommonlocationinfoR16 *utils.BOOLEAN                               `ext`
	IncludebtMeasR16             *utils.Setuprelease[BtNamelistR16]           `ext`
	IncludewlanMeasR16           *utils.Setuprelease[WlanNamelistR16]         `ext`
	IncludesensorMeasR16         *utils.Setuprelease[SensorNamelistR16]       `ext`
	UlDelayvalueconfigR16        *utils.Setuprelease[UlDelayvalueconfigR16]   `ext`
	ReportaddneighmeasR16        *PeriodicalreportconfigReportaddneighmeasR16 `ext`
	UlExcessdelayconfigR17       *utils.Setuprelease[UlExcessdelayconfigR17]  `ext`
	CoarselocationrequestR17     *utils.BOOLEAN                               `ext`
	ReportquantityrelayR17       *SlMeasreportquantityR16                     `ext`
}
