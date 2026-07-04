// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-MRDC-Capability-v1700 (line 25572).

var uEMRDCCapabilityV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measAndMobParametersMRDC-v1700"},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type UE_MRDC_Capability_v1700 struct {
	MeasAndMobParametersMRDC_v1700 MeasAndMobParametersMRDC_v1700
	NonCriticalExtension           *UE_MRDC_Capability_v1730
}

func (ie *UE_MRDC_Capability_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEMRDCCapabilityV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-v1700: ref
	{
		if err := ie.MeasAndMobParametersMRDC_v1700.Encode(e); err != nil {
			return err
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

func (ie *UE_MRDC_Capability_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEMRDCCapabilityV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-v1700: ref
	{
		if err := ie.MeasAndMobParametersMRDC_v1700.Decode(d); err != nil {
			return err
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(UE_MRDC_Capability_v1730)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
