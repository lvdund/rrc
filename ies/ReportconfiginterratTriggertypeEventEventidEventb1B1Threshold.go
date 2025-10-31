package ies

// ReportConfigInterRAT-triggerType-event-eventId-eventB1-b1-Threshold ::= CHOICE
const (
	ReportconfiginterratTriggertypeEventEventidEventb1B1ThresholdChoiceNothing = iota
	ReportconfiginterratTriggertypeEventEventidEventb1B1ThresholdChoiceB1Thresholdutra
	ReportconfiginterratTriggertypeEventEventidEventb1B1ThresholdChoiceB1Thresholdgeran
	ReportconfiginterratTriggertypeEventEventidEventb1B1ThresholdChoiceB1Thresholdcdma2000
)

type ReportconfiginterratTriggertypeEventEventidEventb1B1Threshold struct {
	Choice              uint64
	B1Thresholdutra     *Thresholdutra
	B1Thresholdgeran    *Thresholdgeran
	B1Thresholdcdma2000 *Thresholdcdma2000
}
