// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TDD-UL-DL-ConfigCommon (line 16067).

var tDDULDLConfigCommonConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "referenceSubcarrierSpacing"},
		{Name: "pattern1"},
		{Name: "pattern2", Optional: true},
	},
}

type TDD_UL_DL_ConfigCommon struct {
	ReferenceSubcarrierSpacing SubcarrierSpacing
	Pattern1                   TDD_UL_DL_Pattern
	Pattern2                   *TDD_UL_DL_Pattern
}

func (ie *TDD_UL_DL_ConfigCommon) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tDDULDLConfigCommonConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pattern2 != nil}); err != nil {
		return err
	}

	// 3. referenceSubcarrierSpacing: ref
	{
		if err := ie.ReferenceSubcarrierSpacing.Encode(e); err != nil {
			return err
		}
	}

	// 4. pattern1: ref
	{
		if err := ie.Pattern1.Encode(e); err != nil {
			return err
		}
	}

	// 5. pattern2: ref
	{
		if ie.Pattern2 != nil {
			if err := ie.Pattern2.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *TDD_UL_DL_ConfigCommon) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tDDULDLConfigCommonConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. referenceSubcarrierSpacing: ref
	{
		if err := ie.ReferenceSubcarrierSpacing.Decode(d); err != nil {
			return err
		}
	}

	// 4. pattern1: ref
	{
		if err := ie.Pattern1.Decode(d); err != nil {
			return err
		}
	}

	// 5. pattern2: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Pattern2 = new(TDD_UL_DL_Pattern)
			if err := ie.Pattern2.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
