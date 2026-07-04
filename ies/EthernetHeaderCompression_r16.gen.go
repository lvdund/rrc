// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EthernetHeaderCompression-r16 (line 11260).

var ethernetHeaderCompressionR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ehc-Common-r16"},
		{Name: "ehc-Downlink-r16", Optional: true},
		{Name: "ehc-Uplink-r16", Optional: true},
	},
}

var ethernetHeaderCompressionR16EhcCommonR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ehc-CID-Length-r16"},
	},
}

const (
	EthernetHeaderCompression_r16_Ehc_Common_r16_Ehc_CID_Length_r16_Bits7  = 0
	EthernetHeaderCompression_r16_Ehc_Common_r16_Ehc_CID_Length_r16_Bits15 = 1
)

var ethernetHeaderCompressionR16EhcCommonR16EhcCIDLengthR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EthernetHeaderCompression_r16_Ehc_Common_r16_Ehc_CID_Length_r16_Bits7, EthernetHeaderCompression_r16_Ehc_Common_r16_Ehc_CID_Length_r16_Bits15},
}

var ethernetHeaderCompressionR16EhcDownlinkR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "drb-ContinueEHC-DL-r16", Optional: true},
	},
}

const (
	EthernetHeaderCompression_r16_Ehc_Downlink_r16_Drb_ContinueEHC_DL_r16_True = 0
)

var ethernetHeaderCompressionR16EhcDownlinkR16DrbContinueEHCDLR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EthernetHeaderCompression_r16_Ehc_Downlink_r16_Drb_ContinueEHC_DL_r16_True},
}

var ethernetHeaderCompressionR16EhcUplinkR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "maxCID-EHC-UL-r16"},
		{Name: "drb-ContinueEHC-UL-r16", Optional: true},
	},
}

const (
	EthernetHeaderCompression_r16_Ehc_Uplink_r16_Drb_ContinueEHC_UL_r16_True = 0
)

var ethernetHeaderCompressionR16EhcUplinkR16DrbContinueEHCULR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EthernetHeaderCompression_r16_Ehc_Uplink_r16_Drb_ContinueEHC_UL_r16_True},
}

type EthernetHeaderCompression_r16 struct {
	Ehc_Common_r16   struct{ Ehc_CID_Length_r16 int64 }
	Ehc_Downlink_r16 *struct{ Drb_ContinueEHC_DL_r16 *int64 }
	Ehc_Uplink_r16   *struct {
		MaxCID_EHC_UL_r16      int64
		Drb_ContinueEHC_UL_r16 *int64
	}
}

