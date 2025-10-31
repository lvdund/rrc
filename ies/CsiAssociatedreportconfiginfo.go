package ies

import "rrc/utils"

// CSI-AssociatedReportConfigInfo ::= SEQUENCE
// Extensible
type CsiAssociatedreportconfiginfo struct {
	Reportconfigid                   CsiReportconfigid
	Resourcesforchannel              *CsiAssociatedreportconfiginfoResourcesforchannel
	CsiImResourcesforinterference    *utils.INTEGER                                        `lb:0,ub:maxNrofCSIImResourcesetsperconfig`
	NzpCsiRsResourcesforinterference *utils.INTEGER                                        `lb:0,ub:maxNrofNZPCsiRsResourcesetsperconfig`
	Resourcesforchannel2R17          *CsiAssociatedreportconfiginfoResourcesforchannel2R17 `ext`
	CsiSsbResourcesetext             *utils.INTEGER                                        `lb:0,ub:maxNrofCSISsbResourcesetsperconfigext,ext`
}
