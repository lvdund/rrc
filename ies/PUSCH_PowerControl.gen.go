// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUSCH-PowerControl (line 12567).

var pUSCHPowerControlConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "tpc-Accumulation", Optional: true},
		{Name: "msg3-Alpha", Optional: true},
		{Name: "p0-NominalWithoutGrant", Optional: true},
		{Name: "p0-AlphaSets", Optional: true},
		{Name: "pathlossReferenceRSToAddModList", Optional: true},
		{Name: "pathlossReferenceRSToReleaseList", Optional: true},
		{Name: "twoPUSCH-PC-AdjustmentStates", Optional: true},
		{Name: "deltaMCS", Optional: true},
		{Name: "sri-PUSCH-MappingToAddModList", Optional: true},
		{Name: "sri-PUSCH-MappingToReleaseList", Optional: true},
	},
}

const (
	PUSCH_PowerControl_Tpc_Accumulation_Disabled = 0
)

var pUSCHPowerControlTpcAccumulationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_PowerControl_Tpc_Accumulation_Disabled},
}

var pUSCHPowerControlP0NominalWithoutGrantConstraints = per.Constrained(-202, 24)

var pUSCHPowerControlP0AlphaSetsConstraints = per.SizeRange(1, common.MaxNrofP0_PUSCH_AlphaSets)

var pUSCHPowerControlPathlossReferenceRSToAddModListConstraints = per.SizeRange(1, common.MaxNrofPUSCH_PathlossReferenceRSs)

var pUSCHPowerControlPathlossReferenceRSToReleaseListConstraints = per.SizeRange(1, common.MaxNrofPUSCH_PathlossReferenceRSs)

const (
	PUSCH_PowerControl_TwoPUSCH_PC_AdjustmentStates_TwoStates = 0
)

var pUSCHPowerControlTwoPUSCHPCAdjustmentStatesConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_PowerControl_TwoPUSCH_PC_AdjustmentStates_TwoStates},
}

const (
	PUSCH_PowerControl_DeltaMCS_Enabled = 0
)

var pUSCHPowerControlDeltaMCSConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_PowerControl_DeltaMCS_Enabled},
}

var pUSCHPowerControlSriPUSCHMappingToAddModListConstraints = per.SizeRange(1, common.MaxNrofSRI_PUSCH_Mappings)

var pUSCHPowerControlSriPUSCHMappingToReleaseListConstraints = per.SizeRange(1, common.MaxNrofSRI_PUSCH_Mappings)

type PUSCH_PowerControl struct {
	Tpc_Accumulation                 *int64
	Msg3_Alpha                       *Alpha
	P0_NominalWithoutGrant           *int64
	P0_AlphaSets                     []P0_PUSCH_AlphaSet
	PathlossReferenceRSToAddModList  []PUSCH_PathlossReferenceRS
	PathlossReferenceRSToReleaseList []PUSCH_PathlossReferenceRS_Id
	TwoPUSCH_PC_AdjustmentStates     *int64
	DeltaMCS                         *int64
	Sri_PUSCH_MappingToAddModList    []SRI_PUSCH_PowerControl
	Sri_PUSCH_MappingToReleaseList   []SRI_PUSCH_PowerControlId
}

