// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetUplink-v1610 (line 20120).

var featureSetUplinkV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pusch-RepetitionTypeB-r16", Optional: true},
		{Name: "ul-CancellationSelfCarrier-r16", Optional: true},
		{Name: "ul-CancellationCrossCarrier-r16", Optional: true},
		{Name: "ul-FullPwrMode2-MaxSRS-ResInSet-r16", Optional: true},
		{Name: "cbgPUSCH-ProcessingType1-DifferentTB-PerSlot-r16", Optional: true},
		{Name: "cbgPUSCH-ProcessingType2-DifferentTB-PerSlot-r16", Optional: true},
		{Name: "supportedSRS-PosResources-r16", Optional: true},
		{Name: "intraFreqDAPS-UL-r16", Optional: true},
		{Name: "intraBandFreqSeparationUL-v1620", Optional: true},
		{Name: "multiPUCCH-r16", Optional: true},
		{Name: "twoPUCCH-Type1-r16", Optional: true},
		{Name: "twoPUCCH-Type2-r16", Optional: true},
		{Name: "twoPUCCH-Type3-r16", Optional: true},
		{Name: "twoPUCCH-Type4-r16", Optional: true},
		{Name: "mux-SR-HARQ-ACK-r16", Optional: true},
		{Name: "dummy1", Optional: true},
		{Name: "dummy2", Optional: true},
		{Name: "twoPUCCH-Type5-r16", Optional: true},
		{Name: "twoPUCCH-Type6-r16", Optional: true},
		{Name: "twoPUCCH-Type7-r16", Optional: true},
		{Name: "twoPUCCH-Type8-r16", Optional: true},
		{Name: "twoPUCCH-Type9-r16", Optional: true},
		{Name: "twoPUCCH-Type10-r16", Optional: true},
		{Name: "twoPUCCH-Type11-r16", Optional: true},
		{Name: "ul-IntraUE-Mux-r16", Optional: true},
		{Name: "ul-FullPwrMode-r16", Optional: true},
		{Name: "crossCarrierSchedulingProcessing-DiffSCS-r16", Optional: true},
		{Name: "ul-FullPwrMode1-r16", Optional: true},
		{Name: "ul-FullPwrMode2-SRSConfig-diffNumSRSPorts-r16", Optional: true},
		{Name: "ul-FullPwrMode2-TPMIGroup-r16", Optional: true},
	},
}

const (
	FeatureSetUplink_v1610_Ul_CancellationSelfCarrier_r16_Supported = 0
)

var featureSetUplinkV1610UlCancellationSelfCarrierR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_Ul_CancellationSelfCarrier_r16_Supported},
}

const (
	FeatureSetUplink_v1610_Ul_CancellationCrossCarrier_r16_Supported = 0
)

var featureSetUplinkV1610UlCancellationCrossCarrierR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_Ul_CancellationCrossCarrier_r16_Supported},
}

const (
	FeatureSetUplink_v1610_Ul_FullPwrMode2_MaxSRS_ResInSet_r16_N1 = 0
	FeatureSetUplink_v1610_Ul_FullPwrMode2_MaxSRS_ResInSet_r16_N2 = 1
	FeatureSetUplink_v1610_Ul_FullPwrMode2_MaxSRS_ResInSet_r16_N4 = 2
)

var featureSetUplinkV1610UlFullPwrMode2MaxSRSResInSetR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_Ul_FullPwrMode2_MaxSRS_ResInSet_r16_N1, FeatureSetUplink_v1610_Ul_FullPwrMode2_MaxSRS_ResInSet_r16_N2, FeatureSetUplink_v1610_Ul_FullPwrMode2_MaxSRS_ResInSet_r16_N4},
}

const (
	FeatureSetUplink_v1610_TwoPUCCH_Type1_r16_Supported = 0
)

var featureSetUplinkV1610TwoPUCCHType1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_TwoPUCCH_Type1_r16_Supported},
}

const (
	FeatureSetUplink_v1610_TwoPUCCH_Type2_r16_Supported = 0
)

var featureSetUplinkV1610TwoPUCCHType2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_TwoPUCCH_Type2_r16_Supported},
}

const (
	FeatureSetUplink_v1610_TwoPUCCH_Type3_r16_Supported = 0
)

var featureSetUplinkV1610TwoPUCCHType3R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_TwoPUCCH_Type3_r16_Supported},
}

const (
	FeatureSetUplink_v1610_TwoPUCCH_Type4_r16_Supported = 0
)

var featureSetUplinkV1610TwoPUCCHType4R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_TwoPUCCH_Type4_r16_Supported},
}

const (
	FeatureSetUplink_v1610_Mux_SR_HARQ_ACK_r16_Supported = 0
)

var featureSetUplinkV1610MuxSRHARQACKR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_Mux_SR_HARQ_ACK_r16_Supported},
}

const (
	FeatureSetUplink_v1610_Dummy1_Supported = 0
)

var featureSetUplinkV1610Dummy1Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_Dummy1_Supported},
}

const (
	FeatureSetUplink_v1610_Dummy2_Supported = 0
)

var featureSetUplinkV1610Dummy2Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_Dummy2_Supported},
}

const (
	FeatureSetUplink_v1610_TwoPUCCH_Type5_r16_Supported = 0
)

var featureSetUplinkV1610TwoPUCCHType5R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_TwoPUCCH_Type5_r16_Supported},
}

const (
	FeatureSetUplink_v1610_TwoPUCCH_Type6_r16_Supported = 0
)

var featureSetUplinkV1610TwoPUCCHType6R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_TwoPUCCH_Type6_r16_Supported},
}

const (
	FeatureSetUplink_v1610_TwoPUCCH_Type7_r16_Supported = 0
)

var featureSetUplinkV1610TwoPUCCHType7R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_TwoPUCCH_Type7_r16_Supported},
}

const (
	FeatureSetUplink_v1610_TwoPUCCH_Type8_r16_Supported = 0
)

var featureSetUplinkV1610TwoPUCCHType8R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_TwoPUCCH_Type8_r16_Supported},
}

const (
	FeatureSetUplink_v1610_TwoPUCCH_Type9_r16_Supported = 0
)

var featureSetUplinkV1610TwoPUCCHType9R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_TwoPUCCH_Type9_r16_Supported},
}

const (
	FeatureSetUplink_v1610_TwoPUCCH_Type10_r16_Supported = 0
)

var featureSetUplinkV1610TwoPUCCHType10R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_TwoPUCCH_Type10_r16_Supported},
}

const (
	FeatureSetUplink_v1610_TwoPUCCH_Type11_r16_Supported = 0
)

var featureSetUplinkV1610TwoPUCCHType11R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_TwoPUCCH_Type11_r16_Supported},
}

const (
	FeatureSetUplink_v1610_Ul_FullPwrMode_r16_Supported = 0
)

var featureSetUplinkV1610UlFullPwrModeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_Ul_FullPwrMode_r16_Supported},
}

const (
	FeatureSetUplink_v1610_Ul_FullPwrMode1_r16_Supported = 0
)

var featureSetUplinkV1610UlFullPwrMode1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_Ul_FullPwrMode1_r16_Supported},
}

