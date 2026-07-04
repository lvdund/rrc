// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SSB-ToMeasureAltitudeBasedList-r18 (line 9596).
// SSB-ToMeasureAltitudeBasedList-r18 ::= SEQUENCE (SIZE (1..maxNrofAltitudeRanges-r18)) OF SSB-ToMeasureAltitudeBased-r18

var sSBToMeasureAltitudeBasedListR18SizeConstraints = per.SizeRange(1, common.MaxNrofAltitudeRanges_r18)

type SSB_ToMeasureAltitudeBasedList_r18 []SSB_ToMeasureAltitudeBased_r18

func (ie *SSB_ToMeasureAltitudeBasedList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sSBToMeasureAltitudeBasedListR18SizeConstraints)
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

func (ie *SSB_ToMeasureAltitudeBasedList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sSBToMeasureAltitudeBasedListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SSB_ToMeasureAltitudeBasedList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
