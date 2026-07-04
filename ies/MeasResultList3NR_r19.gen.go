// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultList3NR-r19 (line 3530).
// MeasResultList3NR-r19 ::=            SEQUENCE (SIZE (1..maxFreq)) OF MeasResult3NR-r19

var measResultList3NRR19SizeConstraints = per.SizeRange(1, common.MaxFreq)

type MeasResultList3NR_r19 []MeasResult3NR_r19

func (ie *MeasResultList3NR_r19) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measResultList3NRR19SizeConstraints)
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

func (ie *MeasResultList3NR_r19) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measResultList3NRR19SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasResultList3NR_r19, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
