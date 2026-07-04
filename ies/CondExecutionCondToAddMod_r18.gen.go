// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CondExecutionCondToAddMod-r18 (line 6491).

var condExecutionCondToAddModR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "subsequentCondReconfigId-r18"},
		{Name: "subsequentCondExecutionCond-r18", Optional: true},
		{Name: "subsequentCondExecutionCondSCG-r18", Optional: true},
	},
}

var condExecutionCondToAddModR18SubsequentCondExecutionCondR18Constraints = per.SizeRange(1, 2)

var condExecutionCondToAddModR18SubsequentCondExecutionCondSCGR18Constraints = per.SizeConstraints{}

type CondExecutionCondToAddMod_r18 struct {
	SubsequentCondReconfigId_r18       CondReconfigId_r16
	SubsequentCondExecutionCond_r18    []MeasId
	SubsequentCondExecutionCondSCG_r18 []byte
}

func (ie *CondExecutionCondToAddMod_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(condExecutionCondToAddModR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SubsequentCondExecutionCond_r18 != nil, ie.SubsequentCondExecutionCondSCG_r18 != nil}); err != nil {
		return err
	}

	// 3. subsequentCondReconfigId-r18: ref
	{
		if err := ie.SubsequentCondReconfigId_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. subsequentCondExecutionCond-r18: sequence-of
	{
		if ie.SubsequentCondExecutionCond_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(condExecutionCondToAddModR18SubsequentCondExecutionCondR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SubsequentCondExecutionCond_r18))); err != nil {
				return err
			}
			for i := range ie.SubsequentCondExecutionCond_r18 {
				if err := ie.SubsequentCondExecutionCond_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. subsequentCondExecutionCondSCG-r18: octet-string
	{
		if ie.SubsequentCondExecutionCondSCG_r18 != nil {
			if err := e.EncodeOctetString(ie.SubsequentCondExecutionCondSCG_r18, condExecutionCondToAddModR18SubsequentCondExecutionCondSCGR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CondExecutionCondToAddMod_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(condExecutionCondToAddModR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. subsequentCondReconfigId-r18: ref
	{
		if err := ie.SubsequentCondReconfigId_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. subsequentCondExecutionCond-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(condExecutionCondToAddModR18SubsequentCondExecutionCondR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SubsequentCondExecutionCond_r18 = make([]MeasId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SubsequentCondExecutionCond_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. subsequentCondExecutionCondSCG-r18: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(condExecutionCondToAddModR18SubsequentCondExecutionCondSCGR18Constraints)
			if err != nil {
				return err
			}
			ie.SubsequentCondExecutionCondSCG_r18 = v
		}
	}

	return nil
}
