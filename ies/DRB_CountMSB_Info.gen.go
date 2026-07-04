// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DRB-CountMSB-Info (line 223).

var dRBCountMSBInfoConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "drb-Identity"},
		{Name: "countMSB-Uplink"},
		{Name: "countMSB-Downlink"},
	},
}

var dRBCountMSBInfoCountMSBUplinkConstraints = per.Constrained(0, 33554431)

var dRBCountMSBInfoCountMSBDownlinkConstraints = per.Constrained(0, 33554431)

type DRB_CountMSB_Info struct {
	Drb_Identity      DRB_Identity
	CountMSB_Uplink   int64
	CountMSB_Downlink int64
}

func (ie *DRB_CountMSB_Info) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRBCountMSBInfoConstraints)
	_ = seq

	// 1. drb-Identity: ref
	{
		if err := ie.Drb_Identity.Encode(e); err != nil {
			return err
		}
	}

	// 2. countMSB-Uplink: integer
	{
		if err := e.EncodeInteger(ie.CountMSB_Uplink, dRBCountMSBInfoCountMSBUplinkConstraints); err != nil {
			return err
		}
	}

	// 3. countMSB-Downlink: integer
	{
		if err := e.EncodeInteger(ie.CountMSB_Downlink, dRBCountMSBInfoCountMSBDownlinkConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DRB_CountMSB_Info) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRBCountMSBInfoConstraints)
	_ = seq

	// 1. drb-Identity: ref
	{
		if err := ie.Drb_Identity.Decode(d); err != nil {
			return err
		}
	}

	// 2. countMSB-Uplink: integer
	{
		v1, err := d.DecodeInteger(dRBCountMSBInfoCountMSBUplinkConstraints)
		if err != nil {
			return err
		}
		ie.CountMSB_Uplink = v1
	}

	// 3. countMSB-Downlink: integer
	{
		v2, err := d.DecodeInteger(dRBCountMSBInfoCountMSBDownlinkConstraints)
		if err != nil {
			return err
		}
		ie.CountMSB_Downlink = v2
	}

	return nil
}
