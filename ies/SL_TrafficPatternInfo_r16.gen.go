// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-TrafficPatternInfo-r16 (line 2702).

var sLTrafficPatternInfoR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "trafficPeriodicity-r16"},
		{Name: "timingOffset-r16"},
		{Name: "messageSize-r16"},
		{Name: "sl-QoS-FlowIdentity-r16"},
	},
}

const (
	SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms20   = 0
	SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms50   = 1
	SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms100  = 2
	SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms200  = 3
	SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms300  = 4
	SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms400  = 5
	SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms500  = 6
	SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms600  = 7
	SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms700  = 8
	SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms800  = 9
	SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms900  = 10
	SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms1000 = 11
)

var sLTrafficPatternInfoR16TrafficPeriodicityR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms20, SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms50, SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms100, SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms200, SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms300, SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms400, SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms500, SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms600, SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms700, SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms800, SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms900, SL_TrafficPatternInfo_r16_TrafficPeriodicity_r16_Ms1000},
}

var sLTrafficPatternInfoR16TimingOffsetR16Constraints = per.Constrained(0, 10239)

var sLTrafficPatternInfoR16MessageSizeR16Constraints = per.FixedSize(8)

type SL_TrafficPatternInfo_r16 struct {
	TrafficPeriodicity_r16  int64
	TimingOffset_r16        int64
	MessageSize_r16         per.BitString
	Sl_QoS_FlowIdentity_r16 SL_QoS_FlowIdentity_r16
}

func (ie *SL_TrafficPatternInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLTrafficPatternInfoR16Constraints)
	_ = seq

	// 1. trafficPeriodicity-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.TrafficPeriodicity_r16, sLTrafficPatternInfoR16TrafficPeriodicityR16Constraints); err != nil {
			return err
		}
	}

	// 2. timingOffset-r16: integer
	{
		if err := e.EncodeInteger(ie.TimingOffset_r16, sLTrafficPatternInfoR16TimingOffsetR16Constraints); err != nil {
			return err
		}
	}

	// 3. messageSize-r16: bit-string
	{
		if err := e.EncodeBitString(ie.MessageSize_r16, sLTrafficPatternInfoR16MessageSizeR16Constraints); err != nil {
			return err
		}
	}

	// 4. sl-QoS-FlowIdentity-r16: ref
	{
		if err := ie.Sl_QoS_FlowIdentity_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_TrafficPatternInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLTrafficPatternInfoR16Constraints)
	_ = seq

	// 1. trafficPeriodicity-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(sLTrafficPatternInfoR16TrafficPeriodicityR16Constraints)
		if err != nil {
			return err
		}
		ie.TrafficPeriodicity_r16 = v0
	}

	// 2. timingOffset-r16: integer
	{
		v1, err := d.DecodeInteger(sLTrafficPatternInfoR16TimingOffsetR16Constraints)
		if err != nil {
			return err
		}
		ie.TimingOffset_r16 = v1
	}

	// 3. messageSize-r16: bit-string
	{
		v2, err := d.DecodeBitString(sLTrafficPatternInfoR16MessageSizeR16Constraints)
		if err != nil {
			return err
		}
		ie.MessageSize_r16 = v2
	}

	// 4. sl-QoS-FlowIdentity-r16: ref
	{
		if err := ie.Sl_QoS_FlowIdentity_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
