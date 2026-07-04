// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookParameterseType2DopplerExt-r19 (line 19202).

var codebookParameterseType2DopplerExtR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "eType2Doppler-64PortExt-r19"},
		{Name: "eType2Doppler-48PortExt-r19", Optional: true},
		{Name: "eType2Doppler-128PortExt-r19", Optional: true},
		{Name: "eType2DopplerN4Ext-r19", Optional: true},
		{Name: "ddUnitSize-A-CSI-RS-CMR-Ext-r19", Optional: true},
		{Name: "maxNumberAperiodicCSI-RS-ResourceExt-r19", Optional: true},
		{Name: "eType2DopplerR2Ext-r19", Optional: true},
		{Name: "eType2DopplerX1Ext-r19", Optional: true},
		{Name: "eType2DopplerX2Ext-r19", Optional: true},
		{Name: "eType2DopplerL-N4D1Ext-r19", Optional: true},
		{Name: "eType2DopplerL6Ext-r19", Optional: true},
		{Name: "eType2DopplerR3R4Ext-r19", Optional: true},
		{Name: "eType2DopplerProcessingTimelineExt-r19", Optional: true},
		{Name: "eType2MaxPeriodicityCMR-r19", Optional: true},
	},
}

const (
	CodebookParameterseType2DopplerExt_r19_DdUnitSize_A_CSI_RS_CMR_Ext_r19_Supported = 0
)

var codebookParameterseType2DopplerExtR19DdUnitSizeACSIRSCMRExtR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_DdUnitSize_A_CSI_RS_CMR_Ext_r19_Supported},
}

const (
	CodebookParameterseType2DopplerExt_r19_MaxNumberAperiodicCSI_RS_ResourceExt_r19_N4  = 0
	CodebookParameterseType2DopplerExt_r19_MaxNumberAperiodicCSI_RS_ResourceExt_r19_N8  = 1
	CodebookParameterseType2DopplerExt_r19_MaxNumberAperiodicCSI_RS_ResourceExt_r19_N12 = 2
)

var codebookParameterseType2DopplerExtR19MaxNumberAperiodicCSIRSResourceExtR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_MaxNumberAperiodicCSI_RS_ResourceExt_r19_N4, CodebookParameterseType2DopplerExt_r19_MaxNumberAperiodicCSI_RS_ResourceExt_r19_N8, CodebookParameterseType2DopplerExt_r19_MaxNumberAperiodicCSI_RS_ResourceExt_r19_N12},
}

var codebookParameterseType2DopplerExtR19EType2DopplerR2ExtR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParameterseType2DopplerExt_r19_EType2DopplerX1Ext_r19_Supported = 0
)

var codebookParameterseType2DopplerExtR19EType2DopplerX1ExtR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_EType2DopplerX1Ext_r19_Supported},
}

const (
	CodebookParameterseType2DopplerExt_r19_EType2DopplerX2Ext_r19_Supported = 0
)

var codebookParameterseType2DopplerExtR19EType2DopplerX2ExtR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_EType2DopplerX2Ext_r19_Supported},
}

const (
	CodebookParameterseType2DopplerExt_r19_EType2DopplerL_N4D1Ext_r19_Supported = 0
)

var codebookParameterseType2DopplerExtR19EType2DopplerLN4D1ExtR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_EType2DopplerL_N4D1Ext_r19_Supported},
}

const (
	CodebookParameterseType2DopplerExt_r19_EType2DopplerL6Ext_r19_Supported = 0
)

var codebookParameterseType2DopplerExtR19EType2DopplerL6ExtR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_EType2DopplerL6Ext_r19_Supported},
}

const (
	CodebookParameterseType2DopplerExt_r19_EType2DopplerR3R4Ext_r19_Supported = 0
)

var codebookParameterseType2DopplerExtR19EType2DopplerR3R4ExtR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_EType2DopplerR3R4Ext_r19_Supported},
}

const (
	CodebookParameterseType2DopplerExt_r19_EType2MaxPeriodicityCMR_r19_Sl4  = 0
	CodebookParameterseType2DopplerExt_r19_EType2MaxPeriodicityCMR_r19_Sl5  = 1
	CodebookParameterseType2DopplerExt_r19_EType2MaxPeriodicityCMR_r19_Sl8  = 2
	CodebookParameterseType2DopplerExt_r19_EType2MaxPeriodicityCMR_r19_Sl10 = 3
	CodebookParameterseType2DopplerExt_r19_EType2MaxPeriodicityCMR_r19_Sl20 = 4
)

var codebookParameterseType2DopplerExtR19EType2MaxPeriodicityCMRR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_EType2MaxPeriodicityCMR_r19_Sl4, CodebookParameterseType2DopplerExt_r19_EType2MaxPeriodicityCMR_r19_Sl5, CodebookParameterseType2DopplerExt_r19_EType2MaxPeriodicityCMR_r19_Sl8, CodebookParameterseType2DopplerExt_r19_EType2MaxPeriodicityCMR_r19_Sl10, CodebookParameterseType2DopplerExt_r19_EType2MaxPeriodicityCMR_r19_Sl20},
}

