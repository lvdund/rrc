// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB14-r16 (line 4372).

var sIB14R16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-V2X-ConfigCommonExt-r16"},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
}

var sIB14R16SlV2XConfigCommonExtR16Constraints = per.SizeConstraints{}

var sIB14R16LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB14_r16 struct {
	Sl_V2X_ConfigCommonExt_r16 []byte
	LateNonCriticalExtension   []byte
}

func (ie *SIB14_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB14R16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. sl-V2X-ConfigCommonExt-r16: octet-string
	{
		if err := e.EncodeOctetString(ie.Sl_V2X_ConfigCommonExt_r16, sIB14R16SlV2XConfigCommonExtR16Constraints); err != nil {
			return err
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB14R16LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB14_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB14R16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-V2X-ConfigCommonExt-r16: octet-string
	{
		v0, err := d.DecodeOctetString(sIB14R16SlV2XConfigCommonExtR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_V2X_ConfigCommonExt_r16 = v0
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(sIB14R16LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	return nil
}
