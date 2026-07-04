// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FeatureSetDownlink-v1800 (line 19658).

var featureSetDownlinkV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dynamicSwitchingA-r18", Optional: true},
		{Name: "dynamicSwitchingB-r18", Optional: true},
		{Name: "aperiodicCSI-TimeRelaxation-r18", Optional: true},
		{Name: "pdsch-TypeA-DMRS-r18", Optional: true},
		{Name: "pdsch-TypeB-DMRS-r18", Optional: true},
		{Name: "pdsch-1SymbolFL-DMRS-Addition2Symbol-r18", Optional: true},
		{Name: "pdsch-AlternativeDMRS-Coexistence-r18", Optional: true},
		{Name: "pdsch-2SymbolFL-DMRS-r18", Optional: true},
		{Name: "pdsch-2SymbolFL-DMRS-Addition2Symbol-r18", Optional: true},
		{Name: "pdsch-1SymbolFL-DMRS-Addition3Symbol-r18", Optional: true},
		{Name: "pdsch-DMRS-Type-r18", Optional: true},
		{Name: "pdsch-1PortDL-PTRS-r18", Optional: true},
		{Name: "pdsch-2PortDL-PTRS-r18", Optional: true},
		{Name: "mappingTypeA-1SymbolFL-DMRS-Addition2Symbol-r18", Optional: true},
		{Name: "maxNumberDMRS-AcrossAllDL-DCI-r18", Optional: true},
		{Name: "pdsch-ReceptionWithoutSchedulingRestriction-r18", Optional: true},
		{Name: "pdsch-ReceptionSchemeA-r18", Optional: true},
		{Name: "pdsch-ReceptionSchemeB-r18", Optional: true},
		{Name: "dmrs-MultiTRP-SingleDCI-r18", Optional: true},
		{Name: "dmrs-MultiTRP-AdditionRows-r18", Optional: true},
		{Name: "dmrs-MultiTRP-MultiDCI-r18", Optional: true},
		{Name: "simulDMRS-PDSCH-r18", Optional: true},
		{Name: "bwpOperationMeasWithoutInterrupt-r18", Optional: true},
		{Name: "pdcch-MonitoringSpan2-2-r18", Optional: true},
		{Name: "pdcch-MonitoringMixed-r18", Optional: true},
		{Name: "mTRP-PDCCH-legacyMonitoring-r18", Optional: true},
		{Name: "scellWithoutSSB-InterBandCA-r18", Optional: true},
		{Name: "dummy", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1800_DynamicSwitchingA_r18_Supported = 0
)

var featureSetDownlinkV1800DynamicSwitchingAR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_DynamicSwitchingA_r18_Supported},
}

const (
	FeatureSetDownlink_v1800_DynamicSwitchingB_r18_Supported = 0
)

var featureSetDownlinkV1800DynamicSwitchingBR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_DynamicSwitchingB_r18_Supported},
}

const (
	FeatureSetDownlink_v1800_Pdsch_TypeA_DMRS_r18_Supported = 0
)

var featureSetDownlinkV1800PdschTypeADMRSR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Pdsch_TypeA_DMRS_r18_Supported},
}

const (
	FeatureSetDownlink_v1800_Pdsch_TypeB_DMRS_r18_Supported = 0
)

var featureSetDownlinkV1800PdschTypeBDMRSR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Pdsch_TypeB_DMRS_r18_Supported},
}

const (
	FeatureSetDownlink_v1800_Pdsch_1SymbolFL_DMRS_Addition2Symbol_r18_Supported = 0
)

var featureSetDownlinkV1800Pdsch1SymbolFLDMRSAddition2SymbolR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Pdsch_1SymbolFL_DMRS_Addition2Symbol_r18_Supported},
}

const (
	FeatureSetDownlink_v1800_Pdsch_AlternativeDMRS_Coexistence_r18_Supported = 0
)

var featureSetDownlinkV1800PdschAlternativeDMRSCoexistenceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Pdsch_AlternativeDMRS_Coexistence_r18_Supported},
}

const (
	FeatureSetDownlink_v1800_Pdsch_2SymbolFL_DMRS_r18_Supported = 0
)

var featureSetDownlinkV1800Pdsch2SymbolFLDMRSR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Pdsch_2SymbolFL_DMRS_r18_Supported},
}

const (
	FeatureSetDownlink_v1800_Pdsch_2SymbolFL_DMRS_Addition2Symbol_r18_Supported = 0
)

var featureSetDownlinkV1800Pdsch2SymbolFLDMRSAddition2SymbolR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Pdsch_2SymbolFL_DMRS_Addition2Symbol_r18_Supported},
}

const (
	FeatureSetDownlink_v1800_Pdsch_1SymbolFL_DMRS_Addition3Symbol_r18_Supported = 0
)

var featureSetDownlinkV1800Pdsch1SymbolFLDMRSAddition3SymbolR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Pdsch_1SymbolFL_DMRS_Addition3Symbol_r18_Supported},
}

const (
	FeatureSetDownlink_v1800_Pdsch_DMRS_Type_r18_Etype1     = 0
	FeatureSetDownlink_v1800_Pdsch_DMRS_Type_r18_Etype1And2 = 1
)

var featureSetDownlinkV1800PdschDMRSTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Pdsch_DMRS_Type_r18_Etype1, FeatureSetDownlink_v1800_Pdsch_DMRS_Type_r18_Etype1And2},
}

const (
	FeatureSetDownlink_v1800_Pdsch_1PortDL_PTRS_r18_Supported = 0
)

var featureSetDownlinkV1800Pdsch1PortDLPTRSR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Pdsch_1PortDL_PTRS_r18_Supported},
}

const (
	FeatureSetDownlink_v1800_Pdsch_2PortDL_PTRS_r18_Supported = 0
)

var featureSetDownlinkV1800Pdsch2PortDLPTRSR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Pdsch_2PortDL_PTRS_r18_Supported},
}

const (
	FeatureSetDownlink_v1800_MappingTypeA_1SymbolFL_DMRS_Addition2Symbol_r18_Supported = 0
)

var featureSetDownlinkV1800MappingTypeA1SymbolFLDMRSAddition2SymbolR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_MappingTypeA_1SymbolFL_DMRS_Addition2Symbol_r18_Supported},
}

