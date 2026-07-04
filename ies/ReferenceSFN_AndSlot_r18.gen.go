// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReferenceSFN-AndSlot-r18 (line 2784).

var referenceSFNAndSlotR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "referenceSFN-r18"},
		{Name: "referenceSlot-r18"},
	},
}

var referenceSFNAndSlotR18ReferenceSFNR18Constraints = per.Constrained(0, 1023)

var referenceSFNAndSlotR18ReferenceSlotR18Constraints = per.Constrained(0, 639)

type ReferenceSFN_AndSlot_r18 struct {
	ReferenceSFN_r18  int64
	ReferenceSlot_r18 int64
}

func (ie *ReferenceSFN_AndSlot_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(referenceSFNAndSlotR18Constraints)
	_ = seq

	// 1. referenceSFN-r18: integer
	{
		if err := e.EncodeInteger(ie.ReferenceSFN_r18, referenceSFNAndSlotR18ReferenceSFNR18Constraints); err != nil {
			return err
		}
	}

	// 2. referenceSlot-r18: integer
	{
		if err := e.EncodeInteger(ie.ReferenceSlot_r18, referenceSFNAndSlotR18ReferenceSlotR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *ReferenceSFN_AndSlot_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(referenceSFNAndSlotR18Constraints)
	_ = seq

	// 1. referenceSFN-r18: integer
	{
		v0, err := d.DecodeInteger(referenceSFNAndSlotR18ReferenceSFNR18Constraints)
		if err != nil {
			return err
		}
		ie.ReferenceSFN_r18 = v0
	}

	// 2. referenceSlot-r18: integer
	{
		v1, err := d.DecodeInteger(referenceSFNAndSlotR18ReferenceSlotR18Constraints)
		if err != nil {
			return err
		}
		ie.ReferenceSlot_r18 = v1
	}

	return nil
}
