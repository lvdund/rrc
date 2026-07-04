// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CoverageAreaInfoList-r18 (line 4676).
// CoverageAreaInfoList-r18 ::=   SEQUENCE (SIZE (1..maxTN-AreaInfo-r18)) OF CoverageAreaInfo-r18

var coverageAreaInfoListR18SizeConstraints = per.SizeRange(1, common.MaxTN_AreaInfo_r18)

type CoverageAreaInfoList_r18 []CoverageAreaInfo_r18

func (ie *CoverageAreaInfoList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(coverageAreaInfoListR18SizeConstraints)
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

func (ie *CoverageAreaInfoList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(coverageAreaInfoListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CoverageAreaInfoList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
