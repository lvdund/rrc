// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultNeighFreqListRSSI-r18 (line 3397).
// MeasResultNeighFreqListRSSI-r18 ::=      SEQUENCE(SIZE (1..maxFreq)) OF MeasResultNeighFreqRSSI-r18

var measResultNeighFreqListRSSIR18SizeConstraints = per.SizeRange(1, common.MaxFreq)

type MeasResultNeighFreqListRSSI_r18 []MeasResultNeighFreqRSSI_r18

func (ie *MeasResultNeighFreqListRSSI_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measResultNeighFreqListRSSIR18SizeConstraints)
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

func (ie *MeasResultNeighFreqListRSSI_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measResultNeighFreqListRSSIR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasResultNeighFreqListRSSI_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
