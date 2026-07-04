// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCP-ParametersSidelink-r18 (line 25314).

var pDCPParametersSidelinkR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcp-DuplicationSRB-sidelink-r18", Optional: true},
		{Name: "pdcp-DuplicationDRB-sidelink-r18", Optional: true},
	},
}

const (
	PDCP_ParametersSidelink_r18_Pdcp_DuplicationSRB_Sidelink_r18_Supported = 0
)

var pDCPParametersSidelinkR18PdcpDuplicationSRBSidelinkR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_ParametersSidelink_r18_Pdcp_DuplicationSRB_Sidelink_r18_Supported},
}

const (
	PDCP_ParametersSidelink_r18_Pdcp_DuplicationDRB_Sidelink_r18_Supported = 0
)

var pDCPParametersSidelinkR18PdcpDuplicationDRBSidelinkR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_ParametersSidelink_r18_Pdcp_DuplicationDRB_Sidelink_r18_Supported},
}

type PDCP_ParametersSidelink_r18 struct {
	Pdcp_DuplicationSRB_Sidelink_r18 *int64
	Pdcp_DuplicationDRB_Sidelink_r18 *int64
}

func (ie *PDCP_ParametersSidelink_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCPParametersSidelinkR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pdcp_DuplicationSRB_Sidelink_r18 != nil, ie.Pdcp_DuplicationDRB_Sidelink_r18 != nil}); err != nil {
		return err
	}

	// 3. pdcp-DuplicationSRB-sidelink-r18: enumerated
	{
		if ie.Pdcp_DuplicationSRB_Sidelink_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdcp_DuplicationSRB_Sidelink_r18, pDCPParametersSidelinkR18PdcpDuplicationSRBSidelinkR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. pdcp-DuplicationDRB-sidelink-r18: enumerated
	{
		if ie.Pdcp_DuplicationDRB_Sidelink_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdcp_DuplicationDRB_Sidelink_r18, pDCPParametersSidelinkR18PdcpDuplicationDRBSidelinkR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PDCP_ParametersSidelink_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCPParametersSidelinkR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. pdcp-DuplicationSRB-sidelink-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(pDCPParametersSidelinkR18PdcpDuplicationSRBSidelinkR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdcp_DuplicationSRB_Sidelink_r18 = &idx
		}
	}

	// 4. pdcp-DuplicationDRB-sidelink-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(pDCPParametersSidelinkR18PdcpDuplicationDRBSidelinkR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdcp_DuplicationDRB_Sidelink_r18 = &idx
		}
	}

	return nil
}
