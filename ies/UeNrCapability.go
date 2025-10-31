package ies

import "rrc/utils"

// UE-NR-Capability ::= SEQUENCE
type UeNrCapability struct {
	Accessstratumrelease     Accessstratumrelease
	PdcpParameters           PdcpParameters
	RlcParameters            *RlcParameters
	MacParameters            *MacParameters
	PhyParameters            PhyParameters
	RfParameters             RfParameters
	Measandmobparameters     *Measandmobparameters
	FddAddUeNrCapabilities   *UeNrCapabilityaddxddMode
	TddAddUeNrCapabilities   *UeNrCapabilityaddxddMode
	Fr1AddUeNrCapabilities   *UeNrCapabilityaddfrxMode
	Fr2AddUeNrCapabilities   *UeNrCapabilityaddfrxMode
	Featuresets              *Featuresets
	Featuresetcombinations   *[]Featuresetcombination `lb:1,ub:maxFeatureSetCombinations`
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeNrCapabilityV1530
}
