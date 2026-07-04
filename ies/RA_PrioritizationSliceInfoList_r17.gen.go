// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RA-PrioritizationSliceInfoList-r17 (line 13097).
// RA-PrioritizationSliceInfoList-r17 ::= SEQUENCE (SIZE (1..maxSliceInfo-r17)) OF RA-PrioritizationSliceInfo-r17

var rAPrioritizationSliceInfoListR17SizeConstraints = per.SizeRange(1, common.MaxSliceInfo_r17)

type RA_PrioritizationSliceInfoList_r17 []RA_PrioritizationSliceInfo_r17

func (ie *RA_PrioritizationSliceInfoList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(rAPrioritizationSliceInfoListR17SizeConstraints)
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

func (ie *RA_PrioritizationSliceInfoList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(rAPrioritizationSliceInfoListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(RA_PrioritizationSliceInfoList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
