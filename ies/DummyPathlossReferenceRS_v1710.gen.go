// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DummyPathlossReferenceRS-v1710 (line 12608).

var dummyPathlossReferenceRSV1710Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pusch-PathlossReferenceRS-Id-r17"},
		{Name: "additionalPCI-r17", Optional: true},
	},
}

type DummyPathlossReferenceRS_v1710 struct {
	Pusch_PathlossReferenceRS_Id_r17 PUSCH_PathlossReferenceRS_Id_r17
	AdditionalPCI_r17                *AdditionalPCIIndex_r17
}

func (ie *DummyPathlossReferenceRS_v1710) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dummyPathlossReferenceRSV1710Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AdditionalPCI_r17 != nil}); err != nil {
		return err
	}

	// 2. pusch-PathlossReferenceRS-Id-r17: ref
	{
		if err := ie.Pusch_PathlossReferenceRS_Id_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. additionalPCI-r17: ref
	{
		if ie.AdditionalPCI_r17 != nil {
			if err := ie.AdditionalPCI_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *DummyPathlossReferenceRS_v1710) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dummyPathlossReferenceRSV1710Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pusch-PathlossReferenceRS-Id-r17: ref
	{
		if err := ie.Pusch_PathlossReferenceRS_Id_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. additionalPCI-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.AdditionalPCI_r17 = new(AdditionalPCIIndex_r17)
			if err := ie.AdditionalPCI_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
