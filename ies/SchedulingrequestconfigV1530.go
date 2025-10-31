package ies

// SchedulingRequestConfig-v1530 ::= CHOICE
const (
	SchedulingrequestconfigV1530ChoiceNothing = iota
	SchedulingrequestconfigV1530ChoiceRelease
	SchedulingrequestconfigV1530ChoiceSetup
)

type SchedulingrequestconfigV1530 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SchedulingrequestconfigV1530Setup
}