var codebookParameterseType2DopplerExtR19EType2Doppler64PortExtR19SupportedCSIRSResourceExtListR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParameterseType2DopplerExt_r19_EType2Doppler_64PortExt_r19_ProcessingCapability_r19_Cap1 = 0
	CodebookParameterseType2DopplerExt_r19_EType2Doppler_64PortExt_r19_ProcessingCapability_r19_Cap2 = 1
)

var codebookParameterseType2DopplerExtR19EType2Doppler64PortExtR19ProcessingCapabilityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_EType2Doppler_64PortExt_r19_ProcessingCapability_r19_Cap1, CodebookParameterseType2DopplerExt_r19_EType2Doppler_64PortExt_r19_ProcessingCapability_r19_Cap2},
}

const (
	CodebookParameterseType2DopplerExt_r19_EType2Doppler_64PortExt_r19_Scalingfactor_r19_N1 = 0
	CodebookParameterseType2DopplerExt_r19_EType2Doppler_64PortExt_r19_Scalingfactor_r19_N2 = 1
	CodebookParameterseType2DopplerExt_r19_EType2Doppler_64PortExt_r19_Scalingfactor_r19_N4 = 2
)

var codebookParameterseType2DopplerExtR19EType2Doppler64PortExtR19ScalingfactorR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_EType2Doppler_64PortExt_r19_Scalingfactor_r19_N1, CodebookParameterseType2DopplerExt_r19_EType2Doppler_64PortExt_r19_Scalingfactor_r19_N2, CodebookParameterseType2DopplerExt_r19_EType2Doppler_64PortExt_r19_Scalingfactor_r19_N4},
}

const (
	CodebookParameterseType2DopplerExt_r19_EType2Doppler_64PortExt_r19_MaxNumberResource_r19_N2 = 0
	CodebookParameterseType2DopplerExt_r19_EType2Doppler_64PortExt_r19_MaxNumberResource_r19_N4 = 1
)

var codebookParameterseType2DopplerExtR19EType2Doppler64PortExtR19MaxNumberResourceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_EType2Doppler_64PortExt_r19_MaxNumberResource_r19_N2, CodebookParameterseType2DopplerExt_r19_EType2Doppler_64PortExt_r19_MaxNumberResource_r19_N4},
}

var codebookParameterseType2DopplerExtR19EType2Doppler64PortExtR19SupportedCSIRSResourceListPerCCR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParameterseType2DopplerExtR19EType2Doppler48PortExtR19SupportedCSIRSResourceExtListR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParameterseType2DopplerExt_r19_EType2Doppler_48PortExt_r19_ProcessingCapability_r19_Cap1 = 0
	CodebookParameterseType2DopplerExt_r19_EType2Doppler_48PortExt_r19_ProcessingCapability_r19_Cap2 = 1
)

var codebookParameterseType2DopplerExtR19EType2Doppler48PortExtR19ProcessingCapabilityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_EType2Doppler_48PortExt_r19_ProcessingCapability_r19_Cap1, CodebookParameterseType2DopplerExt_r19_EType2Doppler_48PortExt_r19_ProcessingCapability_r19_Cap2},
}

const (
	CodebookParameterseType2DopplerExt_r19_EType2Doppler_48PortExt_r19_Scalingfactor_r19_N1 = 0
	CodebookParameterseType2DopplerExt_r19_EType2Doppler_48PortExt_r19_Scalingfactor_r19_N2 = 1
	CodebookParameterseType2DopplerExt_r19_EType2Doppler_48PortExt_r19_Scalingfactor_r19_N4 = 2
)

var codebookParameterseType2DopplerExtR19EType2Doppler48PortExtR19ScalingfactorR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_EType2Doppler_48PortExt_r19_Scalingfactor_r19_N1, CodebookParameterseType2DopplerExt_r19_EType2Doppler_48PortExt_r19_Scalingfactor_r19_N2, CodebookParameterseType2DopplerExt_r19_EType2Doppler_48PortExt_r19_Scalingfactor_r19_N4},
}

const (
	CodebookParameterseType2DopplerExt_r19_EType2Doppler_48PortExt_r19_MaxNumberResource_r19_N2 = 0
	CodebookParameterseType2DopplerExt_r19_EType2Doppler_48PortExt_r19_MaxNumberResource_r19_N3 = 1
)

var codebookParameterseType2DopplerExtR19EType2Doppler48PortExtR19MaxNumberResourceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_EType2Doppler_48PortExt_r19_MaxNumberResource_r19_N2, CodebookParameterseType2DopplerExt_r19_EType2Doppler_48PortExt_r19_MaxNumberResource_r19_N3},
}