func (ie *PUSCH_PowerControl) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUSCHPowerControlConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Tpc_Accumulation != nil, ie.Msg3_Alpha != nil, ie.P0_NominalWithoutGrant != nil, ie.P0_AlphaSets != nil, ie.PathlossReferenceRSToAddModList != nil, ie.PathlossReferenceRSToReleaseList != nil, ie.TwoPUSCH_PC_AdjustmentStates != nil, ie.DeltaMCS != nil, ie.Sri_PUSCH_MappingToAddModList != nil, ie.Sri_PUSCH_MappingToReleaseList != nil}); err != nil {
		return err
	}

	// 2. tpc-Accumulation: enumerated
	{
		if ie.Tpc_Accumulation != nil {
			if err := e.EncodeEnumerated(*ie.Tpc_Accumulation, pUSCHPowerControlTpcAccumulationConstraints); err != nil {
				return err
			}
		}
	}

	// 3. msg3-Alpha: ref
	{
		if ie.Msg3_Alpha != nil {
			if err := ie.Msg3_Alpha.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. p0-NominalWithoutGrant: integer
	{
		if ie.P0_NominalWithoutGrant != nil {
			if err := e.EncodeInteger(*ie.P0_NominalWithoutGrant, pUSCHPowerControlP0NominalWithoutGrantConstraints); err != nil {
				return err
			}
		}
	}

	// 5. p0-AlphaSets: sequence-of
	{
		if ie.P0_AlphaSets != nil {
			seqOf := e.NewSequenceOfEncoder(pUSCHPowerControlP0AlphaSetsConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.P0_AlphaSets))); err != nil {
				return err
			}
			for i := range ie.P0_AlphaSets {
				if err := ie.P0_AlphaSets[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. pathlossReferenceRSToAddModList: sequence-of
	{
		if ie.PathlossReferenceRSToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(pUSCHPowerControlPathlossReferenceRSToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.PathlossReferenceRSToAddModList))); err != nil {
				return err
			}
			for i := range ie.PathlossReferenceRSToAddModList {
				if err := ie.PathlossReferenceRSToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. pathlossReferenceRSToReleaseList: sequence-of
	{
		if ie.PathlossReferenceRSToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(pUSCHPowerControlPathlossReferenceRSToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.PathlossReferenceRSToReleaseList))); err != nil {
				return err
			}
			for i := range ie.PathlossReferenceRSToReleaseList {
				if err := ie.PathlossReferenceRSToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 8. twoPUSCH-PC-AdjustmentStates: enumerated
	{
		if ie.TwoPUSCH_PC_AdjustmentStates != nil {
			if err := e.EncodeEnumerated(*ie.TwoPUSCH_PC_AdjustmentStates, pUSCHPowerControlTwoPUSCHPCAdjustmentStatesConstraints); err != nil {
				return err
			}
		}
	}

	// 9. deltaMCS: enumerated
	{
		if ie.DeltaMCS != nil {
			if err := e.EncodeEnumerated(*ie.DeltaMCS, pUSCHPowerControlDeltaMCSConstraints); err != nil {
				return err
			}
		}
	}

	// 10. sri-PUSCH-MappingToAddModList: sequence-of
	{
		if ie.Sri_PUSCH_MappingToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(pUSCHPowerControlSriPUSCHMappingToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sri_PUSCH_MappingToAddModList))); err != nil {
				return err
			}
			for i := range ie.Sri_PUSCH_MappingToAddModList {
				if err := ie.Sri_PUSCH_MappingToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 11. sri-PUSCH-MappingToReleaseList: sequence-of
	{
		if ie.Sri_PUSCH_MappingToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(pUSCHPowerControlSriPUSCHMappingToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sri_PUSCH_MappingToReleaseList))); err != nil {
				return err
			}
			for i := range ie.Sri_PUSCH_MappingToReleaseList {
				if err := ie.Sri_PUSCH_MappingToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *PUSCH_PowerControl) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUSCHPowerControlConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. tpc-Accumulation: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(pUSCHPowerControlTpcAccumulationConstraints)
			if err != nil {
				return err
			}
			ie.Tpc_Accumulation = &idx
		}
	}

	// 3. msg3-Alpha: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Msg3_Alpha = new(Alpha)
			if err := ie.Msg3_Alpha.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. p0-NominalWithoutGrant: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(pUSCHPowerControlP0NominalWithoutGrantConstraints)
			if err != nil {
				return err
			}
			ie.P0_NominalWithoutGrant = &v
		}
	}

	// 5. p0-AlphaSets: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(pUSCHPowerControlP0AlphaSetsConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.P0_AlphaSets = make([]P0_PUSCH_AlphaSet, n)
			for i := int64(0); i < n; i++ {
				if err := ie.P0_AlphaSets[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. pathlossReferenceRSToAddModList: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(pUSCHPowerControlPathlossReferenceRSToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.PathlossReferenceRSToAddModList = make([]PUSCH_PathlossReferenceRS, n)
			for i := int64(0); i < n; i++ {
				if err := ie.PathlossReferenceRSToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. pathlossReferenceRSToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(pUSCHPowerControlPathlossReferenceRSToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.PathlossReferenceRSToReleaseList = make([]PUSCH_PathlossReferenceRS_Id, n)
			for i := int64(0); i < n; i++ {
				if err := ie.PathlossReferenceRSToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. twoPUSCH-PC-AdjustmentStates: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(pUSCHPowerControlTwoPUSCHPCAdjustmentStatesConstraints)
			if err != nil {
				return err
			}
			ie.TwoPUSCH_PC_AdjustmentStates = &idx
		}
	}

	// 9. deltaMCS: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(pUSCHPowerControlDeltaMCSConstraints)
			if err != nil {
				return err
			}
			ie.DeltaMCS = &idx
		}
	}

	// 10. sri-PUSCH-MappingToAddModList: sequence-of
	{
		if seq.IsComponentPresent(8) {
			seqOf := d.NewSequenceOfDecoder(pUSCHPowerControlSriPUSCHMappingToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sri_PUSCH_MappingToAddModList = make([]SRI_PUSCH_PowerControl, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sri_PUSCH_MappingToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 11. sri-PUSCH-MappingToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(9) {
			seqOf := d.NewSequenceOfDecoder(pUSCHPowerControlSriPUSCHMappingToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sri_PUSCH_MappingToReleaseList = make([]SRI_PUSCH_PowerControlId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sri_PUSCH_MappingToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
