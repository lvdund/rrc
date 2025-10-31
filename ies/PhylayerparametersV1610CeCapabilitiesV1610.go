package ies

// PhyLayerParameters-v1610-ce-Capabilities-v1610 ::= SEQUENCE
type PhylayerparametersV1610CeCapabilitiesV1610 struct {
	CeCsiRsFeedbackR16                    *PhylayerparametersV1610CeCapabilitiesV1610CeCsiRsFeedbackR16
	CeCsiRsFeedbackcodebookrestrictionR16 *PhylayerparametersV1610CeCapabilitiesV1610CeCsiRsFeedbackcodebookrestrictionR16
	CrsChestmpdcchCeModeaR16              *PhylayerparametersV1610CeCapabilitiesV1610CrsChestmpdcchCeModeaR16
	CrsChestmpdcchCeModebR16              *PhylayerparametersV1610CeCapabilitiesV1610CrsChestmpdcchCeModebR16
	CrsChestmpdcchCsiR16                  *PhylayerparametersV1610CeCapabilitiesV1610CrsChestmpdcchCsiR16
	CrsChestmpdcchReciprocitytddR16       *PhylayerparametersV1610CeCapabilitiesV1610CrsChestmpdcchReciprocitytddR16
	EtwsCmasRxinconnceModeaR16            *PhylayerparametersV1610CeCapabilitiesV1610EtwsCmasRxinconnceModeaR16
	EtwsCmasRxinconnceModebR16            *PhylayerparametersV1610CeCapabilitiesV1610EtwsCmasRxinconnceModebR16
	MpdcchInltecontrolregionceModeaR16    *PhylayerparametersV1610CeCapabilitiesV1610MpdcchInltecontrolregionceModeaR16
	MpdcchInltecontrolregionceModebR16    *PhylayerparametersV1610CeCapabilitiesV1610MpdcchInltecontrolregionceModebR16
	PdschInltecontrolregionceModeaR16     *PhylayerparametersV1610CeCapabilitiesV1610PdschInltecontrolregionceModeaR16
	PdschInltecontrolregionceModebR16     *PhylayerparametersV1610CeCapabilitiesV1610PdschInltecontrolregionceModebR16
	MultitbParametersR16                  *CeMultitbParametersR16
	ResourceresvparametersR16             *CeResourceresvparametersR16
}
