// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-MeasObjectInfo-r16 (line 27554).

var sLMeasObjectInfoR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-MeasObjectId-r16"},
		{Name: "sl-MeasObject-r16"},
	},
}

type SL_MeasObjectInfo_r16 struct {
	Sl_MeasObjectId_r16 SL_MeasObjectId_r16
	Sl_MeasObject_r16   SL_MeasObject_r16
}

func (ie *SL_MeasObjectInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLMeasObjectInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. sl-MeasObjectId-r16: ref
	{
		if err := ie.Sl_MeasObjectId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. sl-MeasObject-r16: ref
	{
		if err := ie.Sl_MeasObject_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_MeasObjectInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLMeasObjectInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. sl-MeasObjectId-r16: ref
	{
		if err := ie.Sl_MeasObjectId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. sl-MeasObject-r16: ref
	{
		if err := ie.Sl_MeasObject_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
