// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RedirectedCarrierInfo-EUTRA (line 1269).

var redirectedCarrierInfoEUTRAConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "eutraFrequency"},
		{Name: "cnType", Optional: true},
	},
}

const (
	RedirectedCarrierInfo_EUTRA_CnType_Epc    = 0
	RedirectedCarrierInfo_EUTRA_CnType_FiveGC = 1
)

var redirectedCarrierInfoEUTRACnTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RedirectedCarrierInfo_EUTRA_CnType_Epc, RedirectedCarrierInfo_EUTRA_CnType_FiveGC},
}

type RedirectedCarrierInfo_EUTRA struct {
	EutraFrequency ARFCN_ValueEUTRA
	CnType         *int64
}

func (ie *RedirectedCarrierInfo_EUTRA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(redirectedCarrierInfoEUTRAConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CnType != nil}); err != nil {
		return err
	}

	// 2. eutraFrequency: ref
	{
		if err := ie.EutraFrequency.Encode(e); err != nil {
			return err
		}
	}

	// 3. cnType: enumerated
	{
		if ie.CnType != nil {
			if err := e.EncodeEnumerated(*ie.CnType, redirectedCarrierInfoEUTRACnTypeConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RedirectedCarrierInfo_EUTRA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(redirectedCarrierInfoEUTRAConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. eutraFrequency: ref
	{
		if err := ie.EutraFrequency.Decode(d); err != nil {
			return err
		}
	}

	// 3. cnType: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(redirectedCarrierInfoEUTRACnTypeConstraints)
			if err != nil {
				return err
			}
			ie.CnType = &idx
		}
	}

	return nil
}
