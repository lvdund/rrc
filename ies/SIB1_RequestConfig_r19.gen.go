// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SIB1-RequestConfig-r19 (line 4714).

var sIB1RequestConfigR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "totalNumberOfRA-Preambles-r19"},
		{Name: "rach-OccasionsSIB1-r19"},
		{Name: "offsetToCarrier-r19"},
		{Name: "carrierBandwidth-r19"},
		{Name: "sib1-RequestPeriod-r19", Optional: true},
		{Name: "sib1-RequestResources-r19"},
		{Name: "sib1-PDCCH-RestrictionToPRACH-r19", Optional: true},
		{Name: "rach-ConfigSIB1-r19"},
		{Name: "ra-SearchSpace-r19", Optional: true},
		{Name: "n-TimingAdvanceOffset-r19", Optional: true},
		{Name: "msg1-SubcarrierSpacing-r19", Optional: true},
		{Name: "prach-RootSequenceIndex-r19"},
		{Name: "sib1-RSRP-ThresholdSSB-r19", Optional: true},
		{Name: "rsrp-ThresholdSSB-SUL", Optional: true},
		{Name: "locationAndBandwidth-r19"},
		{Name: "absoluteFrequencyPointA-r19", Optional: true},
		{Name: "ul-FrequencyBandList-r19", Optional: true},
		{Name: "ul-SubCarrierSpacing-r19"},
		{Name: "p-Max-r19", Optional: true},
		{Name: "sib1-RestrictedSetConfig-r19"},
	},
}

var sIB1RequestConfigR19TotalNumberOfRAPreamblesR19Constraints = per.Constrained(1, 63)

var sIB1RequestConfigR19OffsetToCarrierR19Constraints = per.Constrained(0, 2199)

var sIB1RequestConfigR19CarrierBandwidthR19Constraints = per.Constrained(1, common.MaxNrofPhysicalResourceBlocks)

const (
	SIB1_RequestConfig_r19_Sib1_RequestPeriod_r19_One     = 0
	SIB1_RequestConfig_r19_Sib1_RequestPeriod_r19_Two     = 1
	SIB1_RequestConfig_r19_Sib1_RequestPeriod_r19_Four    = 2
	SIB1_RequestConfig_r19_Sib1_RequestPeriod_r19_Six     = 3
	SIB1_RequestConfig_r19_Sib1_RequestPeriod_r19_Eight   = 4
	SIB1_RequestConfig_r19_Sib1_RequestPeriod_r19_Ten     = 5
	SIB1_RequestConfig_r19_Sib1_RequestPeriod_r19_Twelve  = 6
	SIB1_RequestConfig_r19_Sib1_RequestPeriod_r19_Sixteen = 7
)

var sIB1RequestConfigR19Sib1RequestPeriodR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_RequestConfig_r19_Sib1_RequestPeriod_r19_One, SIB1_RequestConfig_r19_Sib1_RequestPeriod_r19_Two, SIB1_RequestConfig_r19_Sib1_RequestPeriod_r19_Four, SIB1_RequestConfig_r19_Sib1_RequestPeriod_r19_Six, SIB1_RequestConfig_r19_Sib1_RequestPeriod_r19_Eight, SIB1_RequestConfig_r19_Sib1_RequestPeriod_r19_Ten, SIB1_RequestConfig_r19_Sib1_RequestPeriod_r19_Twelve, SIB1_RequestConfig_r19_Sib1_RequestPeriod_r19_Sixteen},
}

const (
	SIB1_RequestConfig_r19_Sib1_PDCCH_RestrictionToPRACH_r19_True = 0
)

var sIB1RequestConfigR19Sib1PDCCHRestrictionToPRACHR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_RequestConfig_r19_Sib1_PDCCH_RestrictionToPRACH_r19_True},
}

const (
	SIB1_RequestConfig_r19_N_TimingAdvanceOffset_r19_N0     = 0
	SIB1_RequestConfig_r19_N_TimingAdvanceOffset_r19_N25600 = 1
	SIB1_RequestConfig_r19_N_TimingAdvanceOffset_r19_N39936 = 2
	SIB1_RequestConfig_r19_N_TimingAdvanceOffset_r19_Spare1 = 3
)

