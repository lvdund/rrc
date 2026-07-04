// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RLM-RelaxationReportingConfig-r17 (line 26469).

var rLMRelaxationReportingConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rlm-RelaxtionReportingProhibitTimer-r17"},
	},
}

const (
	RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S0       = 0
	RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S0dot5   = 1
	RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S1       = 2
	RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S2       = 3
	RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S5       = 4
	RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S10      = 5
	RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S20      = 6
	RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S30      = 7
	RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S60      = 8
	RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S90      = 9
	RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S120     = 10
	RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S300     = 11
	RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S600     = 12
	RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_Infinity = 13
	RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_Spare2   = 14
	RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_Spare1   = 15
)

var rLMRelaxationReportingConfigR17RlmRelaxtionReportingProhibitTimerR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S0, RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S0dot5, RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S1, RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S2, RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S5, RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S10, RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S20, RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S30, RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S60, RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S90, RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S120, RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S300, RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_S600, RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_Infinity, RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_Spare2, RLM_RelaxationReportingConfig_r17_Rlm_RelaxtionReportingProhibitTimer_r17_Spare1},
}

type RLM_RelaxationReportingConfig_r17 struct {
	Rlm_RelaxtionReportingProhibitTimer_r17 int64
}

func (ie *RLM_RelaxationReportingConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rLMRelaxationReportingConfigR17Constraints)
	_ = seq

	// 1. rlm-RelaxtionReportingProhibitTimer-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Rlm_RelaxtionReportingProhibitTimer_r17, rLMRelaxationReportingConfigR17RlmRelaxtionReportingProhibitTimerR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RLM_RelaxationReportingConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rLMRelaxationReportingConfigR17Constraints)
	_ = seq

	// 1. rlm-RelaxtionReportingProhibitTimer-r17: enumerated
	{
		v0, err := d.DecodeEnumerated(rLMRelaxationReportingConfigR17RlmRelaxtionReportingProhibitTimerR17Constraints)
		if err != nil {
			return err
		}
		ie.Rlm_RelaxtionReportingProhibitTimer_r17 = v0
	}

	return nil
}
