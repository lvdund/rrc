// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRDC-Parameters-v15g0 (line 22569).

var mRDCParametersV15g0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "simultaneousRxTxInterBandENDCPerBandPair", Optional: true},
	},
}

type MRDC_Parameters_V15g0 struct {
	SimultaneousRxTxInterBandENDCPerBandPair *SimultaneousRxTxPerBandPair
}

func (ie *MRDC_Parameters_V15g0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mRDCParametersV15g0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SimultaneousRxTxInterBandENDCPerBandPair != nil}); err != nil {
		return err
	}

	// 2. simultaneousRxTxInterBandENDCPerBandPair: ref
	{
		if ie.SimultaneousRxTxInterBandENDCPerBandPair != nil {
			if err := ie.SimultaneousRxTxInterBandENDCPerBandPair.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MRDC_Parameters_V15g0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mRDCParametersV15g0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. simultaneousRxTxInterBandENDCPerBandPair: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SimultaneousRxTxInterBandENDCPerBandPair = new(SimultaneousRxTxPerBandPair)
			if err := ie.SimultaneousRxTxInterBandENDCPerBandPair.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
