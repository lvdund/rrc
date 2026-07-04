// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DRB-CountInfo (line 249).

var dRBCountInfoConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "drb-Identity"},
		{Name: "count-Uplink"},
		{Name: "count-Downlink"},
	},
}

var dRBCountInfoCountUplinkConstraints = per.Constrained(0, 4294967295)

var dRBCountInfoCountDownlinkConstraints = per.Constrained(0, 4294967295)

type DRB_CountInfo struct {
	Drb_Identity   DRB_Identity
	Count_Uplink   int64
	Count_Downlink int64
}

func (ie *DRB_CountInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRBCountInfoConstraints)
	_ = seq

	// 1. drb-Identity: ref
	{
		if err := ie.Drb_Identity.Encode(e); err != nil {
			return err
		}
	}

	// 2. count-Uplink: integer
	{
		if err := e.EncodeInteger(ie.Count_Uplink, dRBCountInfoCountUplinkConstraints); err != nil {
			return err
		}
	}

	// 3. count-Downlink: integer
	{
		if err := e.EncodeInteger(ie.Count_Downlink, dRBCountInfoCountDownlinkConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DRB_CountInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRBCountInfoConstraints)
	_ = seq

	// 1. drb-Identity: ref
	{
		if err := ie.Drb_Identity.Decode(d); err != nil {
			return err
		}
	}

	// 2. count-Uplink: integer
	{
		v1, err := d.DecodeInteger(dRBCountInfoCountUplinkConstraints)
		if err != nil {
			return err
		}
		ie.Count_Uplink = v1
	}

	// 3. count-Downlink: integer
	{
		v2, err := d.DecodeInteger(dRBCountInfoCountDownlinkConstraints)
		if err != nil {
			return err
		}
		ie.Count_Downlink = v2
	}

	return nil
}
