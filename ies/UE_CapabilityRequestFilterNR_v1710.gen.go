// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-CapabilityRequestFilterNR-v1710 (line 25533).

var uECapabilityRequestFilterNRV1710Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sidelinkRequest-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UE_CapabilityRequestFilterNR_v1710_SidelinkRequest_r17_True = 0
)

var uECapabilityRequestFilterNRV1710SidelinkRequestR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_CapabilityRequestFilterNR_v1710_SidelinkRequest_r17_True},
}

type UE_CapabilityRequestFilterNR_v1710 struct {
	SidelinkRequest_r17  *int64
	NonCriticalExtension *struct{}
}

func (ie *UE_CapabilityRequestFilterNR_v1710) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uECapabilityRequestFilterNRV1710Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SidelinkRequest_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. sidelinkRequest-r17: enumerated
	{
		if ie.SidelinkRequest_r17 != nil {
			if err := e.EncodeEnumerated(*ie.SidelinkRequest_r17, uECapabilityRequestFilterNRV1710SidelinkRequestR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *UE_CapabilityRequestFilterNR_v1710) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uECapabilityRequestFilterNRV1710Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sidelinkRequest-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uECapabilityRequestFilterNRV1710SidelinkRequestR17Constraints)
			if err != nil {
				return err
			}
			ie.SidelinkRequest_r17 = &idx
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
