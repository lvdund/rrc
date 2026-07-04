// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNR-v1700 (line 17436).

var cAParametersNRV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "codebookParametersfetype2PerBC-r17", Optional: true},
		{Name: "demodulationEnhancementCA-r17", Optional: true},
		{Name: "maxUplinkDutyCycle-interBandCA-PC2-r17", Optional: true},
		{Name: "maxUplinkDutyCycle-SULcombination-PC2-r17", Optional: true},
		{Name: "beamManagementType-CBM-r17", Optional: true},
		{Name: "parallelTxPUCCH-PUSCH-r17", Optional: true},
		{Name: "codebookComboParameterMixedTypePerBC-r17", Optional: true},
		{Name: "mTRP-CSI-EnhancementPerBC-r17", Optional: true},
		{Name: "codebookComboParameterMultiTRP-PerBC-r17", Optional: true},
		{Name: "maxCC-32-DL-HARQ-ProcessFR2-2-r17", Optional: true},
		{Name: "maxCC-32-UL-HARQ-ProcessFR2-2-r17", Optional: true},
		{Name: "crossCarrierSchedulingSCell-SpCellTypeB-r17", Optional: true},
		{Name: "crossCarrierSchedulingSCell-SpCellTypeA-r17", Optional: true},
		{Name: "dci-FormatsPCellPSCellUSS-Sets-r17", Optional: true},
		{Name: "disablingScalingFactorDeactSCell-r17", Optional: true},
		{Name: "disablingScalingFactorDormantSCell-r17", Optional: true},
		{Name: "non-AlignedFrameBoundaries-r17", Optional: true},
	},
}

const (
	CA_ParametersNR_v1700_DemodulationEnhancementCA_r17_Supported = 0
)

var cAParametersNRV1700DemodulationEnhancementCAR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1700_DemodulationEnhancementCA_r17_Supported},
}

const (
	CA_ParametersNR_v1700_MaxUplinkDutyCycle_InterBandCA_PC2_r17_N50  = 0
	CA_ParametersNR_v1700_MaxUplinkDutyCycle_InterBandCA_PC2_r17_N60  = 1
	CA_ParametersNR_v1700_MaxUplinkDutyCycle_InterBandCA_PC2_r17_N70  = 2
	CA_ParametersNR_v1700_MaxUplinkDutyCycle_InterBandCA_PC2_r17_N80  = 3
	CA_ParametersNR_v1700_MaxUplinkDutyCycle_InterBandCA_PC2_r17_N90  = 4
	CA_ParametersNR_v1700_MaxUplinkDutyCycle_InterBandCA_PC2_r17_N100 = 5
)

var cAParametersNRV1700MaxUplinkDutyCycleInterBandCAPC2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1700_MaxUplinkDutyCycle_InterBandCA_PC2_r17_N50, CA_ParametersNR_v1700_MaxUplinkDutyCycle_InterBandCA_PC2_r17_N60, CA_ParametersNR_v1700_MaxUplinkDutyCycle_InterBandCA_PC2_r17_N70, CA_ParametersNR_v1700_MaxUplinkDutyCycle_InterBandCA_PC2_r17_N80, CA_ParametersNR_v1700_MaxUplinkDutyCycle_InterBandCA_PC2_r17_N90, CA_ParametersNR_v1700_MaxUplinkDutyCycle_InterBandCA_PC2_r17_N100},
}

const (
	CA_ParametersNR_v1700_MaxUplinkDutyCycle_SULcombination_PC2_r17_N50  = 0
	CA_ParametersNR_v1700_MaxUplinkDutyCycle_SULcombination_PC2_r17_N60  = 1
	CA_ParametersNR_v1700_MaxUplinkDutyCycle_SULcombination_PC2_r17_N70  = 2
	CA_ParametersNR_v1700_MaxUplinkDutyCycle_SULcombination_PC2_r17_N80  = 3
	CA_ParametersNR_v1700_MaxUplinkDutyCycle_SULcombination_PC2_r17_N90  = 4
	CA_ParametersNR_v1700_MaxUplinkDutyCycle_SULcombination_PC2_r17_N100 = 5
)

