// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCSetupComplete-v17b0-IEs (line 1772).

var rRCSetupCompleteV17b0IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ul-RRC-MaxCapaSegments-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	RRCSetupComplete_V17b0_IEs_Ul_RRC_MaxCapaSegments_r17_True = 0
)

var rRCSetupCompleteV17b0IEsUlRRCMaxCapaSegmentsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCSetupComplete_V17b0_IEs_Ul_RRC_MaxCapaSegments_r17_True},
}

type RRCSetupComplete_V17b0_IEs struct {
	Ul_RRC_MaxCapaSegments_r17 *int64
	NonCriticalExtension       *struct{}
}

func (ie *RRCSetupComplete_V17b0_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCSetupCompleteV17b0IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ul_RRC_MaxCapaSegments_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. ul-RRC-MaxCapaSegments-r17: enumerated
	{
		if ie.Ul_RRC_MaxCapaSegments_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_RRC_MaxCapaSegments_r17, rRCSetupCompleteV17b0IEsUlRRCMaxCapaSegmentsR17Constraints); err != nil {
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

func (ie *RRCSetupComplete_V17b0_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCSetupCompleteV17b0IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ul-RRC-MaxCapaSegments-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(rRCSetupCompleteV17b0IEsUlRRCMaxCapaSegmentsR17Constraints)
			if err != nil {
				return err
			}
			ie.Ul_RRC_MaxCapaSegments_r17 = &idx
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
