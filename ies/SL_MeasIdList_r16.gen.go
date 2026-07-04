// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-MeasIdList-r16 (line 27538).
// SL-MeasIdList-r16 ::=               SEQUENCE (SIZE (1..maxNrofSL-MeasId-r16)) OF SL-MeasIdInfo-r16

var sLMeasIdListR16SizeConstraints = per.SizeRange(1, common.MaxNrofSL_MeasId_r16)

type SL_MeasIdList_r16 []SL_MeasIdInfo_r16

func (ie *SL_MeasIdList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLMeasIdListR16SizeConstraints)
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

func (ie *SL_MeasIdList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLMeasIdListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_MeasIdList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
