package ies

// MRDC-Parameters ::= SEQUENCE
// Extensible
type MrdcParameters struct {
	SingleulTransmission          *MrdcParametersSingleulTransmission
	Dynamicpowersharingendc       *MrdcParametersDynamicpowersharingendc
	TdmPattern                    *MrdcParametersTdmPattern
	UlSharingeutraNr              *MrdcParametersUlSharingeutraNr
	UlSwitchingtimeeutraNr        *MrdcParametersUlSwitchingtimeeutraNr
	Simultaneousrxtxinterbandendc *MrdcParametersSimultaneousrxtxinterbandendc
	Asyncintrabandendc            *MrdcParametersAsyncintrabandendc
	DualpaArchitecture            *MrdcParametersDualpaArchitecture       `ext`
	IntrabandendcSupport          *MrdcParametersIntrabandendcSupport     `ext`
	UlTimingalignmenteutraNr      *MrdcParametersUlTimingalignmenteutraNr `ext`
}
