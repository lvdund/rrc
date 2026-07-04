// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultList2EUTRA (line 724).
// MeasResultList2EUTRA ::=          SEQUENCE (SIZE (1..maxFreq)) OF MeasResult2EUTRA-r16

var measResultList2EUTRASizeConstraints = per.SizeRange(1, common.MaxFreq)

type MeasResultList2EUTRA []MeasResult2EUTRA_r16

func (ie *MeasResultList2EUTRA) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measResultList2EUTRASizeConstraints)
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

func (ie *MeasResultList2EUTRA) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measResultList2EUTRASizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasResultList2EUTRA, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
