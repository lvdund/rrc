// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultCBR-NR-r16 (line 10020).

var measResultCBRNRR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-poolReportIdentity-r16"},
		{Name: "sl-CBR-ResultsNR-r16"},
	},
}

type MeasResultCBR_NR_r16 struct {
	Sl_PoolReportIdentity_r16 SL_ResourcePoolID_r16
	Sl_CBR_ResultsNR_r16      SL_CBR_r16
}

func (ie *MeasResultCBR_NR_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultCBRNRR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. sl-poolReportIdentity-r16: ref
	{
		if err := ie.Sl_PoolReportIdentity_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. sl-CBR-ResultsNR-r16: ref
	{
		if err := ie.Sl_CBR_ResultsNR_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasResultCBR_NR_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultCBRNRR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. sl-poolReportIdentity-r16: ref
	{
		if err := ie.Sl_PoolReportIdentity_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. sl-CBR-ResultsNR-r16: ref
	{
		if err := ie.Sl_CBR_ResultsNR_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
