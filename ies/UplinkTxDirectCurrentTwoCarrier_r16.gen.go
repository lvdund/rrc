// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UplinkTxDirectCurrentTwoCarrier-r16 (line 16421).

var uplinkTxDirectCurrentTwoCarrierR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierOneInfo-r16"},
		{Name: "carrierTwoInfo-r16"},
		{Name: "singlePA-TxDirectCurrent-r16"},
		{Name: "secondPA-TxDirectCurrent-r16", Optional: true},
	},
}

type UplinkTxDirectCurrentTwoCarrier_r16 struct {
	CarrierOneInfo_r16           UplinkTxDirectCurrentCarrierInfo_r16
	CarrierTwoInfo_r16           UplinkTxDirectCurrentCarrierInfo_r16
	SinglePA_TxDirectCurrent_r16 UplinkTxDirectCurrentTwoCarrierInfo_r16
	SecondPA_TxDirectCurrent_r16 *UplinkTxDirectCurrentTwoCarrierInfo_r16
}

func (ie *UplinkTxDirectCurrentTwoCarrier_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uplinkTxDirectCurrentTwoCarrierR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SecondPA_TxDirectCurrent_r16 != nil}); err != nil {
		return err
	}

	// 2. carrierOneInfo-r16: ref
	{
		if err := ie.CarrierOneInfo_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. carrierTwoInfo-r16: ref
	{
		if err := ie.CarrierTwoInfo_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. singlePA-TxDirectCurrent-r16: ref
	{
		if err := ie.SinglePA_TxDirectCurrent_r16.Encode(e); err != nil {
			return err
		}
	}

	// 5. secondPA-TxDirectCurrent-r16: ref
	{
		if ie.SecondPA_TxDirectCurrent_r16 != nil {
			if err := ie.SecondPA_TxDirectCurrent_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UplinkTxDirectCurrentTwoCarrier_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uplinkTxDirectCurrentTwoCarrierR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. carrierOneInfo-r16: ref
	{
		if err := ie.CarrierOneInfo_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. carrierTwoInfo-r16: ref
	{
		if err := ie.CarrierTwoInfo_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. singlePA-TxDirectCurrent-r16: ref
	{
		if err := ie.SinglePA_TxDirectCurrent_r16.Decode(d); err != nil {
			return err
		}
	}

	// 5. secondPA-TxDirectCurrent-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.SecondPA_TxDirectCurrent_r16 = new(UplinkTxDirectCurrentTwoCarrierInfo_r16)
			if err := ie.SecondPA_TxDirectCurrent_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
