// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UEAssistanceInformation-IEs (line 2397).

var uEAssistanceInformationIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "delayBudgetReport", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var uEAssistanceInformationIEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type UEAssistanceInformation_IEs struct {
	DelayBudgetReport        *DelayBudgetReport
	LateNonCriticalExtension []byte
	NonCriticalExtension     *UEAssistanceInformation_v1540_IEs
}

func (ie *UEAssistanceInformation_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEAssistanceInformationIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DelayBudgetReport != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. delayBudgetReport: ref
	{
		if ie.DelayBudgetReport != nil {
			if err := ie.DelayBudgetReport.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, uEAssistanceInformationIEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UEAssistanceInformation_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEAssistanceInformationIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. delayBudgetReport: ref
	{
		if seq.IsComponentPresent(0) {
			ie.DelayBudgetReport = new(DelayBudgetReport)
			if err := ie.DelayBudgetReport.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(uEAssistanceInformationIEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = new(UEAssistanceInformation_v1540_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
