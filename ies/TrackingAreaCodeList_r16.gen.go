// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TrackingAreaCodeList-r16 (line 26077).
// TrackingAreaCodeList-r16 ::=     SEQUENCE (SIZE (1..8)) OF TrackingAreaCode

var trackingAreaCodeListR16SizeConstraints = per.SizeRange(1, 8)

type TrackingAreaCodeList_r16 []TrackingAreaCode

func (ie *TrackingAreaCodeList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(trackingAreaCodeListR16SizeConstraints)
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

func (ie *TrackingAreaCodeList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(trackingAreaCodeListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(TrackingAreaCodeList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
