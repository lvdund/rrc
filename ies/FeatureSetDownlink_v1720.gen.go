// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetDownlink-v1720 (line 19627).

var featureSetDownlinkV1720Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rtt-BasedPDC-CSI-RS-ForTracking-r17", Optional: true},
		{Name: "rtt-BasedPDC-PRS-r17", Optional: true},
		{Name: "sps-Multicast-r17", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1720_Rtt_BasedPDC_CSI_RS_ForTracking_r17_Supported = 0
)

var featureSetDownlinkV1720RttBasedPDCCSIRSForTrackingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1720_Rtt_BasedPDC_CSI_RS_ForTracking_r17_Supported},
}

const (
	FeatureSetDownlink_v1720_Sps_Multicast_r17_Supported = 0
)

var featureSetDownlinkV1720SpsMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1720_Sps_Multicast_r17_Supported},
}

const (
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_Resource_r17_N1  = 0
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_Resource_r17_N2  = 1
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_Resource_r17_N4  = 2
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_Resource_r17_N8  = 3
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_Resource_r17_N16 = 4
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_Resource_r17_N32 = 5
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_Resource_r17_N64 = 6
)

var featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_Resource_r17_N1, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_Resource_r17_N2, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_Resource_r17_N4, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_Resource_r17_N8, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_Resource_r17_N16, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_Resource_r17_N32, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_Resource_r17_N64},
}

var featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r17", Optional: true},
		{Name: "scs-30kHz-r17", Optional: true},
		{Name: "scs-60kHz-r17", Optional: true},
		{Name: "scs-120kHz-r17", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N1  = 0
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N2  = 1
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N4  = 2
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N6  = 3
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N8  = 4
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N12 = 5
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N16 = 6
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N24 = 7
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N32 = 8
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N48 = 9
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N64 = 10
)

var featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Scs15kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N1, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N2, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N4, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N6, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N8, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N12, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N16, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N24, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N32, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N48, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_15kHz_r17_N64},
}

const (
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N1  = 0
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N2  = 1
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N4  = 2
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N6  = 3
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N8  = 4
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N12 = 5
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N16 = 6
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N24 = 7
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N32 = 8
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N48 = 9
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N64 = 10
)

var featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Scs30kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N1, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N2, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N4, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N6, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N8, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N12, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N16, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N24, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N32, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N48, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_30kHz_r17_N64},
}

const (
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N1  = 0
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N2  = 1
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N4  = 2
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N6  = 3
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N8  = 4
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N12 = 5
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N16 = 6
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N24 = 7
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N32 = 8
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N48 = 9
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N64 = 10
)

var featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Scs60kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N1, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N2, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N4, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N6, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N8, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N12, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N16, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N24, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N32, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N48, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_60kHz_r17_N64},
}

const (
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N1  = 0
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N2  = 1
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N4  = 2
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N6  = 3
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N8  = 4
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N12 = 5
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N16 = 6
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N24 = 7
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N32 = 8
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N48 = 9
	FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N64 = 10
)

var featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Scs120kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N1, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N2, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N4, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N6, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N8, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N12, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N16, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N24, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N32, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N48, FeatureSetDownlink_v1720_Rtt_BasedPDC_PRS_r17_MaxNumberPRS_ResourceProcessedPerSlot_r17_Scs_120kHz_r17_N64},
}

type FeatureSetDownlink_v1720 struct {
	Rtt_BasedPDC_CSI_RS_ForTracking_r17 *int64
	Rtt_BasedPDC_PRS_r17                *struct {
		MaxNumberPRS_Resource_r17                 int64
		MaxNumberPRS_ResourceProcessedPerSlot_r17 struct {
			Scs_15kHz_r17  *int64
			Scs_30kHz_r17  *int64
			Scs_60kHz_r17  *int64
			Scs_120kHz_r17 *int64
		}
	}
	Sps_Multicast_r17 *int64
}

