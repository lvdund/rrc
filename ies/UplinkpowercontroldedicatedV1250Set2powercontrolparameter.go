package ies

// UplinkPowerControlDedicated-v1250-set2PowerControlParameter ::= CHOICE
const (
	UplinkpowercontroldedicatedV1250Set2powercontrolparameterChoiceNothing = iota
	UplinkpowercontroldedicatedV1250Set2powercontrolparameterChoiceRelease
	UplinkpowercontroldedicatedV1250Set2powercontrolparameterChoiceSetup
)

type UplinkpowercontroldedicatedV1250Set2powercontrolparameter struct {
	Choice  uint64
	Release *struct{}
	Setup   *UplinkpowercontroldedicatedV1250Set2powercontrolparameterSetup
}
