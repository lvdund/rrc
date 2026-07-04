// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultFailedCell-r16 (line 3047).

var measResultFailedCellR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cgi-Info"},
		{Name: "measResult-r16"},
	},
}

type MeasResultFailedCell_r16 struct {
	Cgi_Info       CGI_Info_Logging_r16
	MeasResult_r16 struct {
		CellResults_r16    struct{ ResultsSSB_Cell_r16 MeasQuantityResults }
		RsIndexResults_r16 struct{ ResultsSSB_Indexes_r16 ResultsPerSSB_IndexList }
	}
}

func (ie *MeasResultFailedCell_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultFailedCellR16Constraints)
	_ = seq

	// 1. cgi-Info: ref
	{
		if err := ie.Cgi_Info.Encode(e); err != nil {
			return err
		}
	}

	// 2. measResult-r16: sequence
	{
		{
			c := &ie.MeasResult_r16
			{
				c := &c.CellResults_r16
				if err := c.ResultsSSB_Cell_r16.Encode(e); err != nil {
					return err
				}
			}
			{
				c := &c.RsIndexResults_r16
				if err := c.ResultsSSB_Indexes_r16.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *MeasResultFailedCell_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultFailedCellR16Constraints)
	_ = seq

	// 1. cgi-Info: ref
	{
		if err := ie.Cgi_Info.Decode(d); err != nil {
			return err
		}
	}

	// 2. measResult-r16: sequence
	{
		{
			c := &ie.MeasResult_r16
			{
				c := &c.CellResults_r16
				{
					if err := c.ResultsSSB_Cell_r16.Decode(d); err != nil {
						return err
					}
				}
			}
			{
				c := &c.RsIndexResults_r16
				{
					if err := c.ResultsSSB_Indexes_r16.Decode(d); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