func (ie *FeatureSetDownlink_v1720) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkV1720Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rtt_BasedPDC_CSI_RS_ForTracking_r17 != nil, ie.Rtt_BasedPDC_PRS_r17 != nil, ie.Sps_Multicast_r17 != nil}); err != nil {
		return err
	}

	// 2. rtt-BasedPDC-CSI-RS-ForTracking-r17: enumerated
	{
		if ie.Rtt_BasedPDC_CSI_RS_ForTracking_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Rtt_BasedPDC_CSI_RS_ForTracking_r17, featureSetDownlinkV1720RttBasedPDCCSIRSForTrackingR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. rtt-BasedPDC-PRS-r17: sequence
	{
		if ie.Rtt_BasedPDC_PRS_r17 != nil {
			c := ie.Rtt_BasedPDC_PRS_r17
			if err := e.EncodeEnumerated(c.MaxNumberPRS_Resource_r17, featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceR17Constraints); err != nil {
				return err
			}
			{
				c := &c.MaxNumberPRS_ResourceProcessedPerSlot_r17
				featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Seq := e.NewSequenceEncoder(featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Constraints)
				if err := featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Seq.EncodePreamble([]bool{c.Scs_15kHz_r17 != nil, c.Scs_30kHz_r17 != nil, c.Scs_60kHz_r17 != nil, c.Scs_120kHz_r17 != nil}); err != nil {
					return err
				}
				if c.Scs_15kHz_r17 != nil {
					if err := e.EncodeEnumerated((*c.Scs_15kHz_r17), featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Scs15kHzR17Constraints); err != nil {
						return err
					}
				}
				if c.Scs_30kHz_r17 != nil {
					if err := e.EncodeEnumerated((*c.Scs_30kHz_r17), featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Scs30kHzR17Constraints); err != nil {
						return err
					}
				}
				if c.Scs_60kHz_r17 != nil {
					if err := e.EncodeEnumerated((*c.Scs_60kHz_r17), featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Scs60kHzR17Constraints); err != nil {
						return err
					}
				}
				if c.Scs_120kHz_r17 != nil {
					if err := e.EncodeEnumerated((*c.Scs_120kHz_r17), featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Scs120kHzR17Constraints); err != nil {
						return err
					}
				}
			}
		}
	}

	// 4. sps-Multicast-r17: enumerated
	{
		if ie.Sps_Multicast_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sps_Multicast_r17, featureSetDownlinkV1720SpsMulticastR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlink_v1720) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkV1720Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. rtt-BasedPDC-CSI-RS-ForTracking-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1720RttBasedPDCCSIRSForTrackingR17Constraints)
			if err != nil {
				return err
			}
			ie.Rtt_BasedPDC_CSI_RS_ForTracking_r17 = &idx
		}
	}

	// 3. rtt-BasedPDC-PRS-r17: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.Rtt_BasedPDC_PRS_r17 = &struct {
				MaxNumberPRS_Resource_r17                 int64
				MaxNumberPRS_ResourceProcessedPerSlot_r17 struct {
					Scs_15kHz_r17  *int64
					Scs_30kHz_r17  *int64
					Scs_60kHz_r17  *int64
					Scs_120kHz_r17 *int64
				}
			}{}
			c := ie.Rtt_BasedPDC_PRS_r17
			{
				v, err := d.DecodeEnumerated(featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceR17Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberPRS_Resource_r17 = v
			}
			{
				c := &c.MaxNumberPRS_ResourceProcessedPerSlot_r17
				featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Seq := d.NewSequenceDecoder(featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Constraints)
				if err := featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Seq.DecodePreamble(); err != nil {
					return err
				}
				if featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Seq.IsComponentPresent(0) {
					c.Scs_15kHz_r17 = new(int64)
					v, err := d.DecodeEnumerated(featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Scs15kHzR17Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_15kHz_r17) = v
				}
				if featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Seq.IsComponentPresent(1) {
					c.Scs_30kHz_r17 = new(int64)
					v, err := d.DecodeEnumerated(featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Scs30kHzR17Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_30kHz_r17) = v
				}
				if featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Seq.IsComponentPresent(2) {
					c.Scs_60kHz_r17 = new(int64)
					v, err := d.DecodeEnumerated(featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Scs60kHzR17Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_60kHz_r17) = v
				}
				if featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Seq.IsComponentPresent(3) {
					c.Scs_120kHz_r17 = new(int64)
					v, err := d.DecodeEnumerated(featureSetDownlinkV1720RttBasedPDCPRSR17MaxNumberPRSResourceProcessedPerSlotR17Scs120kHzR17Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_120kHz_r17) = v
				}
			}
		}
	}

	// 4. sps-Multicast-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1720SpsMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.Sps_Multicast_r17 = &idx
		}
	}

	return nil
}
