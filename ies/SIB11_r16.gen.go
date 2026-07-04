// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB11-r16 (line 4263).

var sIB11R16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "measIdleConfigSIB-r16", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
}

var sIB11R16LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB11_r16 struct {
	MeasIdleConfigSIB_r16    *MeasIdleConfigSIB_r16
	LateNonCriticalExtension []byte
}

func (ie *SIB11_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB11R16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasIdleConfigSIB_r16 != nil, ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. measIdleConfigSIB-r16: ref
	{
		if ie.MeasIdleConfigSIB_r16 != nil {
			if err := ie.MeasIdleConfigSIB_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB11R16LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB11_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB11R16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. measIdleConfigSIB-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasIdleConfigSIB_r16 = new(MeasIdleConfigSIB_r16)
			if err := ie.MeasIdleConfigSIB_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(sIB11R16LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	return nil
}
