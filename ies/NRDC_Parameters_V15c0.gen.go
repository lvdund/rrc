// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NRDC-Parameters-v15c0 (line 22660).

var nRDCParametersV15c0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcp-DuplicationSplitSRB", Optional: true},
		{Name: "pdcp-DuplicationSplitDRB", Optional: true},
	},
}

const (
	NRDC_Parameters_V15c0_Pdcp_DuplicationSplitSRB_Supported = 0
)

var nRDCParametersV15c0PdcpDuplicationSplitSRBConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NRDC_Parameters_V15c0_Pdcp_DuplicationSplitSRB_Supported},
}

const (
	NRDC_Parameters_V15c0_Pdcp_DuplicationSplitDRB_Supported = 0
)

var nRDCParametersV15c0PdcpDuplicationSplitDRBConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NRDC_Parameters_V15c0_Pdcp_DuplicationSplitDRB_Supported},
}

type NRDC_Parameters_V15c0 struct {
	Pdcp_DuplicationSplitSRB *int64
	Pdcp_DuplicationSplitDRB *int64
}

func (ie *NRDC_Parameters_V15c0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRDCParametersV15c0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pdcp_DuplicationSplitSRB != nil, ie.Pdcp_DuplicationSplitDRB != nil}); err != nil {
		return err
	}

	// 2. pdcp-DuplicationSplitSRB: enumerated
	{
		if ie.Pdcp_DuplicationSplitSRB != nil {
			if err := e.EncodeEnumerated(*ie.Pdcp_DuplicationSplitSRB, nRDCParametersV15c0PdcpDuplicationSplitSRBConstraints); err != nil {
				return err
			}
		}
	}

	// 3. pdcp-DuplicationSplitDRB: enumerated
	{
		if ie.Pdcp_DuplicationSplitDRB != nil {
			if err := e.EncodeEnumerated(*ie.Pdcp_DuplicationSplitDRB, nRDCParametersV15c0PdcpDuplicationSplitDRBConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NRDC_Parameters_V15c0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRDCParametersV15c0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pdcp-DuplicationSplitSRB: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(nRDCParametersV15c0PdcpDuplicationSplitSRBConstraints)
			if err != nil {
				return err
			}
			ie.Pdcp_DuplicationSplitSRB = &idx
		}
	}

	// 3. pdcp-DuplicationSplitDRB: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(nRDCParametersV15c0PdcpDuplicationSplitDRBConstraints)
			if err != nil {
				return err
			}
			ie.Pdcp_DuplicationSplitDRB = &idx
		}
	}

	return nil
}
