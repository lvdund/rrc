// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LoggedEventTriggerConfig-r16 (line 583).

var loggedEventTriggerConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "eventType-r16"},
		{Name: "loggingInterval-r16"},
	},
}

type LoggedEventTriggerConfig_r16 struct {
	EventType_r16       EventType_r16
	LoggingInterval_r16 LoggingInterval_r16
}

func (ie *LoggedEventTriggerConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(loggedEventTriggerConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. eventType-r16: ref
	{
		if err := ie.EventType_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. loggingInterval-r16: ref
	{
		if err := ie.LoggingInterval_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *LoggedEventTriggerConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(loggedEventTriggerConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. eventType-r16: ref
	{
		if err := ie.EventType_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. loggingInterval-r16: ref
	{
		if err := ie.LoggingInterval_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
