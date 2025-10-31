package ies

import "rrc/utils"

// CSI-AssociatedReportConfigInfo-resourcesForChannel ::= CHOICE
const (
	CsiAssociatedreportconfiginfoResourcesforchannelChoiceNothing = iota
	CsiAssociatedreportconfiginfoResourcesforchannelChoiceNzpCsiRs
	CsiAssociatedreportconfiginfoResourcesforchannelChoiceCsiSsbResourceset
)

type CsiAssociatedreportconfiginfoResourcesforchannel struct {
	Choice            uint64
	NzpCsiRs          *CsiAssociatedreportconfiginfoResourcesforchannelNzpCsiRs
	CsiSsbResourceset *utils.INTEGER `lb:1,ub:maxNrofCSISsbResourcesetsperconfig`
}
