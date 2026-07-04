// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-MRDC-Capability-v1900 (line 25589).

var uEMRDCCapabilityV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measAndMobParametersMRDC-v1900", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type UE_MRDC_Capability_v1900 struct {
	MeasAndMobParametersMRDC_v1900 *MeasAndMobParametersMRDC_v1900
	NonCriticalExtension           *struct{}
}

func (ie *UE_MRDC_Capability_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEMRDCCapabilityV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasAndMobParametersMRDC_v1900 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-v1900: ref
	{
		if ie.MeasAndMobParametersMRDC_v1900 != nil {
			if err := ie.MeasAndMobParametersMRDC_v1900.Encode(e); err != nil {
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

func (ie *UE_MRDC_Capability_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEMRDCCapabilityV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-v1900: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasAndMobParametersMRDC_v1900 = new(MeasAndMobParametersMRDC_v1900)
			if err := ie.MeasAndMobParametersMRDC_v1900.Decode(d); err != nil {
				return err
			}
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
