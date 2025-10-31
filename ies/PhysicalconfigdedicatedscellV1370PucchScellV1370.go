package ies

// PhysicalConfigDedicatedSCell-v1370-pucch-SCell-v1370 ::= CHOICE
const (
	PhysicalconfigdedicatedscellV1370PucchScellV1370ChoiceNothing = iota
	PhysicalconfigdedicatedscellV1370PucchScellV1370ChoiceRelease
	PhysicalconfigdedicatedscellV1370PucchScellV1370ChoiceSetup
)

type PhysicalconfigdedicatedscellV1370PucchScellV1370 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PhysicalconfigdedicatedscellV1370PucchScellV1370Setup
}
