// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-PosTxResourceReqList-r18 (line 2164).
// SL-PosTxResourceReqList-r18 ::=        SEQUENCE (SIZE (1..maxNrofSL-Dest-r16)) OF SL-PosTxResourceReq-r18

var sLPosTxResourceReqListR18SizeConstraints = per.SizeRange(1, common.MaxNrofSL_Dest_r16)

type SL_PosTxResourceReqList_r18 []SL_PosTxResourceReq_r18

func (ie *SL_PosTxResourceReqList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLPosTxResourceReqListR18SizeConstraints)
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

func (ie *SL_PosTxResourceReqList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLPosTxResourceReqListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_PosTxResourceReqList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
