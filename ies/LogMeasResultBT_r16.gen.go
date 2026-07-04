// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LogMeasResultBT-r16 (line 26252).

var logMeasResultBTR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "bt-Addr-r16"},
		{Name: "rssi-BT-r16", Optional: true},
	},
}

var logMeasResultBTR16BtAddrR16Constraints = per.FixedSize(48)

var logMeasResultBTR16RssiBTR16Constraints = per.Constrained(-128, 127)

type LogMeasResultBT_r16 struct {
	Bt_Addr_r16 per.BitString
	Rssi_BT_r16 *int64
}

func (ie *LogMeasResultBT_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(logMeasResultBTR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rssi_BT_r16 != nil}); err != nil {
		return err
	}

	// 3. bt-Addr-r16: bit-string
	{
		if err := e.EncodeBitString(ie.Bt_Addr_r16, logMeasResultBTR16BtAddrR16Constraints); err != nil {
			return err
		}
	}

	// 4. rssi-BT-r16: integer
	{
		if ie.Rssi_BT_r16 != nil {
			if err := e.EncodeInteger(*ie.Rssi_BT_r16, logMeasResultBTR16RssiBTR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LogMeasResultBT_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(logMeasResultBTR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. bt-Addr-r16: bit-string
	{
		v0, err := d.DecodeBitString(logMeasResultBTR16BtAddrR16Constraints)
		if err != nil {
			return err
		}
		ie.Bt_Addr_r16 = v0
	}

	// 4. rssi-BT-r16: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(logMeasResultBTR16RssiBTR16Constraints)
			if err != nil {
				return err
			}
			ie.Rssi_BT_r16 = &v
		}
	}

	return nil
}
