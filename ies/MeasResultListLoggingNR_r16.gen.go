// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultListLoggingNR-r16 (line 3422).
// MeasResultListLoggingNR-r16 ::=      SEQUENCE (SIZE (1..maxCellReport)) OF MeasResultLoggingNR-r16

var measResultListLoggingNRR16SizeConstraints = per.SizeRange(1, common.MaxCellReport)

type MeasResultListLoggingNR_r16 []MeasResultLoggingNR_r16

func (ie *MeasResultListLoggingNR_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measResultListLoggingNRR16SizeConstraints)
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

func (ie *MeasResultListLoggingNR_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measResultListLoggingNRR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasResultListLoggingNR_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
