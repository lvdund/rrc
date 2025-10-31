package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1250-IEs ::= SEQUENCE
type UeEutraCapabilityV1250 struct {
	PhylayerparametersV1250            *PhylayerparametersV1250
	RfParametersV1250                  *RfParametersV1250
	RlcParametersR12                   *RlcParametersR12
	UeBasednetwperfmeasparametersV1250 *UeBasednetwperfmeasparametersV1250
	UeCategorydlR12                    *utils.INTEGER `lb:0,ub:14`
	UeCategoryulR12                    *utils.INTEGER `lb:0,ub:13`
	WlanIwParametersR12                *WlanIwParametersR12
	MeasparametersV1250                *MeasparametersV1250
	DcParametersR12                    *DcParametersR12
	MbmsParametersV1250                *MbmsParametersV1250
	MacParametersR12                   *MacParametersR12
	FddAddUeEutraCapabilitiesV1250     *UeEutraCapabilityaddxddModeV1250
	TddAddUeEutraCapabilitiesV1250     *UeEutraCapabilityaddxddModeV1250
	SlParametersR12                    *SlParametersR12
	Noncriticalextension               *UeEutraCapabilityV1260
}
