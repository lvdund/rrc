// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNR-v15g0 (line 17300).

var cAParametersNRV15g0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "simultaneousRxTxInterBandCAPerBandPair", Optional: true},
		{Name: "simultaneousRxTxSULPerBandPair", Optional: true},
	},
}

type CA_ParametersNR_V15g0 struct {
	SimultaneousRxTxInterBandCAPerBandPair *SimultaneousRxTxPerBandPair
	SimultaneousRxTxSULPerBandPair         *SimultaneousRxTxPerBandPair
}

func (ie *CA_ParametersNR_V15g0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRV15g0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SimultaneousRxTxInterBandCAPerBandPair != nil, ie.SimultaneousRxTxSULPerBandPair != nil}); err != nil {
		return err
	}

	// 2. simultaneousRxTxInterBandCAPerBandPair: ref
	{
		if ie.SimultaneousRxTxInterBandCAPerBandPair != nil {
			if err := ie.SimultaneousRxTxInterBandCAPerBandPair.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. simultaneousRxTxSULPerBandPair: ref
	{
		if ie.SimultaneousRxTxSULPerBandPair != nil {
			if err := ie.SimultaneousRxTxSULPerBandPair.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNR_V15g0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRV15g0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. simultaneousRxTxInterBandCAPerBandPair: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SimultaneousRxTxInterBandCAPerBandPair = new(SimultaneousRxTxPerBandPair)
			if err := ie.SimultaneousRxTxInterBandCAPerBandPair.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. simultaneousRxTxSULPerBandPair: ref
	{
		if seq.IsComponentPresent(1) {
			ie.SimultaneousRxTxSULPerBandPair = new(SimultaneousRxTxPerBandPair)
			if err := ie.SimultaneousRxTxSULPerBandPair.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
