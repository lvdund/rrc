// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-RxDRX-Report-v1700 (line 2189).

var sLRxDRXReportV1700Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-DRX-ConfigFromTx-r17"},
	},
}

type SL_RxDRX_Report_v1700 struct {
	Sl_DRX_ConfigFromTx_r17 SL_DRX_ConfigUC_SemiStatic_r17
}

func (ie *SL_RxDRX_Report_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLRxDRXReportV1700Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. sl-DRX-ConfigFromTx-r17: ref
	{
		if err := ie.Sl_DRX_ConfigFromTx_r17.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_RxDRX_Report_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLRxDRXReportV1700Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. sl-DRX-ConfigFromTx-r17: ref
	{
		if err := ie.Sl_DRX_ConfigFromTx_r17.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
