// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UplinkTxDirectCurrentBWP (line 16381).

var uplinkTxDirectCurrentBWPConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bwp-Id"},
		{Name: "shift7dot5kHz"},
		{Name: "txDirectCurrentLocation"},
	},
}

var uplinkTxDirectCurrentBWPTxDirectCurrentLocationConstraints = per.Constrained(0, 3301)

type UplinkTxDirectCurrentBWP struct {
	Bwp_Id                  BWP_Id
	Shift7dot5kHz           bool
	TxDirectCurrentLocation int64
}

func (ie *UplinkTxDirectCurrentBWP) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uplinkTxDirectCurrentBWPConstraints)
	_ = seq

	// 1. bwp-Id: ref
	{
		if err := ie.Bwp_Id.Encode(e); err != nil {
			return err
		}
	}

	// 2. shift7dot5kHz: boolean
	{
		if err := e.EncodeBoolean(ie.Shift7dot5kHz); err != nil {
			return err
		}
	}

	// 3. txDirectCurrentLocation: integer
	{
		if err := e.EncodeInteger(ie.TxDirectCurrentLocation, uplinkTxDirectCurrentBWPTxDirectCurrentLocationConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UplinkTxDirectCurrentBWP) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uplinkTxDirectCurrentBWPConstraints)
	_ = seq

	// 1. bwp-Id: ref
	{
		if err := ie.Bwp_Id.Decode(d); err != nil {
			return err
		}
	}

	// 2. shift7dot5kHz: boolean
	{
		v1, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.Shift7dot5kHz = v1
	}

	// 3. txDirectCurrentLocation: integer
	{
		v2, err := d.DecodeInteger(uplinkTxDirectCurrentBWPTxDirectCurrentLocationConstraints)
		if err != nil {
			return err
		}
		ie.TxDirectCurrentLocation = v2
	}

	return nil
}
