// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BeamFailureRecoveryConfig (line 5108).

var beamFailureRecoveryConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rootSequenceIndex-BFR", Optional: true},
		{Name: "rach-ConfigBFR", Optional: true},
		{Name: "rsrp-ThresholdSSB", Optional: true},
		{Name: "candidateBeamRSList", Optional: true},
		{Name: "ssb-perRACH-Occasion", Optional: true},
		{Name: "ra-ssb-OccasionMaskIndex", Optional: true},
		{Name: "recoverySearchSpaceId", Optional: true},
		{Name: "ra-Prioritization", Optional: true},
		{Name: "beamFailureRecoveryTimer", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
	},
}

var beamFailureRecoveryConfigRootSequenceIndexBFRConstraints = per.Constrained(0, 137)

var beamFailureRecoveryConfigCandidateBeamRSListConstraints = per.SizeRange(1, common.MaxNrofCandidateBeams)

const (
	BeamFailureRecoveryConfig_Ssb_PerRACH_Occasion_OneEighth = 0
	BeamFailureRecoveryConfig_Ssb_PerRACH_Occasion_OneFourth = 1
	BeamFailureRecoveryConfig_Ssb_PerRACH_Occasion_OneHalf   = 2
	BeamFailureRecoveryConfig_Ssb_PerRACH_Occasion_One       = 3
	BeamFailureRecoveryConfig_Ssb_PerRACH_Occasion_Two       = 4
	BeamFailureRecoveryConfig_Ssb_PerRACH_Occasion_Four      = 5
	BeamFailureRecoveryConfig_Ssb_PerRACH_Occasion_Eight     = 6
	BeamFailureRecoveryConfig_Ssb_PerRACH_Occasion_Sixteen   = 7
)

var beamFailureRecoveryConfigSsbPerRACHOccasionConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BeamFailureRecoveryConfig_Ssb_PerRACH_Occasion_OneEighth, BeamFailureRecoveryConfig_Ssb_PerRACH_Occasion_OneFourth, BeamFailureRecoveryConfig_Ssb_PerRACH_Occasion_OneHalf, BeamFailureRecoveryConfig_Ssb_PerRACH_Occasion_One, BeamFailureRecoveryConfig_Ssb_PerRACH_Occasion_Two, BeamFailureRecoveryConfig_Ssb_PerRACH_Occasion_Four, BeamFailureRecoveryConfig_Ssb_PerRACH_Occasion_Eight, BeamFailureRecoveryConfig_Ssb_PerRACH_Occasion_Sixteen},
}

var beamFailureRecoveryConfigRaSsbOccasionMaskIndexConstraints = per.Constrained(0, 15)

const (
	BeamFailureRecoveryConfig_BeamFailureRecoveryTimer_Ms10  = 0
	BeamFailureRecoveryConfig_BeamFailureRecoveryTimer_Ms20  = 1
	BeamFailureRecoveryConfig_BeamFailureRecoveryTimer_Ms40  = 2
	BeamFailureRecoveryConfig_BeamFailureRecoveryTimer_Ms60  = 3
	BeamFailureRecoveryConfig_BeamFailureRecoveryTimer_Ms80  = 4
	BeamFailureRecoveryConfig_BeamFailureRecoveryTimer_Ms100 = 5
	BeamFailureRecoveryConfig_BeamFailureRecoveryTimer_Ms150 = 6
	BeamFailureRecoveryConfig_BeamFailureRecoveryTimer_Ms200 = 7
)

var beamFailureRecoveryConfigBeamFailureRecoveryTimerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BeamFailureRecoveryConfig_BeamFailureRecoveryTimer_Ms10, BeamFailureRecoveryConfig_BeamFailureRecoveryTimer_Ms20, BeamFailureRecoveryConfig_BeamFailureRecoveryTimer_Ms40, BeamFailureRecoveryConfig_BeamFailureRecoveryTimer_Ms60, BeamFailureRecoveryConfig_BeamFailureRecoveryTimer_Ms80, BeamFailureRecoveryConfig_BeamFailureRecoveryTimer_Ms100, BeamFailureRecoveryConfig_BeamFailureRecoveryTimer_Ms150, BeamFailureRecoveryConfig_BeamFailureRecoveryTimer_Ms200},
}

var beamFailureRecoveryConfigExtCandidateBeamRSListExtV1610Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BeamFailureRecoveryConfig_Ext_CandidateBeamRSListExt_v1610_Release = 0
	BeamFailureRecoveryConfig_Ext_CandidateBeamRSListExt_v1610_Setup   = 1
)

const (
	BeamFailureRecoveryConfig_Ext_SpCell_BFR_CBRA_r16_True = 0
)

var beamFailureRecoveryConfigExtSpCellBFRCBRAR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BeamFailureRecoveryConfig_Ext_SpCell_BFR_CBRA_r16_True},
}

const (
	BeamFailureRecoveryConfig_Ext_Ra_OccasionType_r19_Sbfd = 0
)

