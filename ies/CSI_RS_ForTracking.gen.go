// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CSI-RS-ForTracking (line 22461).

var cSIRSForTrackingConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxBurstLength"},
		{Name: "maxSimultaneousResourceSetsPerCC"},
		{Name: "maxConfiguredResourceSetsPerCC"},
		{Name: "maxConfiguredResourceSetsAllCC"},
	},
}

var cSIRSForTrackingMaxBurstLengthConstraints = per.Constrained(1, 2)

var cSIRSForTrackingMaxSimultaneousResourceSetsPerCCConstraints = per.Constrained(1, 8)

var cSIRSForTrackingMaxConfiguredResourceSetsPerCCConstraints = per.Constrained(1, 64)

var cSIRSForTrackingMaxConfiguredResourceSetsAllCCConstraints = per.Constrained(1, 256)

type CSI_RS_ForTracking struct {
	MaxBurstLength                   int64
	MaxSimultaneousResourceSetsPerCC int64
	MaxConfiguredResourceSetsPerCC   int64
	MaxConfiguredResourceSetsAllCC   int64
}

func (ie *CSI_RS_ForTracking) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIRSForTrackingConstraints)
	_ = seq

	// 1. maxBurstLength: integer
	{
		if err := e.EncodeInteger(ie.MaxBurstLength, cSIRSForTrackingMaxBurstLengthConstraints); err != nil {
			return err
		}
	}

	// 2. maxSimultaneousResourceSetsPerCC: integer
	{
		if err := e.EncodeInteger(ie.MaxSimultaneousResourceSetsPerCC, cSIRSForTrackingMaxSimultaneousResourceSetsPerCCConstraints); err != nil {
			return err
		}
	}

	// 3. maxConfiguredResourceSetsPerCC: integer
	{
		if err := e.EncodeInteger(ie.MaxConfiguredResourceSetsPerCC, cSIRSForTrackingMaxConfiguredResourceSetsPerCCConstraints); err != nil {
			return err
		}
	}

	// 4. maxConfiguredResourceSetsAllCC: integer
	{
		if err := e.EncodeInteger(ie.MaxConfiguredResourceSetsAllCC, cSIRSForTrackingMaxConfiguredResourceSetsAllCCConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CSI_RS_ForTracking) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIRSForTrackingConstraints)
	_ = seq

	// 1. maxBurstLength: integer
	{
		v0, err := d.DecodeInteger(cSIRSForTrackingMaxBurstLengthConstraints)
		if err != nil {
			return err
		}
		ie.MaxBurstLength = v0
	}

	// 2. maxSimultaneousResourceSetsPerCC: integer
	{
		v1, err := d.DecodeInteger(cSIRSForTrackingMaxSimultaneousResourceSetsPerCCConstraints)
		if err != nil {
			return err
		}
		ie.MaxSimultaneousResourceSetsPerCC = v1
	}

	// 3. maxConfiguredResourceSetsPerCC: integer
	{
		v2, err := d.DecodeInteger(cSIRSForTrackingMaxConfiguredResourceSetsPerCCConstraints)
		if err != nil {
			return err
		}
		ie.MaxConfiguredResourceSetsPerCC = v2
	}

	// 4. maxConfiguredResourceSetsAllCC: integer
	{
		v3, err := d.DecodeInteger(cSIRSForTrackingMaxConfiguredResourceSetsAllCCConstraints)
		if err != nil {
			return err
		}
		ie.MaxConfiguredResourceSetsAllCC = v3
	}

	return nil
}
