package ies

// SL-MappedQoS-FlowsListDedicated-r16 ::= SEQUENCE
type SlMappedqosFlowslistdedicatedR16 struct {
	SlMappedqosFlowstoaddlistR16     *[]SlQosFlowidentityR16 `lb:1,ub:maxNrofSLQfisR16`
	SlMappedqosFlowstoreleaselistR16 *[]SlQosFlowidentityR16 `lb:1,ub:maxNrofSLQfisR16`
}