const (
	FeatureSetUplink_v1610_Ul_FullPwrMode2_SRSConfig_DiffNumSRSPorts_r16_P1_2   = 0
	FeatureSetUplink_v1610_Ul_FullPwrMode2_SRSConfig_DiffNumSRSPorts_r16_P1_4   = 1
	FeatureSetUplink_v1610_Ul_FullPwrMode2_SRSConfig_DiffNumSRSPorts_r16_P1_2_4 = 2
)

var featureSetUplinkV1610UlFullPwrMode2SRSConfigDiffNumSRSPortsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_Ul_FullPwrMode2_SRSConfig_DiffNumSRSPorts_r16_P1_2, FeatureSetUplink_v1610_Ul_FullPwrMode2_SRSConfig_DiffNumSRSPorts_r16_P1_4, FeatureSetUplink_v1610_Ul_FullPwrMode2_SRSConfig_DiffNumSRSPorts_r16_P1_2_4},
}

const (
	FeatureSetUplink_v1610_Pusch_RepetitionTypeB_r16_MaxNumberPUSCH_Tx_r16_N2  = 0
	FeatureSetUplink_v1610_Pusch_RepetitionTypeB_r16_MaxNumberPUSCH_Tx_r16_N3  = 1
	FeatureSetUplink_v1610_Pusch_RepetitionTypeB_r16_MaxNumberPUSCH_Tx_r16_N4  = 2
	FeatureSetUplink_v1610_Pusch_RepetitionTypeB_r16_MaxNumberPUSCH_Tx_r16_N7  = 3
	FeatureSetUplink_v1610_Pusch_RepetitionTypeB_r16_MaxNumberPUSCH_Tx_r16_N8  = 4
	FeatureSetUplink_v1610_Pusch_RepetitionTypeB_r16_MaxNumberPUSCH_Tx_r16_N12 = 5
)

var featureSetUplinkV1610PuschRepetitionTypeBR16MaxNumberPUSCHTxR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_Pusch_RepetitionTypeB_r16_MaxNumberPUSCH_Tx_r16_N2, FeatureSetUplink_v1610_Pusch_RepetitionTypeB_r16_MaxNumberPUSCH_Tx_r16_N3, FeatureSetUplink_v1610_Pusch_RepetitionTypeB_r16_MaxNumberPUSCH_Tx_r16_N4, FeatureSetUplink_v1610_Pusch_RepetitionTypeB_r16_MaxNumberPUSCH_Tx_r16_N7, FeatureSetUplink_v1610_Pusch_RepetitionTypeB_r16_MaxNumberPUSCH_Tx_r16_N8, FeatureSetUplink_v1610_Pusch_RepetitionTypeB_r16_MaxNumberPUSCH_Tx_r16_N12},
}

const (
	FeatureSetUplink_v1610_Pusch_RepetitionTypeB_r16_HoppingScheme_r16_InterSlotHopping       = 0
	FeatureSetUplink_v1610_Pusch_RepetitionTypeB_r16_HoppingScheme_r16_InterRepetitionHopping = 1
	FeatureSetUplink_v1610_Pusch_RepetitionTypeB_r16_HoppingScheme_r16_Both                   = 2
)

var featureSetUplinkV1610PuschRepetitionTypeBR16HoppingSchemeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_Pusch_RepetitionTypeB_r16_HoppingScheme_r16_InterSlotHopping, FeatureSetUplink_v1610_Pusch_RepetitionTypeB_r16_HoppingScheme_r16_InterRepetitionHopping, FeatureSetUplink_v1610_Pusch_RepetitionTypeB_r16_HoppingScheme_r16_Both},
}

var featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r16", Optional: true},
		{Name: "scs-30kHz-r16", Optional: true},
		{Name: "scs-60kHz-r16", Optional: true},
		{Name: "scs-120kHz-r16", Optional: true},
	},
}

const (
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_15kHz_r16_One_Pusch = 0
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto2     = 1
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto4     = 2
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto7     = 3
)

var featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Scs15kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_15kHz_r16_One_Pusch, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto2, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto4, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto7},
}

const (
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_30kHz_r16_One_Pusch = 0
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto2     = 1
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto4     = 2
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto7     = 3
)

var featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Scs30kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_30kHz_r16_One_Pusch, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto2, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto4, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto7},
}

const (
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_60kHz_r16_One_Pusch = 0
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto2     = 1
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto4     = 2
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto7     = 3
)

var featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Scs60kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_60kHz_r16_One_Pusch, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto2, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto4, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto7},
}

const (
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_120kHz_r16_One_Pusch = 0
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto2     = 1
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto4     = 2
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto7     = 3
)

var featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Scs120kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_120kHz_r16_One_Pusch, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto2, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto4, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto7},
}

var featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r16", Optional: true},
		{Name: "scs-30kHz-r16", Optional: true},
		{Name: "scs-60kHz-r16", Optional: true},
		{Name: "scs-120kHz-r16", Optional: true},
	},
}

const (
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_15kHz_r16_One_Pusch = 0
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto2     = 1
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto4     = 2
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto7     = 3
)

var featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Scs15kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_15kHz_r16_One_Pusch, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto2, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto4, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto7},
}

const (
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_30kHz_r16_One_Pusch = 0
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto2     = 1
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto4     = 2
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto7     = 3
)

var featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Scs30kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_30kHz_r16_One_Pusch, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto2, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto4, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto7},
}

const (
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_60kHz_r16_One_Pusch = 0
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto2     = 1
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto4     = 2
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto7     = 3
)

var featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Scs60kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_60kHz_r16_One_Pusch, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto2, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto4, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto7},
}

const (
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_120kHz_r16_One_Pusch = 0
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto2     = 1
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto4     = 2
	FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto7     = 3
)

var featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Scs120kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_120kHz_r16_One_Pusch, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto2, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto4, FeatureSetUplink_v1610_CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto7},
}

var featureSetUplinkV1610IntraFreqDAPSULR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dummy", Optional: true},
		{Name: "intraFreqTwoTAGs-DAPS-r16", Optional: true},
		{Name: "dummy1", Optional: true},
		{Name: "dummy2", Optional: true},
		{Name: "dummy3", Optional: true},
	},
}

const (
	FeatureSetUplink_v1610_IntraFreqDAPS_UL_r16_Dummy_Supported = 0
)

var featureSetUplinkV1610IntraFreqDAPSULR16DummyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_IntraFreqDAPS_UL_r16_Dummy_Supported},
}

const (
	FeatureSetUplink_v1610_IntraFreqDAPS_UL_r16_IntraFreqTwoTAGs_DAPS_r16_Supported = 0
)

var featureSetUplinkV1610IntraFreqDAPSULR16IntraFreqTwoTAGsDAPSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_IntraFreqDAPS_UL_r16_IntraFreqTwoTAGs_DAPS_r16_Supported},
}

const (
	FeatureSetUplink_v1610_IntraFreqDAPS_UL_r16_Dummy1_Supported = 0
)

var featureSetUplinkV1610IntraFreqDAPSULR16Dummy1Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_IntraFreqDAPS_UL_r16_Dummy1_Supported},
}

