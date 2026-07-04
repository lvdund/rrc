// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombinationParametersSidelinkEUTRA-NR-v1630 (line 17147).
// BandCombinationParametersSidelinkEUTRA-NR-v1630 ::= SEQUENCE (SIZE (1..maxSimultaneousBands)) OF BandParametersSidelinkEUTRA-NR-v1630

var bandCombinationParametersSidelinkEUTRANRV1630SizeConstraints = per.SizeRange(1, common.MaxSimultaneousBands)

type BandCombinationParametersSidelinkEUTRA_NR_v1630 []BandParametersSidelinkEUTRA_NR_v1630

func (ie *BandCombinationParametersSidelinkEUTRA_NR_v1630) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(bandCombinationParametersSidelinkEUTRANRV1630SizeConstraints)
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

func (ie *BandCombinationParametersSidelinkEUTRA_NR_v1630) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(bandCombinationParametersSidelinkEUTRANRV1630SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(BandCombinationParametersSidelinkEUTRA_NR_v1630, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
