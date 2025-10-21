package ies

import "rrc/utils"

// PDCP-ParametersNR-r15 ::= SEQUENCE
type PdcpParametersnrR15 struct {
	RohcProfilesR15                RohcProfilesupportlistR15
	RohcContextmaxsessionsR15      utils.ENUMERATED
	RohcProfilesulOnlyR15          interface{}
	RohcContextcontinueR15         *utils.ENUMERATED
	OutoforderdeliveryR15          *utils.ENUMERATED
	SnSizeloR15                    *utils.ENUMERATED
	ImsVoiceovernrPdcpMcgBearerR15 *utils.ENUMERATED
	ImsVoiceovernrPdcpScgBearerR15 *utils.ENUMERATED
}
