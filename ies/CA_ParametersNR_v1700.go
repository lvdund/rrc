package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1700 struct {
	CodebookParametersfetype2PerBC_r17          *CodebookParametersfetype2PerBC_r17                              `optional`
	DemodulationEnhancementCA_r17               *CA_ParametersNR_v1700_demodulationEnhancementCA_r17             `optional`
	MaxUplinkDutyCycle_interBandCA_PC2_r17      *CA_ParametersNR_v1700_maxUplinkDutyCycle_interBandCA_PC2_r17    `optional`
	MaxUplinkDutyCycle_SULcombination_PC2_r17   *CA_ParametersNR_v1700_maxUplinkDutyCycle_SULcombination_PC2_r17 `optional`
	BeamManagementType_CBM_r17                  *CA_ParametersNR_v1700_beamManagementType_CBM_r17                `optional`
	ParallelTxPUCCH_PUSCH_r17                   *CA_ParametersNR_v1700_parallelTxPUCCH_PUSCH_r17                 `optional`
	CodebookComboParameterMixedTypePerBC_r17    *CodebookComboParameterMixedTypePerBC_r17                        `optional`
	MTRP_CSI_EnhancementPerBC_r17               *CA_ParametersNR_v1700_mTRP_CSI_EnhancementPerBC_r17             `lb:1,ub:16,optional`
	CodebookComboParameterMultiTRP_PerBC_r17    *CodebookComboParameterMultiTRP_PerBC_r17                        `optional`
	MaxCC_32_DL_HARQ_ProcessFR2_2_r17           *CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17         `optional`
	MaxCC_32_UL_HARQ_ProcessFR2_2_r17           *CA_ParametersNR_v1700_maxCC_32_UL_HARQ_ProcessFR2_2_r17         `optional`
	CrossCarrierSchedulingSCell_SpCellTypeB_r17 *CrossCarrierSchedulingSCell_SpCell_r17                          `optional`
	CrossCarrierSchedulingSCell_SpCellTypeA_r17 *CrossCarrierSchedulingSCell_SpCell_r17                          `optional`
	Dci_FormatsPCellPSCellUSS_Sets_r17          *CA_ParametersNR_v1700_dci_FormatsPCellPSCellUSS_Sets_r17        `optional`
	DisablingScalingFactorDeactSCell_r17        *CA_ParametersNR_v1700_disablingScalingFactorDeactSCell_r17      `optional`
	DisablingScalingFactorDormantSCell_r17      *CA_ParametersNR_v1700_disablingScalingFactorDormantSCell_r17    `optional`
	Non_AlignedFrameBoundaries_r17              *CA_ParametersNR_v1700_non_AlignedFrameBoundaries_r17            `lb:1,ub:496,optional`
}

