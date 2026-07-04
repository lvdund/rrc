// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SSB-MTC-SSBAdapt-r19 (line 9492).
// SSB-MTC-SSBAdapt-r19 ::=            SEQUENCE (SIZE(1..2)) OF SSB-MTC

var sSBMTCSSBAdaptR19SizeConstraints = per.SizeRange(1, 2)

type SSB_MTC_SSBAdapt_r19 []SSB_MTC

func (ie *SSB_MTC_SSBAdapt_r19) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sSBMTCSSBAdaptR19SizeConstraints)
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

func (ie *SSB_MTC_SSBAdapt_r19) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sSBMTCSSBAdaptR19SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SSB_MTC_SSBAdapt_r19, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
