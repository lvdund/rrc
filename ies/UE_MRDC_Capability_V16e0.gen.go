// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-MRDC-Capability-v16e0 (line 25607).

var uEMRDCCapabilityV16e0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rf-ParametersMRDC-v16e0", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type UE_MRDC_Capability_V16e0 struct {
	Rf_ParametersMRDC_V16e0 *RF_ParametersMRDC_V16e0
	NonCriticalExtension    *struct{}
}

func (ie *UE_MRDC_Capability_V16e0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEMRDCCapabilityV16e0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rf_ParametersMRDC_V16e0 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. rf-ParametersMRDC-v16e0: ref
	{
		if ie.Rf_ParametersMRDC_V16e0 != nil {
			if err := ie.Rf_ParametersMRDC_V16e0.Encode(e); err != nil {
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

func (ie *UE_MRDC_Capability_V16e0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEMRDCCapabilityV16e0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. rf-ParametersMRDC-v16e0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Rf_ParametersMRDC_V16e0 = new(RF_ParametersMRDC_V16e0)
			if err := ie.Rf_ParametersMRDC_V16e0.Decode(d); err != nil {
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