const (
	FeatureSetUplink_v1610_IntraFreqDAPS_UL_r16_Dummy2_Supported = 0
)

var featureSetUplinkV1610IntraFreqDAPSULR16Dummy2Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_IntraFreqDAPS_UL_r16_Dummy2_Supported},
}

const (
	FeatureSetUplink_v1610_IntraFreqDAPS_UL_r16_Dummy3_Short = 0
	FeatureSetUplink_v1610_IntraFreqDAPS_UL_r16_Dummy3_Long  = 1
)

var featureSetUplinkV1610IntraFreqDAPSULR16Dummy3Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_IntraFreqDAPS_UL_r16_Dummy3_Short, FeatureSetUplink_v1610_IntraFreqDAPS_UL_r16_Dummy3_Long},
}

var featureSetUplinkV1610MultiPUCCHR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sub-SlotConfig-NCP-r16", Optional: true},
		{Name: "sub-SlotConfig-ECP-r16", Optional: true},
	},
}

const (
	FeatureSetUplink_v1610_MultiPUCCH_r16_Sub_SlotConfig_NCP_r16_Set1 = 0
	FeatureSetUplink_v1610_MultiPUCCH_r16_Sub_SlotConfig_NCP_r16_Set2 = 1
)

var featureSetUplinkV1610MultiPUCCHR16SubSlotConfigNCPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_MultiPUCCH_r16_Sub_SlotConfig_NCP_r16_Set1, FeatureSetUplink_v1610_MultiPUCCH_r16_Sub_SlotConfig_NCP_r16_Set2},
}

const (
	FeatureSetUplink_v1610_MultiPUCCH_r16_Sub_SlotConfig_ECP_r16_Set1 = 0
	FeatureSetUplink_v1610_MultiPUCCH_r16_Sub_SlotConfig_ECP_r16_Set2 = 1
)

var featureSetUplinkV1610MultiPUCCHR16SubSlotConfigECPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_MultiPUCCH_r16_Sub_SlotConfig_ECP_r16_Set1, FeatureSetUplink_v1610_MultiPUCCH_r16_Sub_SlotConfig_ECP_r16_Set2},
}

const (
	FeatureSetUplink_v1610_Ul_IntraUE_Mux_r16_Pusch_PreparationLowPriority_r16_Sym0 = 0
	FeatureSetUplink_v1610_Ul_IntraUE_Mux_r16_Pusch_PreparationLowPriority_r16_Sym1 = 1
	FeatureSetUplink_v1610_Ul_IntraUE_Mux_r16_Pusch_PreparationLowPriority_r16_Sym2 = 2
)

var featureSetUplinkV1610UlIntraUEMuxR16PuschPreparationLowPriorityR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_Ul_IntraUE_Mux_r16_Pusch_PreparationLowPriority_r16_Sym0, FeatureSetUplink_v1610_Ul_IntraUE_Mux_r16_Pusch_PreparationLowPriority_r16_Sym1, FeatureSetUplink_v1610_Ul_IntraUE_Mux_r16_Pusch_PreparationLowPriority_r16_Sym2},
}

const (
	FeatureSetUplink_v1610_Ul_IntraUE_Mux_r16_Pusch_PreparationHighPriority_r16_Sym0 = 0
	FeatureSetUplink_v1610_Ul_IntraUE_Mux_r16_Pusch_PreparationHighPriority_r16_Sym1 = 1
	FeatureSetUplink_v1610_Ul_IntraUE_Mux_r16_Pusch_PreparationHighPriority_r16_Sym2 = 2
)

var featureSetUplinkV1610UlIntraUEMuxR16PuschPreparationHighPriorityR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_Ul_IntraUE_Mux_r16_Pusch_PreparationHighPriority_r16_Sym0, FeatureSetUplink_v1610_Ul_IntraUE_Mux_r16_Pusch_PreparationHighPriority_r16_Sym1, FeatureSetUplink_v1610_Ul_IntraUE_Mux_r16_Pusch_PreparationHighPriority_r16_Sym2},
}

var featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-120kHz-r16", Optional: true},
		{Name: "scs-15kHz-60kHz-r16", Optional: true},
		{Name: "scs-30kHz-120kHz-r16", Optional: true},
		{Name: "scs-15kHz-30kHz-r16", Optional: true},
		{Name: "scs-30kHz-60kHz-r16", Optional: true},
		{Name: "scs-60kHz-120kHz-r16", Optional: true},
	},
}

const (
	FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_120kHz_r16_N1 = 0
	FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_120kHz_r16_N2 = 1
	FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_120kHz_r16_N4 = 2
)

var featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs15kHz120kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_120kHz_r16_N1, FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_120kHz_r16_N2, FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_120kHz_r16_N4},
}

const (
	FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_60kHz_r16_N1 = 0
	FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_60kHz_r16_N2 = 1
	FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_60kHz_r16_N4 = 2
)

var featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs15kHz60kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_60kHz_r16_N1, FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_60kHz_r16_N2, FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_60kHz_r16_N4},
}

const (
	FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_30kHz_120kHz_r16_N1 = 0
	FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_30kHz_120kHz_r16_N2 = 1
	FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_30kHz_120kHz_r16_N4 = 2
)

var featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs30kHz120kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_30kHz_120kHz_r16_N1, FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_30kHz_120kHz_r16_N2, FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_30kHz_120kHz_r16_N4},
}

const (
	FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_30kHz_r16_N2 = 0
)

var featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs15kHz30kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_30kHz_r16_N2},
}

const (
	FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_30kHz_60kHz_r16_N2 = 0
)

var featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs30kHz60kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_30kHz_60kHz_r16_N2},
}

const (
	FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_60kHz_120kHz_r16_N2 = 0
)

var featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs60kHz120kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_60kHz_120kHz_r16_N2},
}

var featureSetUplinkV1610UlFullPwrMode2TPMIGroupR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "twoPorts-r16", Optional: true},
		{Name: "fourPortsNonCoherent-r16", Optional: true},
		{Name: "fourPortsPartialCoherent-r16", Optional: true},
	},
}

const (
	FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsNonCoherent_r16_G0 = 0
	FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsNonCoherent_r16_G1 = 1
	FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsNonCoherent_r16_G2 = 2
	FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsNonCoherent_r16_G3 = 3
)

var featureSetUplinkV1610UlFullPwrMode2TPMIGroupR16FourPortsNonCoherentR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsNonCoherent_r16_G0, FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsNonCoherent_r16_G1, FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsNonCoherent_r16_G2, FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsNonCoherent_r16_G3},
}

const (
	FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsPartialCoherent_r16_G0 = 0
	FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsPartialCoherent_r16_G1 = 1
	FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsPartialCoherent_r16_G2 = 2
	FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsPartialCoherent_r16_G3 = 3
	FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsPartialCoherent_r16_G4 = 4
	FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsPartialCoherent_r16_G5 = 5
	FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsPartialCoherent_r16_G6 = 6
)

