// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-CBR-LevelsDedicatedSL-PRS-RP-r18 (line 26932).
// SL-CBR-LevelsDedicatedSL-PRS-RP-r18 ::= SEQUENCE (SIZE (0..maxCBR-LevelDedSL-PRS-1-r18)) OF SL-CBR-DedicatedSL-PRS-RP-r18

var sLCBRLevelsDedicatedSLPRSRPR18SizeConstraints = per.SizeRange(0, common.MaxCBR_LevelDedSL_PRS_1_r18)

type SL_CBR_LevelsDedicatedSL_PRS_RP_r18 []SL_CBR_DedicatedSL_PRS_RP_r18

func (ie *SL_CBR_LevelsDedicatedSL_PRS_RP_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLCBRLevelsDedicatedSLPRSRPR18SizeConstraints)
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

func (ie *SL_CBR_LevelsDedicatedSL_PRS_RP_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLCBRLevelsDedicatedSLPRSRPR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_CBR_LevelsDedicatedSL_PRS_RP_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
