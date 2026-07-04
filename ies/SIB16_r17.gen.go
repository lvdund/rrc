// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB16-r17 (line 4398).

var sIB16R17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "freqPriorityListSlicing-r17", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
}

var sIB16R17LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB16_r17 struct {
	FreqPriorityListSlicing_r17 *FreqPriorityListSlicing_r17
	LateNonCriticalExtension    []byte
}

func (ie *SIB16_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB16R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FreqPriorityListSlicing_r17 != nil, ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. freqPriorityListSlicing-r17: ref
	{
		if ie.FreqPriorityListSlicing_r17 != nil {
			if err := ie.FreqPriorityListSlicing_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB16R17LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB16_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB16R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. freqPriorityListSlicing-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FreqPriorityListSlicing_r17 = new(FreqPriorityListSlicing_r17)
			if err := ie.FreqPriorityListSlicing_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(sIB16R17LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	return nil
}
