// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FreqPriorityNR (line 1338).

var freqPriorityNRConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreq"},
		{Name: "cellReselectionPriority"},
		{Name: "cellReselectionSubPriority", Optional: true},
	},
}

type FreqPriorityNR struct {
	CarrierFreq                ARFCN_ValueNR
	CellReselectionPriority    CellReselectionPriority
	CellReselectionSubPriority *CellReselectionSubPriority
}

func (ie *FreqPriorityNR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(freqPriorityNRConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CellReselectionSubPriority != nil}); err != nil {
		return err
	}

	// 2. carrierFreq: ref
	{
		if err := ie.CarrierFreq.Encode(e); err != nil {
			return err
		}
	}

	// 3. cellReselectionPriority: ref
	{
		if err := ie.CellReselectionPriority.Encode(e); err != nil {
			return err
		}
	}

	// 4. cellReselectionSubPriority: ref
	{
		if ie.CellReselectionSubPriority != nil {
			if err := ie.CellReselectionSubPriority.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FreqPriorityNR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(freqPriorityNRConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. carrierFreq: ref
	{
		if err := ie.CarrierFreq.Decode(d); err != nil {
			return err
		}
	}

	// 3. cellReselectionPriority: ref
	{
		if err := ie.CellReselectionPriority.Decode(d); err != nil {
			return err
		}
	}

	// 4. cellReselectionSubPriority: ref
	{
		if seq.IsComponentPresent(2) {
			ie.CellReselectionSubPriority = new(CellReselectionSubPriority)
			if err := ie.CellReselectionSubPriority.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
