// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CSI-RS-ProcFrameworkForSRS (line 22476).

var cSIRSProcFrameworkForSRSConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberPeriodicSRS-AssocCSI-RS-PerBWP"},
		{Name: "maxNumberAperiodicSRS-AssocCSI-RS-PerBWP"},
		{Name: "maxNumberSP-SRS-AssocCSI-RS-PerBWP"},
		{Name: "simultaneousSRS-AssocCSI-RS-PerCC"},
	},
}

var cSIRSProcFrameworkForSRSMaxNumberPeriodicSRSAssocCSIRSPerBWPConstraints = per.Constrained(1, 4)

var cSIRSProcFrameworkForSRSMaxNumberAperiodicSRSAssocCSIRSPerBWPConstraints = per.Constrained(1, 4)

var cSIRSProcFrameworkForSRSMaxNumberSPSRSAssocCSIRSPerBWPConstraints = per.Constrained(0, 4)

var cSIRSProcFrameworkForSRSSimultaneousSRSAssocCSIRSPerCCConstraints = per.Constrained(1, 8)

type CSI_RS_ProcFrameworkForSRS struct {
	MaxNumberPeriodicSRS_AssocCSI_RS_PerBWP  int64
	MaxNumberAperiodicSRS_AssocCSI_RS_PerBWP int64
	MaxNumberSP_SRS_AssocCSI_RS_PerBWP       int64
	SimultaneousSRS_AssocCSI_RS_PerCC        int64
}

func (ie *CSI_RS_ProcFrameworkForSRS) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIRSProcFrameworkForSRSConstraints)
	_ = seq

	// 1. maxNumberPeriodicSRS-AssocCSI-RS-PerBWP: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberPeriodicSRS_AssocCSI_RS_PerBWP, cSIRSProcFrameworkForSRSMaxNumberPeriodicSRSAssocCSIRSPerBWPConstraints); err != nil {
			return err
		}
	}

	// 2. maxNumberAperiodicSRS-AssocCSI-RS-PerBWP: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberAperiodicSRS_AssocCSI_RS_PerBWP, cSIRSProcFrameworkForSRSMaxNumberAperiodicSRSAssocCSIRSPerBWPConstraints); err != nil {
			return err
		}
	}

	// 3. maxNumberSP-SRS-AssocCSI-RS-PerBWP: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberSP_SRS_AssocCSI_RS_PerBWP, cSIRSProcFrameworkForSRSMaxNumberSPSRSAssocCSIRSPerBWPConstraints); err != nil {
			return err
		}
	}

	// 4. simultaneousSRS-AssocCSI-RS-PerCC: integer
	{
		if err := e.EncodeInteger(ie.SimultaneousSRS_AssocCSI_RS_PerCC, cSIRSProcFrameworkForSRSSimultaneousSRSAssocCSIRSPerCCConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CSI_RS_ProcFrameworkForSRS) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIRSProcFrameworkForSRSConstraints)
	_ = seq

	// 1. maxNumberPeriodicSRS-AssocCSI-RS-PerBWP: integer
	{
		v0, err := d.DecodeInteger(cSIRSProcFrameworkForSRSMaxNumberPeriodicSRSAssocCSIRSPerBWPConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberPeriodicSRS_AssocCSI_RS_PerBWP = v0
	}

	// 2. maxNumberAperiodicSRS-AssocCSI-RS-PerBWP: integer
	{
		v1, err := d.DecodeInteger(cSIRSProcFrameworkForSRSMaxNumberAperiodicSRSAssocCSIRSPerBWPConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberAperiodicSRS_AssocCSI_RS_PerBWP = v1
	}

	// 3. maxNumberSP-SRS-AssocCSI-RS-PerBWP: integer
	{
		v2, err := d.DecodeInteger(cSIRSProcFrameworkForSRSMaxNumberSPSRSAssocCSIRSPerBWPConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberSP_SRS_AssocCSI_RS_PerBWP = v2
	}

	// 4. simultaneousSRS-AssocCSI-RS-PerCC: integer
	{
		v3, err := d.DecodeInteger(cSIRSProcFrameworkForSRSSimultaneousSRSAssocCSIRSPerCCConstraints)
		if err != nil {
			return err
		}
		ie.SimultaneousSRS_AssocCSI_RS_PerCC = v3
	}

	return nil
}