var featureSetDownlinkV1800MaxNumberDMRSAcrossAllDLDCIR18Constraints = per.Constrained(2, 4)

const (
	FeatureSetDownlink_v1800_Pdsch_ReceptionWithoutSchedulingRestriction_r18_Supported = 0
)

var featureSetDownlinkV1800PdschReceptionWithoutSchedulingRestrictionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Pdsch_ReceptionWithoutSchedulingRestriction_r18_Supported},
}

const (
	FeatureSetDownlink_v1800_Pdsch_ReceptionSchemeA_r18_Supported = 0
)

var featureSetDownlinkV1800PdschReceptionSchemeAR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Pdsch_ReceptionSchemeA_r18_Supported},
}

const (
	FeatureSetDownlink_v1800_Pdsch_ReceptionSchemeB_r18_Supported = 0
)

var featureSetDownlinkV1800PdschReceptionSchemeBR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Pdsch_ReceptionSchemeB_r18_Supported},
}

const (
	FeatureSetDownlink_v1800_Dmrs_MultiTRP_SingleDCI_r18_Supported = 0
)

var featureSetDownlinkV1800DmrsMultiTRPSingleDCIR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Dmrs_MultiTRP_SingleDCI_r18_Supported},
}

const (
	FeatureSetDownlink_v1800_Dmrs_MultiTRP_AdditionRows_r18_Supported = 0
)

var featureSetDownlinkV1800DmrsMultiTRPAdditionRowsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Dmrs_MultiTRP_AdditionRows_r18_Supported},
}

const (
	FeatureSetDownlink_v1800_Dmrs_MultiTRP_MultiDCI_r18_Supported = 0
)

var featureSetDownlinkV1800DmrsMultiTRPMultiDCIR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Dmrs_MultiTRP_MultiDCI_r18_Supported},
}

const (
	FeatureSetDownlink_v1800_BwpOperationMeasWithoutInterrupt_r18_Supported = 0
)

var featureSetDownlinkV1800BwpOperationMeasWithoutInterruptR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_BwpOperationMeasWithoutInterrupt_r18_Supported},
}

const (
	FeatureSetDownlink_v1800_Pdcch_MonitoringMixed_r18_Supported = 0
)

var featureSetDownlinkV1800PdcchMonitoringMixedR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Pdcch_MonitoringMixed_r18_Supported},
}

var featureSetDownlink_v1800ScellWithoutSSBInterBandCAR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "supportOfSingleGroup"},
		{Name: "supportOfMultipleGroups"},
	},
}

const (
	FeatureSetDownlink_v1800_ScellWithoutSSB_InterBandCA_r18_SupportOfSingleGroup    = 0
	FeatureSetDownlink_v1800_ScellWithoutSSB_InterBandCA_r18_SupportOfMultipleGroups = 1
)

var featureSetDownlinkV1800DummyConstraints = per.SizeRange(1, common.MaxBandsMRDC)

var featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz", Optional: true},
		{Name: "scs-30kHz", Optional: true},
		{Name: "scs-60kHz", Optional: true},
		{Name: "scs-120kHz", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1800_AperiodicCSI_TimeRelaxation_r18_ValueW_r18_Scs_15kHz_Value1 = 0
	FeatureSetDownlink_v1800_AperiodicCSI_TimeRelaxation_r18_ValueW_r18_Scs_15kHz_Value2 = 1
)

var featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Scs15kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_AperiodicCSI_TimeRelaxation_r18_ValueW_r18_Scs_15kHz_Value1, FeatureSetDownlink_v1800_AperiodicCSI_TimeRelaxation_r18_ValueW_r18_Scs_15kHz_Value2},
}

const (
	FeatureSetDownlink_v1800_AperiodicCSI_TimeRelaxation_r18_ValueW_r18_Scs_30kHz_Value1 = 0
	FeatureSetDownlink_v1800_AperiodicCSI_TimeRelaxation_r18_ValueW_r18_Scs_30kHz_Value2 = 1
)

var featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Scs30kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_AperiodicCSI_TimeRelaxation_r18_ValueW_r18_Scs_30kHz_Value1, FeatureSetDownlink_v1800_AperiodicCSI_TimeRelaxation_r18_ValueW_r18_Scs_30kHz_Value2},
}

const (
	FeatureSetDownlink_v1800_AperiodicCSI_TimeRelaxation_r18_ValueW_r18_Scs_60kHz_Value1 = 0
	FeatureSetDownlink_v1800_AperiodicCSI_TimeRelaxation_r18_ValueW_r18_Scs_60kHz_Value2 = 1
)

var featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Scs60kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_AperiodicCSI_TimeRelaxation_r18_ValueW_r18_Scs_60kHz_Value1, FeatureSetDownlink_v1800_AperiodicCSI_TimeRelaxation_r18_ValueW_r18_Scs_60kHz_Value2},
}

const (
	FeatureSetDownlink_v1800_AperiodicCSI_TimeRelaxation_r18_ValueW_r18_Scs_120kHz_Value1 = 0
	FeatureSetDownlink_v1800_AperiodicCSI_TimeRelaxation_r18_ValueW_r18_Scs_120kHz_Value2 = 1
)

var featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Scs120kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_AperiodicCSI_TimeRelaxation_r18_ValueW_r18_Scs_120kHz_Value1, FeatureSetDownlink_v1800_AperiodicCSI_TimeRelaxation_r18_ValueW_r18_Scs_120kHz_Value2},
}

const (
	FeatureSetDownlink_v1800_AperiodicCSI_TimeRelaxation_r18_TimeRelaxation_r18_Cap1 = 0
	FeatureSetDownlink_v1800_AperiodicCSI_TimeRelaxation_r18_TimeRelaxation_r18_Cap2 = 1
)

var featureSetDownlinkV1800AperiodicCSITimeRelaxationR18TimeRelaxationR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_AperiodicCSI_TimeRelaxation_r18_TimeRelaxation_r18_Cap1, FeatureSetDownlink_v1800_AperiodicCSI_TimeRelaxation_r18_TimeRelaxation_r18_Cap2},
}

