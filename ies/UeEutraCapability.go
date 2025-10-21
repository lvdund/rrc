package ies

import "rrc/utils"

// UE-EUTRA-Capability ::= SEQUENCE
type UeEutraCapability struct {
	Accessstratumrelease   Accessstratumrelease
	UeCategory             utils.INTEGER
	PdcpParameters         PdcpParameters
	Phylayerparameters     Phylayerparameters
	RfParameters           RfParameters
	Measparameters         Measparameters
	Featuregroupindicators *utils.BITSTRING
	InterratParameters     *interface{}
	Noncriticalextension   *UeEutraCapabilityV920Ies
}
