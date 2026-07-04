// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultListLogging2NR-r16 (line 3415).
// MeasResultListLogging2NR-r16 ::=     SEQUENCE(SIZE (1..maxFreq)) OF MeasResultLogging2NR-r16

var measResultListLogging2NRR16SizeConstraints = per.SizeRange(1, common.MaxFreq)

type MeasResultListLogging2NR_r16 []MeasResultLogging2NR_r16

func (ie *MeasResultListLogging2NR_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measResultListLogging2NRR16SizeConstraints)
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

func (ie *MeasResultListLogging2NR_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measResultListLogging2NRR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasResultListLogging2NR_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
