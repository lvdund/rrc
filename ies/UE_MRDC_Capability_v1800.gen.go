// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-MRDC-Capability-v1800 (line 25582).

var uEMRDCCapabilityV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "requirementTypeIndication-r18", Optional: true},
		{Name: "measAndMobParametersMRDC-v1810", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UE_MRDC_Capability_v1800_RequirementTypeIndication_r18_Supported = 0
)

var uEMRDCCapabilityV1800RequirementTypeIndicationR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_MRDC_Capability_v1800_RequirementTypeIndication_r18_Supported},
}

type UE_MRDC_Capability_v1800 struct {
	RequirementTypeIndication_r18  *int64
	MeasAndMobParametersMRDC_v1810 *MeasAndMobParametersMRDC_v1810
	NonCriticalExtension           *UE_MRDC_Capability_v1900
}

func (ie *UE_MRDC_Capability_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEMRDCCapabilityV1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RequirementTypeIndication_r18 != nil, ie.MeasAndMobParametersMRDC_v1810 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. requirementTypeIndication-r18: enumerated
	{
		if ie.RequirementTypeIndication_r18 != nil {
			if err := e.EncodeEnumerated(*ie.RequirementTypeIndication_r18, uEMRDCCapabilityV1800RequirementTypeIndicationR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. measAndMobParametersMRDC-v1810: ref
	{
		if ie.MeasAndMobParametersMRDC_v1810 != nil {
			if err := ie.MeasAndMobParametersMRDC_v1810.Encode(e); err != nil {
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

func (ie *UE_MRDC_Capability_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEMRDCCapabilityV1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. requirementTypeIndication-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uEMRDCCapabilityV1800RequirementTypeIndicationR18Constraints)
			if err != nil {
				return err
			}
			ie.RequirementTypeIndication_r18 = &idx
		}
	}

	// 3. measAndMobParametersMRDC-v1810: ref
	{
		if seq.IsComponentPresent(1) {
			ie.MeasAndMobParametersMRDC_v1810 = new(MeasAndMobParametersMRDC_v1810)
			if err := ie.MeasAndMobParametersMRDC_v1810.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = new(UE_MRDC_Capability_v1900)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
