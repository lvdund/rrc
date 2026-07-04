// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-MRDC-Capability-v1560 (line 25557).

var uEMRDCCapabilityV1560Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "receivedFilters", Optional: true},
		{Name: "measAndMobParametersMRDC-v1560", Optional: true},
		{Name: "fdd-Add-UE-MRDC-Capabilities-v1560", Optional: true},
		{Name: "tdd-Add-UE-MRDC-Capabilities-v1560", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var uEMRDCCapabilityV1560ReceivedFiltersConstraints = per.SizeConstraints{}

type UE_MRDC_Capability_v1560 struct {
	ReceivedFilters                    []byte
	MeasAndMobParametersMRDC_v1560     *MeasAndMobParametersMRDC_v1560
	Fdd_Add_UE_MRDC_Capabilities_v1560 *UE_MRDC_CapabilityAddXDD_Mode_v1560
	Tdd_Add_UE_MRDC_Capabilities_v1560 *UE_MRDC_CapabilityAddXDD_Mode_v1560
	NonCriticalExtension               *UE_MRDC_Capability_v1610
}

func (ie *UE_MRDC_Capability_v1560) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEMRDCCapabilityV1560Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ReceivedFilters != nil, ie.MeasAndMobParametersMRDC_v1560 != nil, ie.Fdd_Add_UE_MRDC_Capabilities_v1560 != nil, ie.Tdd_Add_UE_MRDC_Capabilities_v1560 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. receivedFilters: octet-string
	{
		if ie.ReceivedFilters != nil {
			if err := e.EncodeOctetString(ie.ReceivedFilters, uEMRDCCapabilityV1560ReceivedFiltersConstraints); err != nil {
				return err
			}
		}
	}

	// 3. measAndMobParametersMRDC-v1560: ref
	{
		if ie.MeasAndMobParametersMRDC_v1560 != nil {
			if err := ie.MeasAndMobParametersMRDC_v1560.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. fdd-Add-UE-MRDC-Capabilities-v1560: ref
	{
		if ie.Fdd_Add_UE_MRDC_Capabilities_v1560 != nil {
			if err := ie.Fdd_Add_UE_MRDC_Capabilities_v1560.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. tdd-Add-UE-MRDC-Capabilities-v1560: ref
	{
		if ie.Tdd_Add_UE_MRDC_Capabilities_v1560 != nil {
			if err := ie.Tdd_Add_UE_MRDC_Capabilities_v1560.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_MRDC_Capability_v1560) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEMRDCCapabilityV1560Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. receivedFilters: octet-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeOctetString(uEMRDCCapabilityV1560ReceivedFiltersConstraints)
			if err != nil {
				return err
			}
			ie.ReceivedFilters = v
		}
	}

	// 3. measAndMobParametersMRDC-v1560: ref
	{
		if seq.IsComponentPresent(1) {
			ie.MeasAndMobParametersMRDC_v1560 = new(MeasAndMobParametersMRDC_v1560)
			if err := ie.MeasAndMobParametersMRDC_v1560.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. fdd-Add-UE-MRDC-Capabilities-v1560: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Fdd_Add_UE_MRDC_Capabilities_v1560 = new(UE_MRDC_CapabilityAddXDD_Mode_v1560)
			if err := ie.Fdd_Add_UE_MRDC_Capabilities_v1560.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. tdd-Add-UE-MRDC-Capabilities-v1560: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Tdd_Add_UE_MRDC_Capabilities_v1560 = new(UE_MRDC_CapabilityAddXDD_Mode_v1560)
			if err := ie.Tdd_Add_UE_MRDC_Capabilities_v1560.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(4) {
			ie.NonCriticalExtension = new(UE_MRDC_Capability_v1610)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
