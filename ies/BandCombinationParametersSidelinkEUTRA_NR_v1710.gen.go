// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombinationParametersSidelinkEUTRA-NR-v1710 (line 17149).
// BandCombinationParametersSidelinkEUTRA-NR-v1710 ::= SEQUENCE (SIZE (1..maxSimultaneousBands)) OF BandParametersSidelinkEUTRA-NR-v1710

var bandCombinationParametersSidelinkEUTRANRV1710SizeConstraints = per.SizeRange(1, common.MaxSimultaneousBands)

type BandCombinationParametersSidelinkEUTRA_NR_v1710 []BandParametersSidelinkEUTRA_NR_v1710

func (ie *BandCombinationParametersSidelinkEUTRA_NR_v1710) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(bandCombinationParametersSidelinkEUTRANRV1710SizeConstraints)
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

func (ie *BandCombinationParametersSidelinkEUTRA_NR_v1710) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(bandCombinationParametersSidelinkEUTRANRV1710SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(BandCombinationParametersSidelinkEUTRA_NR_v1710, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
