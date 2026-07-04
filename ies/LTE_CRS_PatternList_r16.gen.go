// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LTE-CRS-PatternList-r16 (line 13306).
// LTE-CRS-PatternList-r16 ::=         SEQUENCE (SIZE (1..maxLTE-CRS-Patterns-r16)) OF RateMatchPatternLTE-CRS

var lTECRSPatternListR16SizeConstraints = per.SizeRange(1, common.MaxLTE_CRS_Patterns_r16)

type LTE_CRS_PatternList_r16 []RateMatchPatternLTE_CRS

func (ie *LTE_CRS_PatternList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(lTECRSPatternListR16SizeConstraints)
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

func (ie *LTE_CRS_PatternList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(lTECRSPatternListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(LTE_CRS_PatternList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