var beamFailureRecoveryConfigExtRaOccasionTypeR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BeamFailureRecoveryConfig_Ext_Ra_OccasionType_r19_Sbfd},
}

type BeamFailureRecoveryConfig struct {
	RootSequenceIndex_BFR        *int64
	Rach_ConfigBFR               *RACH_ConfigGeneric
	Rsrp_ThresholdSSB            *RSRP_Range
	CandidateBeamRSList          []PRACH_ResourceDedicatedBFR
	Ssb_PerRACH_Occasion         *int64
	Ra_Ssb_OccasionMaskIndex     *int64
	RecoverySearchSpaceId        *SearchSpaceId
	Ra_Prioritization            *RA_Prioritization
	BeamFailureRecoveryTimer     *int64
	Msg1_SubcarrierSpacing       *SubcarrierSpacing
	Ra_PrioritizationTwoStep_r16 *RA_Prioritization
	CandidateBeamRSListExt_v1610 *struct {
		Choice  int
		Release *struct{}
		Setup   *CandidateBeamRSListExt_r16
	}
	SpCell_BFR_CBRA_r16 *int64
	Ra_OccasionType_r19 *int64
}

func (ie *BeamFailureRecoveryConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(beamFailureRecoveryConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Msg1_SubcarrierSpacing != nil
	hasExtGroup1 := ie.Ra_PrioritizationTwoStep_r16 != nil || ie.CandidateBeamRSListExt_v1610 != nil
	hasExtGroup2 := ie.SpCell_BFR_CBRA_r16 != nil
	hasExtGroup3 := ie.Ra_OccasionType_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RootSequenceIndex_BFR != nil, ie.Rach_ConfigBFR != nil, ie.Rsrp_ThresholdSSB != nil, ie.CandidateBeamRSList != nil, ie.Ssb_PerRACH_Occasion != nil, ie.Ra_Ssb_OccasionMaskIndex != nil, ie.RecoverySearchSpaceId != nil, ie.Ra_Prioritization != nil, ie.BeamFailureRecoveryTimer != nil}); err != nil {
		return err
	}

	// 3. rootSequenceIndex-BFR: integer
	{
		if ie.RootSequenceIndex_BFR != nil {
			if err := e.EncodeInteger(*ie.RootSequenceIndex_BFR, beamFailureRecoveryConfigRootSequenceIndexBFRConstraints); err != nil {
				return err
			}
		}
	}

	// 4. rach-ConfigBFR: ref
	{
		if ie.Rach_ConfigBFR != nil {
			if err := ie.Rach_ConfigBFR.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. rsrp-ThresholdSSB: ref
	{
		if ie.Rsrp_ThresholdSSB != nil {
			if err := ie.Rsrp_ThresholdSSB.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. candidateBeamRSList: sequence-of
	{
		if ie.CandidateBeamRSList != nil {
			seqOf := e.NewSequenceOfEncoder(beamFailureRecoveryConfigCandidateBeamRSListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.CandidateBeamRSList))); err != nil {
				return err
			}
			for i := range ie.CandidateBeamRSList {
				if err := ie.CandidateBeamRSList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. ssb-perRACH-Occasion: enumerated
	{
		if ie.Ssb_PerRACH_Occasion != nil {
			if err := e.EncodeEnumerated(*ie.Ssb_PerRACH_Occasion, beamFailureRecoveryConfigSsbPerRACHOccasionConstraints); err != nil {
				return err
			}
		}
	}

	// 8. ra-ssb-OccasionMaskIndex: integer
	{
		if ie.Ra_Ssb_OccasionMaskIndex != nil {
			if err := e.EncodeInteger(*ie.Ra_Ssb_OccasionMaskIndex, beamFailureRecoveryConfigRaSsbOccasionMaskIndexConstraints); err != nil {
				return err
			}
		}
	}

	// 9. recoverySearchSpaceId: ref
	{
		if ie.RecoverySearchSpaceId != nil {
			if err := ie.RecoverySearchSpaceId.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. ra-Prioritization: ref
	{
		if ie.Ra_Prioritization != nil {
			if err := ie.Ra_Prioritization.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. beamFailureRecoveryTimer: enumerated
	{
		if ie.BeamFailureRecoveryTimer != nil {
			if err := e.EncodeEnumerated(*ie.BeamFailureRecoveryTimer, beamFailureRecoveryConfigBeamFailureRecoveryTimerConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "msg1-SubcarrierSpacing", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Msg1_SubcarrierSpacing != nil}); err != nil {
				return err
			}

			if ie.Msg1_SubcarrierSpacing != nil {
				if err := ie.Msg1_SubcarrierSpacing.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "ra-PrioritizationTwoStep-r16", Optional: true},
					{Name: "candidateBeamRSListExt-v1610", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ra_PrioritizationTwoStep_r16 != nil, ie.CandidateBeamRSListExt_v1610 != nil}); err != nil {
				return err
			}

			if ie.Ra_PrioritizationTwoStep_r16 != nil {
				if err := ie.Ra_PrioritizationTwoStep_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CandidateBeamRSListExt_v1610 != nil {
				choiceEnc := ex.NewChoiceEncoder(beamFailureRecoveryConfigExtCandidateBeamRSListExtV1610Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.CandidateBeamRSListExt_v1610).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.CandidateBeamRSListExt_v1610).Choice {
				case BeamFailureRecoveryConfig_Ext_CandidateBeamRSListExt_v1610_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case BeamFailureRecoveryConfig_Ext_CandidateBeamRSListExt_v1610_Setup:
					if err := (*ie.CandidateBeamRSListExt_v1610).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "spCell-BFR-CBRA-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SpCell_BFR_CBRA_r16 != nil}); err != nil {
				return err
			}

			if ie.SpCell_BFR_CBRA_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SpCell_BFR_CBRA_r16, beamFailureRecoveryConfigExtSpCellBFRCBRAR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "ra-OccasionType-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ra_OccasionType_r19 != nil}); err != nil {
				return err
			}

			if ie.Ra_OccasionType_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ra_OccasionType_r19, beamFailureRecoveryConfigExtRaOccasionTypeR19Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *BeamFailureRecoveryConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(beamFailureRecoveryConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. rootSequenceIndex-BFR: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(beamFailureRecoveryConfigRootSequenceIndexBFRConstraints)
			if err != nil {
				return err
			}
			ie.RootSequenceIndex_BFR = &v
		}
	}

	// 4. rach-ConfigBFR: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Rach_ConfigBFR = new(RACH_ConfigGeneric)
			if err := ie.Rach_ConfigBFR.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. rsrp-ThresholdSSB: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Rsrp_ThresholdSSB = new(RSRP_Range)
			if err := ie.Rsrp_ThresholdSSB.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. candidateBeamRSList: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(beamFailureRecoveryConfigCandidateBeamRSListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.CandidateBeamRSList = make([]PRACH_ResourceDedicatedBFR, n)
			for i := int64(0); i < n; i++ {
				if err := ie.CandidateBeamRSList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. ssb-perRACH-Occasion: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(beamFailureRecoveryConfigSsbPerRACHOccasionConstraints)
			if err != nil {
				return err
			}
			ie.Ssb_PerRACH_Occasion = &idx
		}
	}

	// 8. ra-ssb-OccasionMaskIndex: integer
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeInteger(beamFailureRecoveryConfigRaSsbOccasionMaskIndexConstraints)
			if err != nil {
				return err
			}
			ie.Ra_Ssb_OccasionMaskIndex = &v
		}
	}

	// 9. recoverySearchSpaceId: ref
	{
		if seq.IsComponentPresent(6) {
			ie.RecoverySearchSpaceId = new(SearchSpaceId)
			if err := ie.RecoverySearchSpaceId.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. ra-Prioritization: ref
	{
		if seq.IsComponentPresent(7) {
			ie.Ra_Prioritization = new(RA_Prioritization)
			if err := ie.Ra_Prioritization.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. beamFailureRecoveryTimer: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(beamFailureRecoveryConfigBeamFailureRecoveryTimerConstraints)
			if err != nil {
				return err
			}
			ie.BeamFailureRecoveryTimer = &idx
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "msg1-SubcarrierSpacing", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Msg1_SubcarrierSpacing = new(SubcarrierSpacing)
			if err := ie.Msg1_SubcarrierSpacing.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ra-PrioritizationTwoStep-r16", Optional: true},
				{Name: "candidateBeamRSListExt-v1610", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ra_PrioritizationTwoStep_r16 = new(RA_Prioritization)
			if err := ie.Ra_PrioritizationTwoStep_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.CandidateBeamRSListExt_v1610 = &struct {
				Choice  int
				Release *struct{}
				Setup   *CandidateBeamRSListExt_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(beamFailureRecoveryConfigExtCandidateBeamRSListExtV1610Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.CandidateBeamRSListExt_v1610).Choice = int(index)
			switch index {
			case BeamFailureRecoveryConfig_Ext_CandidateBeamRSListExt_v1610_Release:
				(*ie.CandidateBeamRSListExt_v1610).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case BeamFailureRecoveryConfig_Ext_CandidateBeamRSListExt_v1610_Setup:
				(*ie.CandidateBeamRSListExt_v1610).Setup = new(CandidateBeamRSListExt_r16)
				if err := (*ie.CandidateBeamRSListExt_v1610).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "spCell-BFR-CBRA-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(beamFailureRecoveryConfigExtSpCellBFRCBRAR16Constraints)
			if err != nil {
				return err
			}
			ie.SpCell_BFR_CBRA_r16 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ra-OccasionType-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(beamFailureRecoveryConfigExtRaOccasionTypeR19Constraints)
			if err != nil {
				return err
			}
			ie.Ra_OccasionType_r19 = &v
		}
	}

	return nil
}
