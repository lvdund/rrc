// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ThresholdMBS-List-r18 (line 670).
// ThresholdMBS-List-r18 ::= SEQUENCE (SIZE (1..maxNrofThresholdMBS-r18)) OF ThresholdMBS-r18

var thresholdMBSListR18SizeConstraints = per.SizeRange(1, common.MaxNrofThresholdMBS_r18)

type ThresholdMBS_List_r18 []ThresholdMBS_r18

func (ie *ThresholdMBS_List_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(thresholdMBSListR18SizeConstraints)
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

func (ie *ThresholdMBS_List_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(thresholdMBSListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(ThresholdMBS_List_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
