// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UplinkTxDirectCurrentTwoCarrierInfo-r16 (line 16436).

var uplinkTxDirectCurrentTwoCarrierInfoR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "referenceCarrierIndex-r16"},
		{Name: "shift7dot5kHz-r16"},
		{Name: "txDirectCurrentLocation-r16"},
	},
}

var uplinkTxDirectCurrentTwoCarrierInfoR16TxDirectCurrentLocationR16Constraints = per.Constrained(0, 3301)

type UplinkTxDirectCurrentTwoCarrierInfo_r16 struct {
	ReferenceCarrierIndex_r16   ServCellIndex
	Shift7dot5kHz_r16           bool
	TxDirectCurrentLocation_r16 int64
}

func (ie *UplinkTxDirectCurrentTwoCarrierInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uplinkTxDirectCurrentTwoCarrierInfoR16Constraints)
	_ = seq

	// 1. referenceCarrierIndex-r16: ref
	{
		if err := ie.ReferenceCarrierIndex_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. shift7dot5kHz-r16: boolean
	{
		if err := e.EncodeBoolean(ie.Shift7dot5kHz_r16); err != nil {
			return err
		}
	}

	// 3. txDirectCurrentLocation-r16: integer
	{
		if err := e.EncodeInteger(ie.TxDirectCurrentLocation_r16, uplinkTxDirectCurrentTwoCarrierInfoR16TxDirectCurrentLocationR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UplinkTxDirectCurrentTwoCarrierInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uplinkTxDirectCurrentTwoCarrierInfoR16Constraints)
	_ = seq

	// 1. referenceCarrierIndex-r16: ref
	{
		if err := ie.ReferenceCarrierIndex_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. shift7dot5kHz-r16: boolean
	{
		v1, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.Shift7dot5kHz_r16 = v1
	}

	// 3. txDirectCurrentLocation-r16: integer
	{
		v2, err := d.DecodeInteger(uplinkTxDirectCurrentTwoCarrierInfoR16TxDirectCurrentLocationR16Constraints)
		if err != nil {
			return err
		}
		ie.TxDirectCurrentLocation_r16 = v2
	}

	return nil
}
