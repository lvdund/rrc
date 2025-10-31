package ies

// PhysicalConfigDedicatedSCell-v13c0-pucch-SCell-v13c0 ::= CHOICE
const (
	PhysicalconfigdedicatedscellV13c0PucchScellV13c0ChoiceNothing = iota
	PhysicalconfigdedicatedscellV13c0PucchScellV13c0ChoiceRelease
	PhysicalconfigdedicatedscellV13c0PucchScellV13c0ChoiceSetup
)

type PhysicalconfigdedicatedscellV13c0PucchScellV13c0 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PhysicalconfigdedicatedscellV13c0PucchScellV13c0Setup
}
