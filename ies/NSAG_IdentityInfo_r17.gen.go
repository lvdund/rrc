// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NSAG-IdentityInfo-r17 (line 10772).

var nSAGIdentityInfoR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nsag-ID-r17"},
		{Name: "trackingAreaCode-r17", Optional: true},
	},
}

type NSAG_IdentityInfo_r17 struct {
	Nsag_ID_r17          NSAG_ID_r17
	TrackingAreaCode_r17 *TrackingAreaCode
}

func (ie *NSAG_IdentityInfo_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nSAGIdentityInfoR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.TrackingAreaCode_r17 != nil}); err != nil {
		return err
	}

	// 2. nsag-ID-r17: ref
	{
		if err := ie.Nsag_ID_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. trackingAreaCode-r17: ref
	{
		if ie.TrackingAreaCode_r17 != nil {
			if err := ie.TrackingAreaCode_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NSAG_IdentityInfo_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nSAGIdentityInfoR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. nsag-ID-r17: ref
	{
		if err := ie.Nsag_ID_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. trackingAreaCode-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.TrackingAreaCode_r17 = new(TrackingAreaCode)
			if err := ie.TrackingAreaCode_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
