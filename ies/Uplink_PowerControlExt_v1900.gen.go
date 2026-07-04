// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Uplink-powerControlExt-v1900 (line 16344).

var uplinkPowerControlExtV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "p0AlphaSetforPUSCH-SBFD-r19", Optional: true},
		{Name: "p0AlphaSetforPUCCH-SBFD-r19", Optional: true},
		{Name: "p0AlphaSetforSRS-SBFD-r19", Optional: true},
	},
}

type Uplink_PowerControlExt_v1900 struct {
	P0AlphaSetforPUSCH_SBFD_r19 *P0AlphaSet_r17
	P0AlphaSetforPUCCH_SBFD_r19 *P0AlphaSet_r17
	P0AlphaSetforSRS_SBFD_r19   *P0AlphaSet_r17
}

func (ie *Uplink_PowerControlExt_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uplinkPowerControlExtV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.P0AlphaSetforPUSCH_SBFD_r19 != nil, ie.P0AlphaSetforPUCCH_SBFD_r19 != nil, ie.P0AlphaSetforSRS_SBFD_r19 != nil}); err != nil {
		return err
	}

	// 2. p0AlphaSetforPUSCH-SBFD-r19: ref
	{
		if ie.P0AlphaSetforPUSCH_SBFD_r19 != nil {
			if err := ie.P0AlphaSetforPUSCH_SBFD_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. p0AlphaSetforPUCCH-SBFD-r19: ref
	{
		if ie.P0AlphaSetforPUCCH_SBFD_r19 != nil {
			if err := ie.P0AlphaSetforPUCCH_SBFD_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. p0AlphaSetforSRS-SBFD-r19: ref
	{
		if ie.P0AlphaSetforSRS_SBFD_r19 != nil {
			if err := ie.P0AlphaSetforSRS_SBFD_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *Uplink_PowerControlExt_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uplinkPowerControlExtV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. p0AlphaSetforPUSCH-SBFD-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.P0AlphaSetforPUSCH_SBFD_r19 = new(P0AlphaSet_r17)
			if err := ie.P0AlphaSetforPUSCH_SBFD_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. p0AlphaSetforPUCCH-SBFD-r19: ref
	{
		if seq.IsComponentPresent(1) {
			ie.P0AlphaSetforPUCCH_SBFD_r19 = new(P0AlphaSet_r17)
			if err := ie.P0AlphaSetforPUCCH_SBFD_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. p0AlphaSetforSRS-SBFD-r19: ref
	{
		if seq.IsComponentPresent(2) {
			ie.P0AlphaSetforSRS_SBFD_r19 = new(P0AlphaSet_r17)
			if err := ie.P0AlphaSetforSRS_SBFD_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