var featureSetUplinkV1610UlFullPwrMode2TPMIGroupR16FourPortsPartialCoherentR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsPartialCoherent_r16_G0, FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsPartialCoherent_r16_G1, FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsPartialCoherent_r16_G2, FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsPartialCoherent_r16_G3, FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsPartialCoherent_r16_G4, FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsPartialCoherent_r16_G5, FeatureSetUplink_v1610_Ul_FullPwrMode2_TPMIGroup_r16_FourPortsPartialCoherent_r16_G6},
}

type FeatureSetUplink_v1610 struct {
	Pusch_RepetitionTypeB_r16 *struct {
		MaxNumberPUSCH_Tx_r16 int64
		HoppingScheme_r16     int64
	}
	Ul_CancellationSelfCarrier_r16                   *int64
	Ul_CancellationCrossCarrier_r16                  *int64
	Ul_FullPwrMode2_MaxSRS_ResInSet_r16              *int64
	CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16 *struct {
		Scs_15kHz_r16  *int64
		Scs_30kHz_r16  *int64
		Scs_60kHz_r16  *int64
		Scs_120kHz_r16 *int64
	}
	CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16 *struct {
		Scs_15kHz_r16  *int64
		Scs_30kHz_r16  *int64
		Scs_60kHz_r16  *int64
		Scs_120kHz_r16 *int64
	}
	SupportedSRS_PosResources_r16 *SRS_AllPosResources_r16
	IntraFreqDAPS_UL_r16          *struct {
		Dummy                     *int64
		IntraFreqTwoTAGs_DAPS_r16 *int64
		Dummy1                    *int64
		Dummy2                    *int64
		Dummy3                    *int64
	}
	IntraBandFreqSeparationUL_v1620 *FreqSeparationClassUL_v1620
	MultiPUCCH_r16                  *struct {
		Sub_SlotConfig_NCP_r16 *int64
		Sub_SlotConfig_ECP_r16 *int64
	}
	TwoPUCCH_Type1_r16  *int64
	TwoPUCCH_Type2_r16  *int64
	TwoPUCCH_Type3_r16  *int64
	TwoPUCCH_Type4_r16  *int64
	Mux_SR_HARQ_ACK_r16 *int64
	Dummy1              *int64
	Dummy2              *int64
	TwoPUCCH_Type5_r16  *int64
	TwoPUCCH_Type6_r16  *int64
	TwoPUCCH_Type7_r16  *int64
	TwoPUCCH_Type8_r16  *int64
	TwoPUCCH_Type9_r16  *int64
	TwoPUCCH_Type10_r16 *int64
	TwoPUCCH_Type11_r16 *int64
	Ul_IntraUE_Mux_r16  *struct {
		Pusch_PreparationLowPriority_r16  int64
		Pusch_PreparationHighPriority_r16 int64
	}
	Ul_FullPwrMode_r16                           *int64
	CrossCarrierSchedulingProcessing_DiffSCS_r16 *struct {
		Scs_15kHz_120kHz_r16 *int64
		Scs_15kHz_60kHz_r16  *int64
		Scs_30kHz_120kHz_r16 *int64
		Scs_15kHz_30kHz_r16  *int64
		Scs_30kHz_60kHz_r16  *int64
		Scs_60kHz_120kHz_r16 *int64
	}
	Ul_FullPwrMode1_r16                           *int64
	Ul_FullPwrMode2_SRSConfig_DiffNumSRSPorts_r16 *int64
	Ul_FullPwrMode2_TPMIGroup_r16                 *struct {
		TwoPorts_r16                 *per.BitString
		FourPortsNonCoherent_r16     *int64
		FourPortsPartialCoherent_r16 *int64
	}
}

