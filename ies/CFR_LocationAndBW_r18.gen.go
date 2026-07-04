// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CFR-LocationAndBW-r18 (line 28480).

var cFRLocationAndBWR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "locationAndBandwidthMBS-r18", Optional: true},
		{Name: "absoluteFrequencyPointA-MBS-r18", Optional: true},
		{Name: "offsetToCarrierMBS-r18", Optional: true},
	},
}

var cFRLocationAndBWR18LocationAndBandwidthMBSR18Constraints = per.Constrained(0, 37949)

var cFRLocationAndBWR18OffsetToCarrierMBSR18Constraints = per.Constrained(0, 2199)

type CFR_LocationAndBW_r18 struct {
	LocationAndBandwidthMBS_r18     *int64
	AbsoluteFrequencyPointA_MBS_r18 *ARFCN_ValueNR
	OffsetToCarrierMBS_r18          *int64
}

func (ie *CFR_LocationAndBW_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cFRLocationAndBWR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LocationAndBandwidthMBS_r18 != nil, ie.AbsoluteFrequencyPointA_MBS_r18 != nil, ie.OffsetToCarrierMBS_r18 != nil}); err != nil {
		return err
	}

	// 2. locationAndBandwidthMBS-r18: integer
	{
		if ie.LocationAndBandwidthMBS_r18 != nil {
			if err := e.EncodeInteger(*ie.LocationAndBandwidthMBS_r18, cFRLocationAndBWR18LocationAndBandwidthMBSR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. absoluteFrequencyPointA-MBS-r18: ref
	{
		if ie.AbsoluteFrequencyPointA_MBS_r18 != nil {
			if err := ie.AbsoluteFrequencyPointA_MBS_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. offsetToCarrierMBS-r18: integer
	{
		if ie.OffsetToCarrierMBS_r18 != nil {
			if err := e.EncodeInteger(*ie.OffsetToCarrierMBS_r18, cFRLocationAndBWR18OffsetToCarrierMBSR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CFR_LocationAndBW_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cFRLocationAndBWR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. locationAndBandwidthMBS-r18: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(cFRLocationAndBWR18LocationAndBandwidthMBSR18Constraints)
			if err != nil {
				return err
			}
			ie.LocationAndBandwidthMBS_r18 = &v
		}
	}

	// 3. absoluteFrequencyPointA-MBS-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.AbsoluteFrequencyPointA_MBS_r18 = new(ARFCN_ValueNR)
			if err := ie.AbsoluteFrequencyPointA_MBS_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. offsetToCarrierMBS-r18: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(cFRLocationAndBWR18OffsetToCarrierMBSR18Constraints)
			if err != nil {
				return err
			}
			ie.OffsetToCarrierMBS_r18 = &v
		}
	}

	return nil
}
