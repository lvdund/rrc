// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-TypeTxSync-r16 (line 28353).
// SL-TypeTxSync-r16 ::=                     ENUMERATED {gnss, gnbEnb, ue}

const (
	SL_TypeTxSync_r16_Gnss   = 0
	SL_TypeTxSync_r16_GnbEnb = 1
	SL_TypeTxSync_r16_Ue     = 2
)

var sLTypeTxSyncR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_TypeTxSync_r16_Gnss, SL_TypeTxSync_r16_GnbEnb, SL_TypeTxSync_r16_Ue},
}

type SL_TypeTxSync_r16 struct {
	Value int64
}

func (v *SL_TypeTxSync_r16) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, sLTypeTxSyncR16Constraints)
}

func (v *SL_TypeTxSync_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(sLTypeTxSyncR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
