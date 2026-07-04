// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DummyI (line 22526).

var dummyIConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedSRS-TxPortSwitch"},
		{Name: "txSwitchImpactToRx", Optional: true},
	},
}

const (
	DummyI_SupportedSRS_TxPortSwitch_T1r2      = 0
	DummyI_SupportedSRS_TxPortSwitch_T1r4      = 1
	DummyI_SupportedSRS_TxPortSwitch_T2r4      = 2
	DummyI_SupportedSRS_TxPortSwitch_T1r4_T2r4 = 3
	DummyI_SupportedSRS_TxPortSwitch_Tr_Equal  = 4
)

var dummyISupportedSRSTxPortSwitchConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DummyI_SupportedSRS_TxPortSwitch_T1r2, DummyI_SupportedSRS_TxPortSwitch_T1r4, DummyI_SupportedSRS_TxPortSwitch_T2r4, DummyI_SupportedSRS_TxPortSwitch_T1r4_T2r4, DummyI_SupportedSRS_TxPortSwitch_Tr_Equal},
}

const (
	DummyI_TxSwitchImpactToRx_True = 0
)

var dummyITxSwitchImpactToRxConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DummyI_TxSwitchImpactToRx_True},
}

type DummyI struct {
	SupportedSRS_TxPortSwitch int64
	TxSwitchImpactToRx        *int64
}

func (ie *DummyI) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dummyIConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.TxSwitchImpactToRx != nil}); err != nil {
		return err
	}

	// 2. supportedSRS-TxPortSwitch: enumerated
	{
		if err := e.EncodeEnumerated(ie.SupportedSRS_TxPortSwitch, dummyISupportedSRSTxPortSwitchConstraints); err != nil {
			return err
		}
	}

	// 3. txSwitchImpactToRx: enumerated
	{
		if ie.TxSwitchImpactToRx != nil {
			if err := e.EncodeEnumerated(*ie.TxSwitchImpactToRx, dummyITxSwitchImpactToRxConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *DummyI) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dummyIConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedSRS-TxPortSwitch: enumerated
	{
		v0, err := d.DecodeEnumerated(dummyISupportedSRSTxPortSwitchConstraints)
		if err != nil {
			return err
		}
		ie.SupportedSRS_TxPortSwitch = v0
	}

	// 3. txSwitchImpactToRx: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(dummyITxSwitchImpactToRxConstraints)
			if err != nil {
				return err
			}
			ie.TxSwitchImpactToRx = &idx
		}
	}

	return nil
}
