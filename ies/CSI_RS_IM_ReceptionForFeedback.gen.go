// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CSI-RS-IM-ReceptionForFeedback (line 22468).

var cSIRSIMReceptionForFeedbackConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxConfigNumberNZP-CSI-RS-PerCC"},
		{Name: "maxConfigNumberPortsAcrossNZP-CSI-RS-PerCC"},
		{Name: "maxConfigNumberCSI-IM-PerCC"},
		{Name: "maxNumberSimultaneousNZP-CSI-RS-PerCC"},
		{Name: "totalNumberPortsSimultaneousNZP-CSI-RS-PerCC"},
	},
}

var cSIRSIMReceptionForFeedbackMaxConfigNumberNZPCSIRSPerCCConstraints = per.Constrained(1, 64)

var cSIRSIMReceptionForFeedbackMaxConfigNumberPortsAcrossNZPCSIRSPerCCConstraints = per.Constrained(2, 256)

const (
	CSI_RS_IM_ReceptionForFeedback_MaxConfigNumberCSI_IM_PerCC_N1  = 0
	CSI_RS_IM_ReceptionForFeedback_MaxConfigNumberCSI_IM_PerCC_N2  = 1
	CSI_RS_IM_ReceptionForFeedback_MaxConfigNumberCSI_IM_PerCC_N4  = 2
	CSI_RS_IM_ReceptionForFeedback_MaxConfigNumberCSI_IM_PerCC_N8  = 3
	CSI_RS_IM_ReceptionForFeedback_MaxConfigNumberCSI_IM_PerCC_N16 = 4
	CSI_RS_IM_ReceptionForFeedback_MaxConfigNumberCSI_IM_PerCC_N32 = 5
)

var cSIRSIMReceptionForFeedbackMaxConfigNumberCSIIMPerCCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_RS_IM_ReceptionForFeedback_MaxConfigNumberCSI_IM_PerCC_N1, CSI_RS_IM_ReceptionForFeedback_MaxConfigNumberCSI_IM_PerCC_N2, CSI_RS_IM_ReceptionForFeedback_MaxConfigNumberCSI_IM_PerCC_N4, CSI_RS_IM_ReceptionForFeedback_MaxConfigNumberCSI_IM_PerCC_N8, CSI_RS_IM_ReceptionForFeedback_MaxConfigNumberCSI_IM_PerCC_N16, CSI_RS_IM_ReceptionForFeedback_MaxConfigNumberCSI_IM_PerCC_N32},
}

var cSIRSIMReceptionForFeedbackMaxNumberSimultaneousNZPCSIRSPerCCConstraints = per.Constrained(1, 64)

var cSIRSIMReceptionForFeedbackTotalNumberPortsSimultaneousNZPCSIRSPerCCConstraints = per.Constrained(2, 256)

type CSI_RS_IM_ReceptionForFeedback struct {
	MaxConfigNumberNZP_CSI_RS_PerCC              int64
	MaxConfigNumberPortsAcrossNZP_CSI_RS_PerCC   int64
	MaxConfigNumberCSI_IM_PerCC                  int64
	MaxNumberSimultaneousNZP_CSI_RS_PerCC        int64
	TotalNumberPortsSimultaneousNZP_CSI_RS_PerCC int64
}

func (ie *CSI_RS_IM_ReceptionForFeedback) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIRSIMReceptionForFeedbackConstraints)
	_ = seq

	// 1. maxConfigNumberNZP-CSI-RS-PerCC: integer
	{
		if err := e.EncodeInteger(ie.MaxConfigNumberNZP_CSI_RS_PerCC, cSIRSIMReceptionForFeedbackMaxConfigNumberNZPCSIRSPerCCConstraints); err != nil {
			return err
		}
	}

	// 2. maxConfigNumberPortsAcrossNZP-CSI-RS-PerCC: integer
	{
		if err := e.EncodeInteger(ie.MaxConfigNumberPortsAcrossNZP_CSI_RS_PerCC, cSIRSIMReceptionForFeedbackMaxConfigNumberPortsAcrossNZPCSIRSPerCCConstraints); err != nil {
			return err
		}
	}

	// 3. maxConfigNumberCSI-IM-PerCC: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxConfigNumberCSI_IM_PerCC, cSIRSIMReceptionForFeedbackMaxConfigNumberCSIIMPerCCConstraints); err != nil {
			return err
		}
	}

	// 4. maxNumberSimultaneousNZP-CSI-RS-PerCC: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberSimultaneousNZP_CSI_RS_PerCC, cSIRSIMReceptionForFeedbackMaxNumberSimultaneousNZPCSIRSPerCCConstraints); err != nil {
			return err
		}
	}

	// 5. totalNumberPortsSimultaneousNZP-CSI-RS-PerCC: integer
	{
		if err := e.EncodeInteger(ie.TotalNumberPortsSimultaneousNZP_CSI_RS_PerCC, cSIRSIMReceptionForFeedbackTotalNumberPortsSimultaneousNZPCSIRSPerCCConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CSI_RS_IM_ReceptionForFeedback) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIRSIMReceptionForFeedbackConstraints)
	_ = seq

	// 1. maxConfigNumberNZP-CSI-RS-PerCC: integer
	{
		v0, err := d.DecodeInteger(cSIRSIMReceptionForFeedbackMaxConfigNumberNZPCSIRSPerCCConstraints)
		if err != nil {
			return err
		}
		ie.MaxConfigNumberNZP_CSI_RS_PerCC = v0
	}

	// 2. maxConfigNumberPortsAcrossNZP-CSI-RS-PerCC: integer
	{
		v1, err := d.DecodeInteger(cSIRSIMReceptionForFeedbackMaxConfigNumberPortsAcrossNZPCSIRSPerCCConstraints)
		if err != nil {
			return err
		}
		ie.MaxConfigNumberPortsAcrossNZP_CSI_RS_PerCC = v1
	}

	// 3. maxConfigNumberCSI-IM-PerCC: enumerated
	{
		v2, err := d.DecodeEnumerated(cSIRSIMReceptionForFeedbackMaxConfigNumberCSIIMPerCCConstraints)
		if err != nil {
			return err
		}
		ie.MaxConfigNumberCSI_IM_PerCC = v2
	}

	// 4. maxNumberSimultaneousNZP-CSI-RS-PerCC: integer
	{
		v3, err := d.DecodeInteger(cSIRSIMReceptionForFeedbackMaxNumberSimultaneousNZPCSIRSPerCCConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberSimultaneousNZP_CSI_RS_PerCC = v3
	}

	// 5. totalNumberPortsSimultaneousNZP-CSI-RS-PerCC: integer
	{
		v4, err := d.DecodeInteger(cSIRSIMReceptionForFeedbackTotalNumberPortsSimultaneousNZPCSIRSPerCCConstraints)
		if err != nil {
			return err
		}
		ie.TotalNumberPortsSimultaneousNZP_CSI_RS_PerCC = v4
	}

	return nil
}
