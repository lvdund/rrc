package ies

// SchedulingRequestConfig ::= CHOICE
const (
	SchedulingrequestconfigChoiceNothing = iota
	SchedulingrequestconfigChoiceRelease
	SchedulingrequestconfigChoiceSetup
)

type Schedulingrequestconfig struct {
	Choice  uint64
	Release *struct{}
	Setup   *SchedulingrequestconfigSetup
}