var sIB1RequestConfigR19NTimingAdvanceOffsetR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_RequestConfig_r19_N_TimingAdvanceOffset_r19_N0, SIB1_RequestConfig_r19_N_TimingAdvanceOffset_r19_N25600, SIB1_RequestConfig_r19_N_TimingAdvanceOffset_r19_N39936, SIB1_RequestConfig_r19_N_TimingAdvanceOffset_r19_Spare1},
}

const (
	SIB1_RequestConfig_r19_Msg1_SubcarrierSpacing_r19_KHz1dot25 = 0
	SIB1_RequestConfig_r19_Msg1_SubcarrierSpacing_r19_KHz5      = 1
	SIB1_RequestConfig_r19_Msg1_SubcarrierSpacing_r19_KHz15     = 2
	SIB1_RequestConfig_r19_Msg1_SubcarrierSpacing_r19_KHz30     = 3
	SIB1_RequestConfig_r19_Msg1_SubcarrierSpacing_r19_KHz60     = 4
	SIB1_RequestConfig_r19_Msg1_SubcarrierSpacing_r19_KHz120    = 5
	SIB1_RequestConfig_r19_Msg1_SubcarrierSpacing_r19_Spare2    = 6
	SIB1_RequestConfig_r19_Msg1_SubcarrierSpacing_r19_Spare1    = 7
)

var sIB1RequestConfigR19Msg1SubcarrierSpacingR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_RequestConfig_r19_Msg1_SubcarrierSpacing_r19_KHz1dot25, SIB1_RequestConfig_r19_Msg1_SubcarrierSpacing_r19_KHz5, SIB1_RequestConfig_r19_Msg1_SubcarrierSpacing_r19_KHz15, SIB1_RequestConfig_r19_Msg1_SubcarrierSpacing_r19_KHz30, SIB1_RequestConfig_r19_Msg1_SubcarrierSpacing_r19_KHz60, SIB1_RequestConfig_r19_Msg1_SubcarrierSpacing_r19_KHz120, SIB1_RequestConfig_r19_Msg1_SubcarrierSpacing_r19_Spare2, SIB1_RequestConfig_r19_Msg1_SubcarrierSpacing_r19_Spare1},
}

var sIB1_RequestConfig_r19PrachRootSequenceIndexR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "l839"},
		{Name: "l139"},
	},
}

const (
	SIB1_RequestConfig_r19_Prach_RootSequenceIndex_r19_L839 = 0
	SIB1_RequestConfig_r19_Prach_RootSequenceIndex_r19_L139 = 1
)

var sIB1RequestConfigR19LocationAndBandwidthR19Constraints = per.Constrained(0, 37949)

const (
	SIB1_RequestConfig_r19_Ul_SubCarrierSpacing_r19_KHz15  = 0
	SIB1_RequestConfig_r19_Ul_SubCarrierSpacing_r19_KHz30  = 1
	SIB1_RequestConfig_r19_Ul_SubCarrierSpacing_r19_KHz60  = 2
	SIB1_RequestConfig_r19_Ul_SubCarrierSpacing_r19_KHz120 = 3
	SIB1_RequestConfig_r19_Ul_SubCarrierSpacing_r19_KHz240 = 4
	SIB1_RequestConfig_r19_Ul_SubCarrierSpacing_r19_KHz480 = 5
	SIB1_RequestConfig_r19_Ul_SubCarrierSpacing_r19_KHz960 = 6
	SIB1_RequestConfig_r19_Ul_SubCarrierSpacing_r19_Spare1 = 7
)

var sIB1RequestConfigR19UlSubCarrierSpacingR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_RequestConfig_r19_Ul_SubCarrierSpacing_r19_KHz15, SIB1_RequestConfig_r19_Ul_SubCarrierSpacing_r19_KHz30, SIB1_RequestConfig_r19_Ul_SubCarrierSpacing_r19_KHz60, SIB1_RequestConfig_r19_Ul_SubCarrierSpacing_r19_KHz120, SIB1_RequestConfig_r19_Ul_SubCarrierSpacing_r19_KHz240, SIB1_RequestConfig_r19_Ul_SubCarrierSpacing_r19_KHz480, SIB1_RequestConfig_r19_Ul_SubCarrierSpacing_r19_KHz960, SIB1_RequestConfig_r19_Ul_SubCarrierSpacing_r19_Spare1},
}

