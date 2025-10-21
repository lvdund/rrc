package ies

import "rrc/utils"

// PDSCH-ConfigDedicated-v1530 ::= SEQUENCE
type PdschConfigdedicatedV1530 struct {
	QclOperationV1530                     *utils.ENUMERATED
	TbsIndexalt3R15                       *utils.ENUMERATED
	CeCqiAlternativetableconfigR15        *utils.ENUMERATED
	CePdsch64qamConfigR15                 *utils.ENUMERATED
	CePdschFlexiblestartprbAllocconfigR15 *utils.ENUMERATED
	AltmcsTablescalingconfigR15           *utils.ENUMERATED
}
