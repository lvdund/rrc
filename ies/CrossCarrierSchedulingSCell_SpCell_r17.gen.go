// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CrossCarrierSchedulingSCell-SpCell-r17 (line 18229).

var crossCarrierSchedulingSCellSpCellR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedSCS-Combinations-r17"},
		{Name: "pdcch-MonitoringOccasion-r17"},
	},
}

const (
	CrossCarrierSchedulingSCell_SpCell_r17_Pdcch_MonitoringOccasion_r17_Val1 = 0
	CrossCarrierSchedulingSCell_SpCell_r17_Pdcch_MonitoringOccasion_r17_Val2 = 1
)

var crossCarrierSchedulingSCellSpCellR17PdcchMonitoringOccasionR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CrossCarrierSchedulingSCell_SpCell_r17_Pdcch_MonitoringOccasion_r17_Val1, CrossCarrierSchedulingSCell_SpCell_r17_Pdcch_MonitoringOccasion_r17_Val2},
}

var crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs15kHz-15kHz-r17", Optional: true},
		{Name: "scs15kHz-30kHz-r17", Optional: true},
		{Name: "scs15kHz-60kHz-r17", Optional: true},
		{Name: "scs30kHz-30kHz-r17", Optional: true},
		{Name: "scs30kHz-60kHz-r17", Optional: true},
		{Name: "scs60kHz-60kHz-r17", Optional: true},
	},
}

const (
	CrossCarrierSchedulingSCell_SpCell_r17_SupportedSCS_Combinations_r17_Scs15kHz_15kHz_r17_Supported = 0
)

var crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Scs15kHz15kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CrossCarrierSchedulingSCell_SpCell_r17_SupportedSCS_Combinations_r17_Scs15kHz_15kHz_r17_Supported},
}

const (
	CrossCarrierSchedulingSCell_SpCell_r17_SupportedSCS_Combinations_r17_Scs15kHz_30kHz_r17_Supported = 0
)

var crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Scs15kHz30kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CrossCarrierSchedulingSCell_SpCell_r17_SupportedSCS_Combinations_r17_Scs15kHz_30kHz_r17_Supported},
}

const (
	CrossCarrierSchedulingSCell_SpCell_r17_SupportedSCS_Combinations_r17_Scs15kHz_60kHz_r17_Supported = 0
)

var crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Scs15kHz60kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CrossCarrierSchedulingSCell_SpCell_r17_SupportedSCS_Combinations_r17_Scs15kHz_60kHz_r17_Supported},
}

type CrossCarrierSchedulingSCell_SpCell_r17 struct {
	SupportedSCS_Combinations_r17 struct {
		Scs15kHz_15kHz_r17 *int64
		Scs15kHz_30kHz_r17 *int64
		Scs15kHz_60kHz_r17 *int64
		Scs30kHz_30kHz_r17 *per.BitString
		Scs30kHz_60kHz_r17 *per.BitString
		Scs60kHz_60kHz_r17 *per.BitString
	}
	Pdcch_MonitoringOccasion_r17 int64
}