var cAParametersNRV1700MaxUplinkDutyCycleSULcombinationPC2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1700_MaxUplinkDutyCycle_SULcombination_PC2_r17_N50, CA_ParametersNR_v1700_MaxUplinkDutyCycle_SULcombination_PC2_r17_N60, CA_ParametersNR_v1700_MaxUplinkDutyCycle_SULcombination_PC2_r17_N70, CA_ParametersNR_v1700_MaxUplinkDutyCycle_SULcombination_PC2_r17_N80, CA_ParametersNR_v1700_MaxUplinkDutyCycle_SULcombination_PC2_r17_N90, CA_ParametersNR_v1700_MaxUplinkDutyCycle_SULcombination_PC2_r17_N100},
}

const (
	CA_ParametersNR_v1700_BeamManagementType_CBM_r17_Supported = 0
)

var cAParametersNRV1700BeamManagementTypeCBMR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1700_BeamManagementType_CBM_r17_Supported},
}

const (
	CA_ParametersNR_v1700_ParallelTxPUCCH_PUSCH_r17_Supported = 0
)

var cAParametersNRV1700ParallelTxPUCCHPUSCHR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1700_ParallelTxPUCCH_PUSCH_r17_Supported},
}

const (
	CA_ParametersNR_v1700_MaxCC_32_DL_HARQ_ProcessFR2_2_r17_N1  = 0
	CA_ParametersNR_v1700_MaxCC_32_DL_HARQ_ProcessFR2_2_r17_N2  = 1
	CA_ParametersNR_v1700_MaxCC_32_DL_HARQ_ProcessFR2_2_r17_N3  = 2
	CA_ParametersNR_v1700_MaxCC_32_DL_HARQ_ProcessFR2_2_r17_N4  = 3
	CA_ParametersNR_v1700_MaxCC_32_DL_HARQ_ProcessFR2_2_r17_N6  = 4
	CA_ParametersNR_v1700_MaxCC_32_DL_HARQ_ProcessFR2_2_r17_N8  = 5
	CA_ParametersNR_v1700_MaxCC_32_DL_HARQ_ProcessFR2_2_r17_N16 = 6
	CA_ParametersNR_v1700_MaxCC_32_DL_HARQ_ProcessFR2_2_r17_N32 = 7
)

var cAParametersNRV1700MaxCC32DLHARQProcessFR22R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1700_MaxCC_32_DL_HARQ_ProcessFR2_2_r17_N1, CA_ParametersNR_v1700_MaxCC_32_DL_HARQ_ProcessFR2_2_r17_N2, CA_ParametersNR_v1700_MaxCC_32_DL_HARQ_ProcessFR2_2_r17_N3, CA_ParametersNR_v1700_MaxCC_32_DL_HARQ_ProcessFR2_2_r17_N4, CA_ParametersNR_v1700_MaxCC_32_DL_HARQ_ProcessFR2_2_r17_N6, CA_ParametersNR_v1700_MaxCC_32_DL_HARQ_ProcessFR2_2_r17_N8, CA_ParametersNR_v1700_MaxCC_32_DL_HARQ_ProcessFR2_2_r17_N16, CA_ParametersNR_v1700_MaxCC_32_DL_HARQ_ProcessFR2_2_r17_N32},
}

const (
	CA_ParametersNR_v1700_MaxCC_32_UL_HARQ_ProcessFR2_2_r17_N1  = 0
	CA_ParametersNR_v1700_MaxCC_32_UL_HARQ_ProcessFR2_2_r17_N2  = 1
	CA_ParametersNR_v1700_MaxCC_32_UL_HARQ_ProcessFR2_2_r17_N3  = 2
	CA_ParametersNR_v1700_MaxCC_32_UL_HARQ_ProcessFR2_2_r17_N4  = 3
	CA_ParametersNR_v1700_MaxCC_32_UL_HARQ_ProcessFR2_2_r17_N5  = 4
	CA_ParametersNR_v1700_MaxCC_32_UL_HARQ_ProcessFR2_2_r17_N8  = 5
	CA_ParametersNR_v1700_MaxCC_32_UL_HARQ_ProcessFR2_2_r17_N16 = 6
	CA_ParametersNR_v1700_MaxCC_32_UL_HARQ_ProcessFR2_2_r17_N32 = 7
)

