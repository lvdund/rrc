package ies

// CQI-ReportConfig-v1250-csi-SubframePatternConfig-r12 ::= CHOICE
const (
	CqiReportconfigV1250CsiSubframepatternconfigR12ChoiceNothing = iota
	CqiReportconfigV1250CsiSubframepatternconfigR12ChoiceRelease
	CqiReportconfigV1250CsiSubframepatternconfigR12ChoiceSetup
)

type CqiReportconfigV1250CsiSubframepatternconfigR12 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CqiReportconfigV1250CsiSubframepatternconfigR12Setup
}
