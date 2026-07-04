// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SubsequentCondReconfig-r18 (line 6483).

var subsequentCondReconfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "condExecutionCondToReleaseList-r18", Optional: true},
		{Name: "condExecutionCondToAddModList-r18", Optional: true},
	},
}

type SubsequentCondReconfig_r18 struct {
	CondExecutionCondToReleaseList_r18 *CondExecutionCondToReleaseList_r18
	CondExecutionCondToAddModList_r18  *CondExecutionCondToAddModList_r18
}

func (ie *SubsequentCondReconfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(subsequentCondReconfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CondExecutionCondToReleaseList_r18 != nil, ie.CondExecutionCondToAddModList_r18 != nil}); err != nil {
		return err
	}

	// 3. condExecutionCondToReleaseList-r18: ref
	{
		if ie.CondExecutionCondToReleaseList_r18 != nil {
			if err := ie.CondExecutionCondToReleaseList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. condExecutionCondToAddModList-r18: ref
	{
		if ie.CondExecutionCondToAddModList_r18 != nil {
			if err := ie.CondExecutionCondToAddModList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SubsequentCondReconfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(subsequentCondReconfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. condExecutionCondToReleaseList-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.CondExecutionCondToReleaseList_r18 = new(CondExecutionCondToReleaseList_r18)
			if err := ie.CondExecutionCondToReleaseList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. condExecutionCondToAddModList-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.CondExecutionCondToAddModList_r18 = new(CondExecutionCondToAddModList_r18)
			if err := ie.CondExecutionCondToAddModList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
