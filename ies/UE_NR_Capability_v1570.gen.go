// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v1570 (line 25691).

var uENRCapabilityV1570Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nrdc-Parameters-v1570", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type UE_NR_Capability_v1570 struct {
	Nrdc_Parameters_v1570 *NRDC_Parameters_v1570
	NonCriticalExtension  *UE_NR_Capability_v1610
}

func (ie *UE_NR_Capability_v1570) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV1570Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Nrdc_Parameters_v1570 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. nrdc-Parameters-v1570: ref
	{
		if ie.Nrdc_Parameters_v1570 != nil {
			if err := ie.Nrdc_Parameters_v1570.Encode(e); err != nil {
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

func (ie *UE_NR_Capability_v1570) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV1570Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. nrdc-Parameters-v1570: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Nrdc_Parameters_v1570 = new(NRDC_Parameters_v1570)
			if err := ie.Nrdc_Parameters_v1570.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(UE_NR_Capability_v1610)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