func (ie *EthernetHeaderCompression_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(ethernetHeaderCompressionR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ehc_Downlink_r16 != nil, ie.Ehc_Uplink_r16 != nil}); err != nil {
		return err
	}

	// 2. ehc-Common-r16: sequence
	{
		{
			c := &ie.Ehc_Common_r16
			ethernetHeaderCompressionR16EhcCommonR16Seq := e.NewSequenceEncoder(ethernetHeaderCompressionR16EhcCommonR16Constraints)
			if err := ethernetHeaderCompressionR16EhcCommonR16Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Ehc_CID_Length_r16, ethernetHeaderCompressionR16EhcCommonR16EhcCIDLengthR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. ehc-Downlink-r16: sequence
	{
		if ie.Ehc_Downlink_r16 != nil {
			c := ie.Ehc_Downlink_r16
			ethernetHeaderCompressionR16EhcDownlinkR16Seq := e.NewSequenceEncoder(ethernetHeaderCompressionR16EhcDownlinkR16Constraints)
			if err := ethernetHeaderCompressionR16EhcDownlinkR16Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := ethernetHeaderCompressionR16EhcDownlinkR16Seq.EncodePreamble([]bool{c.Drb_ContinueEHC_DL_r16 != nil}); err != nil {
				return err
			}
			if c.Drb_ContinueEHC_DL_r16 != nil {
				if err := e.EncodeEnumerated((*c.Drb_ContinueEHC_DL_r16), ethernetHeaderCompressionR16EhcDownlinkR16DrbContinueEHCDLR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 4. ehc-Uplink-r16: sequence
	{
		if ie.Ehc_Uplink_r16 != nil {
			c := ie.Ehc_Uplink_r16
			ethernetHeaderCompressionR16EhcUplinkR16Seq := e.NewSequenceEncoder(ethernetHeaderCompressionR16EhcUplinkR16Constraints)
			if err := ethernetHeaderCompressionR16EhcUplinkR16Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := ethernetHeaderCompressionR16EhcUplinkR16Seq.EncodePreamble([]bool{c.Drb_ContinueEHC_UL_r16 != nil}); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxCID_EHC_UL_r16, per.Constrained(1, 32767)); err != nil {
				return err
			}
			if c.Drb_ContinueEHC_UL_r16 != nil {
				if err := e.EncodeEnumerated((*c.Drb_ContinueEHC_UL_r16), ethernetHeaderCompressionR16EhcUplinkR16DrbContinueEHCULR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *EthernetHeaderCompression_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(ethernetHeaderCompressionR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ehc-Common-r16: sequence
	{
		{
			c := &ie.Ehc_Common_r16
			ethernetHeaderCompressionR16EhcCommonR16Seq := d.NewSequenceDecoder(ethernetHeaderCompressionR16EhcCommonR16Constraints)
			if err := ethernetHeaderCompressionR16EhcCommonR16Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			{
				v, err := d.DecodeEnumerated(ethernetHeaderCompressionR16EhcCommonR16EhcCIDLengthR16Constraints)
				if err != nil {
					return err
				}
				c.Ehc_CID_Length_r16 = v
			}
		}
	}

	// 3. ehc-Downlink-r16: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.Ehc_Downlink_r16 = &struct{ Drb_ContinueEHC_DL_r16 *int64 }{}
			c := ie.Ehc_Downlink_r16
			ethernetHeaderCompressionR16EhcDownlinkR16Seq := d.NewSequenceDecoder(ethernetHeaderCompressionR16EhcDownlinkR16Constraints)
			if err := ethernetHeaderCompressionR16EhcDownlinkR16Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := ethernetHeaderCompressionR16EhcDownlinkR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if ethernetHeaderCompressionR16EhcDownlinkR16Seq.IsComponentPresent(0) {
				c.Drb_ContinueEHC_DL_r16 = new(int64)
				v, err := d.DecodeEnumerated(ethernetHeaderCompressionR16EhcDownlinkR16DrbContinueEHCDLR16Constraints)
				if err != nil {
					return err
				}
				(*c.Drb_ContinueEHC_DL_r16) = v
			}
		}
	}

	// 4. ehc-Uplink-r16: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.Ehc_Uplink_r16 = &struct {
				MaxCID_EHC_UL_r16      int64
				Drb_ContinueEHC_UL_r16 *int64
			}{}
			c := ie.Ehc_Uplink_r16
			ethernetHeaderCompressionR16EhcUplinkR16Seq := d.NewSequenceDecoder(ethernetHeaderCompressionR16EhcUplinkR16Constraints)
			if err := ethernetHeaderCompressionR16EhcUplinkR16Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := ethernetHeaderCompressionR16EhcUplinkR16Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 32767))
				if err != nil {
					return err
				}
				c.MaxCID_EHC_UL_r16 = v
			}
			if ethernetHeaderCompressionR16EhcUplinkR16Seq.IsComponentPresent(1) {
				c.Drb_ContinueEHC_UL_r16 = new(int64)
				v, err := d.DecodeEnumerated(ethernetHeaderCompressionR16EhcUplinkR16DrbContinueEHCULR16Constraints)
				if err != nil {
					return err
				}
				(*c.Drb_ContinueEHC_UL_r16) = v
			}
		}
	}

	return nil
}
