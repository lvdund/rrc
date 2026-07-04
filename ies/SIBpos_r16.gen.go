// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIBpos-r16 (line 4929).

var sIBposR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "assistanceDataSIB-Element-r16"},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
}

var sIBposR16AssistanceDataSIBElementR16Constraints = per.SizeConstraints{}

var sIBposR16LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIBpos_r16 struct {
	AssistanceDataSIB_Element_r16 []byte
	LateNonCriticalExtension      []byte
}

func (ie *SIBpos_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIBposR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. assistanceDataSIB-Element-r16: octet-string
	{
		if err := e.EncodeOctetString(ie.AssistanceDataSIB_Element_r16, sIBposR16AssistanceDataSIBElementR16Constraints); err != nil {
			return err
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIBposR16LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIBpos_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIBposR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. assistanceDataSIB-Element-r16: octet-string
	{
		v0, err := d.DecodeOctetString(sIBposR16AssistanceDataSIBElementR16Constraints)
		if err != nil {
			return err
		}
		ie.AssistanceDataSIB_Element_r16 = v0
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(sIBposR16LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	return nil
}