func (ie *CrossCarrierSchedulingSCell_SpCell_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(crossCarrierSchedulingSCellSpCellR17Constraints)
	_ = seq

	// 1. supportedSCS-Combinations-r17: sequence
	{
		{
			c := &ie.SupportedSCS_Combinations_r17
			crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Seq := e.NewSequenceEncoder(crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Constraints)
			if err := crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Seq.EncodePreamble([]bool{c.Scs15kHz_15kHz_r17 != nil, c.Scs15kHz_30kHz_r17 != nil, c.Scs15kHz_60kHz_r17 != nil, c.Scs30kHz_30kHz_r17 != nil, c.Scs30kHz_60kHz_r17 != nil, c.Scs60kHz_60kHz_r17 != nil}); err != nil {
				return err
			}
			if c.Scs15kHz_15kHz_r17 != nil {
				if err := e.EncodeEnumerated((*c.Scs15kHz_15kHz_r17), crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Scs15kHz15kHzR17Constraints); err != nil {
					return err
				}
			}
			if c.Scs15kHz_30kHz_r17 != nil {
				if err := e.EncodeEnumerated((*c.Scs15kHz_30kHz_r17), crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Scs15kHz30kHzR17Constraints); err != nil {
					return err
				}
			}
			if c.Scs15kHz_60kHz_r17 != nil {
				if err := e.EncodeEnumerated((*c.Scs15kHz_60kHz_r17), crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Scs15kHz60kHzR17Constraints); err != nil {
					return err
				}
			}
			if c.Scs30kHz_30kHz_r17 != nil {
				if err := e.EncodeBitString((*c.Scs30kHz_30kHz_r17), per.SizeRange(1, 496)); err != nil {
					return err
				}
			}
			if c.Scs30kHz_60kHz_r17 != nil {
				if err := e.EncodeBitString((*c.Scs30kHz_60kHz_r17), per.SizeRange(1, 496)); err != nil {
					return err
				}
			}
			if c.Scs60kHz_60kHz_r17 != nil {
				if err := e.EncodeBitString((*c.Scs60kHz_60kHz_r17), per.SizeRange(1, 496)); err != nil {
					return err
				}
			}
		}
	}

	// 2. pdcch-MonitoringOccasion-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Pdcch_MonitoringOccasion_r17, crossCarrierSchedulingSCellSpCellR17PdcchMonitoringOccasionR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CrossCarrierSchedulingSCell_SpCell_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(crossCarrierSchedulingSCellSpCellR17Constraints)
	_ = seq

	// 1. supportedSCS-Combinations-r17: sequence
	{
		{
			c := &ie.SupportedSCS_Combinations_r17
			crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Seq := d.NewSequenceDecoder(crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Constraints)
			if err := crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Seq.IsComponentPresent(0) {
				c.Scs15kHz_15kHz_r17 = new(int64)
				v, err := d.DecodeEnumerated(crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Scs15kHz15kHzR17Constraints)
				if err != nil {
					return err
				}
				(*c.Scs15kHz_15kHz_r17) = v
			}
			if crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Seq.IsComponentPresent(1) {
				c.Scs15kHz_30kHz_r17 = new(int64)
				v, err := d.DecodeEnumerated(crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Scs15kHz30kHzR17Constraints)
				if err != nil {
					return err
				}
				(*c.Scs15kHz_30kHz_r17) = v
			}
			if crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Seq.IsComponentPresent(2) {
				c.Scs15kHz_60kHz_r17 = new(int64)
				v, err := d.DecodeEnumerated(crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Scs15kHz60kHzR17Constraints)
				if err != nil {
					return err
				}
				(*c.Scs15kHz_60kHz_r17) = v
			}
			if crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Seq.IsComponentPresent(3) {
				c.Scs30kHz_30kHz_r17 = new(per.BitString)
				v, err := d.DecodeBitString(per.SizeRange(1, 496))
				if err != nil {
					return err
				}
				(*c.Scs30kHz_30kHz_r17) = v
			}
			if crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Seq.IsComponentPresent(4) {
				c.Scs30kHz_60kHz_r17 = new(per.BitString)
				v, err := d.DecodeBitString(per.SizeRange(1, 496))
				if err != nil {
					return err
				}
				(*c.Scs30kHz_60kHz_r17) = v
			}
			if crossCarrierSchedulingSCellSpCellR17SupportedSCSCombinationsR17Seq.IsComponentPresent(5) {
				c.Scs60kHz_60kHz_r17 = new(per.BitString)
				v, err := d.DecodeBitString(per.SizeRange(1, 496))
				if err != nil {
					return err
				}
				(*c.Scs60kHz_60kHz_r17) = v
			}
		}
	}

	// 2. pdcch-MonitoringOccasion-r17: enumerated
	{
		v1, err := d.DecodeEnumerated(crossCarrierSchedulingSCellSpCellR17PdcchMonitoringOccasionR17Constraints)
		if err != nil {
			return err
		}
		ie.Pdcch_MonitoringOccasion_r17 = v1
	}

	return nil
}
