package ies

// SL-SDAP-Config-r16-sl-MappedQoS-Flows-r16 ::= CHOICE
const (
	SlSdapConfigR16SlMappedqosFlowsR16ChoiceNothing = iota
	SlSdapConfigR16SlMappedqosFlowsR16ChoiceSlMappedqosFlowslistR16
	SlSdapConfigR16SlMappedqosFlowsR16ChoiceSlMappedqosFlowslistdedicatedR16
)

type SlSdapConfigR16SlMappedqosFlowsR16 struct {
	Choice                           uint64
	SlMappedqosFlowslistR16          *[]SlQosProfileR16 `lb:1,ub:maxNrofSLQfisR16`
	SlMappedqosFlowslistdedicatedR16 *SlMappedqosFlowslistdedicatedR16
}
