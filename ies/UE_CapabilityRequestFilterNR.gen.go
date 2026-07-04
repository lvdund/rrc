// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-CapabilityRequestFilterNR (line 25523).

var uECapabilityRequestFilterNRConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "frequencyBandListFilter", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type UE_CapabilityRequestFilterNR struct {
	FrequencyBandListFilter *FreqBandList
	NonCriticalExtension    *UE_CapabilityRequestFilterNR_v1540
}

func (ie *UE_CapabilityRequestFilterNR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uECapabilityRequestFilterNRConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FrequencyBandListFilter != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. frequencyBandListFilter: ref
	{
		if ie.FrequencyBandListFilter != nil {
			if err := ie.FrequencyBandListFilter.Encode(e); err != nil {
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

func (ie *UE_CapabilityRequestFilterNR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uECapabilityRequestFilterNRConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. frequencyBandListFilter: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FrequencyBandListFilter = new(FreqBandList)
			if err := ie.FrequencyBandListFilter.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(UE_CapabilityRequestFilterNR_v1540)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
