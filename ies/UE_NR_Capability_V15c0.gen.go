// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v15c0 (line 25697).

var uENRCapabilityV15c0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nrdc-Parameters-v15c0", Optional: true},
		{Name: "partialFR2-FallbackRX-Req", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UE_NR_Capability_V15c0_PartialFR2_FallbackRX_Req_True = 0
)

var uENRCapabilityV15c0PartialFR2FallbackRXReqConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_V15c0_PartialFR2_FallbackRX_Req_True},
}

type UE_NR_Capability_V15c0 struct {
	Nrdc_Parameters_V15c0     *NRDC_Parameters_V15c0
	PartialFR2_FallbackRX_Req *int64
	NonCriticalExtension      *UE_NR_Capability_V15g0
}

func (ie *UE_NR_Capability_V15c0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV15c0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Nrdc_Parameters_V15c0 != nil, ie.PartialFR2_FallbackRX_Req != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. nrdc-Parameters-v15c0: ref
	{
		if ie.Nrdc_Parameters_V15c0 != nil {
			if err := ie.Nrdc_Parameters_V15c0.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. partialFR2-FallbackRX-Req: enumerated
	{
		if ie.PartialFR2_FallbackRX_Req != nil {
			if err := e.EncodeEnumerated(*ie.PartialFR2_FallbackRX_Req, uENRCapabilityV15c0PartialFR2FallbackRXReqConstraints); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_NR_Capability_V15c0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV15c0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. nrdc-Parameters-v15c0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Nrdc_Parameters_V15c0 = new(NRDC_Parameters_V15c0)
			if err := ie.Nrdc_Parameters_V15c0.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. partialFR2-FallbackRX-Req: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV15c0PartialFR2FallbackRXReqConstraints)
			if err != nil {
				return err
			}
			ie.PartialFR2_FallbackRX_Req = &idx
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = new(UE_NR_Capability_V15g0)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
