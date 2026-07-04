// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB25-r18 (line 4670).

var sIB25R18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "coverageAreaInfoList-r18", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
}

var sIB25R18LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB25_r18 struct {
	CoverageAreaInfoList_r18 *CoverageAreaInfoList_r18
	LateNonCriticalExtension []byte
}

func (ie *SIB25_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB25R18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CoverageAreaInfoList_r18 != nil, ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. coverageAreaInfoList-r18: ref
	{
		if ie.CoverageAreaInfoList_r18 != nil {
			if err := ie.CoverageAreaInfoList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB25R18LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB25_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB25R18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. coverageAreaInfoList-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.CoverageAreaInfoList_r18 = new(CoverageAreaInfoList_r18)
			if err := ie.CoverageAreaInfoList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(sIB25R18LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	return nil
}
