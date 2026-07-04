// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: InterFreqCarrierFreqInfo-v1610 (line 4006).

var interFreqCarrierFreqInfoV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "interFreqNeighCellList-v1610", Optional: true},
		{Name: "smtc2-LP-r16", Optional: true},
		{Name: "interFreqAllowedCellList-r16", Optional: true},
		{Name: "ssb-PositionQCL-Common-r16", Optional: true},
		{Name: "interFreqCAG-CellList-r16", Optional: true},
	},
}

var interFreqCarrierFreqInfoV1610InterFreqCAGCellListR16Constraints = per.SizeRange(1, common.MaxPLMN)

type InterFreqCarrierFreqInfo_v1610 struct {
	InterFreqNeighCellList_v1610 *InterFreqNeighCellList_v1610
	Smtc2_LP_r16                 *SSB_MTC2_LP_r16
	InterFreqAllowedCellList_r16 *InterFreqAllowedCellList_r16
	Ssb_PositionQCL_Common_r16   *SSB_PositionQCL_Relation_r16
	InterFreqCAG_CellList_r16    []InterFreqCAG_CellListPerPLMN_r16
}

func (ie *InterFreqCarrierFreqInfo_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(interFreqCarrierFreqInfoV1610Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.InterFreqNeighCellList_v1610 != nil, ie.Smtc2_LP_r16 != nil, ie.InterFreqAllowedCellList_r16 != nil, ie.Ssb_PositionQCL_Common_r16 != nil, ie.InterFreqCAG_CellList_r16 != nil}); err != nil {
		return err
	}

	// 2. interFreqNeighCellList-v1610: ref
	{
		if ie.InterFreqNeighCellList_v1610 != nil {
			if err := ie.InterFreqNeighCellList_v1610.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. smtc2-LP-r16: ref
	{
		if ie.Smtc2_LP_r16 != nil {
			if err := ie.Smtc2_LP_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. interFreqAllowedCellList-r16: ref
	{
		if ie.InterFreqAllowedCellList_r16 != nil {
			if err := ie.InterFreqAllowedCellList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. ssb-PositionQCL-Common-r16: ref
	{
		if ie.Ssb_PositionQCL_Common_r16 != nil {
			if err := ie.Ssb_PositionQCL_Common_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. interFreqCAG-CellList-r16: sequence-of
	{
		if ie.InterFreqCAG_CellList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(interFreqCarrierFreqInfoV1610InterFreqCAGCellListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.InterFreqCAG_CellList_r16))); err != nil {
				return err
			}
			for i := range ie.InterFreqCAG_CellList_r16 {
				if err := ie.InterFreqCAG_CellList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *InterFreqCarrierFreqInfo_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(interFreqCarrierFreqInfoV1610Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. interFreqNeighCellList-v1610: ref
	{
		if seq.IsComponentPresent(0) {
			ie.InterFreqNeighCellList_v1610 = new(InterFreqNeighCellList_v1610)
			if err := ie.InterFreqNeighCellList_v1610.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. smtc2-LP-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Smtc2_LP_r16 = new(SSB_MTC2_LP_r16)
			if err := ie.Smtc2_LP_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. interFreqAllowedCellList-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.InterFreqAllowedCellList_r16 = new(InterFreqAllowedCellList_r16)
			if err := ie.InterFreqAllowedCellList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. ssb-PositionQCL-Common-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Ssb_PositionQCL_Common_r16 = new(SSB_PositionQCL_Relation_r16)
			if err := ie.Ssb_PositionQCL_Common_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. interFreqCAG-CellList-r16: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(interFreqCarrierFreqInfoV1610InterFreqCAGCellListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.InterFreqCAG_CellList_r16 = make([]InterFreqCAG_CellListPerPLMN_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.InterFreqCAG_CellList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
