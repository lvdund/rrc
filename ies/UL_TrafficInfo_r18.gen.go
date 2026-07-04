// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UL-TrafficInfo-r18 (line 2761).
// UL-TrafficInfo-r18 ::=                SEQUENCE (SIZE (1..maxNrofPDU-Sessions-r17)) OF PDU-SessionUL-TrafficInfo-r18

var uLTrafficInfoR18SizeConstraints = per.SizeRange(1, common.MaxNrofPDU_Sessions_r17)

type UL_TrafficInfo_r18 []PDU_SessionUL_TrafficInfo_r18

func (ie *UL_TrafficInfo_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uLTrafficInfoR18SizeConstraints)
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

func (ie *UL_TrafficInfo_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uLTrafficInfoR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(UL_TrafficInfo_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
