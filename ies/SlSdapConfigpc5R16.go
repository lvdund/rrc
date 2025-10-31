package ies

// SL-SDAP-ConfigPC5-r16 ::= SEQUENCE
// Extensible
type SlSdapConfigpc5R16 struct {
	SlMappedqosFlowstoaddlistR16     *[]SlPqfiR16 `lb:1,ub:maxNrofSLQfisperdestR16`
	SlMappedqosFlowstoreleaselistR16 *[]SlPqfiR16 `lb:1,ub:maxNrofSLQfisperdestR16`
	SlSdapHeaderR16                  SlSdapConfigpc5R16SlSdapHeaderR16
}
