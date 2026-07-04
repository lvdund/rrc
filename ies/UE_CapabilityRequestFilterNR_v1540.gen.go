// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-CapabilityRequestFilterNR-v1540 (line 25528).

var uECapabilityRequestFilterNRV1540Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-SwitchingTimeRequest", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UE_CapabilityRequestFilterNR_v1540_Srs_SwitchingTimeRequest_True = 0
)

var uECapabilityRequestFilterNRV1540SrsSwitchingTimeRequestConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_CapabilityRequestFilterNR_v1540_Srs_SwitchingTimeRequest_True},
}

type UE_CapabilityRequestFilterNR_v1540 struct {
	Srs_SwitchingTimeRequest *int64
	NonCriticalExtension     *UE_CapabilityRequestFilterNR_v1710
}

func (ie *UE_CapabilityRequestFilterNR_v1540) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uECapabilityRequestFilterNRV1540Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_SwitchingTimeRequest != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. srs-SwitchingTimeRequest: enumerated
	{
		if ie.Srs_SwitchingTimeRequest != nil {
			if err := e.EncodeEnumerated(*ie.Srs_SwitchingTimeRequest, uECapabilityRequestFilterNRV1540SrsSwitchingTimeRequestConstraints); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_CapabilityRequestFilterNR_v1540) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uECapabilityRequestFilterNRV1540Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. srs-SwitchingTimeRequest: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uECapabilityRequestFilterNRV1540SrsSwitchingTimeRequestConstraints)
			if err != nil {
				return err
			}
			ie.Srs_SwitchingTimeRequest = &idx
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(UE_CapabilityRequestFilterNR_v1710)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
