package ies

import "rrc/utils"

// CSI-MeasConfig ::= SEQUENCE
// Extensible
type CsiMeasconfig struct {
	NzpCsiRsResourcetoaddmodlist            *[]NzpCsiRsResource      `lb:1,ub:maxNrofNZPCsiRsResources`
	NzpCsiRsResourcetoreleaselist           *[]NzpCsiRsResourceid    `lb:1,ub:maxNrofNZPCsiRsResources`
	NzpCsiRsResourcesettoaddmodlist         *[]NzpCsiRsResourceset   `lb:1,ub:maxNrofNZPCsiRsResourcesets`
	NzpCsiRsResourcesettoreleaselist        *[]NzpCsiRsResourcesetid `lb:1,ub:maxNrofNZPCsiRsResourcesets`
	CsiImResourcetoaddmodlist               *[]CsiImResource         `lb:1,ub:maxNrofCSIImResources`
	CsiImResourcetoreleaselist              *[]CsiImResourceid       `lb:1,ub:maxNrofCSIImResources`
	CsiImResourcesettoaddmodlist            *[]CsiImResourceset      `lb:1,ub:maxNrofCSIImResourcesets`
	CsiImResourcesettoreleaselist           *[]CsiImResourcesetid    `lb:1,ub:maxNrofCSIImResourcesets`
	CsiSsbResourcesettoaddmodlist           *[]CsiSsbResourceset     `lb:1,ub:maxNrofCSISsbResourcesets`
	CsiSsbResourcesettoreleaselist          *[]CsiSsbResourcesetid   `lb:1,ub:maxNrofCSISsbResourcesets`
	CsiResourceconfigtoaddmodlist           *[]CsiResourceconfig     `lb:1,ub:maxNrofCSIResourceconfigurations`
	CsiResourceconfigtoreleaselist          *[]CsiResourceconfigid   `lb:1,ub:maxNrofCSIResourceconfigurations`
	CsiReportconfigtoaddmodlist             *[]CsiReportconfig       `lb:1,ub:maxNrofCSIReportconfigurations`
	CsiReportconfigtoreleaselist            *[]CsiReportconfigid     `lb:1,ub:maxNrofCSIReportconfigurations`
	Reporttriggersize                       *utils.INTEGER           `lb:0,ub:6`
	Aperiodictriggerstatelist               *utils.Setuprelease[CsiAperiodictriggerstatelist]
	SemipersistentonpuschTriggerstatelist   *utils.Setuprelease[CsiSemipersistentonpuschTriggerstatelist]
	Reporttriggersizedci02R16               *utils.INTEGER                  `lb:0,ub:6,ext`
	ScellactivationrsConfigtoaddmodlistR17  *[]ScellactivationrsConfigR17   `lb:1,ub:maxNrofSCellActRSR17,ext`
	ScellactivationrsConfigtoreleaselistR17 *[]ScellactivationrsConfigidR17 `lb:1,ub:maxNrofSCellActRSR17,ext`
}