func (ie *CA_ParametersNR_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.CodebookParametersfetype2PerBC_r17 != nil, ie.DemodulationEnhancementCA_r17 != nil, ie.MaxUplinkDutyCycle_interBandCA_PC2_r17 != nil, ie.MaxUplinkDutyCycle_SULcombination_PC2_r17 != nil, ie.BeamManagementType_CBM_r17 != nil, ie.ParallelTxPUCCH_PUSCH_r17 != nil, ie.CodebookComboParameterMixedTypePerBC_r17 != nil, ie.MTRP_CSI_EnhancementPerBC_r17 != nil, ie.CodebookComboParameterMultiTRP_PerBC_r17 != nil, ie.MaxCC_32_DL_HARQ_ProcessFR2_2_r17 != nil, ie.MaxCC_32_UL_HARQ_ProcessFR2_2_r17 != nil, ie.CrossCarrierSchedulingSCell_SpCellTypeB_r17 != nil, ie.CrossCarrierSchedulingSCell_SpCellTypeA_r17 != nil, ie.Dci_FormatsPCellPSCellUSS_Sets_r17 != nil, ie.DisablingScalingFactorDeactSCell_r17 != nil, ie.DisablingScalingFactorDormantSCell_r17 != nil, ie.Non_AlignedFrameBoundaries_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CodebookParametersfetype2PerBC_r17 != nil {
		if err = ie.CodebookParametersfetype2PerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CodebookParametersfetype2PerBC_r17", err)
		}
	}
	if ie.DemodulationEnhancementCA_r17 != nil {
		if err = ie.DemodulationEnhancementCA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode DemodulationEnhancementCA_r17", err)
		}
	}
	if ie.MaxUplinkDutyCycle_interBandCA_PC2_r17 != nil {
		if err = ie.MaxUplinkDutyCycle_interBandCA_PC2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxUplinkDutyCycle_interBandCA_PC2_r17", err)
		}
	}
	if ie.MaxUplinkDutyCycle_SULcombination_PC2_r17 != nil {
		if err = ie.MaxUplinkDutyCycle_SULcombination_PC2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxUplinkDutyCycle_SULcombination_PC2_r17", err)
		}
	}
	if ie.BeamManagementType_CBM_r17 != nil {
		if err = ie.BeamManagementType_CBM_r17.Encode(w); err != nil {
			return utils.WrapError("Encode BeamManagementType_CBM_r17", err)
		}
	}
	if ie.ParallelTxPUCCH_PUSCH_r17 != nil {
		if err = ie.ParallelTxPUCCH_PUSCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ParallelTxPUCCH_PUSCH_r17", err)
		}
	}
	if ie.CodebookComboParameterMixedTypePerBC_r17 != nil {
		if err = ie.CodebookComboParameterMixedTypePerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CodebookComboParameterMixedTypePerBC_r17", err)
		}
	}
	if ie.MTRP_CSI_EnhancementPerBC_r17 != nil {
		if err = ie.MTRP_CSI_EnhancementPerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MTRP_CSI_EnhancementPerBC_r17", err)
		}
	}
	if ie.CodebookComboParameterMultiTRP_PerBC_r17 != nil {
		if err = ie.CodebookComboParameterMultiTRP_PerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CodebookComboParameterMultiTRP_PerBC_r17", err)
		}
	}
	if ie.MaxCC_32_DL_HARQ_ProcessFR2_2_r17 != nil {
		if err = ie.MaxCC_32_DL_HARQ_ProcessFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxCC_32_DL_HARQ_ProcessFR2_2_r17", err)
		}
	}
	if ie.MaxCC_32_UL_HARQ_ProcessFR2_2_r17 != nil {
		if err = ie.MaxCC_32_UL_HARQ_ProcessFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxCC_32_UL_HARQ_ProcessFR2_2_r17", err)
		}
	}
	if ie.CrossCarrierSchedulingSCell_SpCellTypeB_r17 != nil {
		if err = ie.CrossCarrierSchedulingSCell_SpCellTypeB_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CrossCarrierSchedulingSCell_SpCellTypeB_r17", err)
		}
	}
	if ie.CrossCarrierSchedulingSCell_SpCellTypeA_r17 != nil {
		if err = ie.CrossCarrierSchedulingSCell_SpCellTypeA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CrossCarrierSchedulingSCell_SpCellTypeA_r17", err)
		}
	}
	if ie.Dci_FormatsPCellPSCellUSS_Sets_r17 != nil {
		if err = ie.Dci_FormatsPCellPSCellUSS_Sets_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Dci_FormatsPCellPSCellUSS_Sets_r17", err)
		}
	}
	if ie.DisablingScalingFactorDeactSCell_r17 != nil {
		if err = ie.DisablingScalingFactorDeactSCell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode DisablingScalingFactorDeactSCell_r17", err)
		}
	}
	if ie.DisablingScalingFactorDormantSCell_r17 != nil {
		if err = ie.DisablingScalingFactorDormantSCell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode DisablingScalingFactorDormantSCell_r17", err)
		}
	}
	if ie.Non_AlignedFrameBoundaries_r17 != nil {
		if err = ie.Non_AlignedFrameBoundaries_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Non_AlignedFrameBoundaries_r17", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1700) Decode(r *aper.AperReader) error {
	var err error
	var CodebookParametersfetype2PerBC_r17Present bool
	if CodebookParametersfetype2PerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DemodulationEnhancementCA_r17Present bool
	if DemodulationEnhancementCA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxUplinkDutyCycle_interBandCA_PC2_r17Present bool
	if MaxUplinkDutyCycle_interBandCA_PC2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxUplinkDutyCycle_SULcombination_PC2_r17Present bool
	if MaxUplinkDutyCycle_SULcombination_PC2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BeamManagementType_CBM_r17Present bool
	if BeamManagementType_CBM_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ParallelTxPUCCH_PUSCH_r17Present bool
	if ParallelTxPUCCH_PUSCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CodebookComboParameterMixedTypePerBC_r17Present bool
	if CodebookComboParameterMixedTypePerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MTRP_CSI_EnhancementPerBC_r17Present bool
	if MTRP_CSI_EnhancementPerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CodebookComboParameterMultiTRP_PerBC_r17Present bool
	if CodebookComboParameterMultiTRP_PerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxCC_32_DL_HARQ_ProcessFR2_2_r17Present bool
	if MaxCC_32_DL_HARQ_ProcessFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxCC_32_UL_HARQ_ProcessFR2_2_r17Present bool
	if MaxCC_32_UL_HARQ_ProcessFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CrossCarrierSchedulingSCell_SpCellTypeB_r17Present bool
	if CrossCarrierSchedulingSCell_SpCellTypeB_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CrossCarrierSchedulingSCell_SpCellTypeA_r17Present bool
	if CrossCarrierSchedulingSCell_SpCellTypeA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dci_FormatsPCellPSCellUSS_Sets_r17Present bool
	if Dci_FormatsPCellPSCellUSS_Sets_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DisablingScalingFactorDeactSCell_r17Present bool
	if DisablingScalingFactorDeactSCell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DisablingScalingFactorDormantSCell_r17Present bool
	if DisablingScalingFactorDormantSCell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Non_AlignedFrameBoundaries_r17Present bool
	if Non_AlignedFrameBoundaries_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if CodebookParametersfetype2PerBC_r17Present {
		ie.CodebookParametersfetype2PerBC_r17 = new(CodebookParametersfetype2PerBC_r17)
		if err = ie.CodebookParametersfetype2PerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CodebookParametersfetype2PerBC_r17", err)
		}
	}
	if DemodulationEnhancementCA_r17Present {
		ie.DemodulationEnhancementCA_r17 = new(CA_ParametersNR_v1700_demodulationEnhancementCA_r17)
		if err = ie.DemodulationEnhancementCA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode DemodulationEnhancementCA_r17", err)
		}
	}
	if MaxUplinkDutyCycle_interBandCA_PC2_r17Present {
		ie.MaxUplinkDutyCycle_interBandCA_PC2_r17 = new(CA_ParametersNR_v1700_maxUplinkDutyCycle_interBandCA_PC2_r17)
		if err = ie.MaxUplinkDutyCycle_interBandCA_PC2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxUplinkDutyCycle_interBandCA_PC2_r17", err)
		}
	}
	if MaxUplinkDutyCycle_SULcombination_PC2_r17Present {
		ie.MaxUplinkDutyCycle_SULcombination_PC2_r17 = new(CA_ParametersNR_v1700_maxUplinkDutyCycle_SULcombination_PC2_r17)
		if err = ie.MaxUplinkDutyCycle_SULcombination_PC2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxUplinkDutyCycle_SULcombination_PC2_r17", err)
		}
	}
	if BeamManagementType_CBM_r17Present {
		ie.BeamManagementType_CBM_r17 = new(CA_ParametersNR_v1700_beamManagementType_CBM_r17)
		if err = ie.BeamManagementType_CBM_r17.Decode(r); err != nil {
			return utils.WrapError("Decode BeamManagementType_CBM_r17", err)
		}
	}
	if ParallelTxPUCCH_PUSCH_r17Present {
		ie.ParallelTxPUCCH_PUSCH_r17 = new(CA_ParametersNR_v1700_parallelTxPUCCH_PUSCH_r17)
		if err = ie.ParallelTxPUCCH_PUSCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ParallelTxPUCCH_PUSCH_r17", err)
		}
	}
	if CodebookComboParameterMixedTypePerBC_r17Present {
		ie.CodebookComboParameterMixedTypePerBC_r17 = new(CodebookComboParameterMixedTypePerBC_r17)
		if err = ie.CodebookComboParameterMixedTypePerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CodebookComboParameterMixedTypePerBC_r17", err)
		}
	}
	if MTRP_CSI_EnhancementPerBC_r17Present {
		ie.MTRP_CSI_EnhancementPerBC_r17 = new(CA_ParametersNR_v1700_mTRP_CSI_EnhancementPerBC_r17)
		if err = ie.MTRP_CSI_EnhancementPerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MTRP_CSI_EnhancementPerBC_r17", err)
		}
	}
	if CodebookComboParameterMultiTRP_PerBC_r17Present {
		ie.CodebookComboParameterMultiTRP_PerBC_r17 = new(CodebookComboParameterMultiTRP_PerBC_r17)
		if err = ie.CodebookComboParameterMultiTRP_PerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CodebookComboParameterMultiTRP_PerBC_r17", err)
		}
	}
	if MaxCC_32_DL_HARQ_ProcessFR2_2_r17Present {
		ie.MaxCC_32_DL_HARQ_ProcessFR2_2_r17 = new(CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17)
		if err = ie.MaxCC_32_DL_HARQ_ProcessFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxCC_32_DL_HARQ_ProcessFR2_2_r17", err)
		}
	}
	if MaxCC_32_UL_HARQ_ProcessFR2_2_r17Present {
		ie.MaxCC_32_UL_HARQ_ProcessFR2_2_r17 = new(CA_ParametersNR_v1700_maxCC_32_UL_HARQ_ProcessFR2_2_r17)
		if err = ie.MaxCC_32_UL_HARQ_ProcessFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxCC_32_UL_HARQ_ProcessFR2_2_r17", err)
		}
	}
	if CrossCarrierSchedulingSCell_SpCellTypeB_r17Present {
		ie.CrossCarrierSchedulingSCell_SpCellTypeB_r17 = new(CrossCarrierSchedulingSCell_SpCell_r17)
		if err = ie.CrossCarrierSchedulingSCell_SpCellTypeB_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CrossCarrierSchedulingSCell_SpCellTypeB_r17", err)
		}
	}
	if CrossCarrierSchedulingSCell_SpCellTypeA_r17Present {
		ie.CrossCarrierSchedulingSCell_SpCellTypeA_r17 = new(CrossCarrierSchedulingSCell_SpCell_r17)
		if err = ie.CrossCarrierSchedulingSCell_SpCellTypeA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CrossCarrierSchedulingSCell_SpCellTypeA_r17", err)
		}
	}
	if Dci_FormatsPCellPSCellUSS_Sets_r17Present {
		ie.Dci_FormatsPCellPSCellUSS_Sets_r17 = new(CA_ParametersNR_v1700_dci_FormatsPCellPSCellUSS_Sets_r17)
		if err = ie.Dci_FormatsPCellPSCellUSS_Sets_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Dci_FormatsPCellPSCellUSS_Sets_r17", err)
		}
	}
	if DisablingScalingFactorDeactSCell_r17Present {
		ie.DisablingScalingFactorDeactSCell_r17 = new(CA_ParametersNR_v1700_disablingScalingFactorDeactSCell_r17)
		if err = ie.DisablingScalingFactorDeactSCell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode DisablingScalingFactorDeactSCell_r17", err)
		}
	}
	if DisablingScalingFactorDormantSCell_r17Present {
		ie.DisablingScalingFactorDormantSCell_r17 = new(CA_ParametersNR_v1700_disablingScalingFactorDormantSCell_r17)
		if err = ie.DisablingScalingFactorDormantSCell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode DisablingScalingFactorDormantSCell_r17", err)
		}
	}
	if Non_AlignedFrameBoundaries_r17Present {
		ie.Non_AlignedFrameBoundaries_r17 = new(CA_ParametersNR_v1700_non_AlignedFrameBoundaries_r17)
		if err = ie.Non_AlignedFrameBoundaries_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Non_AlignedFrameBoundaries_r17", err)
		}
	}
	return nil
}
