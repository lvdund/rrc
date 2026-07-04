// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ResultsPerSSB-Index (line 9823).

var resultsPerSSBIndexConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ssb-Index"},
		{Name: "ssb-Results", Optional: true},
	},
}

type ResultsPerSSB_Index struct {
	Ssb_Index   SSB_Index
	Ssb_Results *MeasQuantityResults
}

func (ie *ResultsPerSSB_Index) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(resultsPerSSBIndexConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ssb_Results != nil}); err != nil {
		return err
	}

	// 2. ssb-Index: ref
	{
		if err := ie.Ssb_Index.Encode(e); err != nil {
			return err
		}
	}

	// 3. ssb-Results: ref
	{
		if ie.Ssb_Results != nil {
			if err := ie.Ssb_Results.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ResultsPerSSB_Index) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(resultsPerSSBIndexConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ssb-Index: ref
	{
		if err := ie.Ssb_Index.Decode(d); err != nil {
			return err
		}
	}

	// 3. ssb-Results: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ssb_Results = new(MeasQuantityResults)
			if err := ie.Ssb_Results.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
