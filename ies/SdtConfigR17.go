package ies

import "rrc/utils"

// SDT-Config-r17 ::= SEQUENCE
type SdtConfigR17 struct {
	SdtDrbListR17         *[]DrbIdentity `lb:0,ub:maxDRB`
	SdtSrb2IndicationR17  *SdtConfigR17SdtSrb2IndicationR17
	SdtMacPhyCgConfigR17  *utils.Setuprelease[SdtCgConfigR17]
	SdtDrbContinuerohcR17 *SdtConfigR17SdtDrbContinuerohcR17
}
