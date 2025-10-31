package ies

import "rrc/utils"

// ReportConfigEUTRA-triggerType-periodical-purpose ::= ENUMERATED
type ReportconfigeutraTriggertypePeriodicalPurpose struct {
	Value utils.ENUMERATED
}

const (
	ReportconfigeutraTriggertypePeriodicalPurposeEnumeratedNothing = iota
	ReportconfigeutraTriggertypePeriodicalPurposeEnumeratedReportstrongestcells
	ReportconfigeutraTriggertypePeriodicalPurposeEnumeratedReportcgi
)
