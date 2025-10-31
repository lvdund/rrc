package ies

import "rrc/utils"

// SDT-MAC-PHY-CG-Config-r17 ::= SEQUENCE
// Extensible
type SdtMacPhyCgConfigR17 struct {
	CgSdtConfiglchRestrictiontoaddmodlistR17  *[]CgSdtConfiglchRestrictionR17 `lb:1,ub:maxLCId`
	CgSdtConfiglchRestrictiontoreleaselistR17 *[]Logicalchannelidentity       `lb:1,ub:maxLCId`
	CgSdtConfiginitialbwpNulR17               *utils.Setuprelease[BwpUplinkdedicatedsdtR17]
	CgSdtConfiginitialbwpSulR17               *utils.Setuprelease[BwpUplinkdedicatedsdtR17]
	CgSdtConfiginitialbwpDlR17                *BwpDownlinkdedicatedsdtR17
	CgSdtTimealignmenttimerR17                *Timealignmenttimer
	CgSdtRsrpThresholdssbR17                  *RsrpRange
	CgSdtTaValidationconfigR17                *utils.Setuprelease[CgSdtTaValidationconfigR17]
	CgSdtCsRntiR17                            *RntiValue
}
