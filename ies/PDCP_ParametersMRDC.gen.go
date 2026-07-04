// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCP-ParametersMRDC (line 22763).

var pDCPParametersMRDCConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcp-DuplicationSplitSRB", Optional: true},
		{Name: "pdcp-DuplicationSplitDRB", Optional: true},
	},
}

const (
	PDCP_ParametersMRDC_Pdcp_DuplicationSplitSRB_Supported = 0
)

var pDCPParametersMRDCPdcpDuplicationSplitSRBConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_ParametersMRDC_Pdcp_DuplicationSplitSRB_Supported},
}

const (
	PDCP_ParametersMRDC_Pdcp_DuplicationSplitDRB_Supported = 0
)

var pDCPParametersMRDCPdcpDuplicationSplitDRBConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_ParametersMRDC_Pdcp_DuplicationSplitDRB_Supported},
}

type PDCP_ParametersMRDC struct {
	Pdcp_DuplicationSplitSRB *int64
	Pdcp_DuplicationSplitDRB *int64
}

func (ie *PDCP_ParametersMRDC) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCPParametersMRDCConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pdcp_DuplicationSplitSRB != nil, ie.Pdcp_DuplicationSplitDRB != nil}); err != nil {
		return err
	}

	// 2. pdcp-DuplicationSplitSRB: enumerated
	{
		if ie.Pdcp_DuplicationSplitSRB != nil {
			if err := e.EncodeEnumerated(*ie.Pdcp_DuplicationSplitSRB, pDCPParametersMRDCPdcpDuplicationSplitSRBConstraints); err != nil {
				return err
			}
		}
	}

	// 3. pdcp-DuplicationSplitDRB: enumerated
	{
		if ie.Pdcp_DuplicationSplitDRB != nil {
			if err := e.EncodeEnumerated(*ie.Pdcp_DuplicationSplitDRB, pDCPParametersMRDCPdcpDuplicationSplitDRBConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PDCP_ParametersMRDC) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCPParametersMRDCConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pdcp-DuplicationSplitSRB: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(pDCPParametersMRDCPdcpDuplicationSplitSRBConstraints)
			if err != nil {
				return err
			}
			ie.Pdcp_DuplicationSplitSRB = &idx
		}
	}

	// 3. pdcp-DuplicationSplitDRB: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(pDCPParametersMRDCPdcpDuplicationSplitDRBConstraints)
			if err != nil {
				return err
			}
			ie.Pdcp_DuplicationSplitDRB = &idx
		}
	}

	return nil
}
