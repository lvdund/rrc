// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Uplink-powerControl-r17 (line 16329).

var uplinkPowerControlR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ul-powercontrolId-r17"},
		{Name: "p0AlphaSetforPUSCH-r17", Optional: true},
		{Name: "p0AlphaSetforPUCCH-r17", Optional: true},
		{Name: "p0AlphaSetforSRS-r17", Optional: true},
	},
}

type Uplink_PowerControl_r17 struct {
	Ul_PowercontrolId_r17  Uplink_PowerControlId_r17
	P0AlphaSetforPUSCH_r17 *P0AlphaSet_r17
	P0AlphaSetforPUCCH_r17 *P0AlphaSet_r17
	P0AlphaSetforSRS_r17   *P0AlphaSet_r17
}

func (ie *Uplink_PowerControl_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uplinkPowerControlR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.P0AlphaSetforPUSCH_r17 != nil, ie.P0AlphaSetforPUCCH_r17 != nil, ie.P0AlphaSetforSRS_r17 != nil}); err != nil {
		return err
	}

	// 2. ul-powercontrolId-r17: ref
	{
		if err := ie.Ul_PowercontrolId_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. p0AlphaSetforPUSCH-r17: ref
	{
		if ie.P0AlphaSetforPUSCH_r17 != nil {
			if err := ie.P0AlphaSetforPUSCH_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. p0AlphaSetforPUCCH-r17: ref
	{
		if ie.P0AlphaSetforPUCCH_r17 != nil {
			if err := ie.P0AlphaSetforPUCCH_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. p0AlphaSetforSRS-r17: ref
	{
		if ie.P0AlphaSetforSRS_r17 != nil {
			if err := ie.P0AlphaSetforSRS_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *Uplink_PowerControl_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uplinkPowerControlR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ul-powercontrolId-r17: ref
	{
		if err := ie.Ul_PowercontrolId_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. p0AlphaSetforPUSCH-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.P0AlphaSetforPUSCH_r17 = new(P0AlphaSet_r17)
			if err := ie.P0AlphaSetforPUSCH_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. p0AlphaSetforPUCCH-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.P0AlphaSetforPUCCH_r17 = new(P0AlphaSet_r17)
			if err := ie.P0AlphaSetforPUCCH_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. p0AlphaSetforSRS-r17: ref
	{
		if seq.IsComponentPresent(3) {
			ie.P0AlphaSetforSRS_r17 = new(P0AlphaSet_r17)
			if err := ie.P0AlphaSetforSRS_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
