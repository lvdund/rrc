// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-MRDC-Capability-v1610 (line 25565).

var uEMRDCCapabilityV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measAndMobParametersMRDC-v1610", Optional: true},
		{Name: "generalParametersMRDC-v1610", Optional: true},
		{Name: "pdcp-ParametersMRDC-v1610", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type UE_MRDC_Capability_v1610 struct {
	MeasAndMobParametersMRDC_v1610 *MeasAndMobParametersMRDC_v1610
	GeneralParametersMRDC_v1610    *GeneralParametersMRDC_v1610
	Pdcp_ParametersMRDC_v1610      *PDCP_ParametersMRDC_v1610
	NonCriticalExtension           *UE_MRDC_Capability_v1700
}

func (ie *UE_MRDC_Capability_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEMRDCCapabilityV1610Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasAndMobParametersMRDC_v1610 != nil, ie.GeneralParametersMRDC_v1610 != nil, ie.Pdcp_ParametersMRDC_v1610 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-v1610: ref
	{
		if ie.MeasAndMobParametersMRDC_v1610 != nil {
			if err := ie.MeasAndMobParametersMRDC_v1610.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. generalParametersMRDC-v1610: ref
	{
		if ie.GeneralParametersMRDC_v1610 != nil {
			if err := ie.GeneralParametersMRDC_v1610.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. pdcp-ParametersMRDC-v1610: ref
	{
		if ie.Pdcp_ParametersMRDC_v1610 != nil {
			if err := ie.Pdcp_ParametersMRDC_v1610.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_MRDC_Capability_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEMRDCCapabilityV1610Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-v1610: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasAndMobParametersMRDC_v1610 = new(MeasAndMobParametersMRDC_v1610)
			if err := ie.MeasAndMobParametersMRDC_v1610.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. generalParametersMRDC-v1610: ref
	{
		if seq.IsComponentPresent(1) {
			ie.GeneralParametersMRDC_v1610 = new(GeneralParametersMRDC_v1610)
			if err := ie.GeneralParametersMRDC_v1610.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. pdcp-ParametersMRDC-v1610: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Pdcp_ParametersMRDC_v1610 = new(PDCP_ParametersMRDC_v1610)
			if err := ie.Pdcp_ParametersMRDC_v1610.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(3) {
			ie.NonCriticalExtension = new(UE_MRDC_Capability_v1700)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
