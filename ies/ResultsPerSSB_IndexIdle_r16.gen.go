// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ResultsPerSSB-IndexIdle-r16 (line 9972).

var resultsPerSSBIndexIdleR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ssb-Index-r16"},
		{Name: "ssb-Results-r16", Optional: true},
	},
}

var resultsPerSSBIndexIdleR16SsbResultsR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ssb-RSRP-Result-r16", Optional: true},
		{Name: "ssb-RSRQ-Result-r16", Optional: true},
	},
}

type ResultsPerSSB_IndexIdle_r16 struct {
	Ssb_Index_r16   SSB_Index
	Ssb_Results_r16 *struct {
		Ssb_RSRP_Result_r16 *RSRP_Range
		Ssb_RSRQ_Result_r16 *RSRQ_Range
	}
}

func (ie *ResultsPerSSB_IndexIdle_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(resultsPerSSBIndexIdleR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ssb_Results_r16 != nil}); err != nil {
		return err
	}

	// 2. ssb-Index-r16: ref
	{
		if err := ie.Ssb_Index_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. ssb-Results-r16: sequence
	{
		if ie.Ssb_Results_r16 != nil {
			c := ie.Ssb_Results_r16
			resultsPerSSBIndexIdleR16SsbResultsR16Seq := e.NewSequenceEncoder(resultsPerSSBIndexIdleR16SsbResultsR16Constraints)
			if err := resultsPerSSBIndexIdleR16SsbResultsR16Seq.EncodePreamble([]bool{c.Ssb_RSRP_Result_r16 != nil, c.Ssb_RSRQ_Result_r16 != nil}); err != nil {
				return err
			}
			if c.Ssb_RSRP_Result_r16 != nil {
				if err := c.Ssb_RSRP_Result_r16.Encode(e); err != nil {
					return err
				}
			}
			if c.Ssb_RSRQ_Result_r16 != nil {
				if err := c.Ssb_RSRQ_Result_r16.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *ResultsPerSSB_IndexIdle_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(resultsPerSSBIndexIdleR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ssb-Index-r16: ref
	{
		if err := ie.Ssb_Index_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. ssb-Results-r16: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.Ssb_Results_r16 = &struct {
				Ssb_RSRP_Result_r16 *RSRP_Range
				Ssb_RSRQ_Result_r16 *RSRQ_Range
			}{}
			c := ie.Ssb_Results_r16
			resultsPerSSBIndexIdleR16SsbResultsR16Seq := d.NewSequenceDecoder(resultsPerSSBIndexIdleR16SsbResultsR16Constraints)
			if err := resultsPerSSBIndexIdleR16SsbResultsR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if resultsPerSSBIndexIdleR16SsbResultsR16Seq.IsComponentPresent(0) {
				c.Ssb_RSRP_Result_r16 = new(RSRP_Range)
				if err := (*c.Ssb_RSRP_Result_r16).Decode(d); err != nil {
					return err
				}
			}
			if resultsPerSSBIndexIdleR16SsbResultsR16Seq.IsComponentPresent(1) {
				c.Ssb_RSRQ_Result_r16 = new(RSRQ_Range)
				if err := (*c.Ssb_RSRQ_Result_r16).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
