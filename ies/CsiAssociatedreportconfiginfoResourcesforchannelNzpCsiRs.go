package ies

import "rrc/utils"

// CSI-AssociatedReportConfigInfo-resourcesForChannel-nzp-CSI-RS ::= SEQUENCE
type CsiAssociatedreportconfiginfoResourcesforchannelNzpCsiRs struct {
	Resourceset utils.INTEGER `lb:0,ub:maxNrofNZPCsiRsResourcesetsperconfig`
	QclInfo     *[]TciStateid `lb:1,ub:maxNrofAPCsiRsResourcesperset`
}
