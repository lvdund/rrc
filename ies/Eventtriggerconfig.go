package ies

import "rrc/utils"

// EventTriggerConfig ::= SEQUENCE
// Extensible
type Eventtriggerconfig struct {
	Eventid                      EventtriggerconfigEventid `ext`
	Rstype                       NrRsType
	Reportinterval               Reportinterval
	Reportamount                 EventtriggerconfigReportamount
	Reportquantitycell           Measreportquantity
	Maxreportcells               utils.INTEGER `lb:0,ub:maxCellReport`
	ReportquantityrsIndexes      *Measreportquantity
	MaxnrofrsIndexestoreport     *utils.INTEGER `lb:0,ub:maxNrofIndexesToReport`
	Includebeammeasurements      utils.BOOLEAN
	Reportaddneighmeas           *EventtriggerconfigReportaddneighmeas
	MeasrssiReportconfigR16      *MeasrssiReportconfigR16               `ext`
	Uset312R16                   *utils.BOOLEAN                         `ext`
	IncludecommonlocationinfoR16 *utils.BOOLEAN                         `ext`
	IncludebtMeasR16             *utils.Setuprelease[BtNamelistR16]     `ext`
	IncludewlanMeasR16           *utils.Setuprelease[WlanNamelistR16]   `ext`
	IncludesensorMeasR16         *utils.Setuprelease[SensorNamelistR16] `ext`
	CoarselocationrequestR17     *utils.BOOLEAN                         `ext`
	ReportquantityrelayR17       *SlMeasreportquantityR16               `ext`
}
