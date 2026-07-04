// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CarrierFreqListEUTRA-v1800 (line 4131).
// CarrierFreqListEUTRA-v1800 ::=      SEQUENCE (SIZE (1..maxEUTRA-Carrier)) OF CarrierFreqEUTRA-v1800

var carrierFreqListEUTRAV1800SizeConstraints = per.SizeRange(1, common.MaxEUTRA_Carrier)

type CarrierFreqListEUTRA_v1800 []CarrierFreqEUTRA_v1800

func (ie *CarrierFreqListEUTRA_v1800) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(carrierFreqListEUTRAV1800SizeConstraints)
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

func (ie *CarrierFreqListEUTRA_v1800) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(carrierFreqListEUTRAV1800SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CarrierFreqListEUTRA_v1800, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
