// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombinationListSidelinkEUTRA-NR-v1630 (line 17141).
// BandCombinationListSidelinkEUTRA-NR-v1630 ::= SEQUENCE (SIZE (1..maxBandComb)) OF BandCombinationParametersSidelinkEUTRA-NR-v1630

var bandCombinationListSidelinkEUTRANRV1630SizeConstraints = per.SizeRange(1, common.MaxBandComb)

type BandCombinationListSidelinkEUTRA_NR_v1630 []BandCombinationParametersSidelinkEUTRA_NR_v1630

func (ie *BandCombinationListSidelinkEUTRA_NR_v1630) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(bandCombinationListSidelinkEUTRANRV1630SizeConstraints)
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

func (ie *BandCombinationListSidelinkEUTRA_NR_v1630) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(bandCombinationListSidelinkEUTRANRV1630SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(BandCombinationListSidelinkEUTRA_NR_v1630, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
