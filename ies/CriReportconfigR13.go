package ies

// CRI-ReportConfig-r13 ::= CHOICE
const (
	CriReportconfigR13ChoiceNothing = iota
	CriReportconfigR13ChoiceRelease
	CriReportconfigR13ChoiceSetup
)

type CriReportconfigR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CriReportconfigR13Setup
}
