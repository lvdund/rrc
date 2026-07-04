// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasReportQuantity (line 13865).

var measReportQuantityConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rsrp"},
		{Name: "rsrq"},
		{Name: "sinr"},
	},
}

type MeasReportQuantity struct {
	Rsrp bool
	Rsrq bool
	Sinr bool
}

func (ie *MeasReportQuantity) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measReportQuantityConstraints)
	_ = seq

	// 1. rsrp: boolean
	{
		if err := e.EncodeBoolean(ie.Rsrp); err != nil {
			return err
		}
	}

	// 2. rsrq: boolean
	{
		if err := e.EncodeBoolean(ie.Rsrq); err != nil {
			return err
		}
	}

	// 3. sinr: boolean
	{
		if err := e.EncodeBoolean(ie.Sinr); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasReportQuantity) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measReportQuantityConstraints)
	_ = seq

	// 1. rsrp: boolean
	{
		v0, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.Rsrp = v0
	}

	// 2. rsrq: boolean
	{
		v1, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.Rsrq = v1
	}

	// 3. sinr: boolean
	{
		v2, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.Sinr = v2
	}

	return nil
}