var codebookParameterseType2DopplerExtR19EType2Doppler48PortExtR19SupportedCSIRSResourceListPerCCR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParameterseType2DopplerExtR19EType2Doppler128PortExtR19SupportedCSIRSResourceExtListR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParameterseType2DopplerExt_r19_EType2Doppler_128PortExt_r19_ProcessingCapability_r19_Cap1 = 0
	CodebookParameterseType2DopplerExt_r19_EType2Doppler_128PortExt_r19_ProcessingCapability_r19_Cap2 = 1
)

var codebookParameterseType2DopplerExtR19EType2Doppler128PortExtR19ProcessingCapabilityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_EType2Doppler_128PortExt_r19_ProcessingCapability_r19_Cap1, CodebookParameterseType2DopplerExt_r19_EType2Doppler_128PortExt_r19_ProcessingCapability_r19_Cap2},
}

const (
	CodebookParameterseType2DopplerExt_r19_EType2Doppler_128PortExt_r19_Scalingfactor_r19_N1 = 0
	CodebookParameterseType2DopplerExt_r19_EType2Doppler_128PortExt_r19_Scalingfactor_r19_N2 = 1
	CodebookParameterseType2DopplerExt_r19_EType2Doppler_128PortExt_r19_Scalingfactor_r19_N4 = 2
)

var codebookParameterseType2DopplerExtR19EType2Doppler128PortExtR19ScalingfactorR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_EType2Doppler_128PortExt_r19_Scalingfactor_r19_N1, CodebookParameterseType2DopplerExt_r19_EType2Doppler_128PortExt_r19_Scalingfactor_r19_N2, CodebookParameterseType2DopplerExt_r19_EType2Doppler_128PortExt_r19_Scalingfactor_r19_N4},
}

var codebookParameterseType2DopplerExtR19EType2Doppler128PortExtR19SupportedCSIRSResourceListPerCCR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParameterseType2DopplerExtR19EType2DopplerN4ExtR19SupportedCSIRSReportSettingList1R19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParameterseType2DopplerExtR19EType2DopplerN4ExtR19SupportedCSIRSReportSettingList2R19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz", Optional: true},
		{Name: "scs-30kHz", Optional: true},
		{Name: "scs-60kHz", Optional: true},
		{Name: "scs-120kHz", Optional: true},
	},
}

const (
	CodebookParameterseType2DopplerExt_r19_EType2DopplerProcessingTimelineExt_r19_ValueW_r19_Scs_15kHz_Value1 = 0
	CodebookParameterseType2DopplerExt_r19_EType2DopplerProcessingTimelineExt_r19_ValueW_r19_Scs_15kHz_Value2 = 1
)

var codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Scs15kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_EType2DopplerProcessingTimelineExt_r19_ValueW_r19_Scs_15kHz_Value1, CodebookParameterseType2DopplerExt_r19_EType2DopplerProcessingTimelineExt_r19_ValueW_r19_Scs_15kHz_Value2},
}

const (
	CodebookParameterseType2DopplerExt_r19_EType2DopplerProcessingTimelineExt_r19_ValueW_r19_Scs_30kHz_Value1 = 0
	CodebookParameterseType2DopplerExt_r19_EType2DopplerProcessingTimelineExt_r19_ValueW_r19_Scs_30kHz_Value2 = 1
)

var codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Scs30kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_EType2DopplerProcessingTimelineExt_r19_ValueW_r19_Scs_30kHz_Value1, CodebookParameterseType2DopplerExt_r19_EType2DopplerProcessingTimelineExt_r19_ValueW_r19_Scs_30kHz_Value2},
}

const (
	CodebookParameterseType2DopplerExt_r19_EType2DopplerProcessingTimelineExt_r19_ValueW_r19_Scs_60kHz_Value1 = 0
	CodebookParameterseType2DopplerExt_r19_EType2DopplerProcessingTimelineExt_r19_ValueW_r19_Scs_60kHz_Value2 = 1
)

var codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Scs60kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_EType2DopplerProcessingTimelineExt_r19_ValueW_r19_Scs_60kHz_Value1, CodebookParameterseType2DopplerExt_r19_EType2DopplerProcessingTimelineExt_r19_ValueW_r19_Scs_60kHz_Value2},
}

const (
	CodebookParameterseType2DopplerExt_r19_EType2DopplerProcessingTimelineExt_r19_ValueW_r19_Scs_120kHz_Value1 = 0
	CodebookParameterseType2DopplerExt_r19_EType2DopplerProcessingTimelineExt_r19_ValueW_r19_Scs_120kHz_Value2 = 1
)

var codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Scs120kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_EType2DopplerProcessingTimelineExt_r19_ValueW_r19_Scs_120kHz_Value1, CodebookParameterseType2DopplerExt_r19_EType2DopplerProcessingTimelineExt_r19_ValueW_r19_Scs_120kHz_Value2},
}

