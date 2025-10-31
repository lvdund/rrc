package ies

import "rrc/utils"

// CSI-AssociatedReportConfigInfo-resourcesForChannel2-r17 ::= CHOICE
const (
	CsiAssociatedreportconfiginfoResourcesforchannel2R17ChoiceNothing = iota
	CsiAssociatedreportconfiginfoResourcesforchannel2R17ChoiceNzpCsiRs2R17
	CsiAssociatedreportconfiginfoResourcesforchannel2R17ChoiceCsiSsbResourceset2R17
)

type CsiAssociatedreportconfiginfoResourcesforchannel2R17 struct {
	Choice                uint64
	NzpCsiRs2R17          *CsiAssociatedreportconfiginfoResourcesforchannel2R17NzpCsiRs2R17
	CsiSsbResourceset2R17 *utils.INTEGER `lb:1,ub:maxNrofCSISsbResourcesetsperconfigext`
}
