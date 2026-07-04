// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-AllPosResources-r16 (line 20454).

var sRSAllPosResourcesR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-PosResources-r16"},
		{Name: "srs-PosResourceAP-r16", Optional: true},
		{Name: "srs-PosResourceSP-r16", Optional: true},
	},
}

type SRS_AllPosResources_r16 struct {
	Srs_PosResources_r16  SRS_PosResources_r16
	Srs_PosResourceAP_r16 *SRS_PosResourceAP_r16
	Srs_PosResourceSP_r16 *SRS_PosResourceSP_r16
}

func (ie *SRS_AllPosResources_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSAllPosResourcesR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_PosResourceAP_r16 != nil, ie.Srs_PosResourceSP_r16 != nil}); err != nil {
		return err
	}

	// 2. srs-PosResources-r16: ref
	{
		if err := ie.Srs_PosResources_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. srs-PosResourceAP-r16: ref
	{
		if ie.Srs_PosResourceAP_r16 != nil {
			if err := ie.Srs_PosResourceAP_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. srs-PosResourceSP-r16: ref
	{
		if ie.Srs_PosResourceSP_r16 != nil {
			if err := ie.Srs_PosResourceSP_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SRS_AllPosResources_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSAllPosResourcesR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. srs-PosResources-r16: ref
	{
		if err := ie.Srs_PosResources_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. srs-PosResourceAP-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Srs_PosResourceAP_r16 = new(SRS_PosResourceAP_r16)
			if err := ie.Srs_PosResourceAP_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. srs-PosResourceSP-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Srs_PosResourceSP_r16 = new(SRS_PosResourceSP_r16)
			if err := ie.Srs_PosResourceSP_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
