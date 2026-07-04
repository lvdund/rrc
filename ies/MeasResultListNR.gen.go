// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultListNR (line 9760).
// MeasResultListNR ::=                    SEQUENCE (SIZE (1..maxCellReport)) OF MeasResultNR

var measResultListNRSizeConstraints = per.SizeRange(1, common.MaxCellReport)

type MeasResultListNR []MeasResultNR

func (ie *MeasResultListNR) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measResultListNRSizeConstraints)
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

func (ie *MeasResultListNR) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measResultListNRSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasResultListNR, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
