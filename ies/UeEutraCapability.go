package ies

import "rrc/utils"

// UE-EUTRA-Capability ::= SEQUENCE
type UeEutraCapability struct {
	Accessstratumrelease   Accessstratumrelease
	UeCategory             utils.INTEGER `lb:0,ub:5`
	PdcpParameters         PdcpParameters
	Phylayerparameters     Phylayerparameters
	RfParameters           RfParameters
	Measparameters         Measparameters
	Featuregroupindicators *utils.BITSTRING `lb:32,ub:32`
	InterratParameters     *UeEutraCapabilityInterratParameters
	Noncriticalextension   *UeEutraCapabilityV920
}
