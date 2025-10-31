package ies

// Phy-ParametersSharedSpectrumChAccess-r16 ::= SEQUENCE
// Extensible
type PhyParameterssharedspectrumchaccessR16 struct {
	SsSinrMeasR16                       *PhyParameterssharedspectrumchaccessR16SsSinrMeasR16
	SpCsiReportpucchR16                 *PhyParameterssharedspectrumchaccessR16SpCsiReportpucchR16
	SpCsiReportpuschR16                 *PhyParameterssharedspectrumchaccessR16SpCsiReportpuschR16
	DynamicsfiR16                       *PhyParameterssharedspectrumchaccessR16DynamicsfiR16
	MuxSrHarqAckCsiPucchOnceperslotR16  *PhyParameterssharedspectrumchaccessR16MuxSrHarqAckCsiPucchOnceperslotR16
	MuxSrHarqAckPucchR16                *PhyParameterssharedspectrumchaccessR16MuxSrHarqAckPucchR16
	MuxSrHarqAckCsiPucchMultiperslotR16 *PhyParameterssharedspectrumchaccessR16MuxSrHarqAckCsiPucchMultiperslotR16
	MuxHarqAckPuschDiffsymbolR16        *PhyParameterssharedspectrumchaccessR16MuxHarqAckPuschDiffsymbolR16
	PucchRepetitionF134R16              *PhyParameterssharedspectrumchaccessR16PucchRepetitionF134R16
	Type1PuschRepetitionmultislotsR16   *PhyParameterssharedspectrumchaccessR16Type1PuschRepetitionmultislotsR16
	Type2PuschRepetitionmultislotsR16   *PhyParameterssharedspectrumchaccessR16Type2PuschRepetitionmultislotsR16
	PuschRepetitionmultislotsR16        *PhyParameterssharedspectrumchaccessR16PuschRepetitionmultislotsR16
	PdschRepetitionmultislotsR16        *PhyParameterssharedspectrumchaccessR16PdschRepetitionmultislotsR16
	DownlinkspsR16                      *PhyParameterssharedspectrumchaccessR16DownlinkspsR16
	ConfiguredulGranttype1R16           *PhyParameterssharedspectrumchaccessR16ConfiguredulGranttype1R16
	ConfiguredulGranttype2R16           *PhyParameterssharedspectrumchaccessR16ConfiguredulGranttype2R16
	PreEmptindicationDlR16              *PhyParameterssharedspectrumchaccessR16PreEmptindicationDlR16
}