var cAParametersNRV1700MaxCC32ULHARQProcessFR22R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1700_MaxCC_32_UL_HARQ_ProcessFR2_2_r17_N1, CA_ParametersNR_v1700_MaxCC_32_UL_HARQ_ProcessFR2_2_r17_N2, CA_ParametersNR_v1700_MaxCC_32_UL_HARQ_ProcessFR2_2_r17_N3, CA_ParametersNR_v1700_MaxCC_32_UL_HARQ_ProcessFR2_2_r17_N4, CA_ParametersNR_v1700_MaxCC_32_UL_HARQ_ProcessFR2_2_r17_N5, CA_ParametersNR_v1700_MaxCC_32_UL_HARQ_ProcessFR2_2_r17_N8, CA_ParametersNR_v1700_MaxCC_32_UL_HARQ_ProcessFR2_2_r17_N16, CA_ParametersNR_v1700_MaxCC_32_UL_HARQ_ProcessFR2_2_r17_N32},
}

const (
	CA_ParametersNR_v1700_Dci_FormatsPCellPSCellUSS_Sets_r17_Supported = 0
)

var cAParametersNRV1700DciFormatsPCellPSCellUSSSetsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1700_Dci_FormatsPCellPSCellUSS_Sets_r17_Supported},
}

const (
	CA_ParametersNR_v1700_DisablingScalingFactorDeactSCell_r17_Supported = 0
)

var cAParametersNRV1700DisablingScalingFactorDeactSCellR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1700_DisablingScalingFactorDeactSCell_r17_Supported},
}

const (
	CA_ParametersNR_v1700_DisablingScalingFactorDormantSCell_r17_Supported = 0
)

var cAParametersNRV1700DisablingScalingFactorDormantSCellR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1700_DisablingScalingFactorDormantSCell_r17_Supported},
}

const (
	CA_ParametersNR_v1700_MTRP_CSI_EnhancementPerBC_r17_CSI_Report_Mode_r17_Mode1 = 0
	CA_ParametersNR_v1700_MTRP_CSI_EnhancementPerBC_r17_CSI_Report_Mode_r17_Mode2 = 1
	CA_ParametersNR_v1700_MTRP_CSI_EnhancementPerBC_r17_CSI_Report_Mode_r17_Both  = 2
)

var cAParametersNRV1700MTRPCSIEnhancementPerBCR17CSIReportModeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1700_MTRP_CSI_EnhancementPerBC_r17_CSI_Report_Mode_r17_Mode1, CA_ParametersNR_v1700_MTRP_CSI_EnhancementPerBC_r17_CSI_Report_Mode_r17_Mode2, CA_ParametersNR_v1700_MTRP_CSI_EnhancementPerBC_r17_CSI_Report_Mode_r17_Both},
}

var cAParametersNRV1700MTRPCSIEnhancementPerBCR17SupportedComboAcrossCCsR17Constraints = per.SizeRange(1, 16)

const (
	CA_ParametersNR_v1700_MTRP_CSI_EnhancementPerBC_r17_CodebookMode_NCJT_r17_Mode1     = 0
	CA_ParametersNR_v1700_MTRP_CSI_EnhancementPerBC_r17_CodebookMode_NCJT_r17_Mode1And2 = 1
)

var cAParametersNRV1700MTRPCSIEnhancementPerBCR17CodebookModeNCJTR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1700_MTRP_CSI_EnhancementPerBC_r17_CodebookMode_NCJT_r17_Mode1, CA_ParametersNR_v1700_MTRP_CSI_EnhancementPerBC_r17_CodebookMode_NCJT_r17_Mode1And2},
}

var cAParametersNRV1700NonAlignedFrameBoundariesR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs15kHz-15kHz-r17", Optional: true},
		{Name: "scs15kHz-30kHz-r17", Optional: true},
		{Name: "scs15kHz-60kHz-r17", Optional: true},
		{Name: "scs30kHz-30kHz-r17", Optional: true},
		{Name: "scs30kHz-60kHz-r17", Optional: true},
		{Name: "scs60kHz-60kHz-r17", Optional: true},
	},
}

