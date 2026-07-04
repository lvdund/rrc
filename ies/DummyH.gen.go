// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DummyH (line 22454).

var dummyHConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "burstLength"},
		{Name: "maxSimultaneousResourceSetsPerCC"},
		{Name: "maxConfiguredResourceSetsPerCC"},
		{Name: "maxConfiguredResourceSetsAllCC"},
	},
}

var dummyHBurstLengthConstraints = per.Constrained(1, 2)

var dummyHMaxSimultaneousResourceSetsPerCCConstraints = per.Constrained(1, 8)

var dummyHMaxConfiguredResourceSetsPerCCConstraints = per.Constrained(1, 64)

var dummyHMaxConfiguredResourceSetsAllCCConstraints = per.Constrained(1, 128)

type DummyH struct {
	BurstLength                      int64
	MaxSimultaneousResourceSetsPerCC int64
	MaxConfiguredResourceSetsPerCC   int64
	MaxConfiguredResourceSetsAllCC   int64
}

func (ie *DummyH) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dummyHConstraints)
	_ = seq

	// 1. burstLength: integer
	{
		if err := e.EncodeInteger(ie.BurstLength, dummyHBurstLengthConstraints); err != nil {
			return err
		}
	}

	// 2. maxSimultaneousResourceSetsPerCC: integer
	{
		if err := e.EncodeInteger(ie.MaxSimultaneousResourceSetsPerCC, dummyHMaxSimultaneousResourceSetsPerCCConstraints); err != nil {
			return err
		}
	}

	// 3. maxConfiguredResourceSetsPerCC: integer
	{
		if err := e.EncodeInteger(ie.MaxConfiguredResourceSetsPerCC, dummyHMaxConfiguredResourceSetsPerCCConstraints); err != nil {
			return err
		}
	}

	// 4. maxConfiguredResourceSetsAllCC: integer
	{
		if err := e.EncodeInteger(ie.MaxConfiguredResourceSetsAllCC, dummyHMaxConfiguredResourceSetsAllCCConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DummyH) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dummyHConstraints)
	_ = seq

	// 1. burstLength: integer
	{
		v0, err := d.DecodeInteger(dummyHBurstLengthConstraints)
		if err != nil {
			return err
		}
		ie.BurstLength = v0
	}

	// 2. maxSimultaneousResourceSetsPerCC: integer
	{
		v1, err := d.DecodeInteger(dummyHMaxSimultaneousResourceSetsPerCCConstraints)
		if err != nil {
			return err
		}
		ie.MaxSimultaneousResourceSetsPerCC = v1
	}

	// 3. maxConfiguredResourceSetsPerCC: integer
	{
		v2, err := d.DecodeInteger(dummyHMaxConfiguredResourceSetsPerCCConstraints)
		if err != nil {
			return err
		}
		ie.MaxConfiguredResourceSetsPerCC = v2
	}

	// 4. maxConfiguredResourceSetsAllCC: integer
	{
		v3, err := d.DecodeInteger(dummyHMaxConfiguredResourceSetsAllCCConstraints)
		if err != nil {
			return err
		}
		ie.MaxConfiguredResourceSetsAllCC = v3
	}

	return nil
}