const (
	CodebookParameterseType2DopplerExt_r19_EType2DopplerProcessingTimelineExt_r19_TimeRelaxation_r19_Cap1 = 0
	CodebookParameterseType2DopplerExt_r19_EType2DopplerProcessingTimelineExt_r19_TimeRelaxation_r19_Cap2 = 1
)

var codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19TimeRelaxationR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2DopplerExt_r19_EType2DopplerProcessingTimelineExt_r19_TimeRelaxation_r19_Cap1, CodebookParameterseType2DopplerExt_r19_EType2DopplerProcessingTimelineExt_r19_TimeRelaxation_r19_Cap2},
}

type CodebookParameterseType2DopplerExt_r19 struct {
	EType2Doppler_64PortExt_r19 struct {
		SupportedCSI_RS_ResourceExtList_r19   []int64
		ProcessingCapability_r19              int64
		ValueY_P_SP_CSI_RS_r19                int64
		ValueY_A_CSI_RS_r19                   int64
		Scalingfactor_r19                     int64
		MaxNumberResource_r19                 int64
		SupportedCSI_RS_ResourceListPerCC_r19 []int64
	}
	EType2Doppler_48PortExt_r19 *struct {
		SupportedCSI_RS_ResourceExtList_r19   []int64
		ProcessingCapability_r19              int64
		ValueY_P_SP_CSI_RS_r19                int64
		ValueY_A_CSI_RS_r19                   int64
		Scalingfactor_r19                     int64
		MaxNumberResource_r19                 int64
		SupportedCSI_RS_ResourceListPerCC_r19 []int64
	}
	EType2Doppler_128PortExt_r19 *struct {
		SupportedCSI_RS_ResourceExtList_r19   []int64
		ProcessingCapability_r19              int64
		ValueY_P_SP_CSI_RS_r19                int64
		ValueY_A_CSI_RS_r19                   int64
		Scalingfactor_r19                     int64
		SupportedCSI_RS_ResourceListPerCC_r19 []int64
	}
	EType2DopplerN4Ext_r19 *struct {
		SupportedCSI_RS_ReportSettingList1_r19 []SupportedCSI_RS_ReportSettingExt_r19
		SupportedCSI_RS_ReportSettingList2_r19 []SupportedCSI_RS_ReportSettingExt_r19
	}
	DdUnitSize_A_CSI_RS_CMR_Ext_r19          *int64
	MaxNumberAperiodicCSI_RS_ResourceExt_r19 *int64
	EType2DopplerR2Ext_r19                   []int64
	EType2DopplerX1Ext_r19                   *int64
	EType2DopplerX2Ext_r19                   *int64
	EType2DopplerL_N4D1Ext_r19               *int64
	EType2DopplerL6Ext_r19                   *int64
	EType2DopplerR3R4Ext_r19                 *int64
	EType2DopplerProcessingTimelineExt_r19   *struct {
		ValueW_r19 struct {
			Scs_15kHz  *int64
			Scs_30kHz  *int64
			Scs_60kHz  *int64
			Scs_120kHz *int64
		}
		TimeRelaxation_r19 int64
	}
	EType2MaxPeriodicityCMR_r19 *int64
}

