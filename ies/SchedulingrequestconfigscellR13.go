package ies

// SchedulingRequestConfigSCell-r13 ::= CHOICE
const (
	SchedulingrequestconfigscellR13ChoiceNothing = iota
	SchedulingrequestconfigscellR13ChoiceRelease
	SchedulingrequestconfigscellR13ChoiceSetup
)

type SchedulingrequestconfigscellR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SchedulingrequestconfigscellR13Setup
}
