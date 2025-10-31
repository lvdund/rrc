package ies

// PDSCH-ConfigDedicated-v1530 ::= SEQUENCE
type PdschConfigdedicatedV1530 struct {
	QclOperationV1530                     *PdschConfigdedicatedV1530QclOperationV1530
	TbsIndexalt3R15                       *PdschConfigdedicatedV1530TbsIndexalt3R15
	CeCqiAlternativetableconfigR15        *PdschConfigdedicatedV1530CeCqiAlternativetableconfigR15
	CePdsch64qamConfigR15                 *PdschConfigdedicatedV1530CePdsch64qamConfigR15
	CePdschFlexiblestartprbAllocconfigR15 *PdschConfigdedicatedV1530CePdschFlexiblestartprbAllocconfigR15
	AltmcsTablescalingconfigR15           *PdschConfigdedicatedV1530AltmcsTablescalingconfigR15
}
