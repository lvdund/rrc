// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultList2NR (line 9999).
// MeasResultList2NR ::=               SEQUENCE (SIZE (1..maxFreq)) OF MeasResult2NR

var measResultList2NRSizeConstraints = per.SizeRange(1, common.MaxFreq)

type MeasResultList2NR []MeasResult2NR

func (ie *MeasResultList2NR) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measResultList2NRSizeConstraints)
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

func (ie *MeasResultList2NR) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measResultList2NRSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasResultList2NR, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
