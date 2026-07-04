// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultList2UTRA (line 717).
// MeasResultList2UTRA ::=    SEQUENCE (SIZE (1..maxFreq)) OF MeasResult2UTRA-FDD-r16

var measResultList2UTRASizeConstraints = per.SizeRange(1, common.MaxFreq)

type MeasResultList2UTRA []MeasResult2UTRA_FDD_r16

func (ie *MeasResultList2UTRA) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measResultList2UTRASizeConstraints)
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

func (ie *MeasResultList2UTRA) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measResultList2UTRASizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasResultList2UTRA, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
