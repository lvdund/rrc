// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UL-TrafficInfoReportingConfig-r18 (line 26516).

var uLTrafficInfoReportingConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdu-SessionsToReportUL-TrafficInfoList-r18"},
		{Name: "ul-TrafficInfoProhibitTimer-r18"},
	},
}

var uLTrafficInfoReportingConfigR18PduSessionsToReportULTrafficInfoListR18Constraints = per.SizeRange(1, common.MaxNrofPDU_Sessions_r17)

const (
	UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S0     = 0
	UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S0dot5 = 1
	UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S1     = 2
	UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S2     = 3
	UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S5     = 4
	UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S10    = 5
	UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S20    = 6
	UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S30    = 7
	UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S60    = 8
	UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S90    = 9
	UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S120   = 10
	UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S300   = 11
	UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S600   = 12
	UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_Spare3 = 13
	UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_Spare2 = 14
	UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_Spare1 = 15
)

var uLTrafficInfoReportingConfigR18UlTrafficInfoProhibitTimerR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S0, UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S0dot5, UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S1, UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S2, UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S5, UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S10, UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S20, UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S30, UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S60, UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S90, UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S120, UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S300, UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_S600, UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_Spare3, UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_Spare2, UL_TrafficInfoReportingConfig_r18_Ul_TrafficInfoProhibitTimer_r18_Spare1},
}

type UL_TrafficInfoReportingConfig_r18 struct {
	Pdu_SessionsToReportUL_TrafficInfoList_r18 []PDU_SessionToReportUL_TrafficInfo_r18
	Ul_TrafficInfoProhibitTimer_r18            int64
}

func (ie *UL_TrafficInfoReportingConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLTrafficInfoReportingConfigR18Constraints)
	_ = seq

	// 1. pdu-SessionsToReportUL-TrafficInfoList-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(uLTrafficInfoReportingConfigR18PduSessionsToReportULTrafficInfoListR18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Pdu_SessionsToReportUL_TrafficInfoList_r18))); err != nil {
			return err
		}
		for i := range ie.Pdu_SessionsToReportUL_TrafficInfoList_r18 {
			if err := ie.Pdu_SessionsToReportUL_TrafficInfoList_r18[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 2. ul-TrafficInfoProhibitTimer-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.Ul_TrafficInfoProhibitTimer_r18, uLTrafficInfoReportingConfigR18UlTrafficInfoProhibitTimerR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UL_TrafficInfoReportingConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLTrafficInfoReportingConfigR18Constraints)
	_ = seq

	// 1. pdu-SessionsToReportUL-TrafficInfoList-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(uLTrafficInfoReportingConfigR18PduSessionsToReportULTrafficInfoListR18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Pdu_SessionsToReportUL_TrafficInfoList_r18 = make([]PDU_SessionToReportUL_TrafficInfo_r18, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Pdu_SessionsToReportUL_TrafficInfoList_r18[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 2. ul-TrafficInfoProhibitTimer-r18: enumerated
	{
		v1, err := d.DecodeEnumerated(uLTrafficInfoReportingConfigR18UlTrafficInfoProhibitTimerR18Constraints)
		if err != nil {
			return err
		}
		ie.Ul_TrafficInfoProhibitTimer_r18 = v1
	}

	return nil
}
