package ies

import "rrc/utils"

// CSI-AssociatedReportConfigInfo-resourcesForChannel2-r17-nzp-CSI-RS2-r17 ::= SEQUENCE
type CsiAssociatedreportconfiginfoResourcesforchannel2R17NzpCsiRs2R17 struct {
	Resourceset2R17 utils.INTEGER `lb:0,ub:maxNrofNZPCsiRsResourcesetsperconfig`
	QclInfo2R17     *[]TciStateid `lb:1,ub:maxNrofAPCsiRsResourcesperset`
}