type CA_ParametersNR_v1700 struct {
	CodebookParametersfetype2PerBC_r17        *CodebookParametersfetype2PerBC_r17
	DemodulationEnhancementCA_r17             *int64
	MaxUplinkDutyCycle_InterBandCA_PC2_r17    *int64
	MaxUplinkDutyCycle_SULcombination_PC2_r17 *int64
	BeamManagementType_CBM_r17                *int64
	ParallelTxPUCCH_PUSCH_r17                 *int64
	CodebookComboParameterMixedTypePerBC_r17  *CodebookComboParameterMixedTypePerBC_r17
	MTRP_CSI_EnhancementPerBC_r17             *struct {
		MaxNumNZP_CSI_RS_r17        int64
		CSI_Report_Mode_r17         int64
		SupportedComboAcrossCCs_r17 []CSI_MultiTRP_SupportedCombinations_r17
		CodebookMode_NCJT_r17       int64
	}
	CodebookComboParameterMultiTRP_PerBC_r17    *CodebookComboParameterMultiTRP_PerBC_r17
	MaxCC_32_DL_HARQ_ProcessFR2_2_r17           *int64
	MaxCC_32_UL_HARQ_ProcessFR2_2_r17           *int64
	CrossCarrierSchedulingSCell_SpCellTypeB_r17 *CrossCarrierSchedulingSCell_SpCell_r17
	CrossCarrierSchedulingSCell_SpCellTypeA_r17 *CrossCarrierSchedulingSCell_SpCell_r17
	Dci_FormatsPCellPSCellUSS_Sets_r17          *int64
	DisablingScalingFactorDeactSCell_r17        *int64
	DisablingScalingFactorDormantSCell_r17      *int64
	Non_AlignedFrameBoundaries_r17              *struct {
		Scs15kHz_15kHz_r17 *per.BitString
		Scs15kHz_30kHz_r17 *per.BitString
		Scs15kHz_60kHz_r17 *per.BitString
		Scs30kHz_30kHz_r17 *per.BitString
		Scs30kHz_60kHz_r17 *per.BitString
		Scs60kHz_60kHz_r17 *per.BitString
	}
}

