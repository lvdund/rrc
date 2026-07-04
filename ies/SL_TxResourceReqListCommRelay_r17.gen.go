// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-TxResourceReqListCommRelay-r17 (line 2216).
// SL-TxResourceReqListCommRelay-r17 ::=  SEQUENCE (SIZE (1..maxNrofSL-Dest-r16)) OF SL-TxResourceReqCommRelayInfo-r17

var sLTxResourceReqListCommRelayR17SizeConstraints = per.SizeRange(1, common.MaxNrofSL_Dest_r16)

type SL_TxResourceReqListCommRelay_r17 []SL_TxResourceReqCommRelayInfo_r17

func (ie *SL_TxResourceReqListCommRelay_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLTxResourceReqListCommRelayR17SizeConstraints)
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

func (ie *SL_TxResourceReqListCommRelay_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLTxResourceReqListCommRelayR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_TxResourceReqListCommRelay_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
