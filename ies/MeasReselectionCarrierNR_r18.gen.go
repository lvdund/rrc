// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasReselectionCarrierNR-r18 (line 9306).

var measReselectionCarrierNRR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreq-r18"},
	},
}

type MeasReselectionCarrierNR_r18 struct {
	CarrierFreq_r18 ARFCN_ValueNR
}

func (ie *MeasReselectionCarrierNR_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measReselectionCarrierNRR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. carrierFreq-r18: ref
	{
		if err := ie.CarrierFreq_r18.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasReselectionCarrierNR_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measReselectionCarrierNRR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. carrierFreq-r18: ref
	{
		if err := ie.CarrierFreq_r18.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