func (ie *CA_ParametersNR_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CodebookParametersfetype2PerBC_r17 != nil, ie.DemodulationEnhancementCA_r17 != nil, ie.MaxUplinkDutyCycle_InterBandCA_PC2_r17 != nil, ie.MaxUplinkDutyCycle_SULcombination_PC2_r17 != nil, ie.BeamManagementType_CBM_r17 != nil, ie.ParallelTxPUCCH_PUSCH_r17 != nil, ie.CodebookComboParameterMixedTypePerBC_r17 != nil, ie.MTRP_CSI_EnhancementPerBC_r17 != nil, ie.CodebookComboParameterMultiTRP_PerBC_r17 != nil, ie.MaxCC_32_DL_HARQ_ProcessFR2_2_r17 != nil, ie.MaxCC_32_UL_HARQ_ProcessFR2_2_r17 != nil, ie.CrossCarrierSchedulingSCell_SpCellTypeB_r17 != nil, ie.CrossCarrierSchedulingSCell_SpCellTypeA_r17 != nil, ie.Dci_FormatsPCellPSCellUSS_Sets_r17 != nil, ie.DisablingScalingFactorDeactSCell_r17 != nil, ie.DisablingScalingFactorDormantSCell_r17 != nil, ie.Non_AlignedFrameBoundaries_r17 != nil}); err != nil {
		return err
	}

	// 2. codebookParametersfetype2PerBC-r17: ref
	{
		if ie.CodebookParametersfetype2PerBC_r17 != nil {
			if err := ie.CodebookParametersfetype2PerBC_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. demodulationEnhancementCA-r17: enumerated
	{
		if ie.DemodulationEnhancementCA_r17 != nil {
			if err := e.EncodeEnumerated(*ie.DemodulationEnhancementCA_r17, cAParametersNRV1700DemodulationEnhancementCAR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. maxUplinkDutyCycle-interBandCA-PC2-r17: enumerated
	{
		if ie.MaxUplinkDutyCycle_InterBandCA_PC2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MaxUplinkDutyCycle_InterBandCA_PC2_r17, cAParametersNRV1700MaxUplinkDutyCycleInterBandCAPC2R17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. maxUplinkDutyCycle-SULcombination-PC2-r17: enumerated
	{
		if ie.MaxUplinkDutyCycle_SULcombination_PC2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MaxUplinkDutyCycle_SULcombination_PC2_r17, cAParametersNRV1700MaxUplinkDutyCycleSULcombinationPC2R17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. beamManagementType-CBM-r17: enumerated
	{
		if ie.BeamManagementType_CBM_r17 != nil {
			if err := e.EncodeEnumerated(*ie.BeamManagementType_CBM_r17, cAParametersNRV1700BeamManagementTypeCBMR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. parallelTxPUCCH-PUSCH-r17: enumerated
	{
		if ie.ParallelTxPUCCH_PUSCH_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ParallelTxPUCCH_PUSCH_r17, cAParametersNRV1700ParallelTxPUCCHPUSCHR17Constraints); err != nil {
				return err
			}
		}
	}

	// 8. codebookComboParameterMixedTypePerBC-r17: ref
	{
		if ie.CodebookComboParameterMixedTypePerBC_r17 != nil {
			if err := ie.CodebookComboParameterMixedTypePerBC_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. mTRP-CSI-EnhancementPerBC-r17: sequence
	{
		if ie.MTRP_CSI_EnhancementPerBC_r17 != nil {
			c := ie.MTRP_CSI_EnhancementPerBC_r17
			if err := e.EncodeInteger(c.MaxNumNZP_CSI_RS_r17, per.Constrained(2, 8)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.CSI_Report_Mode_r17, cAParametersNRV1700MTRPCSIEnhancementPerBCR17CSIReportModeR17Constraints); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(cAParametersNRV1700MTRPCSIEnhancementPerBCR17SupportedComboAcrossCCsR17Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedComboAcrossCCs_r17))); err != nil {
					return err
				}
				for i := range c.SupportedComboAcrossCCs_r17 {
					if err := c.SupportedComboAcrossCCs_r17[i].Encode(e); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeEnumerated(c.CodebookMode_NCJT_r17, cAParametersNRV1700MTRPCSIEnhancementPerBCR17CodebookModeNCJTR17Constraints); err != nil {
				return err
			}
		}
	}

	// 10. codebookComboParameterMultiTRP-PerBC-r17: ref
	{
		if ie.CodebookComboParameterMultiTRP_PerBC_r17 != nil {
			if err := ie.CodebookComboParameterMultiTRP_PerBC_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. maxCC-32-DL-HARQ-ProcessFR2-2-r17: enumerated
	{
		if ie.MaxCC_32_DL_HARQ_ProcessFR2_2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MaxCC_32_DL_HARQ_ProcessFR2_2_r17, cAParametersNRV1700MaxCC32DLHARQProcessFR22R17Constraints); err != nil {
				return err
			}
		}
	}

	// 12. maxCC-32-UL-HARQ-ProcessFR2-2-r17: enumerated
	{
		if ie.MaxCC_32_UL_HARQ_ProcessFR2_2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MaxCC_32_UL_HARQ_ProcessFR2_2_r17, cAParametersNRV1700MaxCC32ULHARQProcessFR22R17Constraints); err != nil {
				return err
			}
		}
	}

	// 13. crossCarrierSchedulingSCell-SpCellTypeB-r17: ref
	{
		if ie.CrossCarrierSchedulingSCell_SpCellTypeB_r17 != nil {
			if err := ie.CrossCarrierSchedulingSCell_SpCellTypeB_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 14. crossCarrierSchedulingSCell-SpCellTypeA-r17: ref
	{
		if ie.CrossCarrierSchedulingSCell_SpCellTypeA_r17 != nil {
			if err := ie.CrossCarrierSchedulingSCell_SpCellTypeA_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 15. dci-FormatsPCellPSCellUSS-Sets-r17: enumerated
	{
		if ie.Dci_FormatsPCellPSCellUSS_Sets_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Dci_FormatsPCellPSCellUSS_Sets_r17, cAParametersNRV1700DciFormatsPCellPSCellUSSSetsR17Constraints); err != nil {
				return err
			}
		}
	}

	// 16. disablingScalingFactorDeactSCell-r17: enumerated
	{
		if ie.DisablingScalingFactorDeactSCell_r17 != nil {
			if err := e.EncodeEnumerated(*ie.DisablingScalingFactorDeactSCell_r17, cAParametersNRV1700DisablingScalingFactorDeactSCellR17Constraints); err != nil {
				return err
			}
		}
	}

	// 17. disablingScalingFactorDormantSCell-r17: enumerated
	{
		if ie.DisablingScalingFactorDormantSCell_r17 != nil {
			if err := e.EncodeEnumerated(*ie.DisablingScalingFactorDormantSCell_r17, cAParametersNRV1700DisablingScalingFactorDormantSCellR17Constraints); err != nil {
				return err
			}
		}
	}

	// 18. non-AlignedFrameBoundaries-r17: sequence
	{
		if ie.Non_AlignedFrameBoundaries_r17 != nil {
			c := ie.Non_AlignedFrameBoundaries_r17
			cAParametersNRV1700NonAlignedFrameBoundariesR17Seq := e.NewSequenceEncoder(cAParametersNRV1700NonAlignedFrameBoundariesR17Constraints)
			if err := cAParametersNRV1700NonAlignedFrameBoundariesR17Seq.EncodePreamble([]bool{c.Scs15kHz_15kHz_r17 != nil, c.Scs15kHz_30kHz_r17 != nil, c.Scs15kHz_60kHz_r17 != nil, c.Scs30kHz_30kHz_r17 != nil, c.Scs30kHz_60kHz_r17 != nil, c.Scs60kHz_60kHz_r17 != nil}); err != nil {
				return err
			}
			if c.Scs15kHz_15kHz_r17 != nil {
				if err := e.EncodeBitString((*c.Scs15kHz_15kHz_r17), per.SizeRange(1, 496)); err != nil {
					return err
				}
			}
			if c.Scs15kHz_30kHz_r17 != nil {
				if err := e.EncodeBitString((*c.Scs15kHz_30kHz_r17), per.SizeRange(1, 496)); err != nil {
					return err
				}
			}
			if c.Scs15kHz_60kHz_r17 != nil {
				if err := e.EncodeBitString((*c.Scs15kHz_60kHz_r17), per.SizeRange(1, 496)); err != nil {
					return err
				}
			}
			if c.Scs30kHz_30kHz_r17 != nil {
				if err := e.EncodeBitString((*c.Scs30kHz_30kHz_r17), per.SizeRange(1, 496)); err != nil {
					return err
				}
			}
			if c.Scs30kHz_60kHz_r17 != nil {
				if err := e.EncodeBitString((*c.Scs30kHz_60kHz_r17), per.SizeRange(1, 496)); err != nil {
					return err
				}
			}
			if c.Scs60kHz_60kHz_r17 != nil {
				if err := e.EncodeBitString((*c.Scs60kHz_60kHz_r17), per.SizeRange(1, 496)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNR_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. codebookParametersfetype2PerBC-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.CodebookParametersfetype2PerBC_r17 = new(CodebookParametersfetype2PerBC_r17)
			if err := ie.CodebookParametersfetype2PerBC_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. demodulationEnhancementCA-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1700DemodulationEnhancementCAR17Constraints)
			if err != nil {
				return err
			}
			ie.DemodulationEnhancementCA_r17 = &idx
		}
	}

	// 4. maxUplinkDutyCycle-interBandCA-PC2-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1700MaxUplinkDutyCycleInterBandCAPC2R17Constraints)
			if err != nil {
				return err
			}
			ie.MaxUplinkDutyCycle_InterBandCA_PC2_r17 = &idx
		}
	}

	// 5. maxUplinkDutyCycle-SULcombination-PC2-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1700MaxUplinkDutyCycleSULcombinationPC2R17Constraints)
			if err != nil {
				return err
			}
			ie.MaxUplinkDutyCycle_SULcombination_PC2_r17 = &idx
		}
	}

	// 6. beamManagementType-CBM-r17: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1700BeamManagementTypeCBMR17Constraints)
			if err != nil {
				return err
			}
			ie.BeamManagementType_CBM_r17 = &idx
		}
	}

	// 7. parallelTxPUCCH-PUSCH-r17: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1700ParallelTxPUCCHPUSCHR17Constraints)
			if err != nil {
				return err
			}
			ie.ParallelTxPUCCH_PUSCH_r17 = &idx
		}
	}

	// 8. codebookComboParameterMixedTypePerBC-r17: ref
	{
		if seq.IsComponentPresent(6) {
			ie.CodebookComboParameterMixedTypePerBC_r17 = new(CodebookComboParameterMixedTypePerBC_r17)
			if err := ie.CodebookComboParameterMixedTypePerBC_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. mTRP-CSI-EnhancementPerBC-r17: sequence
	{
		if seq.IsComponentPresent(7) {
			ie.MTRP_CSI_EnhancementPerBC_r17 = &struct {
				MaxNumNZP_CSI_RS_r17        int64
				CSI_Report_Mode_r17         int64
				SupportedComboAcrossCCs_r17 []CSI_MultiTRP_SupportedCombinations_r17
				CodebookMode_NCJT_r17       int64
			}{}
			c := ie.MTRP_CSI_EnhancementPerBC_r17
			{
				v, err := d.DecodeInteger(per.Constrained(2, 8))
				if err != nil {
					return err
				}
				c.MaxNumNZP_CSI_RS_r17 = v
			}
			{
				v, err := d.DecodeEnumerated(cAParametersNRV1700MTRPCSIEnhancementPerBCR17CSIReportModeR17Constraints)
				if err != nil {
					return err
				}
				c.CSI_Report_Mode_r17 = v
			}
			{
				seqOf := d.NewSequenceOfDecoder(cAParametersNRV1700MTRPCSIEnhancementPerBCR17SupportedComboAcrossCCsR17Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedComboAcrossCCs_r17 = make([]CSI_MultiTRP_SupportedCombinations_r17, n)
				for i := int64(0); i < n; i++ {
					if err := c.SupportedComboAcrossCCs_r17[i].Decode(d); err != nil {
						return err
					}
				}
			}
			{
				v, err := d.DecodeEnumerated(cAParametersNRV1700MTRPCSIEnhancementPerBCR17CodebookModeNCJTR17Constraints)
				if err != nil {
					return err
				}
				c.CodebookMode_NCJT_r17 = v
			}
		}
	}

	// 10. codebookComboParameterMultiTRP-PerBC-r17: ref
	{
		if seq.IsComponentPresent(8) {
			ie.CodebookComboParameterMultiTRP_PerBC_r17 = new(CodebookComboParameterMultiTRP_PerBC_r17)
			if err := ie.CodebookComboParameterMultiTRP_PerBC_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. maxCC-32-DL-HARQ-ProcessFR2-2-r17: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1700MaxCC32DLHARQProcessFR22R17Constraints)
			if err != nil {
				return err
			}
			ie.MaxCC_32_DL_HARQ_ProcessFR2_2_r17 = &idx
		}
	}

	// 12. maxCC-32-UL-HARQ-ProcessFR2-2-r17: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1700MaxCC32ULHARQProcessFR22R17Constraints)
			if err != nil {
				return err
			}
			ie.MaxCC_32_UL_HARQ_ProcessFR2_2_r17 = &idx
		}
	}

	// 13. crossCarrierSchedulingSCell-SpCellTypeB-r17: ref
	{
		if seq.IsComponentPresent(11) {
			ie.CrossCarrierSchedulingSCell_SpCellTypeB_r17 = new(CrossCarrierSchedulingSCell_SpCell_r17)
			if err := ie.CrossCarrierSchedulingSCell_SpCellTypeB_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 14. crossCarrierSchedulingSCell-SpCellTypeA-r17: ref
	{
		if seq.IsComponentPresent(12) {
			ie.CrossCarrierSchedulingSCell_SpCellTypeA_r17 = new(CrossCarrierSchedulingSCell_SpCell_r17)
			if err := ie.CrossCarrierSchedulingSCell_SpCellTypeA_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 15. dci-FormatsPCellPSCellUSS-Sets-r17: enumerated
	{
		if seq.IsComponentPresent(13) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1700DciFormatsPCellPSCellUSSSetsR17Constraints)
			if err != nil {
				return err
			}
			ie.Dci_FormatsPCellPSCellUSS_Sets_r17 = &idx
		}
	}

	// 16. disablingScalingFactorDeactSCell-r17: enumerated
	{
		if seq.IsComponentPresent(14) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1700DisablingScalingFactorDeactSCellR17Constraints)
			if err != nil {
				return err
			}
			ie.DisablingScalingFactorDeactSCell_r17 = &idx
		}
	}

	// 17. disablingScalingFactorDormantSCell-r17: enumerated
	{
		if seq.IsComponentPresent(15) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1700DisablingScalingFactorDormantSCellR17Constraints)
			if err != nil {
				return err
			}
			ie.DisablingScalingFactorDormantSCell_r17 = &idx
		}
	}

	// 18. non-AlignedFrameBoundaries-r17: sequence
	{
		if seq.IsComponentPresent(16) {
			ie.Non_AlignedFrameBoundaries_r17 = &struct {
				Scs15kHz_15kHz_r17 *per.BitString
				Scs15kHz_30kHz_r17 *per.BitString
				Scs15kHz_60kHz_r17 *per.BitString
				Scs30kHz_30kHz_r17 *per.BitString
				Scs30kHz_60kHz_r17 *per.BitString
				Scs60kHz_60kHz_r17 *per.BitString
			}{}
			c := ie.Non_AlignedFrameBoundaries_r17
			cAParametersNRV1700NonAlignedFrameBoundariesR17Seq := d.NewSequenceDecoder(cAParametersNRV1700NonAlignedFrameBoundariesR17Constraints)
			if err := cAParametersNRV1700NonAlignedFrameBoundariesR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if cAParametersNRV1700NonAlignedFrameBoundariesR17Seq.IsComponentPresent(0) {
				c.Scs15kHz_15kHz_r17 = new(per.BitString)
				v, err := d.DecodeBitString(per.SizeRange(1, 496))
				if err != nil {
					return err
				}
				(*c.Scs15kHz_15kHz_r17) = v
			}
			if cAParametersNRV1700NonAlignedFrameBoundariesR17Seq.IsComponentPresent(1) {
				c.Scs15kHz_30kHz_r17 = new(per.BitString)
				v, err := d.DecodeBitString(per.SizeRange(1, 496))
				if err != nil {
					return err
				}
				(*c.Scs15kHz_30kHz_r17) = v
			}
			if cAParametersNRV1700NonAlignedFrameBoundariesR17Seq.IsComponentPresent(2) {
				c.Scs15kHz_60kHz_r17 = new(per.BitString)
				v, err := d.DecodeBitString(per.SizeRange(1, 496))
				if err != nil {
					return err
				}
				(*c.Scs15kHz_60kHz_r17) = v
			}
			if cAParametersNRV1700NonAlignedFrameBoundariesR17Seq.IsComponentPresent(3) {
				c.Scs30kHz_30kHz_r17 = new(per.BitString)
				v, err := d.DecodeBitString(per.SizeRange(1, 496))
				if err != nil {
					return err
				}
				(*c.Scs30kHz_30kHz_r17) = v
			}
			if cAParametersNRV1700NonAlignedFrameBoundariesR17Seq.IsComponentPresent(4) {
				c.Scs30kHz_60kHz_r17 = new(per.BitString)
				v, err := d.DecodeBitString(per.SizeRange(1, 496))
				if err != nil {
					return err
				}
				(*c.Scs30kHz_60kHz_r17) = v
			}
			if cAParametersNRV1700NonAlignedFrameBoundariesR17Seq.IsComponentPresent(5) {
				c.Scs60kHz_60kHz_r17 = new(per.BitString)
				v, err := d.DecodeBitString(per.SizeRange(1, 496))
				if err != nil {
					return err
				}
				(*c.Scs60kHz_60kHz_r17) = v
			}
		}
	}

	return nil
}
