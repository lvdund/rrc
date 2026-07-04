// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultFreqListFailMRDC (line 1920).
// MeasResultFreqListFailMRDC ::=      SEQUENCE (SIZE (1.. maxFreq)) OF MeasResult2EUTRA

var measResultFreqListFailMRDCSizeConstraints = per.SizeRange(1, common.MaxFreq)

type MeasResultFreqListFailMRDC []MeasResult2EUTRA

func (ie *MeasResultFreqListFailMRDC) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measResultFreqListFailMRDCSizeConstraints)
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

func (ie *MeasResultFreqListFailMRDC) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measResultFreqListFailMRDCSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasResultFreqListFailMRDC, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
