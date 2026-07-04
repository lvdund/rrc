// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultListEUTRA (line 9794).
// MeasResultListEUTRA ::=                 SEQUENCE (SIZE (1..maxCellReport)) OF MeasResultEUTRA

var measResultListEUTRASizeConstraints = per.SizeRange(1, common.MaxCellReport)

type MeasResultListEUTRA []MeasResultEUTRA

func (ie *MeasResultListEUTRA) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measResultListEUTRASizeConstraints)
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

func (ie *MeasResultListEUTRA) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measResultListEUTRASizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasResultListEUTRA, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
