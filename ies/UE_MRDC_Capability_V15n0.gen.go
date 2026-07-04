// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-MRDC-Capability-v15n0 (line 25600).

var uEMRDCCapabilityV15n0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rf-ParametersMRDC-v15n0", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var uEMRDCCapabilityV15n0LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type UE_MRDC_Capability_V15n0 struct {
	Rf_ParametersMRDC_V15n0  *RF_ParametersMRDC_V15n0
	LateNonCriticalExtension []byte
	NonCriticalExtension     *UE_MRDC_Capability_V16e0
}

func (ie *UE_MRDC_Capability_V15n0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEMRDCCapabilityV15n0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rf_ParametersMRDC_V15n0 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. rf-ParametersMRDC-v15n0: ref
	{
		if ie.Rf_ParametersMRDC_V15n0 != nil {
			if err := ie.Rf_ParametersMRDC_V15n0.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, uEMRDCCapabilityV15n0LateNonCriticalExtensionConstraints); err != nil {
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

func (ie *UE_MRDC_Capability_V15n0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEMRDCCapabilityV15n0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. rf-ParametersMRDC-v15n0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Rf_ParametersMRDC_V15n0 = new(RF_ParametersMRDC_V15n0)
			if err := ie.Rf_ParametersMRDC_V15n0.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(uEMRDCCapabilityV15n0LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = new(UE_MRDC_Capability_V16e0)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
