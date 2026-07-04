// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: InterFreqCarrierFreqInfo-v1700 (line 4014).

var interFreqCarrierFreqInfoV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "interFreqNeighHSDN-CellList-r17", Optional: true},
		{Name: "highSpeedMeasInterFreq-r17", Optional: true},
		{Name: "redCapAccessAllowed-r17", Optional: true},
		{Name: "ssb-PositionQCL-Common-r17", Optional: true},
		{Name: "interFreqNeighCellList-v1710", Optional: true},
	},
}

const (
	InterFreqCarrierFreqInfo_v1700_HighSpeedMeasInterFreq_r17_True = 0
)

var interFreqCarrierFreqInfoV1700HighSpeedMeasInterFreqR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{InterFreqCarrierFreqInfo_v1700_HighSpeedMeasInterFreq_r17_True},
}

const (
	InterFreqCarrierFreqInfo_v1700_RedCapAccessAllowed_r17_True = 0
)

var interFreqCarrierFreqInfoV1700RedCapAccessAllowedR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{InterFreqCarrierFreqInfo_v1700_RedCapAccessAllowed_r17_True},
}

type InterFreqCarrierFreqInfo_v1700 struct {
	InterFreqNeighHSDN_CellList_r17 *InterFreqNeighHSDN_CellList_r17
	HighSpeedMeasInterFreq_r17      *int64
	RedCapAccessAllowed_r17         *int64
	Ssb_PositionQCL_Common_r17      *SSB_PositionQCL_Relation_r17
	InterFreqNeighCellList_v1710    *InterFreqNeighCellList_v1710
}

func (ie *InterFreqCarrierFreqInfo_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(interFreqCarrierFreqInfoV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.InterFreqNeighHSDN_CellList_r17 != nil, ie.HighSpeedMeasInterFreq_r17 != nil, ie.RedCapAccessAllowed_r17 != nil, ie.Ssb_PositionQCL_Common_r17 != nil, ie.InterFreqNeighCellList_v1710 != nil}); err != nil {
		return err
	}

	// 2. interFreqNeighHSDN-CellList-r17: ref
	{
		if ie.InterFreqNeighHSDN_CellList_r17 != nil {
			if err := ie.InterFreqNeighHSDN_CellList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. highSpeedMeasInterFreq-r17: enumerated
	{
		if ie.HighSpeedMeasInterFreq_r17 != nil {
			if err := e.EncodeEnumerated(*ie.HighSpeedMeasInterFreq_r17, interFreqCarrierFreqInfoV1700HighSpeedMeasInterFreqR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. redCapAccessAllowed-r17: enumerated
	{
		if ie.RedCapAccessAllowed_r17 != nil {
			if err := e.EncodeEnumerated(*ie.RedCapAccessAllowed_r17, interFreqCarrierFreqInfoV1700RedCapAccessAllowedR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. ssb-PositionQCL-Common-r17: ref
	{
		if ie.Ssb_PositionQCL_Common_r17 != nil {
			if err := ie.Ssb_PositionQCL_Common_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. interFreqNeighCellList-v1710: ref
	{
		if ie.InterFreqNeighCellList_v1710 != nil {
			if err := ie.InterFreqNeighCellList_v1710.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *InterFreqCarrierFreqInfo_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(interFreqCarrierFreqInfoV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. interFreqNeighHSDN-CellList-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.InterFreqNeighHSDN_CellList_r17 = new(InterFreqNeighHSDN_CellList_r17)
			if err := ie.InterFreqNeighHSDN_CellList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. highSpeedMeasInterFreq-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(interFreqCarrierFreqInfoV1700HighSpeedMeasInterFreqR17Constraints)
			if err != nil {
				return err
			}
			ie.HighSpeedMeasInterFreq_r17 = &idx
		}
	}

	// 4. redCapAccessAllowed-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(interFreqCarrierFreqInfoV1700RedCapAccessAllowedR17Constraints)
			if err != nil {
				return err
			}
			ie.RedCapAccessAllowed_r17 = &idx
		}
	}

	// 5. ssb-PositionQCL-Common-r17: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Ssb_PositionQCL_Common_r17 = new(SSB_PositionQCL_Relation_r17)
			if err := ie.Ssb_PositionQCL_Common_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. interFreqNeighCellList-v1710: ref
	{
		if seq.IsComponentPresent(4) {
			ie.InterFreqNeighCellList_v1710 = new(InterFreqNeighCellList_v1710)
			if err := ie.InterFreqNeighCellList_v1710.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