const (
	SIB1_RequestConfig_r19_Sib1_RestrictedSetConfig_r19_UnrestrictedSet    = 0
	SIB1_RequestConfig_r19_Sib1_RestrictedSetConfig_r19_RestrictedSetTypeA = 1
	SIB1_RequestConfig_r19_Sib1_RestrictedSetConfig_r19_RestrictedSetTypeB = 2
	SIB1_RequestConfig_r19_Sib1_RestrictedSetConfig_r19_Spare1             = 3
)

var sIB1RequestConfigR19Sib1RestrictedSetConfigR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_RequestConfig_r19_Sib1_RestrictedSetConfig_r19_UnrestrictedSet, SIB1_RequestConfig_r19_Sib1_RestrictedSetConfig_r19_RestrictedSetTypeA, SIB1_RequestConfig_r19_Sib1_RestrictedSetConfig_r19_RestrictedSetTypeB, SIB1_RequestConfig_r19_Sib1_RestrictedSetConfig_r19_Spare1},
}

const (
	SIB1_RequestConfig_r19_Rach_OccasionsSIB1_r19_Ssb_PerRACH_Occasion_r19_OneEighth = 0
	SIB1_RequestConfig_r19_Rach_OccasionsSIB1_r19_Ssb_PerRACH_Occasion_r19_OneFourth = 1
	SIB1_RequestConfig_r19_Rach_OccasionsSIB1_r19_Ssb_PerRACH_Occasion_r19_OneHalf   = 2
	SIB1_RequestConfig_r19_Rach_OccasionsSIB1_r19_Ssb_PerRACH_Occasion_r19_One       = 3
	SIB1_RequestConfig_r19_Rach_OccasionsSIB1_r19_Ssb_PerRACH_Occasion_r19_Two       = 4
	SIB1_RequestConfig_r19_Rach_OccasionsSIB1_r19_Ssb_PerRACH_Occasion_r19_Four      = 5
	SIB1_RequestConfig_r19_Rach_OccasionsSIB1_r19_Ssb_PerRACH_Occasion_r19_Eight     = 6
	SIB1_RequestConfig_r19_Rach_OccasionsSIB1_r19_Ssb_PerRACH_Occasion_r19_Sixteen   = 7
)

var sIB1RequestConfigR19RachOccasionsSIB1R19SsbPerRACHOccasionR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_RequestConfig_r19_Rach_OccasionsSIB1_r19_Ssb_PerRACH_Occasion_r19_OneEighth, SIB1_RequestConfig_r19_Rach_OccasionsSIB1_r19_Ssb_PerRACH_Occasion_r19_OneFourth, SIB1_RequestConfig_r19_Rach_OccasionsSIB1_r19_Ssb_PerRACH_Occasion_r19_OneHalf, SIB1_RequestConfig_r19_Rach_OccasionsSIB1_r19_Ssb_PerRACH_Occasion_r19_One, SIB1_RequestConfig_r19_Rach_OccasionsSIB1_r19_Ssb_PerRACH_Occasion_r19_Two, SIB1_RequestConfig_r19_Rach_OccasionsSIB1_r19_Ssb_PerRACH_Occasion_r19_Four, SIB1_RequestConfig_r19_Rach_OccasionsSIB1_r19_Ssb_PerRACH_Occasion_r19_Eight, SIB1_RequestConfig_r19_Rach_OccasionsSIB1_r19_Ssb_PerRACH_Occasion_r19_Sixteen},
}

type SIB1_RequestConfig_r19 struct {
	TotalNumberOfRA_Preambles_r19     int64
	Rach_OccasionsSIB1_r19            struct{ Ssb_PerRACH_Occasion_r19 int64 }
	OffsetToCarrier_r19               int64
	CarrierBandwidth_r19              int64
	Sib1_RequestPeriod_r19            *int64
	Sib1_RequestResources_r19         SIB1_RequestResources_r19
	Sib1_PDCCH_RestrictionToPRACH_r19 *int64
	Rach_ConfigSIB1_r19               RACH_ConfigGeneric
	Ra_SearchSpace_r19                *SearchSpace
	N_TimingAdvanceOffset_r19         *int64
	Msg1_SubcarrierSpacing_r19        *int64
	Prach_RootSequenceIndex_r19       struct {
		Choice int
		L839   *int64
		L139   *int64
	}
	Sib1_RSRP_ThresholdSSB_r19   *RSRP_Range
	Rsrp_ThresholdSSB_SUL        *RSRP_Range
	LocationAndBandwidth_r19     int64
	AbsoluteFrequencyPointA_r19  *ARFCN_ValueNR
	Ul_FrequencyBandList_r19     *MultiFrequencyBandListNR_SIB
	Ul_SubCarrierSpacing_r19     int64
	P_Max_r19                    *P_Max
	Sib1_RestrictedSetConfig_r19 int64
}

