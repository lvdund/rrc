// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombinationParametersSidelinkEUTRA-NR-r16 (line 17145).
// BandCombinationParametersSidelinkEUTRA-NR-r16 ::= SEQUENCE (SIZE (1..maxSimultaneousBands)) OF BandParametersSidelinkEUTRA-NR-r16

var bandCombinationParametersSidelinkEUTRANRR16SizeConstraints = per.SizeRange(1, common.MaxSimultaneousBands)

type BandCombinationParametersSidelinkEUTRA_NR_r16 []BandParametersSidelinkEUTRA_NR_r16

func (ie *BandCombinationParametersSidelinkEUTRA_NR_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(bandCombinationParametersSidelinkEUTRANRR16SizeConstraints)
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

func (ie *BandCombinationParametersSidelinkEUTRA_NR_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(bandCombinationParametersSidelinkEUTRANRR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(BandCombinationParametersSidelinkEUTRA_NR_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
