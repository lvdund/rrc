// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BWP-Downlink (line 5239).

var bWPDownlinkConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "bwp-Id"},
		{Name: "bwp-Common", Optional: true},
		{Name: "bwp-Dedicated", Optional: true},
	},
}

type BWP_Downlink struct {
	Bwp_Id        BWP_Id
	Bwp_Common    *BWP_DownlinkCommon
	Bwp_Dedicated *BWP_DownlinkDedicated
}

func (ie *BWP_Downlink) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bWPDownlinkConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Bwp_Common != nil, ie.Bwp_Dedicated != nil}); err != nil {
		return err
	}

	// 3. bwp-Id: ref
	{
		if err := ie.Bwp_Id.Encode(e); err != nil {
			return err
		}
	}

	// 4. bwp-Common: ref
	{
		if ie.Bwp_Common != nil {
			if err := ie.Bwp_Common.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. bwp-Dedicated: ref
	{
		if ie.Bwp_Dedicated != nil {
			if err := ie.Bwp_Dedicated.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BWP_Downlink) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bWPDownlinkConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. bwp-Id: ref
	{
		if err := ie.Bwp_Id.Decode(d); err != nil {
			return err
		}
	}

	// 4. bwp-Common: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Bwp_Common = new(BWP_DownlinkCommon)
			if err := ie.Bwp_Common.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. bwp-Dedicated: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Bwp_Dedicated = new(BWP_DownlinkDedicated)
			if err := ie.Bwp_Dedicated.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
