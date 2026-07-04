// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-TxResourceReqListDisc-r17 (line 2201).
// SL-TxResourceReqListDisc-r17 ::=       SEQUENCE (SIZE (1..maxNrofSL-Dest-r16)) OF SL-TxResourceReqDisc-r17

var sLTxResourceReqListDiscR17SizeConstraints = per.SizeRange(1, common.MaxNrofSL_Dest_r16)

type SL_TxResourceReqListDisc_r17 []SL_TxResourceReqDisc_r17

func (ie *SL_TxResourceReqListDisc_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLTxResourceReqListDiscR17SizeConstraints)
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

func (ie *SL_TxResourceReqListDisc_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLTxResourceReqListDiscR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_TxResourceReqListDisc_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
