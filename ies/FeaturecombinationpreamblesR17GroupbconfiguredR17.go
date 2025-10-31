package ies

import "rrc/utils"

// FeatureCombinationPreambles-r17-groupBconfigured-r17 ::= SEQUENCE
type FeaturecombinationpreamblesR17GroupbconfiguredR17 struct {
	RaSizegroupaR17              FeaturecombinationpreamblesR17GroupbconfiguredR17RaSizegroupaR17
	MessagepoweroffsetgroupbR17  FeaturecombinationpreamblesR17GroupbconfiguredR17MessagepoweroffsetgroupbR17
	NumberofraPreamblesgroupaR17 utils.INTEGER `lb:0,ub:64`
}