var featureSetDownlinkV1800SimulDMRSPDSCHR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r18", Optional: true},
		{Name: "scs-30kHz-r18", Optional: true},
		{Name: "scs-60kHz-r18", Optional: true},
	},
}

var featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType1R18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r18", Optional: true},
		{Name: "scs-30kHz-r18", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1800_Pdcch_MonitoringSpan2_2_r18_Pdsch_ProcessingType1_r18_Scs_15kHz_r18_Supported = 0
)

var featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType1R18Scs15kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Pdcch_MonitoringSpan2_2_r18_Pdsch_ProcessingType1_r18_Scs_15kHz_r18_Supported},
}

const (
	FeatureSetDownlink_v1800_Pdcch_MonitoringSpan2_2_r18_Pdsch_ProcessingType1_r18_Scs_30kHz_r18_Supported = 0
)

var featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType1R18Scs30kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Pdcch_MonitoringSpan2_2_r18_Pdsch_ProcessingType1_r18_Scs_30kHz_r18_Supported},
}

var featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType2R18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r18", Optional: true},
		{Name: "scs-30kHz-r18", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1800_Pdcch_MonitoringSpan2_2_r18_Pdsch_ProcessingType2_r18_Scs_15kHz_r18_Supported = 0
)

var featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType2R18Scs15kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Pdcch_MonitoringSpan2_2_r18_Pdsch_ProcessingType2_r18_Scs_15kHz_r18_Supported},
}

const (
	FeatureSetDownlink_v1800_Pdcch_MonitoringSpan2_2_r18_Pdsch_ProcessingType2_r18_Scs_30kHz_r18_Supported = 0
)

var featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType2R18Scs30kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_Pdcch_MonitoringSpan2_2_r18_Pdsch_ProcessingType2_r18_Scs_30kHz_r18_Supported},
}

var featureSetDownlinkV1800MTRPPDCCHLegacyMonitoringR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r18", Optional: true},
		{Name: "scs-30kHz-r18", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1800_ScellWithoutSSB_InterBandCA_r18_SupportOfSingleGroup_ReferenceBand   = 0
	FeatureSetDownlink_v1800_ScellWithoutSSB_InterBandCA_r18_SupportOfSingleGroup_ScellWithoutSSB = 1
	FeatureSetDownlink_v1800_ScellWithoutSSB_InterBandCA_r18_SupportOfSingleGroup_Both            = 2
)

var featureSetDownlinkV1800ScellWithoutSSBInterBandCAR18SupportOfSingleGroupConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_ScellWithoutSSB_InterBandCA_r18_SupportOfSingleGroup_ReferenceBand, FeatureSetDownlink_v1800_ScellWithoutSSB_InterBandCA_r18_SupportOfSingleGroup_ScellWithoutSSB, FeatureSetDownlink_v1800_ScellWithoutSSB_InterBandCA_r18_SupportOfSingleGroup_Both},
}

const (
	FeatureSetDownlink_v1800_ScellWithoutSSB_InterBandCA_r18_SupportOfMultipleGroups_ReferenceBand1   = 0
	FeatureSetDownlink_v1800_ScellWithoutSSB_InterBandCA_r18_SupportOfMultipleGroups_ScellWithoutSSB1 = 1
	FeatureSetDownlink_v1800_ScellWithoutSSB_InterBandCA_r18_SupportOfMultipleGroups_ReferenceBand2   = 2
	FeatureSetDownlink_v1800_ScellWithoutSSB_InterBandCA_r18_SupportOfMultipleGroups_ScellWithoutSSB2 = 3
)

var featureSetDownlinkV1800ScellWithoutSSBInterBandCAR18SupportOfMultipleGroupsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1800_ScellWithoutSSB_InterBandCA_r18_SupportOfMultipleGroups_ReferenceBand1, FeatureSetDownlink_v1800_ScellWithoutSSB_InterBandCA_r18_SupportOfMultipleGroups_ScellWithoutSSB1, FeatureSetDownlink_v1800_ScellWithoutSSB_InterBandCA_r18_SupportOfMultipleGroups_ReferenceBand2, FeatureSetDownlink_v1800_ScellWithoutSSB_InterBandCA_r18_SupportOfMultipleGroups_ScellWithoutSSB2},
}

type FeatureSetDownlink_v1800 struct {
	DynamicSwitchingA_r18           *int64
	DynamicSwitchingB_r18           *int64
	AperiodicCSI_TimeRelaxation_r18 *struct {
		ValueW_r18 struct {
			Scs_15kHz  *int64
			Scs_30kHz  *int64
			Scs_60kHz  *int64
			Scs_120kHz *int64
		}
		TimeRelaxation_r18 int64
	}
	Pdsch_TypeA_DMRS_r18                            *int64
	Pdsch_TypeB_DMRS_r18                            *int64
	Pdsch_1SymbolFL_DMRS_Addition2Symbol_r18        *int64
	Pdsch_AlternativeDMRS_Coexistence_r18           *int64
	Pdsch_2SymbolFL_DMRS_r18                        *int64
	Pdsch_2SymbolFL_DMRS_Addition2Symbol_r18        *int64
	Pdsch_1SymbolFL_DMRS_Addition3Symbol_r18        *int64
	Pdsch_DMRS_Type_r18                             *int64
	Pdsch_1PortDL_PTRS_r18                          *int64
	Pdsch_2PortDL_PTRS_r18                          *int64
	MappingTypeA_1SymbolFL_DMRS_Addition2Symbol_r18 *int64
	MaxNumberDMRS_AcrossAllDL_DCI_r18               *int64
	Pdsch_ReceptionWithoutSchedulingRestriction_r18 *int64
	Pdsch_ReceptionSchemeA_r18                      *int64
	Pdsch_ReceptionSchemeB_r18                      *int64
	Dmrs_MultiTRP_SingleDCI_r18                     *int64
	Dmrs_MultiTRP_AdditionRows_r18                  *int64
	Dmrs_MultiTRP_MultiDCI_r18                      *int64
	SimulDMRS_PDSCH_r18                             *struct {
		Scs_15kHz_r18 *int64
		Scs_30kHz_r18 *int64
		Scs_60kHz_r18 *int64
	}
	BwpOperationMeasWithoutInterrupt_r18 *int64
	Pdcch_MonitoringSpan2_2_r18          *struct {
		Pdsch_ProcessingType1_r18 struct {
			Scs_15kHz_r18 *int64
			Scs_30kHz_r18 *int64
		}
		Pdsch_ProcessingType2_r18 struct {
			Scs_15kHz_r18 *int64
			Scs_30kHz_r18 *int64
		}
	}
	Pdcch_MonitoringMixed_r18       *int64
	MTRP_PDCCH_LegacyMonitoring_r18 *struct {
		Scs_15kHz_r18 *PDCCH_RepetitionParameters_r17
		Scs_30kHz_r18 *PDCCH_RepetitionParameters_r17
	}
	ScellWithoutSSB_InterBandCA_r18 *struct {
		Choice                  int
		SupportOfSingleGroup    *int64
		SupportOfMultipleGroups *int64
	}
	Dummy []Dummy_PDCCH_RACH_DL_Info_r18
}

