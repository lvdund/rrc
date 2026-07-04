// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BeamFailureDetection-r17 (line 13208).

var beamFailureDetectionR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "failureDetectionSet1-r17", Optional: true},
		{Name: "failureDetectionSet2-r17", Optional: true},
		{Name: "additionalPCI-r17", Optional: true},
	},
}

type BeamFailureDetection_r17 struct {
	FailureDetectionSet1_r17 *BeamFailureDetectionSet_r17
	FailureDetectionSet2_r17 *BeamFailureDetectionSet_r17
	AdditionalPCI_r17        *AdditionalPCIIndex_r17
}

func (ie *BeamFailureDetection_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(beamFailureDetectionR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FailureDetectionSet1_r17 != nil, ie.FailureDetectionSet2_r17 != nil, ie.AdditionalPCI_r17 != nil}); err != nil {
		return err
	}

	// 2. failureDetectionSet1-r17: ref
	{
		if ie.FailureDetectionSet1_r17 != nil {
			if err := ie.FailureDetectionSet1_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. failureDetectionSet2-r17: ref
	{
		if ie.FailureDetectionSet2_r17 != nil {
			if err := ie.FailureDetectionSet2_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. additionalPCI-r17: ref
	{
		if ie.AdditionalPCI_r17 != nil {
			if err := ie.AdditionalPCI_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BeamFailureDetection_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(beamFailureDetectionR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. failureDetectionSet1-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FailureDetectionSet1_r17 = new(BeamFailureDetectionSet_r17)
			if err := ie.FailureDetectionSet1_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. failureDetectionSet2-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.FailureDetectionSet2_r17 = new(BeamFailureDetectionSet_r17)
			if err := ie.FailureDetectionSet2_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. additionalPCI-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.AdditionalPCI_r17 = new(AdditionalPCIIndex_r17)
			if err := ie.AdditionalPCI_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
