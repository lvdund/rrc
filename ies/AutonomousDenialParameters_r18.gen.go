// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AutonomousDenialParameters-r18 (line 5802).

var autonomousDenialParametersR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "autonomousDenialSlots-r18"},
		{Name: "autonomousDenialValidity-r18"},
	},
}

const (
	AutonomousDenialParameters_r18_AutonomousDenialSlots_r18_N2     = 0
	AutonomousDenialParameters_r18_AutonomousDenialSlots_r18_N5     = 1
	AutonomousDenialParameters_r18_AutonomousDenialSlots_r18_N10    = 2
	AutonomousDenialParameters_r18_AutonomousDenialSlots_r18_N15    = 3
	AutonomousDenialParameters_r18_AutonomousDenialSlots_r18_N20    = 4
	AutonomousDenialParameters_r18_AutonomousDenialSlots_r18_N30    = 5
	AutonomousDenialParameters_r18_AutonomousDenialSlots_r18_Spare2 = 6
	AutonomousDenialParameters_r18_AutonomousDenialSlots_r18_Spare1 = 7
)

var autonomousDenialParametersR18AutonomousDenialSlotsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AutonomousDenialParameters_r18_AutonomousDenialSlots_r18_N2, AutonomousDenialParameters_r18_AutonomousDenialSlots_r18_N5, AutonomousDenialParameters_r18_AutonomousDenialSlots_r18_N10, AutonomousDenialParameters_r18_AutonomousDenialSlots_r18_N15, AutonomousDenialParameters_r18_AutonomousDenialSlots_r18_N20, AutonomousDenialParameters_r18_AutonomousDenialSlots_r18_N30, AutonomousDenialParameters_r18_AutonomousDenialSlots_r18_Spare2, AutonomousDenialParameters_r18_AutonomousDenialSlots_r18_Spare1},
}

const (
	AutonomousDenialParameters_r18_AutonomousDenialValidity_r18_N200  = 0
	AutonomousDenialParameters_r18_AutonomousDenialValidity_r18_N500  = 1
	AutonomousDenialParameters_r18_AutonomousDenialValidity_r18_N1000 = 2
	AutonomousDenialParameters_r18_AutonomousDenialValidity_r18_N2000 = 3
)

var autonomousDenialParametersR18AutonomousDenialValidityR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AutonomousDenialParameters_r18_AutonomousDenialValidity_r18_N200, AutonomousDenialParameters_r18_AutonomousDenialValidity_r18_N500, AutonomousDenialParameters_r18_AutonomousDenialValidity_r18_N1000, AutonomousDenialParameters_r18_AutonomousDenialValidity_r18_N2000},
}

type AutonomousDenialParameters_r18 struct {
	AutonomousDenialSlots_r18    int64
	AutonomousDenialValidity_r18 int64
}

func (ie *AutonomousDenialParameters_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(autonomousDenialParametersR18Constraints)
	_ = seq

	// 1. autonomousDenialSlots-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.AutonomousDenialSlots_r18, autonomousDenialParametersR18AutonomousDenialSlotsR18Constraints); err != nil {
			return err
		}
	}

	// 2. autonomousDenialValidity-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.AutonomousDenialValidity_r18, autonomousDenialParametersR18AutonomousDenialValidityR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *AutonomousDenialParameters_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(autonomousDenialParametersR18Constraints)
	_ = seq

	// 1. autonomousDenialSlots-r18: enumerated
	{
		v0, err := d.DecodeEnumerated(autonomousDenialParametersR18AutonomousDenialSlotsR18Constraints)
		if err != nil {
			return err
		}
		ie.AutonomousDenialSlots_r18 = v0
	}

	// 2. autonomousDenialValidity-r18: enumerated
	{
		v1, err := d.DecodeEnumerated(autonomousDenialParametersR18AutonomousDenialValidityR18Constraints)
		if err != nil {
			return err
		}
		ie.AutonomousDenialValidity_r18 = v1
	}

	return nil
}
