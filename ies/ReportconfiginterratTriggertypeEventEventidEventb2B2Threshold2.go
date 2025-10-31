package ies

// ReportConfigInterRAT-triggerType-event-eventId-eventB2-b2-Threshold2 ::= CHOICE
const (
	ReportconfiginterratTriggertypeEventEventidEventb2B2Threshold2ChoiceNothing = iota
	ReportconfiginterratTriggertypeEventEventidEventb2B2Threshold2ChoiceB2Threshold2utra
	ReportconfiginterratTriggertypeEventEventidEventb2B2Threshold2ChoiceB2Threshold2geran
	ReportconfiginterratTriggertypeEventEventidEventb2B2Threshold2ChoiceB2Threshold2cdma2000
)

type ReportconfiginterratTriggertypeEventEventidEventb2B2Threshold2 struct {
	Choice               uint64
	B2Threshold2utra     *Thresholdutra
	B2Threshold2geran    *Thresholdgeran
	B2Threshold2cdma2000 *Thresholdcdma2000
}
