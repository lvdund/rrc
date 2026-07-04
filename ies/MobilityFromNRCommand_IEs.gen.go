// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MobilityFromNRCommand-IEs (line 820).

var mobilityFromNRCommandIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "targetRAT-Type"},
		{Name: "targetRAT-MessageContainer"},
		{Name: "nas-SecurityParamFromNR", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	MobilityFromNRCommand_IEs_TargetRAT_Type_Eutra          = 0
	MobilityFromNRCommand_IEs_TargetRAT_Type_Utra_Fdd_v1610 = 1
	MobilityFromNRCommand_IEs_TargetRAT_Type_Spare2         = 2
	MobilityFromNRCommand_IEs_TargetRAT_Type_Spare1         = 3
)

var mobilityFromNRCommandIEsTargetRATTypeConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{MobilityFromNRCommand_IEs_TargetRAT_Type_Eutra, MobilityFromNRCommand_IEs_TargetRAT_Type_Utra_Fdd_v1610, MobilityFromNRCommand_IEs_TargetRAT_Type_Spare2, MobilityFromNRCommand_IEs_TargetRAT_Type_Spare1},
}

var mobilityFromNRCommandIEsTargetRATMessageContainerConstraints = per.SizeConstraints{}

var mobilityFromNRCommandIEsNasSecurityParamFromNRConstraints = per.SizeConstraints{}

var mobilityFromNRCommandIEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type MobilityFromNRCommand_IEs struct {
	TargetRAT_Type             int64
	TargetRAT_MessageContainer []byte
	Nas_SecurityParamFromNR    []byte
	LateNonCriticalExtension   []byte
	NonCriticalExtension       *MobilityFromNRCommand_v1610_IEs
}

func (ie *MobilityFromNRCommand_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mobilityFromNRCommandIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Nas_SecurityParamFromNR != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. targetRAT-Type: enumerated
	{
		if err := e.EncodeEnumerated(ie.TargetRAT_Type, mobilityFromNRCommandIEsTargetRATTypeConstraints); err != nil {
			return err
		}
	}

	// 3. targetRAT-MessageContainer: octet-string
	{
		if err := e.EncodeOctetString(ie.TargetRAT_MessageContainer, mobilityFromNRCommandIEsTargetRATMessageContainerConstraints); err != nil {
			return err
		}
	}

	// 4. nas-SecurityParamFromNR: octet-string
	{
		if ie.Nas_SecurityParamFromNR != nil {
			if err := e.EncodeOctetString(ie.Nas_SecurityParamFromNR, mobilityFromNRCommandIEsNasSecurityParamFromNRConstraints); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, mobilityFromNRCommandIEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 6. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MobilityFromNRCommand_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mobilityFromNRCommandIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. targetRAT-Type: enumerated
	{
		v0, err := d.DecodeEnumerated(mobilityFromNRCommandIEsTargetRATTypeConstraints)
		if err != nil {
			return err
		}
		ie.TargetRAT_Type = v0
	}

	// 3. targetRAT-MessageContainer: octet-string
	{
		v1, err := d.DecodeOctetString(mobilityFromNRCommandIEsTargetRATMessageContainerConstraints)
		if err != nil {
			return err
		}
		ie.TargetRAT_MessageContainer = v1
	}

	// 4. nas-SecurityParamFromNR: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(mobilityFromNRCommandIEsNasSecurityParamFromNRConstraints)
			if err != nil {
				return err
			}
			ie.Nas_SecurityParamFromNR = v
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeOctetString(mobilityFromNRCommandIEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 6. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(4) {
			ie.NonCriticalExtension = new(MobilityFromNRCommand_v1610_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
