// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v1560 (line 25685).

var uENRCapabilityV1560Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nrdc-Parameters", Optional: true},
		{Name: "receivedFilters", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var uENRCapabilityV1560ReceivedFiltersConstraints = per.SizeConstraints{}

type UE_NR_Capability_v1560 struct {
	Nrdc_Parameters      *NRDC_Parameters
	ReceivedFilters      []byte
	NonCriticalExtension *UE_NR_Capability_v1570
}

func (ie *UE_NR_Capability_v1560) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV1560Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Nrdc_Parameters != nil, ie.ReceivedFilters != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. nrdc-Parameters: ref
	{
		if ie.Nrdc_Parameters != nil {
			if err := ie.Nrdc_Parameters.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. receivedFilters: octet-string
	{
		if ie.ReceivedFilters != nil {
			if err := e.EncodeOctetString(ie.ReceivedFilters, uENRCapabilityV1560ReceivedFiltersConstraints); err != nil {
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

func (ie *UE_NR_Capability_v1560) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV1560Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. nrdc-Parameters: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Nrdc_Parameters = new(NRDC_Parameters)
			if err := ie.Nrdc_Parameters.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. receivedFilters: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(uENRCapabilityV1560ReceivedFiltersConstraints)
			if err != nil {
				return err
			}
			ie.ReceivedFilters = v
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = new(UE_NR_Capability_v1570)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