func (ie *SIB1_RequestConfig_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB1RequestConfigR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sib1_RequestPeriod_r19 != nil, ie.Sib1_PDCCH_RestrictionToPRACH_r19 != nil, ie.Ra_SearchSpace_r19 != nil, ie.N_TimingAdvanceOffset_r19 != nil, ie.Msg1_SubcarrierSpacing_r19 != nil, ie.Sib1_RSRP_ThresholdSSB_r19 != nil, ie.Rsrp_ThresholdSSB_SUL != nil, ie.AbsoluteFrequencyPointA_r19 != nil, ie.Ul_FrequencyBandList_r19 != nil, ie.P_Max_r19 != nil}); err != nil {
		return err
	}

	// 2. totalNumberOfRA-Preambles-r19: integer
	{
		if err := e.EncodeInteger(ie.TotalNumberOfRA_Preambles_r19, sIB1RequestConfigR19TotalNumberOfRAPreamblesR19Constraints); err != nil {
			return err
		}
	}

	// 3. rach-OccasionsSIB1-r19: sequence
	{
		{
			c := &ie.Rach_OccasionsSIB1_r19
			if err := e.EncodeEnumerated(c.Ssb_PerRACH_Occasion_r19, sIB1RequestConfigR19RachOccasionsSIB1R19SsbPerRACHOccasionR19Constraints); err != nil {
				return err
			}
		}
	}

	// 4. offsetToCarrier-r19: integer
	{
		if err := e.EncodeInteger(ie.OffsetToCarrier_r19, sIB1RequestConfigR19OffsetToCarrierR19Constraints); err != nil {
			return err
		}
	}

	// 5. carrierBandwidth-r19: integer
	{
		if err := e.EncodeInteger(ie.CarrierBandwidth_r19, sIB1RequestConfigR19CarrierBandwidthR19Constraints); err != nil {
			return err
		}
	}

	// 6. sib1-RequestPeriod-r19: enumerated
	{
		if ie.Sib1_RequestPeriod_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Sib1_RequestPeriod_r19, sIB1RequestConfigR19Sib1RequestPeriodR19Constraints); err != nil {
				return err
			}
		}
	}

	// 7. sib1-RequestResources-r19: ref
	{
		if err := ie.Sib1_RequestResources_r19.Encode(e); err != nil {
			return err
		}
	}

	// 8. sib1-PDCCH-RestrictionToPRACH-r19: enumerated
	{
		if ie.Sib1_PDCCH_RestrictionToPRACH_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Sib1_PDCCH_RestrictionToPRACH_r19, sIB1RequestConfigR19Sib1PDCCHRestrictionToPRACHR19Constraints); err != nil {
				return err
			}
		}
	}

	// 9. rach-ConfigSIB1-r19: ref
	{
		if err := ie.Rach_ConfigSIB1_r19.Encode(e); err != nil {
			return err
		}
	}

	// 10. ra-SearchSpace-r19: ref
	{
		if ie.Ra_SearchSpace_r19 != nil {
			if err := ie.Ra_SearchSpace_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. n-TimingAdvanceOffset-r19: enumerated
	{
		if ie.N_TimingAdvanceOffset_r19 != nil {
			if err := e.EncodeEnumerated(*ie.N_TimingAdvanceOffset_r19, sIB1RequestConfigR19NTimingAdvanceOffsetR19Constraints); err != nil {
				return err
			}
		}
	}

	// 12. msg1-SubcarrierSpacing-r19: enumerated
	{
		if ie.Msg1_SubcarrierSpacing_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Msg1_SubcarrierSpacing_r19, sIB1RequestConfigR19Msg1SubcarrierSpacingR19Constraints); err != nil {
				return err
			}
		}
	}

	// 13. prach-RootSequenceIndex-r19: choice
	{
		choiceEnc := e.NewChoiceEncoder(sIB1_RequestConfig_r19PrachRootSequenceIndexR19Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Prach_RootSequenceIndex_r19.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Prach_RootSequenceIndex_r19.Choice {
		case SIB1_RequestConfig_r19_Prach_RootSequenceIndex_r19_L839:
			if err := e.EncodeInteger((*ie.Prach_RootSequenceIndex_r19.L839), per.Constrained(0, 837)); err != nil {
				return err
			}
		case SIB1_RequestConfig_r19_Prach_RootSequenceIndex_r19_L139:
			if err := e.EncodeInteger((*ie.Prach_RootSequenceIndex_r19.L139), per.Constrained(0, 137)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Prach_RootSequenceIndex_r19.Choice), Constraint: "index out of range"}
		}
	}

	// 14. sib1-RSRP-ThresholdSSB-r19: ref
	{
		if ie.Sib1_RSRP_ThresholdSSB_r19 != nil {
			if err := ie.Sib1_RSRP_ThresholdSSB_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 15. rsrp-ThresholdSSB-SUL: ref
	{
		if ie.Rsrp_ThresholdSSB_SUL != nil {
			if err := ie.Rsrp_ThresholdSSB_SUL.Encode(e); err != nil {
				return err
			}
		}
	}

	// 16. locationAndBandwidth-r19: integer
	{
		if err := e.EncodeInteger(ie.LocationAndBandwidth_r19, sIB1RequestConfigR19LocationAndBandwidthR19Constraints); err != nil {
			return err
		}
	}

	// 17. absoluteFrequencyPointA-r19: ref
	{
		if ie.AbsoluteFrequencyPointA_r19 != nil {
			if err := ie.AbsoluteFrequencyPointA_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 18. ul-FrequencyBandList-r19: ref
	{
		if ie.Ul_FrequencyBandList_r19 != nil {
			if err := ie.Ul_FrequencyBandList_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 19. ul-SubCarrierSpacing-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.Ul_SubCarrierSpacing_r19, sIB1RequestConfigR19UlSubCarrierSpacingR19Constraints); err != nil {
			return err
		}
	}

	// 20. p-Max-r19: ref
	{
		if ie.P_Max_r19 != nil {
			if err := ie.P_Max_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 21. sib1-RestrictedSetConfig-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sib1_RestrictedSetConfig_r19, sIB1RequestConfigR19Sib1RestrictedSetConfigR19Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SIB1_RequestConfig_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB1RequestConfigR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. totalNumberOfRA-Preambles-r19: integer
	{
		v0, err := d.DecodeInteger(sIB1RequestConfigR19TotalNumberOfRAPreamblesR19Constraints)
		if err != nil {
			return err
		}
		ie.TotalNumberOfRA_Preambles_r19 = v0
	}

	// 3. rach-OccasionsSIB1-r19: sequence
	{
		{
			c := &ie.Rach_OccasionsSIB1_r19
			{
				v, err := d.DecodeEnumerated(sIB1RequestConfigR19RachOccasionsSIB1R19SsbPerRACHOccasionR19Constraints)
				if err != nil {
					return err
				}
				c.Ssb_PerRACH_Occasion_r19 = v
			}
		}
	}

	// 4. offsetToCarrier-r19: integer
	{
		v2, err := d.DecodeInteger(sIB1RequestConfigR19OffsetToCarrierR19Constraints)
		if err != nil {
			return err
		}
		ie.OffsetToCarrier_r19 = v2
	}

	// 5. carrierBandwidth-r19: integer
	{
		v3, err := d.DecodeInteger(sIB1RequestConfigR19CarrierBandwidthR19Constraints)
		if err != nil {
			return err
		}
		ie.CarrierBandwidth_r19 = v3
	}

	// 6. sib1-RequestPeriod-r19: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(sIB1RequestConfigR19Sib1RequestPeriodR19Constraints)
			if err != nil {
				return err
			}
			ie.Sib1_RequestPeriod_r19 = &idx
		}
	}

	// 7. sib1-RequestResources-r19: ref
	{
		if err := ie.Sib1_RequestResources_r19.Decode(d); err != nil {
			return err
		}
	}

	// 8. sib1-PDCCH-RestrictionToPRACH-r19: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(sIB1RequestConfigR19Sib1PDCCHRestrictionToPRACHR19Constraints)
			if err != nil {
				return err
			}
			ie.Sib1_PDCCH_RestrictionToPRACH_r19 = &idx
		}
	}

	// 9. rach-ConfigSIB1-r19: ref
	{
		if err := ie.Rach_ConfigSIB1_r19.Decode(d); err != nil {
			return err
		}
	}

	// 10. ra-SearchSpace-r19: ref
	{
		if seq.IsComponentPresent(8) {
			ie.Ra_SearchSpace_r19 = new(SearchSpace)
			if err := ie.Ra_SearchSpace_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. n-TimingAdvanceOffset-r19: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(sIB1RequestConfigR19NTimingAdvanceOffsetR19Constraints)
			if err != nil {
				return err
			}
			ie.N_TimingAdvanceOffset_r19 = &idx
		}
	}

	// 12. msg1-SubcarrierSpacing-r19: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(sIB1RequestConfigR19Msg1SubcarrierSpacingR19Constraints)
			if err != nil {
				return err
			}
			ie.Msg1_SubcarrierSpacing_r19 = &idx
		}
	}

	// 13. prach-RootSequenceIndex-r19: choice
	{
		choiceDec := d.NewChoiceDecoder(sIB1_RequestConfig_r19PrachRootSequenceIndexR19Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Prach_RootSequenceIndex_r19.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SIB1_RequestConfig_r19_Prach_RootSequenceIndex_r19_L839:
			ie.Prach_RootSequenceIndex_r19.L839 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 837))
			if err != nil {
				return err
			}
			(*ie.Prach_RootSequenceIndex_r19.L839) = v
		case SIB1_RequestConfig_r19_Prach_RootSequenceIndex_r19_L139:
			ie.Prach_RootSequenceIndex_r19.L139 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 137))
			if err != nil {
				return err
			}
			(*ie.Prach_RootSequenceIndex_r19.L139) = v
		}
	}

	// 14. sib1-RSRP-ThresholdSSB-r19: ref
	{
		if seq.IsComponentPresent(12) {
			ie.Sib1_RSRP_ThresholdSSB_r19 = new(RSRP_Range)
			if err := ie.Sib1_RSRP_ThresholdSSB_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 15. rsrp-ThresholdSSB-SUL: ref
	{
		if seq.IsComponentPresent(13) {
			ie.Rsrp_ThresholdSSB_SUL = new(RSRP_Range)
			if err := ie.Rsrp_ThresholdSSB_SUL.Decode(d); err != nil {
				return err
			}
		}
	}

	// 16. locationAndBandwidth-r19: integer
	{
		v14, err := d.DecodeInteger(sIB1RequestConfigR19LocationAndBandwidthR19Constraints)
		if err != nil {
			return err
		}
		ie.LocationAndBandwidth_r19 = v14
	}

	// 17. absoluteFrequencyPointA-r19: ref
	{
		if seq.IsComponentPresent(15) {
			ie.AbsoluteFrequencyPointA_r19 = new(ARFCN_ValueNR)
			if err := ie.AbsoluteFrequencyPointA_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 18. ul-FrequencyBandList-r19: ref
	{
		if seq.IsComponentPresent(16) {
			ie.Ul_FrequencyBandList_r19 = new(MultiFrequencyBandListNR_SIB)
			if err := ie.Ul_FrequencyBandList_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 19. ul-SubCarrierSpacing-r19: enumerated
	{
		v17, err := d.DecodeEnumerated(sIB1RequestConfigR19UlSubCarrierSpacingR19Constraints)
		if err != nil {
			return err
		}
		ie.Ul_SubCarrierSpacing_r19 = v17
	}

	// 20. p-Max-r19: ref
	{
		if seq.IsComponentPresent(18) {
			ie.P_Max_r19 = new(P_Max)
			if err := ie.P_Max_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 21. sib1-RestrictedSetConfig-r19: enumerated
	{
		v19, err := d.DecodeEnumerated(sIB1RequestConfigR19Sib1RestrictedSetConfigR19Constraints)
		if err != nil {
			return err
		}
		ie.Sib1_RestrictedSetConfig_r19 = v19
	}

	return nil
}
