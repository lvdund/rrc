// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultRLFNR-r16 (line 3435).

var measResultRLFNRR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measResult-r16"},
	},
}

var measResultRLFNRR16MeasResultR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cellResults-r16"},
		{Name: "rsIndexResults-r16", Optional: true},
	},
}

var measResultRLFNRR16MeasResultR16CellResultsR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resultsSSB-Cell-r16", Optional: true},
		{Name: "resultsCSI-RS-Cell-r16", Optional: true},
	},
}

var measResultRLFNRR16MeasResultR16RsIndexResultsR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resultsSSB-Indexes-r16", Optional: true},
		{Name: "ssbRLMConfigBitmap-r16", Optional: true},
		{Name: "resultsCSI-RS-Indexes-r16", Optional: true},
		{Name: "csi-rsRLMConfigBitmap-r16", Optional: true},
	},
}

type MeasResultRLFNR_r16 struct {
	MeasResult_r16 struct {
		CellResults_r16 struct {
			ResultsSSB_Cell_r16    *MeasQuantityResults
			ResultsCSI_RS_Cell_r16 *MeasQuantityResults
		}
		RsIndexResults_r16 *struct {
			ResultsSSB_Indexes_r16    *ResultsPerSSB_IndexList
			SsbRLMConfigBitmap_r16    *per.BitString
			ResultsCSI_RS_Indexes_r16 *ResultsPerCSI_RS_IndexList
			Csi_RsRLMConfigBitmap_r16 *per.BitString
		}
	}
}

