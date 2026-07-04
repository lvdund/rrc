// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-CC-SetIndex (line 15302).

var sRSCCSetIndexConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cc-SetIndex", Optional: true},
		{Name: "cc-IndexInOneCC-Set", Optional: true},
	},
}

var sRSCCSetIndexCcSetIndexConstraints = per.Constrained(0, 3)

var sRSCCSetIndexCcIndexInOneCCSetConstraints = per.Constrained(0, 7)

type SRS_CC_SetIndex struct {
	Cc_SetIndex         *int64
	Cc_IndexInOneCC_Set *int64
}

func (ie *SRS_CC_SetIndex) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSCCSetIndexConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Cc_SetIndex != nil, ie.Cc_IndexInOneCC_Set != nil}); err != nil {
		return err
	}

	// 2. cc-SetIndex: integer
	{
		if ie.Cc_SetIndex != nil {
			if err := e.EncodeInteger(*ie.Cc_SetIndex, sRSCCSetIndexCcSetIndexConstraints); err != nil {
				return err
			}
		}
	}

	// 3. cc-IndexInOneCC-Set: integer
	{
		if ie.Cc_IndexInOneCC_Set != nil {
			if err := e.EncodeInteger(*ie.Cc_IndexInOneCC_Set, sRSCCSetIndexCcIndexInOneCCSetConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SRS_CC_SetIndex) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSCCSetIndexConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. cc-SetIndex: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(sRSCCSetIndexCcSetIndexConstraints)
			if err != nil {
				return err
			}
			ie.Cc_SetIndex = &v
		}
	}

	// 3. cc-IndexInOneCC-Set: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sRSCCSetIndexCcIndexInOneCCSetConstraints)
			if err != nil {
				return err
			}
			ie.Cc_IndexInOneCC_Set = &v
		}
	}

	return nil
}