func (ie *FeatureSetUplink_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetUplinkV1610Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pusch_RepetitionTypeB_r16 != nil, ie.Ul_CancellationSelfCarrier_r16 != nil, ie.Ul_CancellationCrossCarrier_r16 != nil, ie.Ul_FullPwrMode2_MaxSRS_ResInSet_r16 != nil, ie.CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16 != nil, ie.CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16 != nil, ie.SupportedSRS_PosResources_r16 != nil, ie.IntraFreqDAPS_UL_r16 != nil, ie.IntraBandFreqSeparationUL_v1620 != nil, ie.MultiPUCCH_r16 != nil, ie.TwoPUCCH_Type1_r16 != nil, ie.TwoPUCCH_Type2_r16 != nil, ie.TwoPUCCH_Type3_r16 != nil, ie.TwoPUCCH_Type4_r16 != nil, ie.Mux_SR_HARQ_ACK_r16 != nil, ie.Dummy1 != nil, ie.Dummy2 != nil, ie.TwoPUCCH_Type5_r16 != nil, ie.TwoPUCCH_Type6_r16 != nil, ie.TwoPUCCH_Type7_r16 != nil, ie.TwoPUCCH_Type8_r16 != nil, ie.TwoPUCCH_Type9_r16 != nil, ie.TwoPUCCH_Type10_r16 != nil, ie.TwoPUCCH_Type11_r16 != nil, ie.Ul_IntraUE_Mux_r16 != nil, ie.Ul_FullPwrMode_r16 != nil, ie.CrossCarrierSchedulingProcessing_DiffSCS_r16 != nil, ie.Ul_FullPwrMode1_r16 != nil, ie.Ul_FullPwrMode2_SRSConfig_DiffNumSRSPorts_r16 != nil, ie.Ul_FullPwrMode2_TPMIGroup_r16 != nil}); err != nil {
		return err
	}

	// 2. pusch-RepetitionTypeB-r16: sequence
	{
		if ie.Pusch_RepetitionTypeB_r16 != nil {
			c := ie.Pusch_RepetitionTypeB_r16
			if err := e.EncodeEnumerated(c.MaxNumberPUSCH_Tx_r16, featureSetUplinkV1610PuschRepetitionTypeBR16MaxNumberPUSCHTxR16Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.HoppingScheme_r16, featureSetUplinkV1610PuschRepetitionTypeBR16HoppingSchemeR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. ul-CancellationSelfCarrier-r16: enumerated
	{
		if ie.Ul_CancellationSelfCarrier_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_CancellationSelfCarrier_r16, featureSetUplinkV1610UlCancellationSelfCarrierR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. ul-CancellationCrossCarrier-r16: enumerated
	{
		if ie.Ul_CancellationCrossCarrier_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_CancellationCrossCarrier_r16, featureSetUplinkV1610UlCancellationCrossCarrierR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. ul-FullPwrMode2-MaxSRS-ResInSet-r16: enumerated
	{
		if ie.Ul_FullPwrMode2_MaxSRS_ResInSet_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_FullPwrMode2_MaxSRS_ResInSet_r16, featureSetUplinkV1610UlFullPwrMode2MaxSRSResInSetR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. cbgPUSCH-ProcessingType1-DifferentTB-PerSlot-r16: sequence
	{
		if ie.CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16 != nil {
			c := ie.CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16
			featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Seq := e.NewSequenceEncoder(featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Constraints)
			if err := featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Seq.EncodePreamble([]bool{c.Scs_15kHz_r16 != nil, c.Scs_30kHz_r16 != nil, c.Scs_60kHz_r16 != nil, c.Scs_120kHz_r16 != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_15kHz_r16), featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Scs15kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_30kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_30kHz_r16), featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Scs30kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_60kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_60kHz_r16), featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Scs60kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_120kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_120kHz_r16), featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Scs120kHzR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 7. cbgPUSCH-ProcessingType2-DifferentTB-PerSlot-r16: sequence
	{
		if ie.CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16 != nil {
			c := ie.CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16
			featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Seq := e.NewSequenceEncoder(featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Constraints)
			if err := featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Seq.EncodePreamble([]bool{c.Scs_15kHz_r16 != nil, c.Scs_30kHz_r16 != nil, c.Scs_60kHz_r16 != nil, c.Scs_120kHz_r16 != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_15kHz_r16), featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Scs15kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_30kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_30kHz_r16), featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Scs30kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_60kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_60kHz_r16), featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Scs60kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_120kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_120kHz_r16), featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Scs120kHzR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 8. supportedSRS-PosResources-r16: ref
	{
		if ie.SupportedSRS_PosResources_r16 != nil {
			if err := ie.SupportedSRS_PosResources_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. intraFreqDAPS-UL-r16: sequence
	{
		if ie.IntraFreqDAPS_UL_r16 != nil {
			c := ie.IntraFreqDAPS_UL_r16
			featureSetUplinkV1610IntraFreqDAPSULR16Seq := e.NewSequenceEncoder(featureSetUplinkV1610IntraFreqDAPSULR16Constraints)
			if err := featureSetUplinkV1610IntraFreqDAPSULR16Seq.EncodePreamble([]bool{c.Dummy != nil, c.IntraFreqTwoTAGs_DAPS_r16 != nil, c.Dummy1 != nil, c.Dummy2 != nil, c.Dummy3 != nil}); err != nil {
				return err
			}
			if c.Dummy != nil {
				if err := e.EncodeEnumerated((*c.Dummy), featureSetUplinkV1610IntraFreqDAPSULR16DummyConstraints); err != nil {
					return err
				}
			}
			if c.IntraFreqTwoTAGs_DAPS_r16 != nil {
				if err := e.EncodeEnumerated((*c.IntraFreqTwoTAGs_DAPS_r16), featureSetUplinkV1610IntraFreqDAPSULR16IntraFreqTwoTAGsDAPSR16Constraints); err != nil {
					return err
				}
			}
			if c.Dummy1 != nil {
				if err := e.EncodeEnumerated((*c.Dummy1), featureSetUplinkV1610IntraFreqDAPSULR16Dummy1Constraints); err != nil {
					return err
				}
			}
			if c.Dummy2 != nil {
				if err := e.EncodeEnumerated((*c.Dummy2), featureSetUplinkV1610IntraFreqDAPSULR16Dummy2Constraints); err != nil {
					return err
				}
			}
			if c.Dummy3 != nil {
				if err := e.EncodeEnumerated((*c.Dummy3), featureSetUplinkV1610IntraFreqDAPSULR16Dummy3Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 10. intraBandFreqSeparationUL-v1620: ref
	{
		if ie.IntraBandFreqSeparationUL_v1620 != nil {
			if err := ie.IntraBandFreqSeparationUL_v1620.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. multiPUCCH-r16: sequence
	{
		if ie.MultiPUCCH_r16 != nil {
			c := ie.MultiPUCCH_r16
			featureSetUplinkV1610MultiPUCCHR16Seq := e.NewSequenceEncoder(featureSetUplinkV1610MultiPUCCHR16Constraints)
			if err := featureSetUplinkV1610MultiPUCCHR16Seq.EncodePreamble([]bool{c.Sub_SlotConfig_NCP_r16 != nil, c.Sub_SlotConfig_ECP_r16 != nil}); err != nil {
				return err
			}
			if c.Sub_SlotConfig_NCP_r16 != nil {
				if err := e.EncodeEnumerated((*c.Sub_SlotConfig_NCP_r16), featureSetUplinkV1610MultiPUCCHR16SubSlotConfigNCPR16Constraints); err != nil {
					return err
				}
			}
			if c.Sub_SlotConfig_ECP_r16 != nil {
				if err := e.EncodeEnumerated((*c.Sub_SlotConfig_ECP_r16), featureSetUplinkV1610MultiPUCCHR16SubSlotConfigECPR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 12. twoPUCCH-Type1-r16: enumerated
	{
		if ie.TwoPUCCH_Type1_r16 != nil {
			if err := e.EncodeEnumerated(*ie.TwoPUCCH_Type1_r16, featureSetUplinkV1610TwoPUCCHType1R16Constraints); err != nil {
				return err
			}
		}
	}

	// 13. twoPUCCH-Type2-r16: enumerated
	{
		if ie.TwoPUCCH_Type2_r16 != nil {
			if err := e.EncodeEnumerated(*ie.TwoPUCCH_Type2_r16, featureSetUplinkV1610TwoPUCCHType2R16Constraints); err != nil {
				return err
			}
		}
	}

	// 14. twoPUCCH-Type3-r16: enumerated
	{
		if ie.TwoPUCCH_Type3_r16 != nil {
			if err := e.EncodeEnumerated(*ie.TwoPUCCH_Type3_r16, featureSetUplinkV1610TwoPUCCHType3R16Constraints); err != nil {
				return err
			}
		}
	}

	// 15. twoPUCCH-Type4-r16: enumerated
	{
		if ie.TwoPUCCH_Type4_r16 != nil {
			if err := e.EncodeEnumerated(*ie.TwoPUCCH_Type4_r16, featureSetUplinkV1610TwoPUCCHType4R16Constraints); err != nil {
				return err
			}
		}
	}

	// 16. mux-SR-HARQ-ACK-r16: enumerated
	{
		if ie.Mux_SR_HARQ_ACK_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Mux_SR_HARQ_ACK_r16, featureSetUplinkV1610MuxSRHARQACKR16Constraints); err != nil {
				return err
			}
		}
	}

	// 17. dummy1: enumerated
	{
		if ie.Dummy1 != nil {
			if err := e.EncodeEnumerated(*ie.Dummy1, featureSetUplinkV1610Dummy1Constraints); err != nil {
				return err
			}
		}
	}

	// 18. dummy2: enumerated
	{
		if ie.Dummy2 != nil {
			if err := e.EncodeEnumerated(*ie.Dummy2, featureSetUplinkV1610Dummy2Constraints); err != nil {
				return err
			}
		}
	}

	// 19. twoPUCCH-Type5-r16: enumerated
	{
		if ie.TwoPUCCH_Type5_r16 != nil {
			if err := e.EncodeEnumerated(*ie.TwoPUCCH_Type5_r16, featureSetUplinkV1610TwoPUCCHType5R16Constraints); err != nil {
				return err
			}
		}
	}

	// 20. twoPUCCH-Type6-r16: enumerated
	{
		if ie.TwoPUCCH_Type6_r16 != nil {
			if err := e.EncodeEnumerated(*ie.TwoPUCCH_Type6_r16, featureSetUplinkV1610TwoPUCCHType6R16Constraints); err != nil {
				return err
			}
		}
	}

	// 21. twoPUCCH-Type7-r16: enumerated
	{
		if ie.TwoPUCCH_Type7_r16 != nil {
			if err := e.EncodeEnumerated(*ie.TwoPUCCH_Type7_r16, featureSetUplinkV1610TwoPUCCHType7R16Constraints); err != nil {
				return err
			}
		}
	}

	// 22. twoPUCCH-Type8-r16: enumerated
	{
		if ie.TwoPUCCH_Type8_r16 != nil {
			if err := e.EncodeEnumerated(*ie.TwoPUCCH_Type8_r16, featureSetUplinkV1610TwoPUCCHType8R16Constraints); err != nil {
				return err
			}
		}
	}

	// 23. twoPUCCH-Type9-r16: enumerated
	{
		if ie.TwoPUCCH_Type9_r16 != nil {
			if err := e.EncodeEnumerated(*ie.TwoPUCCH_Type9_r16, featureSetUplinkV1610TwoPUCCHType9R16Constraints); err != nil {
				return err
			}
		}
	}

	// 24. twoPUCCH-Type10-r16: enumerated
	{
		if ie.TwoPUCCH_Type10_r16 != nil {
			if err := e.EncodeEnumerated(*ie.TwoPUCCH_Type10_r16, featureSetUplinkV1610TwoPUCCHType10R16Constraints); err != nil {
				return err
			}
		}
	}

	// 25. twoPUCCH-Type11-r16: enumerated
	{
		if ie.TwoPUCCH_Type11_r16 != nil {
			if err := e.EncodeEnumerated(*ie.TwoPUCCH_Type11_r16, featureSetUplinkV1610TwoPUCCHType11R16Constraints); err != nil {
				return err
			}
		}
	}

	// 26. ul-IntraUE-Mux-r16: sequence
	{
		if ie.Ul_IntraUE_Mux_r16 != nil {
			c := ie.Ul_IntraUE_Mux_r16
			if err := e.EncodeEnumerated(c.Pusch_PreparationLowPriority_r16, featureSetUplinkV1610UlIntraUEMuxR16PuschPreparationLowPriorityR16Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Pusch_PreparationHighPriority_r16, featureSetUplinkV1610UlIntraUEMuxR16PuschPreparationHighPriorityR16Constraints); err != nil {
				return err
			}
		}
	}

	// 27. ul-FullPwrMode-r16: enumerated
	{
		if ie.Ul_FullPwrMode_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_FullPwrMode_r16, featureSetUplinkV1610UlFullPwrModeR16Constraints); err != nil {
				return err
			}
		}
	}

	// 28. crossCarrierSchedulingProcessing-DiffSCS-r16: sequence
	{
		if ie.CrossCarrierSchedulingProcessing_DiffSCS_r16 != nil {
			c := ie.CrossCarrierSchedulingProcessing_DiffSCS_r16
			featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Seq := e.NewSequenceEncoder(featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Constraints)
			if err := featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Seq.EncodePreamble([]bool{c.Scs_15kHz_120kHz_r16 != nil, c.Scs_15kHz_60kHz_r16 != nil, c.Scs_30kHz_120kHz_r16 != nil, c.Scs_15kHz_30kHz_r16 != nil, c.Scs_30kHz_60kHz_r16 != nil, c.Scs_60kHz_120kHz_r16 != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz_120kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_15kHz_120kHz_r16), featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs15kHz120kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_15kHz_60kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_15kHz_60kHz_r16), featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs15kHz60kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_30kHz_120kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_30kHz_120kHz_r16), featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs30kHz120kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_15kHz_30kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_15kHz_30kHz_r16), featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs15kHz30kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_30kHz_60kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_30kHz_60kHz_r16), featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs30kHz60kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_60kHz_120kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_60kHz_120kHz_r16), featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs60kHz120kHzR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 29. ul-FullPwrMode1-r16: enumerated
	{
		if ie.Ul_FullPwrMode1_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_FullPwrMode1_r16, featureSetUplinkV1610UlFullPwrMode1R16Constraints); err != nil {
				return err
			}
		}
	}

	// 30. ul-FullPwrMode2-SRSConfig-diffNumSRSPorts-r16: enumerated
	{
		if ie.Ul_FullPwrMode2_SRSConfig_DiffNumSRSPorts_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_FullPwrMode2_SRSConfig_DiffNumSRSPorts_r16, featureSetUplinkV1610UlFullPwrMode2SRSConfigDiffNumSRSPortsR16Constraints); err != nil {
				return err
			}
		}
	}

	// 31. ul-FullPwrMode2-TPMIGroup-r16: sequence
	{
		if ie.Ul_FullPwrMode2_TPMIGroup_r16 != nil {
			c := ie.Ul_FullPwrMode2_TPMIGroup_r16
			featureSetUplinkV1610UlFullPwrMode2TPMIGroupR16Seq := e.NewSequenceEncoder(featureSetUplinkV1610UlFullPwrMode2TPMIGroupR16Constraints)
			if err := featureSetUplinkV1610UlFullPwrMode2TPMIGroupR16Seq.EncodePreamble([]bool{c.TwoPorts_r16 != nil, c.FourPortsNonCoherent_r16 != nil, c.FourPortsPartialCoherent_r16 != nil}); err != nil {
				return err
			}
			if c.TwoPorts_r16 != nil {
				if err := e.EncodeBitString((*c.TwoPorts_r16), per.FixedSize(2)); err != nil {
					return err
				}
			}
			if c.FourPortsNonCoherent_r16 != nil {
				if err := e.EncodeEnumerated((*c.FourPortsNonCoherent_r16), featureSetUplinkV1610UlFullPwrMode2TPMIGroupR16FourPortsNonCoherentR16Constraints); err != nil {
					return err
				}
			}
			if c.FourPortsPartialCoherent_r16 != nil {
				if err := e.EncodeEnumerated((*c.FourPortsPartialCoherent_r16), featureSetUplinkV1610UlFullPwrMode2TPMIGroupR16FourPortsPartialCoherentR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *FeatureSetUplink_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetUplinkV1610Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pusch-RepetitionTypeB-r16: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.Pusch_RepetitionTypeB_r16 = &struct {
				MaxNumberPUSCH_Tx_r16 int64
				HoppingScheme_r16     int64
			}{}
			c := ie.Pusch_RepetitionTypeB_r16
			{
				v, err := d.DecodeEnumerated(featureSetUplinkV1610PuschRepetitionTypeBR16MaxNumberPUSCHTxR16Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberPUSCH_Tx_r16 = v
			}
			{
				v, err := d.DecodeEnumerated(featureSetUplinkV1610PuschRepetitionTypeBR16HoppingSchemeR16Constraints)
				if err != nil {
					return err
				}
				c.HoppingScheme_r16 = v
			}
		}
	}

	// 3. ul-CancellationSelfCarrier-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1610UlCancellationSelfCarrierR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_CancellationSelfCarrier_r16 = &idx
		}
	}

	// 4. ul-CancellationCrossCarrier-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1610UlCancellationCrossCarrierR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_CancellationCrossCarrier_r16 = &idx
		}
	}

	// 5. ul-FullPwrMode2-MaxSRS-ResInSet-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1610UlFullPwrMode2MaxSRSResInSetR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_FullPwrMode2_MaxSRS_ResInSet_r16 = &idx
		}
	}

	// 6. cbgPUSCH-ProcessingType1-DifferentTB-PerSlot-r16: sequence
	{
		if seq.IsComponentPresent(4) {
			ie.CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16 = &struct {
				Scs_15kHz_r16  *int64
				Scs_30kHz_r16  *int64
				Scs_60kHz_r16  *int64
				Scs_120kHz_r16 *int64
			}{}
			c := ie.CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16
			featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Seq := d.NewSequenceDecoder(featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Constraints)
			if err := featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Seq.IsComponentPresent(0) {
				c.Scs_15kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Scs15kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_15kHz_r16) = v
			}
			if featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Seq.IsComponentPresent(1) {
				c.Scs_30kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Scs30kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_30kHz_r16) = v
			}
			if featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Seq.IsComponentPresent(2) {
				c.Scs_60kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Scs60kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_60kHz_r16) = v
			}
			if featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Seq.IsComponentPresent(3) {
				c.Scs_120kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610CbgPUSCHProcessingType1DifferentTBPerSlotR16Scs120kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_120kHz_r16) = v
			}
		}
	}

	// 7. cbgPUSCH-ProcessingType2-DifferentTB-PerSlot-r16: sequence
	{
		if seq.IsComponentPresent(5) {
			ie.CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16 = &struct {
				Scs_15kHz_r16  *int64
				Scs_30kHz_r16  *int64
				Scs_60kHz_r16  *int64
				Scs_120kHz_r16 *int64
			}{}
			c := ie.CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16
			featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Seq := d.NewSequenceDecoder(featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Constraints)
			if err := featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Seq.IsComponentPresent(0) {
				c.Scs_15kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Scs15kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_15kHz_r16) = v
			}
			if featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Seq.IsComponentPresent(1) {
				c.Scs_30kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Scs30kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_30kHz_r16) = v
			}
			if featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Seq.IsComponentPresent(2) {
				c.Scs_60kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Scs60kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_60kHz_r16) = v
			}
			if featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Seq.IsComponentPresent(3) {
				c.Scs_120kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610CbgPUSCHProcessingType2DifferentTBPerSlotR16Scs120kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_120kHz_r16) = v
			}
		}
	}

	// 8. supportedSRS-PosResources-r16: ref
	{
		if seq.IsComponentPresent(6) {
			ie.SupportedSRS_PosResources_r16 = new(SRS_AllPosResources_r16)
			if err := ie.SupportedSRS_PosResources_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. intraFreqDAPS-UL-r16: sequence
	{
		if seq.IsComponentPresent(7) {
			ie.IntraFreqDAPS_UL_r16 = &struct {
				Dummy                     *int64
				IntraFreqTwoTAGs_DAPS_r16 *int64
				Dummy1                    *int64
				Dummy2                    *int64
				Dummy3                    *int64
			}{}
			c := ie.IntraFreqDAPS_UL_r16
			featureSetUplinkV1610IntraFreqDAPSULR16Seq := d.NewSequenceDecoder(featureSetUplinkV1610IntraFreqDAPSULR16Constraints)
			if err := featureSetUplinkV1610IntraFreqDAPSULR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetUplinkV1610IntraFreqDAPSULR16Seq.IsComponentPresent(0) {
				c.Dummy = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610IntraFreqDAPSULR16DummyConstraints)
				if err != nil {
					return err
				}
				(*c.Dummy) = v
			}
			if featureSetUplinkV1610IntraFreqDAPSULR16Seq.IsComponentPresent(1) {
				c.IntraFreqTwoTAGs_DAPS_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610IntraFreqDAPSULR16IntraFreqTwoTAGsDAPSR16Constraints)
				if err != nil {
					return err
				}
				(*c.IntraFreqTwoTAGs_DAPS_r16) = v
			}
			if featureSetUplinkV1610IntraFreqDAPSULR16Seq.IsComponentPresent(2) {
				c.Dummy1 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610IntraFreqDAPSULR16Dummy1Constraints)
				if err != nil {
					return err
				}
				(*c.Dummy1) = v
			}
			if featureSetUplinkV1610IntraFreqDAPSULR16Seq.IsComponentPresent(3) {
				c.Dummy2 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610IntraFreqDAPSULR16Dummy2Constraints)
				if err != nil {
					return err
				}
				(*c.Dummy2) = v
			}
			if featureSetUplinkV1610IntraFreqDAPSULR16Seq.IsComponentPresent(4) {
				c.Dummy3 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610IntraFreqDAPSULR16Dummy3Constraints)
				if err != nil {
					return err
				}
				(*c.Dummy3) = v
			}
		}
	}

	// 10. intraBandFreqSeparationUL-v1620: ref
	{
		if seq.IsComponentPresent(8) {
			ie.IntraBandFreqSeparationUL_v1620 = new(FreqSeparationClassUL_v1620)
			if err := ie.IntraBandFreqSeparationUL_v1620.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. multiPUCCH-r16: sequence
	{
		if seq.IsComponentPresent(9) {
			ie.MultiPUCCH_r16 = &struct {
				Sub_SlotConfig_NCP_r16 *int64
				Sub_SlotConfig_ECP_r16 *int64
			}{}
			c := ie.MultiPUCCH_r16
			featureSetUplinkV1610MultiPUCCHR16Seq := d.NewSequenceDecoder(featureSetUplinkV1610MultiPUCCHR16Constraints)
			if err := featureSetUplinkV1610MultiPUCCHR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetUplinkV1610MultiPUCCHR16Seq.IsComponentPresent(0) {
				c.Sub_SlotConfig_NCP_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610MultiPUCCHR16SubSlotConfigNCPR16Constraints)
				if err != nil {
					return err
				}
				(*c.Sub_SlotConfig_NCP_r16) = v
			}
			if featureSetUplinkV1610MultiPUCCHR16Seq.IsComponentPresent(1) {
				c.Sub_SlotConfig_ECP_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610MultiPUCCHR16SubSlotConfigECPR16Constraints)
				if err != nil {
					return err
				}
				(*c.Sub_SlotConfig_ECP_r16) = v
			}
		}
	}

	// 12. twoPUCCH-Type1-r16: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1610TwoPUCCHType1R16Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUCCH_Type1_r16 = &idx
		}
	}

	// 13. twoPUCCH-Type2-r16: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1610TwoPUCCHType2R16Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUCCH_Type2_r16 = &idx
		}
	}

	// 14. twoPUCCH-Type3-r16: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1610TwoPUCCHType3R16Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUCCH_Type3_r16 = &idx
		}
	}

	// 15. twoPUCCH-Type4-r16: enumerated
	{
		if seq.IsComponentPresent(13) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1610TwoPUCCHType4R16Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUCCH_Type4_r16 = &idx
		}
	}

	// 16. mux-SR-HARQ-ACK-r16: enumerated
	{
		if seq.IsComponentPresent(14) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1610MuxSRHARQACKR16Constraints)
			if err != nil {
				return err
			}
			ie.Mux_SR_HARQ_ACK_r16 = &idx
		}
	}

	// 17. dummy1: enumerated
	{
		if seq.IsComponentPresent(15) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1610Dummy1Constraints)
			if err != nil {
				return err
			}
			ie.Dummy1 = &idx
		}
	}

	// 18. dummy2: enumerated
	{
		if seq.IsComponentPresent(16) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1610Dummy2Constraints)
			if err != nil {
				return err
			}
			ie.Dummy2 = &idx
		}
	}

	// 19. twoPUCCH-Type5-r16: enumerated
	{
		if seq.IsComponentPresent(17) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1610TwoPUCCHType5R16Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUCCH_Type5_r16 = &idx
		}
	}

	// 20. twoPUCCH-Type6-r16: enumerated
	{
		if seq.IsComponentPresent(18) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1610TwoPUCCHType6R16Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUCCH_Type6_r16 = &idx
		}
	}

	// 21. twoPUCCH-Type7-r16: enumerated
	{
		if seq.IsComponentPresent(19) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1610TwoPUCCHType7R16Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUCCH_Type7_r16 = &idx
		}
	}

	// 22. twoPUCCH-Type8-r16: enumerated
	{
		if seq.IsComponentPresent(20) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1610TwoPUCCHType8R16Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUCCH_Type8_r16 = &idx
		}
	}

	// 23. twoPUCCH-Type9-r16: enumerated
	{
		if seq.IsComponentPresent(21) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1610TwoPUCCHType9R16Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUCCH_Type9_r16 = &idx
		}
	}

	// 24. twoPUCCH-Type10-r16: enumerated
	{
		if seq.IsComponentPresent(22) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1610TwoPUCCHType10R16Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUCCH_Type10_r16 = &idx
		}
	}

	// 25. twoPUCCH-Type11-r16: enumerated
	{
		if seq.IsComponentPresent(23) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1610TwoPUCCHType11R16Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUCCH_Type11_r16 = &idx
		}
	}

	// 26. ul-IntraUE-Mux-r16: sequence
	{
		if seq.IsComponentPresent(24) {
			ie.Ul_IntraUE_Mux_r16 = &struct {
				Pusch_PreparationLowPriority_r16  int64
				Pusch_PreparationHighPriority_r16 int64
			}{}
			c := ie.Ul_IntraUE_Mux_r16
			{
				v, err := d.DecodeEnumerated(featureSetUplinkV1610UlIntraUEMuxR16PuschPreparationLowPriorityR16Constraints)
				if err != nil {
					return err
				}
				c.Pusch_PreparationLowPriority_r16 = v
			}
			{
				v, err := d.DecodeEnumerated(featureSetUplinkV1610UlIntraUEMuxR16PuschPreparationHighPriorityR16Constraints)
				if err != nil {
					return err
				}
				c.Pusch_PreparationHighPriority_r16 = v
			}
		}
	}

	// 27. ul-FullPwrMode-r16: enumerated
	{
		if seq.IsComponentPresent(25) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1610UlFullPwrModeR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_FullPwrMode_r16 = &idx
		}
	}

	// 28. crossCarrierSchedulingProcessing-DiffSCS-r16: sequence
	{
		if seq.IsComponentPresent(26) {
			ie.CrossCarrierSchedulingProcessing_DiffSCS_r16 = &struct {
				Scs_15kHz_120kHz_r16 *int64
				Scs_15kHz_60kHz_r16  *int64
				Scs_30kHz_120kHz_r16 *int64
				Scs_15kHz_30kHz_r16  *int64
				Scs_30kHz_60kHz_r16  *int64
				Scs_60kHz_120kHz_r16 *int64
			}{}
			c := ie.CrossCarrierSchedulingProcessing_DiffSCS_r16
			featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Seq := d.NewSequenceDecoder(featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Constraints)
			if err := featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Seq.IsComponentPresent(0) {
				c.Scs_15kHz_120kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs15kHz120kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_15kHz_120kHz_r16) = v
			}
			if featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Seq.IsComponentPresent(1) {
				c.Scs_15kHz_60kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs15kHz60kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_15kHz_60kHz_r16) = v
			}
			if featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Seq.IsComponentPresent(2) {
				c.Scs_30kHz_120kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs30kHz120kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_30kHz_120kHz_r16) = v
			}
			if featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Seq.IsComponentPresent(3) {
				c.Scs_15kHz_30kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs15kHz30kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_15kHz_30kHz_r16) = v
			}
			if featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Seq.IsComponentPresent(4) {
				c.Scs_30kHz_60kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs30kHz60kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_30kHz_60kHz_r16) = v
			}
			if featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Seq.IsComponentPresent(5) {
				c.Scs_60kHz_120kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs60kHz120kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_60kHz_120kHz_r16) = v
			}
		}
	}

	// 29. ul-FullPwrMode1-r16: enumerated
	{
		if seq.IsComponentPresent(27) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1610UlFullPwrMode1R16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_FullPwrMode1_r16 = &idx
		}
	}

	// 30. ul-FullPwrMode2-SRSConfig-diffNumSRSPorts-r16: enumerated
	{
		if seq.IsComponentPresent(28) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1610UlFullPwrMode2SRSConfigDiffNumSRSPortsR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_FullPwrMode2_SRSConfig_DiffNumSRSPorts_r16 = &idx
		}
	}

	// 31. ul-FullPwrMode2-TPMIGroup-r16: sequence
	{
		if seq.IsComponentPresent(29) {
			ie.Ul_FullPwrMode2_TPMIGroup_r16 = &struct {
				TwoPorts_r16                 *per.BitString
				FourPortsNonCoherent_r16     *int64
				FourPortsPartialCoherent_r16 *int64
			}{}
			c := ie.Ul_FullPwrMode2_TPMIGroup_r16
			featureSetUplinkV1610UlFullPwrMode2TPMIGroupR16Seq := d.NewSequenceDecoder(featureSetUplinkV1610UlFullPwrMode2TPMIGroupR16Constraints)
			if err := featureSetUplinkV1610UlFullPwrMode2TPMIGroupR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetUplinkV1610UlFullPwrMode2TPMIGroupR16Seq.IsComponentPresent(0) {
				c.TwoPorts_r16 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(2))
				if err != nil {
					return err
				}
				(*c.TwoPorts_r16) = v
			}
			if featureSetUplinkV1610UlFullPwrMode2TPMIGroupR16Seq.IsComponentPresent(1) {
				c.FourPortsNonCoherent_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610UlFullPwrMode2TPMIGroupR16FourPortsNonCoherentR16Constraints)
				if err != nil {
					return err
				}
				(*c.FourPortsNonCoherent_r16) = v
			}
			if featureSetUplinkV1610UlFullPwrMode2TPMIGroupR16Seq.IsComponentPresent(2) {
				c.FourPortsPartialCoherent_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1610UlFullPwrMode2TPMIGroupR16FourPortsPartialCoherentR16Constraints)
				if err != nil {
					return err
				}
				(*c.FourPortsPartialCoherent_r16) = v
			}
		}
	}

	return nil
}
