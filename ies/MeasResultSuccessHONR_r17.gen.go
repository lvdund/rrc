// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultSuccessHONR-r17 (line 3450).

var measResultSuccessHONRR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measResult-r17"},
	},
}

var measResultSuccessHONRR17MeasResultR17CellResultsR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resultsSSB-Cell-r17", Optional: true},
		{Name: "resultsCSI-RS-Cell-r17", Optional: true},
	},
}

var measResultSuccessHONRR17MeasResultR17RsIndexResultsR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resultsSSB-Indexes-r17", Optional: true},
		{Name: "resultsCSI-RS-Indexes-r17", Optional: true},
	},
}

type MeasResultSuccessHONR_r17 struct {
	MeasResult_r17 struct {
		CellResults_r17 struct {
			ResultsSSB_Cell_r17    *MeasQuantityResults
			ResultsCSI_RS_Cell_r17 *MeasQuantityResults
		}
		RsIndexResults_r17 struct {
			ResultsSSB_Indexes_r17    *ResultsPerSSB_IndexList
			ResultsCSI_RS_Indexes_r17 *ResultsPerCSI_RS_IndexList
		}
	}
}

func (ie *MeasResultSuccessHONR_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultSuccessHONRR17Constraints)
	_ = seq

	// 1. measResult-r17: sequence
	{
		{
			c := &ie.MeasResult_r17
			{
				c := &c.CellResults_r17
				measResultSuccessHONRR17MeasResultR17CellResultsR17Seq := e.NewSequenceEncoder(measResultSuccessHONRR17MeasResultR17CellResultsR17Constraints)
				if err := measResultSuccessHONRR17MeasResultR17CellResultsR17Seq.EncodePreamble([]bool{c.ResultsSSB_Cell_r17 != nil, c.ResultsCSI_RS_Cell_r17 != nil}); err != nil {
					return err
				}
				if c.ResultsSSB_Cell_r17 != nil {
					if err := c.ResultsSSB_Cell_r17.Encode(e); err != nil {
						return err
					}
				}
				if c.ResultsCSI_RS_Cell_r17 != nil {
					if err := c.ResultsCSI_RS_Cell_r17.Encode(e); err != nil {
						return err
					}
				}
			}
			{
				c := &c.RsIndexResults_r17
				measResultSuccessHONRR17MeasResultR17RsIndexResultsR17Seq := e.NewSequenceEncoder(measResultSuccessHONRR17MeasResultR17RsIndexResultsR17Constraints)
				if err := measResultSuccessHONRR17MeasResultR17RsIndexResultsR17Seq.EncodePreamble([]bool{c.ResultsSSB_Indexes_r17 != nil, c.ResultsCSI_RS_Indexes_r17 != nil}); err != nil {
					return err
				}
				if c.ResultsSSB_Indexes_r17 != nil {
					if err := c.ResultsSSB_Indexes_r17.Encode(e); err != nil {
						return err
					}
				}
				if c.ResultsCSI_RS_Indexes_r17 != nil {
					if err := c.ResultsCSI_RS_Indexes_r17.Encode(e); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (ie *MeasResultSuccessHONR_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultSuccessHONRR17Constraints)
	_ = seq

	// 1. measResult-r17: sequence
	{
		{
			c := &ie.MeasResult_r17
			{
				c := &c.CellResults_r17
				measResultSuccessHONRR17MeasResultR17CellResultsR17Seq := d.NewSequenceDecoder(measResultSuccessHONRR17MeasResultR17CellResultsR17Constraints)
				if err := measResultSuccessHONRR17MeasResultR17CellResultsR17Seq.DecodePreamble(); err != nil {
					return err
				}
				if measResultSuccessHONRR17MeasResultR17CellResultsR17Seq.IsComponentPresent(0) {
					c.ResultsSSB_Cell_r17 = new(MeasQuantityResults)
					if err := (*c.ResultsSSB_Cell_r17).Decode(d); err != nil {
						return err
					}
				}
				if measResultSuccessHONRR17MeasResultR17CellResultsR17Seq.IsComponentPresent(1) {
					c.ResultsCSI_RS_Cell_r17 = new(MeasQuantityResults)
					if err := (*c.ResultsCSI_RS_Cell_r17).Decode(d); err != nil {
						return err
					}
				}
			}
			{
				c := &c.RsIndexResults_r17
				measResultSuccessHONRR17MeasResultR17RsIndexResultsR17Seq := d.NewSequenceDecoder(measResultSuccessHONRR17MeasResultR17RsIndexResultsR17Constraints)
				if err := measResultSuccessHONRR17MeasResultR17RsIndexResultsR17Seq.DecodePreamble(); err != nil {
					return err
				}
				if measResultSuccessHONRR17MeasResultR17RsIndexResultsR17Seq.IsComponentPresent(0) {
					c.ResultsSSB_Indexes_r17 = new(ResultsPerSSB_IndexList)
					if err := (*c.ResultsSSB_Indexes_r17).Decode(d); err != nil {
						return err
					}
				}
				if measResultSuccessHONRR17MeasResultR17RsIndexResultsR17Seq.IsComponentPresent(1) {
					c.ResultsCSI_RS_Indexes_r17 = new(ResultsPerCSI_RS_IndexList)
					if err := (*c.ResultsCSI_RS_Indexes_r17).Decode(d); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
