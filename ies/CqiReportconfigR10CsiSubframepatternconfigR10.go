package ies

// CQI-ReportConfig-r10-csi-SubframePatternConfig-r10 ::= CHOICE
const (
	CqiReportconfigR10CsiSubframepatternconfigR10ChoiceNothing = iota
	CqiReportconfigR10CsiSubframepatternconfigR10ChoiceRelease
	CqiReportconfigR10CsiSubframepatternconfigR10ChoiceSetup
)

type CqiReportconfigR10CsiSubframepatternconfigR10 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CqiReportconfigR10CsiSubframepatternconfigR10Setup
}
