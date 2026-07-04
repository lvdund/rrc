// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LogMeasResultListBT-r16 (line 26250).
// LogMeasResultListBT-r16 ::= SEQUENCE (SIZE (1..maxBT-IdReport-r16)) OF LogMeasResultBT-r16

var logMeasResultListBTR16SizeConstraints = per.SizeRange(1, common.MaxBT_IdReport_r16)

type LogMeasResultListBT_r16 []LogMeasResultBT_r16

func (ie *LogMeasResultListBT_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(logMeasResultListBTR16SizeConstraints)
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

func (ie *LogMeasResultListBT_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(logMeasResultListBTR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(LogMeasResultListBT_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
