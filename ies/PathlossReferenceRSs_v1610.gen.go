// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PathlossReferenceRSs-v1610 (line 12257).
// PathlossReferenceRSs-v1610 ::=          SEQUENCE (SIZE (1..maxNrofPUCCH-PathlossReferenceRSsDiff-r16)) OF PUCCH-PathlossReferenceRS-r16

var pathlossReferenceRSsV1610SizeConstraints = per.SizeRange(1, common.MaxNrofPUCCH_PathlossReferenceRSsDiff_r16)

type PathlossReferenceRSs_v1610 []PUCCH_PathlossReferenceRS_r16

func (ie *PathlossReferenceRSs_v1610) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pathlossReferenceRSsV1610SizeConstraints)
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

func (ie *PathlossReferenceRSs_v1610) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pathlossReferenceRSsV1610SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PathlossReferenceRSs_v1610, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
