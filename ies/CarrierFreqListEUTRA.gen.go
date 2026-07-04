// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CarrierFreqListEUTRA (line 4125).
// CarrierFreqListEUTRA ::=            SEQUENCE (SIZE (1..maxEUTRA-Carrier)) OF CarrierFreqEUTRA

var carrierFreqListEUTRASizeConstraints = per.SizeRange(1, common.MaxEUTRA_Carrier)

type CarrierFreqListEUTRA []CarrierFreqEUTRA

func (ie *CarrierFreqListEUTRA) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(carrierFreqListEUTRASizeConstraints)
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

func (ie *CarrierFreqListEUTRA) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(carrierFreqListEUTRASizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CarrierFreqListEUTRA, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
