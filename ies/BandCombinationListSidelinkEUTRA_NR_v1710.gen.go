// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombinationListSidelinkEUTRA-NR-v1710 (line 17143).
// BandCombinationListSidelinkEUTRA-NR-v1710 ::= SEQUENCE (SIZE (1..maxBandComb)) OF BandCombinationParametersSidelinkEUTRA-NR-v1710

var bandCombinationListSidelinkEUTRANRV1710SizeConstraints = per.SizeRange(1, common.MaxBandComb)

type BandCombinationListSidelinkEUTRA_NR_v1710 []BandCombinationParametersSidelinkEUTRA_NR_v1710

func (ie *BandCombinationListSidelinkEUTRA_NR_v1710) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(bandCombinationListSidelinkEUTRANRV1710SizeConstraints)
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

func (ie *BandCombinationListSidelinkEUTRA_NR_v1710) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(bandCombinationListSidelinkEUTRANRV1710SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(BandCombinationListSidelinkEUTRA_NR_v1710, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
