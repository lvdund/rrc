// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasQuantityResultsEUTRA (line 9815).

var measQuantityResultsEUTRAConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rsrp", Optional: true},
		{Name: "rsrq", Optional: true},
		{Name: "sinr", Optional: true},
	},
}

type MeasQuantityResultsEUTRA struct {
	Rsrp *RSRP_RangeEUTRA
	Rsrq *RSRQ_RangeEUTRA
	Sinr *SINR_RangeEUTRA
}

func (ie *MeasQuantityResultsEUTRA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measQuantityResultsEUTRAConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rsrp != nil, ie.Rsrq != nil, ie.Sinr != nil}); err != nil {
		return err
	}

	// 2. rsrp: ref
	{
		if ie.Rsrp != nil {
			if err := ie.Rsrp.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. rsrq: ref
	{
		if ie.Rsrq != nil {
			if err := ie.Rsrq.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sinr: ref
	{
		if ie.Sinr != nil {
			if err := ie.Sinr.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasQuantityResultsEUTRA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measQuantityResultsEUTRAConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. rsrp: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Rsrp = new(RSRP_RangeEUTRA)
			if err := ie.Rsrp.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. rsrq: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Rsrq = new(RSRQ_RangeEUTRA)
			if err := ie.Rsrq.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sinr: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sinr = new(SINR_RangeEUTRA)
			if err := ie.Sinr.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
