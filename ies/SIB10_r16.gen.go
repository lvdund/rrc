// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB10-r16 (line 4248).

var sIB10R16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "hrnn-List-r16", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
}

var sIB10R16LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB10_r16 struct {
	Hrnn_List_r16            *HRNN_List_r16
	LateNonCriticalExtension []byte
}

func (ie *SIB10_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB10R16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Hrnn_List_r16 != nil, ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. hrnn-List-r16: ref
	{
		if ie.Hrnn_List_r16 != nil {
			if err := ie.Hrnn_List_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB10R16LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB10_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB10R16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. hrnn-List-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Hrnn_List_r16 = new(HRNN_List_r16)
			if err := ie.Hrnn_List_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(sIB10R16LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	return nil
}
