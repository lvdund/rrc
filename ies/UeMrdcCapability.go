package ies

import "rrc/utils"

// UE-MRDC-Capability ::= SEQUENCE
type UeMrdcCapability struct {
	Measandmobparametersmrdc *Measandmobparametersmrdc
	PhyParametersmrdcV1530   *PhyParametersmrdc
	RfParametersmrdc         RfParametersmrdc
	Generalparametersmrdc    *GeneralparametersmrdcXddDiff
	FddAddUeMrdcCapabilities *UeMrdcCapabilityaddxddMode
	TddAddUeMrdcCapabilities *UeMrdcCapabilityaddxddMode
	Fr1AddUeMrdcCapabilities *UeMrdcCapabilityaddfrxMode
	Fr2AddUeMrdcCapabilities *UeMrdcCapabilityaddfrxMode
	Featuresetcombinations   *[]Featuresetcombination `lb:1,ub:maxFeatureSetCombinations`
	PdcpParametersmrdcV1530  *PdcpParametersmrdc
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeMrdcCapabilityV1560
}