func (ie *CodebookParameterseType2DopplerExt_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookParameterseType2DopplerExtR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.EType2Doppler_48PortExt_r19 != nil, ie.EType2Doppler_128PortExt_r19 != nil, ie.EType2DopplerN4Ext_r19 != nil, ie.DdUnitSize_A_CSI_RS_CMR_Ext_r19 != nil, ie.MaxNumberAperiodicCSI_RS_ResourceExt_r19 != nil, ie.EType2DopplerR2Ext_r19 != nil, ie.EType2DopplerX1Ext_r19 != nil, ie.EType2DopplerX2Ext_r19 != nil, ie.EType2DopplerL_N4D1Ext_r19 != nil, ie.EType2DopplerL6Ext_r19 != nil, ie.EType2DopplerR3R4Ext_r19 != nil, ie.EType2DopplerProcessingTimelineExt_r19 != nil, ie.EType2MaxPeriodicityCMR_r19 != nil}); err != nil {
		return err
	}

	// 2. eType2Doppler-64PortExt-r19: sequence
	{
		{
			c := &ie.EType2Doppler_64PortExt_r19
			{
				seqOf := e.NewSequenceOfEncoder(codebookParameterseType2DopplerExtR19EType2Doppler64PortExtR19SupportedCSIRSResourceExtListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceExtList_r19))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceExtList_r19 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceExtList_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeEnumerated(c.ProcessingCapability_r19, codebookParameterseType2DopplerExtR19EType2Doppler64PortExtR19ProcessingCapabilityR19Constraints); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.ValueY_P_SP_CSI_RS_r19, per.Constrained(1, 3)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.ValueY_A_CSI_RS_r19, per.Constrained(1, 3)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Scalingfactor_r19, codebookParameterseType2DopplerExtR19EType2Doppler64PortExtR19ScalingfactorR19Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MaxNumberResource_r19, codebookParameterseType2DopplerExtR19EType2Doppler64PortExtR19MaxNumberResourceR19Constraints); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(codebookParameterseType2DopplerExtR19EType2Doppler64PortExtR19SupportedCSIRSResourceListPerCCR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceListPerCC_r19))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceListPerCC_r19 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceListPerCC_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 3. eType2Doppler-48PortExt-r19: sequence
	{
		if ie.EType2Doppler_48PortExt_r19 != nil {
			c := ie.EType2Doppler_48PortExt_r19
			{
				seqOf := e.NewSequenceOfEncoder(codebookParameterseType2DopplerExtR19EType2Doppler48PortExtR19SupportedCSIRSResourceExtListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceExtList_r19))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceExtList_r19 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceExtList_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeEnumerated(c.ProcessingCapability_r19, codebookParameterseType2DopplerExtR19EType2Doppler48PortExtR19ProcessingCapabilityR19Constraints); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.ValueY_P_SP_CSI_RS_r19, per.Constrained(1, 3)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.ValueY_A_CSI_RS_r19, per.Constrained(1, 3)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Scalingfactor_r19, codebookParameterseType2DopplerExtR19EType2Doppler48PortExtR19ScalingfactorR19Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MaxNumberResource_r19, codebookParameterseType2DopplerExtR19EType2Doppler48PortExtR19MaxNumberResourceR19Constraints); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(codebookParameterseType2DopplerExtR19EType2Doppler48PortExtR19SupportedCSIRSResourceListPerCCR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceListPerCC_r19))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceListPerCC_r19 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceListPerCC_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 4. eType2Doppler-128PortExt-r19: sequence
	{
		if ie.EType2Doppler_128PortExt_r19 != nil {
			c := ie.EType2Doppler_128PortExt_r19
			{
				seqOf := e.NewSequenceOfEncoder(codebookParameterseType2DopplerExtR19EType2Doppler128PortExtR19SupportedCSIRSResourceExtListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceExtList_r19))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceExtList_r19 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceExtList_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeEnumerated(c.ProcessingCapability_r19, codebookParameterseType2DopplerExtR19EType2Doppler128PortExtR19ProcessingCapabilityR19Constraints); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.ValueY_P_SP_CSI_RS_r19, per.Constrained(1, 3)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.ValueY_A_CSI_RS_r19, per.Constrained(1, 3)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Scalingfactor_r19, codebookParameterseType2DopplerExtR19EType2Doppler128PortExtR19ScalingfactorR19Constraints); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(codebookParameterseType2DopplerExtR19EType2Doppler128PortExtR19SupportedCSIRSResourceListPerCCR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceListPerCC_r19))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceListPerCC_r19 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceListPerCC_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 5. eType2DopplerN4Ext-r19: sequence
	{
		if ie.EType2DopplerN4Ext_r19 != nil {
			c := ie.EType2DopplerN4Ext_r19
			{
				seqOf := e.NewSequenceOfEncoder(codebookParameterseType2DopplerExtR19EType2DopplerN4ExtR19SupportedCSIRSReportSettingList1R19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ReportSettingList1_r19))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ReportSettingList1_r19 {
					if err := c.SupportedCSI_RS_ReportSettingList1_r19[i].Encode(e); err != nil {
						return err
					}
				}
			}
			{
				seqOf := e.NewSequenceOfEncoder(codebookParameterseType2DopplerExtR19EType2DopplerN4ExtR19SupportedCSIRSReportSettingList2R19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ReportSettingList2_r19))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ReportSettingList2_r19 {
					if err := c.SupportedCSI_RS_ReportSettingList2_r19[i].Encode(e); err != nil {
						return err
					}
				}
			}
		}
	}

	// 6. ddUnitSize-A-CSI-RS-CMR-Ext-r19: enumerated
	{
		if ie.DdUnitSize_A_CSI_RS_CMR_Ext_r19 != nil {
			if err := e.EncodeEnumerated(*ie.DdUnitSize_A_CSI_RS_CMR_Ext_r19, codebookParameterseType2DopplerExtR19DdUnitSizeACSIRSCMRExtR19Constraints); err != nil {
				return err
			}
		}
	}

	// 7. maxNumberAperiodicCSI-RS-ResourceExt-r19: enumerated
	{
		if ie.MaxNumberAperiodicCSI_RS_ResourceExt_r19 != nil {
			if err := e.EncodeEnumerated(*ie.MaxNumberAperiodicCSI_RS_ResourceExt_r19, codebookParameterseType2DopplerExtR19MaxNumberAperiodicCSIRSResourceExtR19Constraints); err != nil {
				return err
			}
		}
	}

	// 8. eType2DopplerR2Ext-r19: sequence-of
	{
		if ie.EType2DopplerR2Ext_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookParameterseType2DopplerExtR19EType2DopplerR2ExtR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.EType2DopplerR2Ext_r19))); err != nil {
				return err
			}
			for i := range ie.EType2DopplerR2Ext_r19 {
				if err := e.EncodeInteger(ie.EType2DopplerR2Ext_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 9. eType2DopplerX1Ext-r19: enumerated
	{
		if ie.EType2DopplerX1Ext_r19 != nil {
			if err := e.EncodeEnumerated(*ie.EType2DopplerX1Ext_r19, codebookParameterseType2DopplerExtR19EType2DopplerX1ExtR19Constraints); err != nil {
				return err
			}
		}
	}

	// 10. eType2DopplerX2Ext-r19: enumerated
	{
		if ie.EType2DopplerX2Ext_r19 != nil {
			if err := e.EncodeEnumerated(*ie.EType2DopplerX2Ext_r19, codebookParameterseType2DopplerExtR19EType2DopplerX2ExtR19Constraints); err != nil {
				return err
			}
		}
	}

	// 11. eType2DopplerL-N4D1Ext-r19: enumerated
	{
		if ie.EType2DopplerL_N4D1Ext_r19 != nil {
			if err := e.EncodeEnumerated(*ie.EType2DopplerL_N4D1Ext_r19, codebookParameterseType2DopplerExtR19EType2DopplerLN4D1ExtR19Constraints); err != nil {
				return err
			}
		}
	}

	// 12. eType2DopplerL6Ext-r19: enumerated
	{
		if ie.EType2DopplerL6Ext_r19 != nil {
			if err := e.EncodeEnumerated(*ie.EType2DopplerL6Ext_r19, codebookParameterseType2DopplerExtR19EType2DopplerL6ExtR19Constraints); err != nil {
				return err
			}
		}
	}

	// 13. eType2DopplerR3R4Ext-r19: enumerated
	{
		if ie.EType2DopplerR3R4Ext_r19 != nil {
			if err := e.EncodeEnumerated(*ie.EType2DopplerR3R4Ext_r19, codebookParameterseType2DopplerExtR19EType2DopplerR3R4ExtR19Constraints); err != nil {
				return err
			}
		}
	}

	// 14. eType2DopplerProcessingTimelineExt-r19: sequence
	{
		if ie.EType2DopplerProcessingTimelineExt_r19 != nil {
			c := ie.EType2DopplerProcessingTimelineExt_r19
			{
				c := &c.ValueW_r19
				codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Seq := e.NewSequenceEncoder(codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Constraints)
				if err := codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Seq.EncodePreamble([]bool{c.Scs_15kHz != nil, c.Scs_30kHz != nil, c.Scs_60kHz != nil, c.Scs_120kHz != nil}); err != nil {
					return err
				}
				if c.Scs_15kHz != nil {
					if err := e.EncodeEnumerated((*c.Scs_15kHz), codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Scs15kHzConstraints); err != nil {
						return err
					}
				}
				if c.Scs_30kHz != nil {
					if err := e.EncodeEnumerated((*c.Scs_30kHz), codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Scs30kHzConstraints); err != nil {
						return err
					}
				}
				if c.Scs_60kHz != nil {
					if err := e.EncodeEnumerated((*c.Scs_60kHz), codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Scs60kHzConstraints); err != nil {
						return err
					}
				}
				if c.Scs_120kHz != nil {
					if err := e.EncodeEnumerated((*c.Scs_120kHz), codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Scs120kHzConstraints); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeEnumerated(c.TimeRelaxation_r19, codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19TimeRelaxationR19Constraints); err != nil {
				return err
			}
		}
	}

	// 15. eType2MaxPeriodicityCMR-r19: enumerated
	{
		if ie.EType2MaxPeriodicityCMR_r19 != nil {
			if err := e.EncodeEnumerated(*ie.EType2MaxPeriodicityCMR_r19, codebookParameterseType2DopplerExtR19EType2MaxPeriodicityCMRR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CodebookParameterseType2DopplerExt_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookParameterseType2DopplerExtR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. eType2Doppler-64PortExt-r19: sequence
	{
		{
			c := &ie.EType2Doppler_64PortExt_r19
			{
				seqOf := d.NewSequenceOfDecoder(codebookParameterseType2DopplerExtR19EType2Doppler64PortExtR19SupportedCSIRSResourceExtListR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceExtList_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceExtList_r19[i] = v
				}
			}
			{
				v, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19EType2Doppler64PortExtR19ProcessingCapabilityR19Constraints)
				if err != nil {
					return err
				}
				c.ProcessingCapability_r19 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 3))
				if err != nil {
					return err
				}
				c.ValueY_P_SP_CSI_RS_r19 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 3))
				if err != nil {
					return err
				}
				c.ValueY_A_CSI_RS_r19 = v
			}
			{
				v, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19EType2Doppler64PortExtR19ScalingfactorR19Constraints)
				if err != nil {
					return err
				}
				c.Scalingfactor_r19 = v
			}
			{
				v, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19EType2Doppler64PortExtR19MaxNumberResourceR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberResource_r19 = v
			}
			{
				seqOf := d.NewSequenceOfDecoder(codebookParameterseType2DopplerExtR19EType2Doppler64PortExtR19SupportedCSIRSResourceListPerCCR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceListPerCC_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceListPerCC_r19[i] = v
				}
			}
		}
	}

	// 3. eType2Doppler-48PortExt-r19: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.EType2Doppler_48PortExt_r19 = &struct {
				SupportedCSI_RS_ResourceExtList_r19   []int64
				ProcessingCapability_r19              int64
				ValueY_P_SP_CSI_RS_r19                int64
				ValueY_A_CSI_RS_r19                   int64
				Scalingfactor_r19                     int64
				MaxNumberResource_r19                 int64
				SupportedCSI_RS_ResourceListPerCC_r19 []int64
			}{}
			c := ie.EType2Doppler_48PortExt_r19
			{
				seqOf := d.NewSequenceOfDecoder(codebookParameterseType2DopplerExtR19EType2Doppler48PortExtR19SupportedCSIRSResourceExtListR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceExtList_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceExtList_r19[i] = v
				}
			}
			{
				v, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19EType2Doppler48PortExtR19ProcessingCapabilityR19Constraints)
				if err != nil {
					return err
				}
				c.ProcessingCapability_r19 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 3))
				if err != nil {
					return err
				}
				c.ValueY_P_SP_CSI_RS_r19 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 3))
				if err != nil {
					return err
				}
				c.ValueY_A_CSI_RS_r19 = v
			}
			{
				v, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19EType2Doppler48PortExtR19ScalingfactorR19Constraints)
				if err != nil {
					return err
				}
				c.Scalingfactor_r19 = v
			}
			{
				v, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19EType2Doppler48PortExtR19MaxNumberResourceR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberResource_r19 = v
			}
			{
				seqOf := d.NewSequenceOfDecoder(codebookParameterseType2DopplerExtR19EType2Doppler48PortExtR19SupportedCSIRSResourceListPerCCR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceListPerCC_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceListPerCC_r19[i] = v
				}
			}
		}
	}

	// 4. eType2Doppler-128PortExt-r19: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.EType2Doppler_128PortExt_r19 = &struct {
				SupportedCSI_RS_ResourceExtList_r19   []int64
				ProcessingCapability_r19              int64
				ValueY_P_SP_CSI_RS_r19                int64
				ValueY_A_CSI_RS_r19                   int64
				Scalingfactor_r19                     int64
				SupportedCSI_RS_ResourceListPerCC_r19 []int64
			}{}
			c := ie.EType2Doppler_128PortExt_r19
			{
				seqOf := d.NewSequenceOfDecoder(codebookParameterseType2DopplerExtR19EType2Doppler128PortExtR19SupportedCSIRSResourceExtListR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceExtList_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceExtList_r19[i] = v
				}
			}
			{
				v, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19EType2Doppler128PortExtR19ProcessingCapabilityR19Constraints)
				if err != nil {
					return err
				}
				c.ProcessingCapability_r19 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 3))
				if err != nil {
					return err
				}
				c.ValueY_P_SP_CSI_RS_r19 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 3))
				if err != nil {
					return err
				}
				c.ValueY_A_CSI_RS_r19 = v
			}
			{
				v, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19EType2Doppler128PortExtR19ScalingfactorR19Constraints)
				if err != nil {
					return err
				}
				c.Scalingfactor_r19 = v
			}
			{
				seqOf := d.NewSequenceOfDecoder(codebookParameterseType2DopplerExtR19EType2Doppler128PortExtR19SupportedCSIRSResourceListPerCCR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceListPerCC_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceListPerCC_r19[i] = v
				}
			}
		}
	}

	// 5. eType2DopplerN4Ext-r19: sequence
	{
		if seq.IsComponentPresent(3) {
			ie.EType2DopplerN4Ext_r19 = &struct {
				SupportedCSI_RS_ReportSettingList1_r19 []SupportedCSI_RS_ReportSettingExt_r19
				SupportedCSI_RS_ReportSettingList2_r19 []SupportedCSI_RS_ReportSettingExt_r19
			}{}
			c := ie.EType2DopplerN4Ext_r19
			{
				seqOf := d.NewSequenceOfDecoder(codebookParameterseType2DopplerExtR19EType2DopplerN4ExtR19SupportedCSIRSReportSettingList1R19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ReportSettingList1_r19 = make([]SupportedCSI_RS_ReportSettingExt_r19, n)
				for i := int64(0); i < n; i++ {
					if err := c.SupportedCSI_RS_ReportSettingList1_r19[i].Decode(d); err != nil {
						return err
					}
				}
			}
			{
				seqOf := d.NewSequenceOfDecoder(codebookParameterseType2DopplerExtR19EType2DopplerN4ExtR19SupportedCSIRSReportSettingList2R19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ReportSettingList2_r19 = make([]SupportedCSI_RS_ReportSettingExt_r19, n)
				for i := int64(0); i < n; i++ {
					if err := c.SupportedCSI_RS_ReportSettingList2_r19[i].Decode(d); err != nil {
						return err
					}
				}
			}
		}
	}

	// 6. ddUnitSize-A-CSI-RS-CMR-Ext-r19: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19DdUnitSizeACSIRSCMRExtR19Constraints)
			if err != nil {
				return err
			}
			ie.DdUnitSize_A_CSI_RS_CMR_Ext_r19 = &idx
		}
	}

	// 7. maxNumberAperiodicCSI-RS-ResourceExt-r19: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19MaxNumberAperiodicCSIRSResourceExtR19Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberAperiodicCSI_RS_ResourceExt_r19 = &idx
		}
	}

	// 8. eType2DopplerR2Ext-r19: sequence-of
	{
		if seq.IsComponentPresent(6) {
			seqOf := d.NewSequenceOfDecoder(codebookParameterseType2DopplerExtR19EType2DopplerR2ExtR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.EType2DopplerR2Ext_r19 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.EType2DopplerR2Ext_r19[i] = v
			}
		}
	}

	// 9. eType2DopplerX1Ext-r19: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19EType2DopplerX1ExtR19Constraints)
			if err != nil {
				return err
			}
			ie.EType2DopplerX1Ext_r19 = &idx
		}
	}

	// 10. eType2DopplerX2Ext-r19: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19EType2DopplerX2ExtR19Constraints)
			if err != nil {
				return err
			}
			ie.EType2DopplerX2Ext_r19 = &idx
		}
	}

	// 11. eType2DopplerL-N4D1Ext-r19: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19EType2DopplerLN4D1ExtR19Constraints)
			if err != nil {
				return err
			}
			ie.EType2DopplerL_N4D1Ext_r19 = &idx
		}
	}

	// 12. eType2DopplerL6Ext-r19: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19EType2DopplerL6ExtR19Constraints)
			if err != nil {
				return err
			}
			ie.EType2DopplerL6Ext_r19 = &idx
		}
	}

	// 13. eType2DopplerR3R4Ext-r19: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19EType2DopplerR3R4ExtR19Constraints)
			if err != nil {
				return err
			}
			ie.EType2DopplerR3R4Ext_r19 = &idx
		}
	}

	// 14. eType2DopplerProcessingTimelineExt-r19: sequence
	{
		if seq.IsComponentPresent(12) {
			ie.EType2DopplerProcessingTimelineExt_r19 = &struct {
				ValueW_r19 struct {
					Scs_15kHz  *int64
					Scs_30kHz  *int64
					Scs_60kHz  *int64
					Scs_120kHz *int64
				}
				TimeRelaxation_r19 int64
			}{}
			c := ie.EType2DopplerProcessingTimelineExt_r19
			{
				c := &c.ValueW_r19
				codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Seq := d.NewSequenceDecoder(codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Constraints)
				if err := codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Seq.DecodePreamble(); err != nil {
					return err
				}
				if codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Seq.IsComponentPresent(0) {
					c.Scs_15kHz = new(int64)
					v, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Scs15kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_15kHz) = v
				}
				if codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Seq.IsComponentPresent(1) {
					c.Scs_30kHz = new(int64)
					v, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Scs30kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_30kHz) = v
				}
				if codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Seq.IsComponentPresent(2) {
					c.Scs_60kHz = new(int64)
					v, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Scs60kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_60kHz) = v
				}
				if codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Seq.IsComponentPresent(3) {
					c.Scs_120kHz = new(int64)
					v, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19ValueWR19Scs120kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_120kHz) = v
				}
			}
			{
				v, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19EType2DopplerProcessingTimelineExtR19TimeRelaxationR19Constraints)
				if err != nil {
					return err
				}
				c.TimeRelaxation_r19 = v
			}
		}
	}

	// 15. eType2MaxPeriodicityCMR-r19: enumerated
	{
		if seq.IsComponentPresent(13) {
			idx, err := d.DecodeEnumerated(codebookParameterseType2DopplerExtR19EType2MaxPeriodicityCMRR19Constraints)
			if err != nil {
				return err
			}
			ie.EType2MaxPeriodicityCMR_r19 = &idx
		}
	}

	return nil
}
