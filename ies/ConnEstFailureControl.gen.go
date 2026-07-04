// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ConnEstFailureControl (line 6752).

var connEstFailureControlConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "connEstFailCount"},
		{Name: "connEstFailOffsetValidity"},
		{Name: "connEstFailOffset", Optional: true},
	},
}

const (
	ConnEstFailureControl_ConnEstFailCount_N1 = 0
	ConnEstFailureControl_ConnEstFailCount_N2 = 1
	ConnEstFailureControl_ConnEstFailCount_N3 = 2
	ConnEstFailureControl_ConnEstFailCount_N4 = 3
)

var connEstFailureControlConnEstFailCountConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConnEstFailureControl_ConnEstFailCount_N1, ConnEstFailureControl_ConnEstFailCount_N2, ConnEstFailureControl_ConnEstFailCount_N3, ConnEstFailureControl_ConnEstFailCount_N4},
}

const (
	ConnEstFailureControl_ConnEstFailOffsetValidity_S30  = 0
	ConnEstFailureControl_ConnEstFailOffsetValidity_S60  = 1
	ConnEstFailureControl_ConnEstFailOffsetValidity_S120 = 2
	ConnEstFailureControl_ConnEstFailOffsetValidity_S240 = 3
	ConnEstFailureControl_ConnEstFailOffsetValidity_S300 = 4
	ConnEstFailureControl_ConnEstFailOffsetValidity_S420 = 5
	ConnEstFailureControl_ConnEstFailOffsetValidity_S600 = 6
	ConnEstFailureControl_ConnEstFailOffsetValidity_S900 = 7
)

var connEstFailureControlConnEstFailOffsetValidityConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConnEstFailureControl_ConnEstFailOffsetValidity_S30, ConnEstFailureControl_ConnEstFailOffsetValidity_S60, ConnEstFailureControl_ConnEstFailOffsetValidity_S120, ConnEstFailureControl_ConnEstFailOffsetValidity_S240, ConnEstFailureControl_ConnEstFailOffsetValidity_S300, ConnEstFailureControl_ConnEstFailOffsetValidity_S420, ConnEstFailureControl_ConnEstFailOffsetValidity_S600, ConnEstFailureControl_ConnEstFailOffsetValidity_S900},
}

var connEstFailureControlConnEstFailOffsetConstraints = per.Constrained(0, 15)

type ConnEstFailureControl struct {
	ConnEstFailCount          int64
	ConnEstFailOffsetValidity int64
	ConnEstFailOffset         *int64
}

func (ie *ConnEstFailureControl) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(connEstFailureControlConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ConnEstFailOffset != nil}); err != nil {
		return err
	}

	// 2. connEstFailCount: enumerated
	{
		if err := e.EncodeEnumerated(ie.ConnEstFailCount, connEstFailureControlConnEstFailCountConstraints); err != nil {
			return err
		}
	}

	// 3. connEstFailOffsetValidity: enumerated
	{
		if err := e.EncodeEnumerated(ie.ConnEstFailOffsetValidity, connEstFailureControlConnEstFailOffsetValidityConstraints); err != nil {
			return err
		}
	}

	// 4. connEstFailOffset: integer
	{
		if ie.ConnEstFailOffset != nil {
			if err := e.EncodeInteger(*ie.ConnEstFailOffset, connEstFailureControlConnEstFailOffsetConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ConnEstFailureControl) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(connEstFailureControlConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. connEstFailCount: enumerated
	{
		v0, err := d.DecodeEnumerated(connEstFailureControlConnEstFailCountConstraints)
		if err != nil {
			return err
		}
		ie.ConnEstFailCount = v0
	}

	// 3. connEstFailOffsetValidity: enumerated
	{
		v1, err := d.DecodeEnumerated(connEstFailureControlConnEstFailOffsetValidityConstraints)
		if err != nil {
			return err
		}
		ie.ConnEstFailOffsetValidity = v1
	}

	// 4. connEstFailOffset: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(connEstFailureControlConnEstFailOffsetConstraints)
			if err != nil {
				return err
			}
			ie.ConnEstFailOffset = &v
		}
	}

	return nil
}
