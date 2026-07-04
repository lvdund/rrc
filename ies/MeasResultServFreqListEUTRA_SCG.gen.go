// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultServFreqListEUTRA-SCG (line 9834).
// MeasResultServFreqListEUTRA-SCG ::= SEQUENCE (SIZE (1..maxNrofServingCellsEUTRA)) OF MeasResult2EUTRA

var measResultServFreqListEUTRASCGSizeConstraints = per.SizeRange(1, common.MaxNrofServingCellsEUTRA)

type MeasResultServFreqListEUTRA_SCG []MeasResult2EUTRA

func (ie *MeasResultServFreqListEUTRA_SCG) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measResultServFreqListEUTRASCGSizeConstraints)
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

func (ie *MeasResultServFreqListEUTRA_SCG) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measResultServFreqListEUTRASCGSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasResultServFreqListEUTRA_SCG, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
