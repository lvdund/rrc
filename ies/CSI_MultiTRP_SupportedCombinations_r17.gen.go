// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CSI-MultiTRP-SupportedCombinations-r17 (line 22531).

var cSIMultiTRPSupportedCombinationsR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumTx-Ports-r17"},
		{Name: "maxTotalNumCMR-r17"},
		{Name: "maxTotalNumTx-PortsNZP-CSI-RS-r17"},
	},
}

const (
	CSI_MultiTRP_SupportedCombinations_r17_MaxNumTx_Ports_r17_N2  = 0
	CSI_MultiTRP_SupportedCombinations_r17_MaxNumTx_Ports_r17_N4  = 1
	CSI_MultiTRP_SupportedCombinations_r17_MaxNumTx_Ports_r17_N8  = 2
	CSI_MultiTRP_SupportedCombinations_r17_MaxNumTx_Ports_r17_N12 = 3
	CSI_MultiTRP_SupportedCombinations_r17_MaxNumTx_Ports_r17_N16 = 4
	CSI_MultiTRP_SupportedCombinations_r17_MaxNumTx_Ports_r17_N24 = 5
	CSI_MultiTRP_SupportedCombinations_r17_MaxNumTx_Ports_r17_N32 = 6
)

var cSIMultiTRPSupportedCombinationsR17MaxNumTxPortsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_MultiTRP_SupportedCombinations_r17_MaxNumTx_Ports_r17_N2, CSI_MultiTRP_SupportedCombinations_r17_MaxNumTx_Ports_r17_N4, CSI_MultiTRP_SupportedCombinations_r17_MaxNumTx_Ports_r17_N8, CSI_MultiTRP_SupportedCombinations_r17_MaxNumTx_Ports_r17_N12, CSI_MultiTRP_SupportedCombinations_r17_MaxNumTx_Ports_r17_N16, CSI_MultiTRP_SupportedCombinations_r17_MaxNumTx_Ports_r17_N24, CSI_MultiTRP_SupportedCombinations_r17_MaxNumTx_Ports_r17_N32},
}

var cSIMultiTRPSupportedCombinationsR17MaxTotalNumCMRR17Constraints = per.Constrained(2, 64)

var cSIMultiTRPSupportedCombinationsR17MaxTotalNumTxPortsNZPCSIRSR17Constraints = per.Constrained(2, 256)

type CSI_MultiTRP_SupportedCombinations_r17 struct {
	MaxNumTx_Ports_r17                int64
	MaxTotalNumCMR_r17                int64
	MaxTotalNumTx_PortsNZP_CSI_RS_r17 int64
}

func (ie *CSI_MultiTRP_SupportedCombinations_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIMultiTRPSupportedCombinationsR17Constraints)
	_ = seq

	// 1. maxNumTx-Ports-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumTx_Ports_r17, cSIMultiTRPSupportedCombinationsR17MaxNumTxPortsR17Constraints); err != nil {
			return err
		}
	}

	// 2. maxTotalNumCMR-r17: integer
	{
		if err := e.EncodeInteger(ie.MaxTotalNumCMR_r17, cSIMultiTRPSupportedCombinationsR17MaxTotalNumCMRR17Constraints); err != nil {
			return err
		}
	}

	// 3. maxTotalNumTx-PortsNZP-CSI-RS-r17: integer
	{
		if err := e.EncodeInteger(ie.MaxTotalNumTx_PortsNZP_CSI_RS_r17, cSIMultiTRPSupportedCombinationsR17MaxTotalNumTxPortsNZPCSIRSR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CSI_MultiTRP_SupportedCombinations_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIMultiTRPSupportedCombinationsR17Constraints)
	_ = seq

	// 1. maxNumTx-Ports-r17: enumerated
	{
		v0, err := d.DecodeEnumerated(cSIMultiTRPSupportedCombinationsR17MaxNumTxPortsR17Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumTx_Ports_r17 = v0
	}

	// 2. maxTotalNumCMR-r17: integer
	{
		v1, err := d.DecodeInteger(cSIMultiTRPSupportedCombinationsR17MaxTotalNumCMRR17Constraints)
		if err != nil {
			return err
		}
		ie.MaxTotalNumCMR_r17 = v1
	}

	// 3. maxTotalNumTx-PortsNZP-CSI-RS-r17: integer
	{
		v2, err := d.DecodeInteger(cSIMultiTRPSupportedCombinationsR17MaxTotalNumTxPortsNZPCSIRSR17Constraints)
		if err != nil {
			return err
		}
		ie.MaxTotalNumTx_PortsNZP_CSI_RS_r17 = v2
	}

	return nil
}
