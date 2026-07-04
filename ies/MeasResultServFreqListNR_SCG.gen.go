// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultServFreqListNR-SCG (line 9836).
// MeasResultServFreqListNR-SCG ::= SEQUENCE (SIZE (1..maxNrofServingCells)) OF MeasResult2NR

var measResultServFreqListNRSCGSizeConstraints = per.SizeRange(1, common.MaxNrofServingCells)

type MeasResultServFreqListNR_SCG []MeasResult2NR

func (ie *MeasResultServFreqListNR_SCG) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measResultServFreqListNRSCGSizeConstraints)
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

func (ie *MeasResultServFreqListNR_SCG) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measResultServFreqListNRSCGSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasResultServFreqListNR_SCG, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
