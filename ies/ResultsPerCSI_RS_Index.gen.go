// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ResultsPerCSI-RS-Index (line 9830).

var resultsPerCSIRSIndexConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "csi-RS-Index"},
		{Name: "csi-RS-Results", Optional: true},
	},
}

type ResultsPerCSI_RS_Index struct {
	Csi_RS_Index   CSI_RS_Index
	Csi_RS_Results *MeasQuantityResults
}

func (ie *ResultsPerCSI_RS_Index) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(resultsPerCSIRSIndexConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Csi_RS_Results != nil}); err != nil {
		return err
	}

	// 2. csi-RS-Index: ref
	{
		if err := ie.Csi_RS_Index.Encode(e); err != nil {
			return err
		}
	}

	// 3. csi-RS-Results: ref
	{
		if ie.Csi_RS_Results != nil {
			if err := ie.Csi_RS_Results.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ResultsPerCSI_RS_Index) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(resultsPerCSIRSIndexConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. csi-RS-Index: ref
	{
		if err := ie.Csi_RS_Index.Decode(d); err != nil {
			return err
		}
	}

	// 3. csi-RS-Results: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Csi_RS_Results = new(MeasQuantityResults)
			if err := ie.Csi_RS_Results.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
