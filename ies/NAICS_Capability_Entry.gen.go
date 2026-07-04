// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NAICS-Capability-Entry (line 23272).

var nAICSCapabilityEntryConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "numberOfNAICS-CapableCC"},
		{Name: "numberOfAggregatedPRB"},
	},
}

var nAICSCapabilityEntryNumberOfNAICSCapableCCConstraints = per.Constrained(1, 5)

const (
	NAICS_Capability_Entry_NumberOfAggregatedPRB_N50   = 0
	NAICS_Capability_Entry_NumberOfAggregatedPRB_N75   = 1
	NAICS_Capability_Entry_NumberOfAggregatedPRB_N100  = 2
	NAICS_Capability_Entry_NumberOfAggregatedPRB_N125  = 3
	NAICS_Capability_Entry_NumberOfAggregatedPRB_N150  = 4
	NAICS_Capability_Entry_NumberOfAggregatedPRB_N175  = 5
	NAICS_Capability_Entry_NumberOfAggregatedPRB_N200  = 6
	NAICS_Capability_Entry_NumberOfAggregatedPRB_N225  = 7
	NAICS_Capability_Entry_NumberOfAggregatedPRB_N250  = 8
	NAICS_Capability_Entry_NumberOfAggregatedPRB_N275  = 9
	NAICS_Capability_Entry_NumberOfAggregatedPRB_N300  = 10
	NAICS_Capability_Entry_NumberOfAggregatedPRB_N350  = 11
	NAICS_Capability_Entry_NumberOfAggregatedPRB_N400  = 12
	NAICS_Capability_Entry_NumberOfAggregatedPRB_N450  = 13
	NAICS_Capability_Entry_NumberOfAggregatedPRB_N500  = 14
	NAICS_Capability_Entry_NumberOfAggregatedPRB_Spare = 15
)

var nAICSCapabilityEntryNumberOfAggregatedPRBConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NAICS_Capability_Entry_NumberOfAggregatedPRB_N50, NAICS_Capability_Entry_NumberOfAggregatedPRB_N75, NAICS_Capability_Entry_NumberOfAggregatedPRB_N100, NAICS_Capability_Entry_NumberOfAggregatedPRB_N125, NAICS_Capability_Entry_NumberOfAggregatedPRB_N150, NAICS_Capability_Entry_NumberOfAggregatedPRB_N175, NAICS_Capability_Entry_NumberOfAggregatedPRB_N200, NAICS_Capability_Entry_NumberOfAggregatedPRB_N225, NAICS_Capability_Entry_NumberOfAggregatedPRB_N250, NAICS_Capability_Entry_NumberOfAggregatedPRB_N275, NAICS_Capability_Entry_NumberOfAggregatedPRB_N300, NAICS_Capability_Entry_NumberOfAggregatedPRB_N350, NAICS_Capability_Entry_NumberOfAggregatedPRB_N400, NAICS_Capability_Entry_NumberOfAggregatedPRB_N450, NAICS_Capability_Entry_NumberOfAggregatedPRB_N500, NAICS_Capability_Entry_NumberOfAggregatedPRB_Spare},
}

type NAICS_Capability_Entry struct {
	NumberOfNAICS_CapableCC int64
	NumberOfAggregatedPRB   int64
}

func (ie *NAICS_Capability_Entry) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nAICSCapabilityEntryConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. numberOfNAICS-CapableCC: integer
	{
		if err := e.EncodeInteger(ie.NumberOfNAICS_CapableCC, nAICSCapabilityEntryNumberOfNAICSCapableCCConstraints); err != nil {
			return err
		}
	}

	// 3. numberOfAggregatedPRB: enumerated
	{
		if err := e.EncodeEnumerated(ie.NumberOfAggregatedPRB, nAICSCapabilityEntryNumberOfAggregatedPRBConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *NAICS_Capability_Entry) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nAICSCapabilityEntryConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. numberOfNAICS-CapableCC: integer
	{
		v0, err := d.DecodeInteger(nAICSCapabilityEntryNumberOfNAICSCapableCCConstraints)
		if err != nil {
			return err
		}
		ie.NumberOfNAICS_CapableCC = v0
	}

	// 3. numberOfAggregatedPRB: enumerated
	{
		v1, err := d.DecodeEnumerated(nAICSCapabilityEntryNumberOfAggregatedPRBConstraints)
		if err != nil {
			return err
		}
		ie.NumberOfAggregatedPRB = v1
	}

	return nil
}
