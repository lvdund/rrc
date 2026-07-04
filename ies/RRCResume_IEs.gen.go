// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCResume-IEs (line 1537).

var rRCResumeIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "radioBearerConfig", Optional: true},
		{Name: "masterCellGroup", Optional: true},
		{Name: "measConfig", Optional: true},
		{Name: "fullConfig", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var rRCResumeIEsMasterCellGroupConstraints = per.SizeConstraints{}

const (
	RRCResume_IEs_FullConfig_True = 0
)

var rRCResumeIEsFullConfigConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCResume_IEs_FullConfig_True},
}

var rRCResumeIEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type RRCResume_IEs struct {
	RadioBearerConfig        *RadioBearerConfig
	MasterCellGroup          []byte
	MeasConfig               *MeasConfig
	FullConfig               *int64
	LateNonCriticalExtension []byte
	NonCriticalExtension     *RRCResume_v1560_IEs
}

func (ie *RRCResume_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCResumeIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RadioBearerConfig != nil, ie.MasterCellGroup != nil, ie.MeasConfig != nil, ie.FullConfig != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. radioBearerConfig: ref
	{
		if ie.RadioBearerConfig != nil {
			if err := ie.RadioBearerConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. masterCellGroup: octet-string
	{
		if ie.MasterCellGroup != nil {
			if err := e.EncodeOctetString(ie.MasterCellGroup, rRCResumeIEsMasterCellGroupConstraints); err != nil {
				return err
			}
		}
	}

	// 4. measConfig: ref
	{
		if ie.MeasConfig != nil {
			if err := ie.MeasConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. fullConfig: enumerated
	{
		if ie.FullConfig != nil {
			if err := e.EncodeEnumerated(*ie.FullConfig, rRCResumeIEsFullConfigConstraints); err != nil {
				return err
			}
		}
	}

	// 6. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, rRCResumeIEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 7. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCResume_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCResumeIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. radioBearerConfig: ref
	{
		if seq.IsComponentPresent(0) {
			ie.RadioBearerConfig = new(RadioBearerConfig)
			if err := ie.RadioBearerConfig.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. masterCellGroup: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(rRCResumeIEsMasterCellGroupConstraints)
			if err != nil {
				return err
			}
			ie.MasterCellGroup = v
		}
	}

	// 4. measConfig: ref
	{
		if seq.IsComponentPresent(2) {
			ie.MeasConfig = new(MeasConfig)
			if err := ie.MeasConfig.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. fullConfig: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(rRCResumeIEsFullConfigConstraints)
			if err != nil {
				return err
			}
			ie.FullConfig = &idx
		}
	}

	// 6. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeOctetString(rRCResumeIEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 7. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(5) {
			ie.NonCriticalExtension = new(RRCResume_v1560_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
