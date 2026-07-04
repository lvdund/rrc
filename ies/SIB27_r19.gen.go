// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB27-r19 (line 4751).

var sIB27R19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "intendedServiceAreaList-r19", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
}

var sIB27R19LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB27_r19 struct {
	IntendedServiceAreaList_r19 *IntendedServiceAreaList_r19
	LateNonCriticalExtension    []byte
}

func (ie *SIB27_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB27R19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IntendedServiceAreaList_r19 != nil, ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. intendedServiceAreaList-r19: ref
	{
		if ie.IntendedServiceAreaList_r19 != nil {
			if err := ie.IntendedServiceAreaList_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB27R19LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB27_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB27R19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. intendedServiceAreaList-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.IntendedServiceAreaList_r19 = new(IntendedServiceAreaList_r19)
			if err := ie.IntendedServiceAreaList_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(sIB27R19LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	return nil
}
