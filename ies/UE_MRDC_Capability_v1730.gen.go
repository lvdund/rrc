// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-MRDC-Capability-v1730 (line 25577).

var uEMRDCCapabilityV1730Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measAndMobParametersMRDC-v1730", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type UE_MRDC_Capability_v1730 struct {
	MeasAndMobParametersMRDC_v1730 *MeasAndMobParametersMRDC_v1730
	NonCriticalExtension           *UE_MRDC_Capability_v1800
}

func (ie *UE_MRDC_Capability_v1730) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEMRDCCapabilityV1730Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasAndMobParametersMRDC_v1730 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-v1730: ref
	{
		if ie.MeasAndMobParametersMRDC_v1730 != nil {
			if err := ie.MeasAndMobParametersMRDC_v1730.Encode(e); err != nil {
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

func (ie *UE_MRDC_Capability_v1730) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEMRDCCapabilityV1730Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-v1730: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasAndMobParametersMRDC_v1730 = new(MeasAndMobParametersMRDC_v1730)
			if err := ie.MeasAndMobParametersMRDC_v1730.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(UE_MRDC_Capability_v1800)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
