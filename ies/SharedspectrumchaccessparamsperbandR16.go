package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16 ::= SEQUENCE
type SharedspectrumchaccessparamsperbandR16 struct {
	UlDynamicchaccessR16                *SharedspectrumchaccessparamsperbandR16UlDynamicchaccessR16
	UlSemiStaticchaccessR16             *SharedspectrumchaccessparamsperbandR16UlSemiStaticchaccessR16
	SsbRrmDynamicchaccessR16            *SharedspectrumchaccessparamsperbandR16SsbRrmDynamicchaccessR16
	SsbRrmSemiStaticchaccessR16         *SharedspectrumchaccessparamsperbandR16SsbRrmSemiStaticchaccessR16
	MibAcquisitionR16                   *SharedspectrumchaccessparamsperbandR16MibAcquisitionR16
	SsbRlmDynamicchaccessR16            *SharedspectrumchaccessparamsperbandR16SsbRlmDynamicchaccessR16
	SsbRlmSemiStaticchaccessR16         *SharedspectrumchaccessparamsperbandR16SsbRlmSemiStaticchaccessR16
	Sib1AcquisitionR16                  *SharedspectrumchaccessparamsperbandR16Sib1AcquisitionR16
	ExtraResponsewindowR16              *SharedspectrumchaccessparamsperbandR16ExtraResponsewindowR16
	SsbBfdCbdDynamicchannelaccessR16    *SharedspectrumchaccessparamsperbandR16SsbBfdCbdDynamicchannelaccessR16
	SsbBfdCbdSemiStaticchannelaccessR16 *SharedspectrumchaccessparamsperbandR16SsbBfdCbdSemiStaticchannelaccessR16
	CsiRsBfdCbdR16                      *SharedspectrumchaccessparamsperbandR16CsiRsBfdCbdR16
	UlChannelbwScell10mhzR16            *SharedspectrumchaccessparamsperbandR16UlChannelbwScell10mhzR16
	RssiChanneloccupancyreportingR16    *SharedspectrumchaccessparamsperbandR16RssiChanneloccupancyreportingR16
	SrsStartanyofdmSymbolR16            *SharedspectrumchaccessparamsperbandR16SrsStartanyofdmSymbolR16
	SearchspacefreqmonitorlocationR16   *utils.INTEGER `lb:0,ub:5`
	CoresetRbOffsetR16                  *SharedspectrumchaccessparamsperbandR16CoresetRbOffsetR16
	CgiAcquisitionR16                   *SharedspectrumchaccessparamsperbandR16CgiAcquisitionR16
	ConfiguredulTxR16                   *SharedspectrumchaccessparamsperbandR16ConfiguredulTxR16
	PrachWidebandR16                    *SharedspectrumchaccessparamsperbandR16PrachWidebandR16
	DciAvailablerbSetR16                *SharedspectrumchaccessparamsperbandR16DciAvailablerbSetR16
	DciChoccupancydurationR16           *SharedspectrumchaccessparamsperbandR16DciChoccupancydurationR16
	TypebPdschLengthR16                 *SharedspectrumchaccessparamsperbandR16TypebPdschLengthR16
	SearchspaceswitchwithdciR16         *SharedspectrumchaccessparamsperbandR16SearchspaceswitchwithdciR16
	SearchspaceswitchwithoutdciR16      *SharedspectrumchaccessparamsperbandR16SearchspaceswitchwithoutdciR16
	Searchspaceswitchcapability2R16     *SharedspectrumchaccessparamsperbandR16Searchspaceswitchcapability2R16
	NonNumericalpdschHarqTimingR16      *SharedspectrumchaccessparamsperbandR16NonNumericalpdschHarqTimingR16
	EnhanceddynamicharqCodebookR16      *SharedspectrumchaccessparamsperbandR16EnhanceddynamicharqCodebookR16
	OneshotharqFeedbackR16              *SharedspectrumchaccessparamsperbandR16OneshotharqFeedbackR16
	MultipuschUlGrantR16                *SharedspectrumchaccessparamsperbandR16MultipuschUlGrantR16
	CsiRsRlmR16                         *SharedspectrumchaccessparamsperbandR16CsiRsRlmR16
	Dummy                               *SharedspectrumchaccessparamsperbandR16Dummy
	PeriodicandsemiPersistentcsiRsR16   *SharedspectrumchaccessparamsperbandR16PeriodicandsemiPersistentcsiRsR16
	PuschPrbInterlaceR16                *SharedspectrumchaccessparamsperbandR16PuschPrbInterlaceR16
	PucchF0F1PrbInterlaceR16            *SharedspectrumchaccessparamsperbandR16PucchF0F1PrbInterlaceR16
	OccPrbPf2Pf3R16                     *SharedspectrumchaccessparamsperbandR16OccPrbPf2Pf3R16
	ExtcpRangecgPuschR16                *SharedspectrumchaccessparamsperbandR16ExtcpRangecgPuschR16
	ConfiguredgrantwithretxR16          *SharedspectrumchaccessparamsperbandR16ConfiguredgrantwithretxR16
	EdThresholdR16                      *SharedspectrumchaccessparamsperbandR16EdThresholdR16
	UlDlCotSharingR16                   *SharedspectrumchaccessparamsperbandR16UlDlCotSharingR16
	MuxCgUciHarqAckR16                  *SharedspectrumchaccessparamsperbandR16MuxCgUciHarqAckR16
	CgResourceconfigR16                 *SharedspectrumchaccessparamsperbandR16CgResourceconfigR16
}
