// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-MeasIdToRemoveList-r16 (line 27533).
// SL-MeasIdToRemoveList-r16 ::=       SEQUENCE (SIZE (1..maxNrofSL-MeasId-r16)) OF SL-MeasId-r16

var sLMeasIdToRemoveListR16SizeConstraints = per.SizeRange(1, common.MaxNrofSL_MeasId_r16)

type SL_MeasIdToRemoveList_r16 []SL_MeasId_r16

func (ie *SL_MeasIdToRemoveList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLMeasIdToRemoveListR16SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := (*ie)[i].Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SL_MeasIdToRemoveList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLMeasIdToRemoveListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_MeasIdToRemoveList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