func (ie *MeasResultRLFNR_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultRLFNRR16Constraints)
	_ = seq

	// 1. measResult-r16: sequence
	{
		{
			c := &ie.MeasResult_r16
			measResultRLFNRR16MeasResultR16Seq := e.NewSequenceEncoder(measResultRLFNRR16MeasResultR16Constraints)
			if err := measResultRLFNRR16MeasResultR16Seq.EncodePreamble([]bool{c.RsIndexResults_r16 != nil}); err != nil {
				return err
			}
			{
				c := &c.CellResults_r16
				measResultRLFNRR16MeasResultR16CellResultsR16Seq := e.NewSequenceEncoder(measResultRLFNRR16MeasResultR16CellResultsR16Constraints)
				if err := measResultRLFNRR16MeasResultR16CellResultsR16Seq.EncodePreamble([]bool{c.ResultsSSB_Cell_r16 != nil, c.ResultsCSI_RS_Cell_r16 != nil}); err != nil {
					return err
				}
				if c.ResultsSSB_Cell_r16 != nil {
					if err := c.ResultsSSB_Cell_r16.Encode(e); err != nil {
						return err
					}
				}
				if c.ResultsCSI_RS_Cell_r16 != nil {
					if err := c.ResultsCSI_RS_Cell_r16.Encode(e); err != nil {
						return err
					}
				}
			}
			if c.RsIndexResults_r16 != nil {
				c := (*c.RsIndexResults_r16)
				measResultRLFNRR16MeasResultR16RsIndexResultsR16Seq := e.NewSequenceEncoder(measResultRLFNRR16MeasResultR16RsIndexResultsR16Constraints)
				if err := measResultRLFNRR16MeasResultR16RsIndexResultsR16Seq.EncodePreamble([]bool{c.ResultsSSB_Indexes_r16 != nil, c.SsbRLMConfigBitmap_r16 != nil, c.ResultsCSI_RS_Indexes_r16 != nil, c.Csi_RsRLMConfigBitmap_r16 != nil}); err != nil {
					return err
				}
				if c.ResultsSSB_Indexes_r16 != nil {
					if err := c.ResultsSSB_Indexes_r16.Encode(e); err != nil {
						return err
					}
				}
				if c.SsbRLMConfigBitmap_r16 != nil {
					if err := e.EncodeBitString((*c.SsbRLMConfigBitmap_r16), per.FixedSize(64)); err != nil {
						return err
					}
				}
				if c.ResultsCSI_RS_Indexes_r16 != nil {
					if err := c.ResultsCSI_RS_Indexes_r16.Encode(e); err != nil {
						return err
					}
				}
				if c.Csi_RsRLMConfigBitmap_r16 != nil {
					if err := e.EncodeBitString((*c.Csi_RsRLMConfigBitmap_r16), per.FixedSize(96)); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (ie *MeasResultRLFNR_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultRLFNRR16Constraints)
	_ = seq

	// 1. measResult-r16: sequence
	{
		{
			c := &ie.MeasResult_r16
			measResultRLFNRR16MeasResultR16Seq := d.NewSequenceDecoder(measResultRLFNRR16MeasResultR16Constraints)
			if err := measResultRLFNRR16MeasResultR16Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				c := &c.CellResults_r16
				measResultRLFNRR16MeasResultR16CellResultsR16Seq := d.NewSequenceDecoder(measResultRLFNRR16MeasResultR16CellResultsR16Constraints)
				if err := measResultRLFNRR16MeasResultR16CellResultsR16Seq.DecodePreamble(); err != nil {
					return err
				}
				if measResultRLFNRR16MeasResultR16CellResultsR16Seq.IsComponentPresent(0) {
					c.ResultsSSB_Cell_r16 = new(MeasQuantityResults)
					if err := (*c.ResultsSSB_Cell_r16).Decode(d); err != nil {
						return err
					}
				}
				if measResultRLFNRR16MeasResultR16CellResultsR16Seq.IsComponentPresent(1) {
					c.ResultsCSI_RS_Cell_r16 = new(MeasQuantityResults)
					if err := (*c.ResultsCSI_RS_Cell_r16).Decode(d); err != nil {
						return err
					}
				}
			}
			if measResultRLFNRR16MeasResultR16Seq.IsComponentPresent(1) {
				c.RsIndexResults_r16 = &struct {
					ResultsSSB_Indexes_r16    *ResultsPerSSB_IndexList
					SsbRLMConfigBitmap_r16    *per.BitString
					ResultsCSI_RS_Indexes_r16 *ResultsPerCSI_RS_IndexList
					Csi_RsRLMConfigBitmap_r16 *per.BitString
				}{}
				c := (*c.RsIndexResults_r16)
				measResultRLFNRR16MeasResultR16RsIndexResultsR16Seq := d.NewSequenceDecoder(measResultRLFNRR16MeasResultR16RsIndexResultsR16Constraints)
				if err := measResultRLFNRR16MeasResultR16RsIndexResultsR16Seq.DecodePreamble(); err != nil {
					return err
				}
				if measResultRLFNRR16MeasResultR16RsIndexResultsR16Seq.IsComponentPresent(0) {
					c.ResultsSSB_Indexes_r16 = new(ResultsPerSSB_IndexList)
					if err := (*c.ResultsSSB_Indexes_r16).Decode(d); err != nil {
						return err
					}
				}
				if measResultRLFNRR16MeasResultR16RsIndexResultsR16Seq.IsComponentPresent(1) {
					c.SsbRLMConfigBitmap_r16 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(64))
					if err != nil {
						return err
					}
					(*c.SsbRLMConfigBitmap_r16) = v
				}
				if measResultRLFNRR16MeasResultR16RsIndexResultsR16Seq.IsComponentPresent(2) {
					c.ResultsCSI_RS_Indexes_r16 = new(ResultsPerCSI_RS_IndexList)
					if err := (*c.ResultsCSI_RS_Indexes_r16).Decode(d); err != nil {
						return err
					}
				}
				if measResultRLFNRR16MeasResultR16RsIndexResultsR16Seq.IsComponentPresent(3) {
					c.Csi_RsRLMConfigBitmap_r16 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(96))
					if err != nil {
						return err
					}
					(*c.Csi_RsRLMConfigBitmap_r16) = v
				}
			}
		}
	}

	return nil
}
