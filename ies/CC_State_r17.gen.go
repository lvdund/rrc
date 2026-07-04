// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CC-State-r17 (line 5792).

var cCStateR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dlCarrier-r17", Optional: true},
		{Name: "ulCarrier-r17", Optional: true},
	},
}

type CC_State_r17 struct {
	DlCarrier_r17 *CarrierState_r17
	UlCarrier_r17 *CarrierState_r17
}

func (ie *CC_State_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cCStateR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DlCarrier_r17 != nil, ie.UlCarrier_r17 != nil}); err != nil {
		return err
	}

	// 2. dlCarrier-r17: ref
	{
		if ie.DlCarrier_r17 != nil {
			if err := ie.DlCarrier_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ulCarrier-r17: ref
	{
		if ie.UlCarrier_r17 != nil {
			if err := ie.UlCarrier_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CC_State_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cCStateR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dlCarrier-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.DlCarrier_r17 = new(CarrierState_r17)
			if err := ie.DlCarrier_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ulCarrier-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.UlCarrier_r17 = new(CarrierState_r17)
			if err := ie.UlCarrier_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
