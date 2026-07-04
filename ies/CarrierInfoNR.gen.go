// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CarrierInfoNR (line 1274).

var carrierInfoNRConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreq"},
		{Name: "ssbSubcarrierSpacing"},
		{Name: "smtc", Optional: true},
	},
}

type CarrierInfoNR struct {
	CarrierFreq          ARFCN_ValueNR
	SsbSubcarrierSpacing SubcarrierSpacing
	Smtc                 *SSB_MTC
}

func (ie *CarrierInfoNR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(carrierInfoNRConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Smtc != nil}); err != nil {
		return err
	}

	// 3. carrierFreq: ref
	{
		if err := ie.CarrierFreq.Encode(e); err != nil {
			return err
		}
	}

	// 4. ssbSubcarrierSpacing: ref
	{
		if err := ie.SsbSubcarrierSpacing.Encode(e); err != nil {
			return err
		}
	}

	// 5. smtc: ref
	{
		if ie.Smtc != nil {
			if err := ie.Smtc.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CarrierInfoNR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(carrierInfoNRConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. carrierFreq: ref
	{
		if err := ie.CarrierFreq.Decode(d); err != nil {
			return err
		}
	}

	// 4. ssbSubcarrierSpacing: ref
	{
		if err := ie.SsbSubcarrierSpacing.Decode(d); err != nil {
			return err
		}
	}

	// 5. smtc: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Smtc = new(SSB_MTC)
			if err := ie.Smtc.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
