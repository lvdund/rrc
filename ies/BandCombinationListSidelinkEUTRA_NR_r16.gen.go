// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombinationListSidelinkEUTRA-NR-r16 (line 17139).
// BandCombinationListSidelinkEUTRA-NR-r16 ::= SEQUENCE (SIZE (1..maxBandComb)) OF BandCombinationParametersSidelinkEUTRA-NR-r16

var bandCombinationListSidelinkEUTRANRR16SizeConstraints = per.SizeRange(1, common.MaxBandComb)

type BandCombinationListSidelinkEUTRA_NR_r16 []BandCombinationParametersSidelinkEUTRA_NR_r16

func (ie *BandCombinationListSidelinkEUTRA_NR_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(bandCombinationListSidelinkEUTRANRR16SizeConstraints)
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

func (ie *BandCombinationListSidelinkEUTRA_NR_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(bandCombinationListSidelinkEUTRANRR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(BandCombinationListSidelinkEUTRA_NR_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