func (ie *FeatureSetDownlink_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkV1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DynamicSwitchingA_r18 != nil, ie.DynamicSwitchingB_r18 != nil, ie.AperiodicCSI_TimeRelaxation_r18 != nil, ie.Pdsch_TypeA_DMRS_r18 != nil, ie.Pdsch_TypeB_DMRS_r18 != nil, ie.Pdsch_1SymbolFL_DMRS_Addition2Symbol_r18 != nil, ie.Pdsch_AlternativeDMRS_Coexistence_r18 != nil, ie.Pdsch_2SymbolFL_DMRS_r18 != nil, ie.Pdsch_2SymbolFL_DMRS_Addition2Symbol_r18 != nil, ie.Pdsch_1SymbolFL_DMRS_Addition3Symbol_r18 != nil, ie.Pdsch_DMRS_Type_r18 != nil, ie.Pdsch_1PortDL_PTRS_r18 != nil, ie.Pdsch_2PortDL_PTRS_r18 != nil, ie.MappingTypeA_1SymbolFL_DMRS_Addition2Symbol_r18 != nil, ie.MaxNumberDMRS_AcrossAllDL_DCI_r18 != nil, ie.Pdsch_ReceptionWithoutSchedulingRestriction_r18 != nil, ie.Pdsch_ReceptionSchemeA_r18 != nil, ie.Pdsch_ReceptionSchemeB_r18 != nil, ie.Dmrs_MultiTRP_SingleDCI_r18 != nil, ie.Dmrs_MultiTRP_AdditionRows_r18 != nil, ie.Dmrs_MultiTRP_MultiDCI_r18 != nil, ie.SimulDMRS_PDSCH_r18 != nil, ie.BwpOperationMeasWithoutInterrupt_r18 != nil, ie.Pdcch_MonitoringSpan2_2_r18 != nil, ie.Pdcch_MonitoringMixed_r18 != nil, ie.MTRP_PDCCH_LegacyMonitoring_r18 != nil, ie.ScellWithoutSSB_InterBandCA_r18 != nil, ie.Dummy != nil}); err != nil {
		return err
	}

	// 2. dynamicSwitchingA-r18: enumerated
	{
		if ie.DynamicSwitchingA_r18 != nil {
			if err := e.EncodeEnumerated(*ie.DynamicSwitchingA_r18, featureSetDownlinkV1800DynamicSwitchingAR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. dynamicSwitchingB-r18: enumerated
	{
		if ie.DynamicSwitchingB_r18 != nil {
			if err := e.EncodeEnumerated(*ie.DynamicSwitchingB_r18, featureSetDownlinkV1800DynamicSwitchingBR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. aperiodicCSI-TimeRelaxation-r18: sequence
	{
		if ie.AperiodicCSI_TimeRelaxation_r18 != nil {
			c := ie.AperiodicCSI_TimeRelaxation_r18
			{
				c := &c.ValueW_r18
				featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Seq := e.NewSequenceEncoder(featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Constraints)
				if err := featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Seq.EncodePreamble([]bool{c.Scs_15kHz != nil, c.Scs_30kHz != nil, c.Scs_60kHz != nil, c.Scs_120kHz != nil}); err != nil {
					return err
				}
				if c.Scs_15kHz != nil {
					if err := e.EncodeEnumerated((*c.Scs_15kHz), featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Scs15kHzConstraints); err != nil {
						return err
					}
				}
				if c.Scs_30kHz != nil {
					if err := e.EncodeEnumerated((*c.Scs_30kHz), featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Scs30kHzConstraints); err != nil {
						return err
					}
				}
				if c.Scs_60kHz != nil {
					if err := e.EncodeEnumerated((*c.Scs_60kHz), featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Scs60kHzConstraints); err != nil {
						return err
					}
				}
				if c.Scs_120kHz != nil {
					if err := e.EncodeEnumerated((*c.Scs_120kHz), featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Scs120kHzConstraints); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeEnumerated(c.TimeRelaxation_r18, featureSetDownlinkV1800AperiodicCSITimeRelaxationR18TimeRelaxationR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. pdsch-TypeA-DMRS-r18: enumerated
	{
		if ie.Pdsch_TypeA_DMRS_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_TypeA_DMRS_r18, featureSetDownlinkV1800PdschTypeADMRSR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. pdsch-TypeB-DMRS-r18: enumerated
	{
		if ie.Pdsch_TypeB_DMRS_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_TypeB_DMRS_r18, featureSetDownlinkV1800PdschTypeBDMRSR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. pdsch-1SymbolFL-DMRS-Addition2Symbol-r18: enumerated
	{
		if ie.Pdsch_1SymbolFL_DMRS_Addition2Symbol_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_1SymbolFL_DMRS_Addition2Symbol_r18, featureSetDownlinkV1800Pdsch1SymbolFLDMRSAddition2SymbolR18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. pdsch-AlternativeDMRS-Coexistence-r18: enumerated
	{
		if ie.Pdsch_AlternativeDMRS_Coexistence_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_AlternativeDMRS_Coexistence_r18, featureSetDownlinkV1800PdschAlternativeDMRSCoexistenceR18Constraints); err != nil {
				return err
			}
		}
	}

	// 9. pdsch-2SymbolFL-DMRS-r18: enumerated
	{
		if ie.Pdsch_2SymbolFL_DMRS_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_2SymbolFL_DMRS_r18, featureSetDownlinkV1800Pdsch2SymbolFLDMRSR18Constraints); err != nil {
				return err
			}
		}
	}

	// 10. pdsch-2SymbolFL-DMRS-Addition2Symbol-r18: enumerated
	{
		if ie.Pdsch_2SymbolFL_DMRS_Addition2Symbol_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_2SymbolFL_DMRS_Addition2Symbol_r18, featureSetDownlinkV1800Pdsch2SymbolFLDMRSAddition2SymbolR18Constraints); err != nil {
				return err
			}
		}
	}

	// 11. pdsch-1SymbolFL-DMRS-Addition3Symbol-r18: enumerated
	{
		if ie.Pdsch_1SymbolFL_DMRS_Addition3Symbol_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_1SymbolFL_DMRS_Addition3Symbol_r18, featureSetDownlinkV1800Pdsch1SymbolFLDMRSAddition3SymbolR18Constraints); err != nil {
				return err
			}
		}
	}

	// 12. pdsch-DMRS-Type-r18: enumerated
	{
		if ie.Pdsch_DMRS_Type_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_DMRS_Type_r18, featureSetDownlinkV1800PdschDMRSTypeR18Constraints); err != nil {
				return err
			}
		}
	}

	// 13. pdsch-1PortDL-PTRS-r18: enumerated
	{
		if ie.Pdsch_1PortDL_PTRS_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_1PortDL_PTRS_r18, featureSetDownlinkV1800Pdsch1PortDLPTRSR18Constraints); err != nil {
				return err
			}
		}
	}

	// 14. pdsch-2PortDL-PTRS-r18: enumerated
	{
		if ie.Pdsch_2PortDL_PTRS_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_2PortDL_PTRS_r18, featureSetDownlinkV1800Pdsch2PortDLPTRSR18Constraints); err != nil {
				return err
			}
		}
	}

	// 15. mappingTypeA-1SymbolFL-DMRS-Addition2Symbol-r18: enumerated
	{
		if ie.MappingTypeA_1SymbolFL_DMRS_Addition2Symbol_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MappingTypeA_1SymbolFL_DMRS_Addition2Symbol_r18, featureSetDownlinkV1800MappingTypeA1SymbolFLDMRSAddition2SymbolR18Constraints); err != nil {
				return err
			}
		}
	}

	// 16. maxNumberDMRS-AcrossAllDL-DCI-r18: integer
	{
		if ie.MaxNumberDMRS_AcrossAllDL_DCI_r18 != nil {
			if err := e.EncodeInteger(*ie.MaxNumberDMRS_AcrossAllDL_DCI_r18, featureSetDownlinkV1800MaxNumberDMRSAcrossAllDLDCIR18Constraints); err != nil {
				return err
			}
		}
	}

	// 17. pdsch-ReceptionWithoutSchedulingRestriction-r18: enumerated
	{
		if ie.Pdsch_ReceptionWithoutSchedulingRestriction_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_ReceptionWithoutSchedulingRestriction_r18, featureSetDownlinkV1800PdschReceptionWithoutSchedulingRestrictionR18Constraints); err != nil {
				return err
			}
		}
	}

	// 18. pdsch-ReceptionSchemeA-r18: enumerated
	{
		if ie.Pdsch_ReceptionSchemeA_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_ReceptionSchemeA_r18, featureSetDownlinkV1800PdschReceptionSchemeAR18Constraints); err != nil {
				return err
			}
		}
	}

	// 19. pdsch-ReceptionSchemeB-r18: enumerated
	{
		if ie.Pdsch_ReceptionSchemeB_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_ReceptionSchemeB_r18, featureSetDownlinkV1800PdschReceptionSchemeBR18Constraints); err != nil {
				return err
			}
		}
	}

	// 20. dmrs-MultiTRP-SingleDCI-r18: enumerated
	{
		if ie.Dmrs_MultiTRP_SingleDCI_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Dmrs_MultiTRP_SingleDCI_r18, featureSetDownlinkV1800DmrsMultiTRPSingleDCIR18Constraints); err != nil {
				return err
			}
		}
	}

	// 21. dmrs-MultiTRP-AdditionRows-r18: enumerated
	{
		if ie.Dmrs_MultiTRP_AdditionRows_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Dmrs_MultiTRP_AdditionRows_r18, featureSetDownlinkV1800DmrsMultiTRPAdditionRowsR18Constraints); err != nil {
				return err
			}
		}
	}

	// 22. dmrs-MultiTRP-MultiDCI-r18: enumerated
	{
		if ie.Dmrs_MultiTRP_MultiDCI_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Dmrs_MultiTRP_MultiDCI_r18, featureSetDownlinkV1800DmrsMultiTRPMultiDCIR18Constraints); err != nil {
				return err
			}
		}
	}

	// 23. simulDMRS-PDSCH-r18: sequence
	{
		if ie.SimulDMRS_PDSCH_r18 != nil {
			c := ie.SimulDMRS_PDSCH_r18
			featureSetDownlinkV1800SimulDMRSPDSCHR18Seq := e.NewSequenceEncoder(featureSetDownlinkV1800SimulDMRSPDSCHR18Constraints)
			if err := featureSetDownlinkV1800SimulDMRSPDSCHR18Seq.EncodePreamble([]bool{c.Scs_15kHz_r18 != nil, c.Scs_30kHz_r18 != nil, c.Scs_60kHz_r18 != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz_r18 != nil {
				if err := e.EncodeInteger((*c.Scs_15kHz_r18), per.Constrained(0, 4)); err != nil {
					return err
				}
			}
			if c.Scs_30kHz_r18 != nil {
				if err := e.EncodeInteger((*c.Scs_30kHz_r18), per.Constrained(0, 5)); err != nil {
					return err
				}
			}
			if c.Scs_60kHz_r18 != nil {
				if err := e.EncodeInteger((*c.Scs_60kHz_r18), per.Constrained(0, 7)); err != nil {
					return err
				}
			}
		}
	}

	// 24. bwpOperationMeasWithoutInterrupt-r18: enumerated
	{
		if ie.BwpOperationMeasWithoutInterrupt_r18 != nil {
			if err := e.EncodeEnumerated(*ie.BwpOperationMeasWithoutInterrupt_r18, featureSetDownlinkV1800BwpOperationMeasWithoutInterruptR18Constraints); err != nil {
				return err
			}
		}
	}

	// 25. pdcch-MonitoringSpan2-2-r18: sequence
	{
		if ie.Pdcch_MonitoringSpan2_2_r18 != nil {
			c := ie.Pdcch_MonitoringSpan2_2_r18
			{
				c := &c.Pdsch_ProcessingType1_r18
				featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType1R18Seq := e.NewSequenceEncoder(featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType1R18Constraints)
				if err := featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType1R18Seq.EncodePreamble([]bool{c.Scs_15kHz_r18 != nil, c.Scs_30kHz_r18 != nil}); err != nil {
					return err
				}
				if c.Scs_15kHz_r18 != nil {
					if err := e.EncodeEnumerated((*c.Scs_15kHz_r18), featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType1R18Scs15kHzR18Constraints); err != nil {
						return err
					}
				}
				if c.Scs_30kHz_r18 != nil {
					if err := e.EncodeEnumerated((*c.Scs_30kHz_r18), featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType1R18Scs30kHzR18Constraints); err != nil {
						return err
					}
				}
			}
			{
				c := &c.Pdsch_ProcessingType2_r18
				featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType2R18Seq := e.NewSequenceEncoder(featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType2R18Constraints)
				if err := featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType2R18Seq.EncodePreamble([]bool{c.Scs_15kHz_r18 != nil, c.Scs_30kHz_r18 != nil}); err != nil {
					return err
				}
				if c.Scs_15kHz_r18 != nil {
					if err := e.EncodeEnumerated((*c.Scs_15kHz_r18), featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType2R18Scs15kHzR18Constraints); err != nil {
						return err
					}
				}
				if c.Scs_30kHz_r18 != nil {
					if err := e.EncodeEnumerated((*c.Scs_30kHz_r18), featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType2R18Scs30kHzR18Constraints); err != nil {
						return err
					}
				}
			}
		}
	}

	// 26. pdcch-MonitoringMixed-r18: enumerated
	{
		if ie.Pdcch_MonitoringMixed_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdcch_MonitoringMixed_r18, featureSetDownlinkV1800PdcchMonitoringMixedR18Constraints); err != nil {
				return err
			}
		}
	}

	// 27. mTRP-PDCCH-legacyMonitoring-r18: sequence
	{
		if ie.MTRP_PDCCH_LegacyMonitoring_r18 != nil {
			c := ie.MTRP_PDCCH_LegacyMonitoring_r18
			featureSetDownlinkV1800MTRPPDCCHLegacyMonitoringR18Seq := e.NewSequenceEncoder(featureSetDownlinkV1800MTRPPDCCHLegacyMonitoringR18Constraints)
			if err := featureSetDownlinkV1800MTRPPDCCHLegacyMonitoringR18Seq.EncodePreamble([]bool{c.Scs_15kHz_r18 != nil, c.Scs_30kHz_r18 != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz_r18 != nil {
				if err := c.Scs_15kHz_r18.Encode(e); err != nil {
					return err
				}
			}
			if c.Scs_30kHz_r18 != nil {
				if err := c.Scs_30kHz_r18.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 28. scellWithoutSSB-InterBandCA-r18: choice
	{
		if ie.ScellWithoutSSB_InterBandCA_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(featureSetDownlink_v1800ScellWithoutSSBInterBandCAR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.ScellWithoutSSB_InterBandCA_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.ScellWithoutSSB_InterBandCA_r18).Choice {
			case FeatureSetDownlink_v1800_ScellWithoutSSB_InterBandCA_r18_SupportOfSingleGroup:
				if err := e.EncodeEnumerated((*(*ie.ScellWithoutSSB_InterBandCA_r18).SupportOfSingleGroup), featureSetDownlinkV1800ScellWithoutSSBInterBandCAR18SupportOfSingleGroupConstraints); err != nil {
					return err
				}
			case FeatureSetDownlink_v1800_ScellWithoutSSB_InterBandCA_r18_SupportOfMultipleGroups:
				if err := e.EncodeEnumerated((*(*ie.ScellWithoutSSB_InterBandCA_r18).SupportOfMultipleGroups), featureSetDownlinkV1800ScellWithoutSSBInterBandCAR18SupportOfMultipleGroupsConstraints); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.ScellWithoutSSB_InterBandCA_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 29. dummy: sequence-of
	{
		if ie.Dummy != nil {
			seqOf := e.NewSequenceOfEncoder(featureSetDownlinkV1800DummyConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Dummy))); err != nil {
				return err
			}
			for i := range ie.Dummy {
				if err := ie.Dummy[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlink_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkV1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dynamicSwitchingA-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800DynamicSwitchingAR18Constraints)
			if err != nil {
				return err
			}
			ie.DynamicSwitchingA_r18 = &idx
		}
	}

	// 3. dynamicSwitchingB-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800DynamicSwitchingBR18Constraints)
			if err != nil {
				return err
			}
			ie.DynamicSwitchingB_r18 = &idx
		}
	}

	// 4. aperiodicCSI-TimeRelaxation-r18: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.AperiodicCSI_TimeRelaxation_r18 = &struct {
				ValueW_r18 struct {
					Scs_15kHz  *int64
					Scs_30kHz  *int64
					Scs_60kHz  *int64
					Scs_120kHz *int64
				}
				TimeRelaxation_r18 int64
			}{}
			c := ie.AperiodicCSI_TimeRelaxation_r18
			{
				c := &c.ValueW_r18
				featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Seq := d.NewSequenceDecoder(featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Constraints)
				if err := featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Seq.DecodePreamble(); err != nil {
					return err
				}
				if featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Seq.IsComponentPresent(0) {
					c.Scs_15kHz = new(int64)
					v, err := d.DecodeEnumerated(featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Scs15kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_15kHz) = v
				}
				if featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Seq.IsComponentPresent(1) {
					c.Scs_30kHz = new(int64)
					v, err := d.DecodeEnumerated(featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Scs30kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_30kHz) = v
				}
				if featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Seq.IsComponentPresent(2) {
					c.Scs_60kHz = new(int64)
					v, err := d.DecodeEnumerated(featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Scs60kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_60kHz) = v
				}
				if featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Seq.IsComponentPresent(3) {
					c.Scs_120kHz = new(int64)
					v, err := d.DecodeEnumerated(featureSetDownlinkV1800AperiodicCSITimeRelaxationR18ValueWR18Scs120kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_120kHz) = v
				}
			}
			{
				v, err := d.DecodeEnumerated(featureSetDownlinkV1800AperiodicCSITimeRelaxationR18TimeRelaxationR18Constraints)
				if err != nil {
					return err
				}
				c.TimeRelaxation_r18 = v
			}
		}
	}

	// 5. pdsch-TypeA-DMRS-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800PdschTypeADMRSR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_TypeA_DMRS_r18 = &idx
		}
	}

	// 6. pdsch-TypeB-DMRS-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800PdschTypeBDMRSR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_TypeB_DMRS_r18 = &idx
		}
	}

	// 7. pdsch-1SymbolFL-DMRS-Addition2Symbol-r18: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800Pdsch1SymbolFLDMRSAddition2SymbolR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_1SymbolFL_DMRS_Addition2Symbol_r18 = &idx
		}
	}

	// 8. pdsch-AlternativeDMRS-Coexistence-r18: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800PdschAlternativeDMRSCoexistenceR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_AlternativeDMRS_Coexistence_r18 = &idx
		}
	}

	// 9. pdsch-2SymbolFL-DMRS-r18: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800Pdsch2SymbolFLDMRSR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_2SymbolFL_DMRS_r18 = &idx
		}
	}

	// 10. pdsch-2SymbolFL-DMRS-Addition2Symbol-r18: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800Pdsch2SymbolFLDMRSAddition2SymbolR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_2SymbolFL_DMRS_Addition2Symbol_r18 = &idx
		}
	}

	// 11. pdsch-1SymbolFL-DMRS-Addition3Symbol-r18: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800Pdsch1SymbolFLDMRSAddition3SymbolR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_1SymbolFL_DMRS_Addition3Symbol_r18 = &idx
		}
	}

	// 12. pdsch-DMRS-Type-r18: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800PdschDMRSTypeR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_DMRS_Type_r18 = &idx
		}
	}

	// 13. pdsch-1PortDL-PTRS-r18: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800Pdsch1PortDLPTRSR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_1PortDL_PTRS_r18 = &idx
		}
	}

	// 14. pdsch-2PortDL-PTRS-r18: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800Pdsch2PortDLPTRSR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_2PortDL_PTRS_r18 = &idx
		}
	}

	// 15. mappingTypeA-1SymbolFL-DMRS-Addition2Symbol-r18: enumerated
	{
		if seq.IsComponentPresent(13) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800MappingTypeA1SymbolFLDMRSAddition2SymbolR18Constraints)
			if err != nil {
				return err
			}
			ie.MappingTypeA_1SymbolFL_DMRS_Addition2Symbol_r18 = &idx
		}
	}

	// 16. maxNumberDMRS-AcrossAllDL-DCI-r18: integer
	{
		if seq.IsComponentPresent(14) {
			v, err := d.DecodeInteger(featureSetDownlinkV1800MaxNumberDMRSAcrossAllDLDCIR18Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberDMRS_AcrossAllDL_DCI_r18 = &v
		}
	}

	// 17. pdsch-ReceptionWithoutSchedulingRestriction-r18: enumerated
	{
		if seq.IsComponentPresent(15) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800PdschReceptionWithoutSchedulingRestrictionR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_ReceptionWithoutSchedulingRestriction_r18 = &idx
		}
	}

	// 18. pdsch-ReceptionSchemeA-r18: enumerated
	{
		if seq.IsComponentPresent(16) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800PdschReceptionSchemeAR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_ReceptionSchemeA_r18 = &idx
		}
	}

	// 19. pdsch-ReceptionSchemeB-r18: enumerated
	{
		if seq.IsComponentPresent(17) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800PdschReceptionSchemeBR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_ReceptionSchemeB_r18 = &idx
		}
	}

	// 20. dmrs-MultiTRP-SingleDCI-r18: enumerated
	{
		if seq.IsComponentPresent(18) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800DmrsMultiTRPSingleDCIR18Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_MultiTRP_SingleDCI_r18 = &idx
		}
	}

	// 21. dmrs-MultiTRP-AdditionRows-r18: enumerated
	{
		if seq.IsComponentPresent(19) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800DmrsMultiTRPAdditionRowsR18Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_MultiTRP_AdditionRows_r18 = &idx
		}
	}

	// 22. dmrs-MultiTRP-MultiDCI-r18: enumerated
	{
		if seq.IsComponentPresent(20) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800DmrsMultiTRPMultiDCIR18Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_MultiTRP_MultiDCI_r18 = &idx
		}
	}

	// 23. simulDMRS-PDSCH-r18: sequence
	{
		if seq.IsComponentPresent(21) {
			ie.SimulDMRS_PDSCH_r18 = &struct {
				Scs_15kHz_r18 *int64
				Scs_30kHz_r18 *int64
				Scs_60kHz_r18 *int64
			}{}
			c := ie.SimulDMRS_PDSCH_r18
			featureSetDownlinkV1800SimulDMRSPDSCHR18Seq := d.NewSequenceDecoder(featureSetDownlinkV1800SimulDMRSPDSCHR18Constraints)
			if err := featureSetDownlinkV1800SimulDMRSPDSCHR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetDownlinkV1800SimulDMRSPDSCHR18Seq.IsComponentPresent(0) {
				c.Scs_15kHz_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 4))
				if err != nil {
					return err
				}
				(*c.Scs_15kHz_r18) = v
			}
			if featureSetDownlinkV1800SimulDMRSPDSCHR18Seq.IsComponentPresent(1) {
				c.Scs_30kHz_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 5))
				if err != nil {
					return err
				}
				(*c.Scs_30kHz_r18) = v
			}
			if featureSetDownlinkV1800SimulDMRSPDSCHR18Seq.IsComponentPresent(2) {
				c.Scs_60kHz_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 7))
				if err != nil {
					return err
				}
				(*c.Scs_60kHz_r18) = v
			}
		}
	}

	// 24. bwpOperationMeasWithoutInterrupt-r18: enumerated
	{
		if seq.IsComponentPresent(22) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800BwpOperationMeasWithoutInterruptR18Constraints)
			if err != nil {
				return err
			}
			ie.BwpOperationMeasWithoutInterrupt_r18 = &idx
		}
	}

	// 25. pdcch-MonitoringSpan2-2-r18: sequence
	{
		if seq.IsComponentPresent(23) {
			ie.Pdcch_MonitoringSpan2_2_r18 = &struct {
				Pdsch_ProcessingType1_r18 struct {
					Scs_15kHz_r18 *int64
					Scs_30kHz_r18 *int64
				}
				Pdsch_ProcessingType2_r18 struct {
					Scs_15kHz_r18 *int64
					Scs_30kHz_r18 *int64
				}
			}{}
			c := ie.Pdcch_MonitoringSpan2_2_r18
			{
				c := &c.Pdsch_ProcessingType1_r18
				featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType1R18Seq := d.NewSequenceDecoder(featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType1R18Constraints)
				if err := featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType1R18Seq.DecodePreamble(); err != nil {
					return err
				}
				if featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType1R18Seq.IsComponentPresent(0) {
					c.Scs_15kHz_r18 = new(int64)
					v, err := d.DecodeEnumerated(featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType1R18Scs15kHzR18Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_15kHz_r18) = v
				}
				if featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType1R18Seq.IsComponentPresent(1) {
					c.Scs_30kHz_r18 = new(int64)
					v, err := d.DecodeEnumerated(featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType1R18Scs30kHzR18Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_30kHz_r18) = v
				}
			}
			{
				c := &c.Pdsch_ProcessingType2_r18
				featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType2R18Seq := d.NewSequenceDecoder(featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType2R18Constraints)
				if err := featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType2R18Seq.DecodePreamble(); err != nil {
					return err
				}
				if featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType2R18Seq.IsComponentPresent(0) {
					c.Scs_15kHz_r18 = new(int64)
					v, err := d.DecodeEnumerated(featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType2R18Scs15kHzR18Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_15kHz_r18) = v
				}
				if featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType2R18Seq.IsComponentPresent(1) {
					c.Scs_30kHz_r18 = new(int64)
					v, err := d.DecodeEnumerated(featureSetDownlinkV1800PdcchMonitoringSpan22R18PdschProcessingType2R18Scs30kHzR18Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_30kHz_r18) = v
				}
			}
		}
	}

	// 26. pdcch-MonitoringMixed-r18: enumerated
	{
		if seq.IsComponentPresent(24) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1800PdcchMonitoringMixedR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdcch_MonitoringMixed_r18 = &idx
		}
	}

	// 27. mTRP-PDCCH-legacyMonitoring-r18: sequence
	{
		if seq.IsComponentPresent(25) {
			ie.MTRP_PDCCH_LegacyMonitoring_r18 = &struct {
				Scs_15kHz_r18 *PDCCH_RepetitionParameters_r17
				Scs_30kHz_r18 *PDCCH_RepetitionParameters_r17
			}{}
			c := ie.MTRP_PDCCH_LegacyMonitoring_r18
			featureSetDownlinkV1800MTRPPDCCHLegacyMonitoringR18Seq := d.NewSequenceDecoder(featureSetDownlinkV1800MTRPPDCCHLegacyMonitoringR18Constraints)
			if err := featureSetDownlinkV1800MTRPPDCCHLegacyMonitoringR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetDownlinkV1800MTRPPDCCHLegacyMonitoringR18Seq.IsComponentPresent(0) {
				c.Scs_15kHz_r18 = new(PDCCH_RepetitionParameters_r17)
				if err := (*c.Scs_15kHz_r18).Decode(d); err != nil {
					return err
				}
			}
			if featureSetDownlinkV1800MTRPPDCCHLegacyMonitoringR18Seq.IsComponentPresent(1) {
				c.Scs_30kHz_r18 = new(PDCCH_RepetitionParameters_r17)
				if err := (*c.Scs_30kHz_r18).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 28. scellWithoutSSB-InterBandCA-r18: choice
	{
		if seq.IsComponentPresent(26) {
			ie.ScellWithoutSSB_InterBandCA_r18 = &struct {
				Choice                  int
				SupportOfSingleGroup    *int64
				SupportOfMultipleGroups *int64
			}{}
			choiceDec := d.NewChoiceDecoder(featureSetDownlink_v1800ScellWithoutSSBInterBandCAR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ScellWithoutSSB_InterBandCA_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case FeatureSetDownlink_v1800_ScellWithoutSSB_InterBandCA_r18_SupportOfSingleGroup:
				(*ie.ScellWithoutSSB_InterBandCA_r18).SupportOfSingleGroup = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1800ScellWithoutSSBInterBandCAR18SupportOfSingleGroupConstraints)
				if err != nil {
					return err
				}
				(*(*ie.ScellWithoutSSB_InterBandCA_r18).SupportOfSingleGroup) = v
			case FeatureSetDownlink_v1800_ScellWithoutSSB_InterBandCA_r18_SupportOfMultipleGroups:
				(*ie.ScellWithoutSSB_InterBandCA_r18).SupportOfMultipleGroups = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1800ScellWithoutSSBInterBandCAR18SupportOfMultipleGroupsConstraints)
				if err != nil {
					return err
				}
				(*(*ie.ScellWithoutSSB_InterBandCA_r18).SupportOfMultipleGroups) = v
			}
		}
	}

	// 29. dummy: sequence-of
	{
		if seq.IsComponentPresent(27) {
			seqOf := d.NewSequenceOfDecoder(featureSetDownlinkV1800DummyConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Dummy = make([]Dummy_PDCCH_RACH_DL_Info_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Dummy[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
