// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultListUTRA-FDD-r16 (line 9838).
// MeasResultListUTRA-FDD-r16 ::=          SEQUENCE (SIZE (1..maxCellReport)) OF MeasResultUTRA-FDD-r16

var measResultListUTRAFDDR16SizeConstraints = per.SizeRange(1, common.MaxCellReport)

type MeasResultListUTRA_FDD_r16 []MeasResultUTRA_FDD_r16

func (ie *MeasResultListUTRA_FDD_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measResultListUTRAFDDR16SizeConstraints)
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

func (ie *MeasResultListUTRA_FDD_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measResultListUTRAFDDR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasResultListUTRA_FDD_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
