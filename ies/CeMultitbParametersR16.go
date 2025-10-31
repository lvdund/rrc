package ies

// CE-MultiTB-Parameters-r16 ::= SEQUENCE
type CeMultitbParametersR16 struct {
	PdschMultitbCeModeaR16       *CeMultitbParametersR16PdschMultitbCeModeaR16
	PdschMultitbCeModebR16       *CeMultitbParametersR16PdschMultitbCeModebR16
	PuschMultitbCeModeaR16       *CeMultitbParametersR16PuschMultitbCeModeaR16
	PuschMultitbCeModebR16       *CeMultitbParametersR16PuschMultitbCeModebR16
	CeMultitb64qamR16            *CeMultitbParametersR16CeMultitb64qamR16
	CeMultitbEarlyterminationR16 *CeMultitbParametersR16CeMultitbEarlyterminationR16
	CeMultitbFrequencyhoppingR16 *CeMultitbParametersR16CeMultitbFrequencyhoppingR16
	CeMultitbHarqAckbundlingR16  *CeMultitbParametersR16CeMultitbHarqAckbundlingR16
	CeMultitbInterleavingR16     *CeMultitbParametersR16CeMultitbInterleavingR16
	CeMultitbSubprbR16           *CeMultitbParametersR16CeMultitbSubprbR16
}
