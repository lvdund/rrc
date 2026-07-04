// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DLInformationTransferMRDC-r16 (line 377).

var dLInformationTransferMRDCR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var dLInformationTransferMRDC_r16CriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "c1"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	DLInformationTransferMRDC_r16_CriticalExtensions_C1                       = 0
	DLInformationTransferMRDC_r16_CriticalExtensions_CriticalExtensionsFuture = 1
)

var dLInformationTransferMRDCR16CriticalExtensionsC1Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "dlInformationTransferMRDC-r16"},
		{Name: "spare3"},
		{Name: "spare2"},
		{Name: "spare1"},
	},
}

const (
	DLInformationTransferMRDC_r16_CriticalExtensions_C1_DlInformationTransferMRDC_r16 = 0
	DLInformationTransferMRDC_r16_CriticalExtensions_C1_Spare3                        = 1
	DLInformationTransferMRDC_r16_CriticalExtensions_C1_Spare2                        = 2
	DLInformationTransferMRDC_r16_CriticalExtensions_C1_Spare1                        = 3
)

type DLInformationTransferMRDC_r16 struct {
	CriticalExtensions struct {
		Choice int
		C1     *struct {
			Choice                        int
			DlInformationTransferMRDC_r16 *DLInformationTransferMRDC_r16_IEs
			Spare3                        *struct{}
			Spare2                        *struct{}
			Spare1                        *struct{}
		}
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *DLInformationTransferMRDC_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dLInformationTransferMRDCR16Constraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(dLInformationTransferMRDC_r16CriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case DLInformationTransferMRDC_r16_CriticalExtensions_C1:
			choiceEnc := e.NewChoiceEncoder(dLInformationTransferMRDCR16CriticalExtensionsC1Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.CriticalExtensions.C1).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.CriticalExtensions.C1).Choice {
			case DLInformationTransferMRDC_r16_CriticalExtensions_C1_DlInformationTransferMRDC_r16:
				if err := (*ie.CriticalExtensions.C1).DlInformationTransferMRDC_r16.Encode(e); err != nil {
					return err
				}
			case DLInformationTransferMRDC_r16_CriticalExtensions_C1_Spare3:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case DLInformationTransferMRDC_r16_CriticalExtensions_C1_Spare2:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case DLInformationTransferMRDC_r16_CriticalExtensions_C1_Spare1:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			}
		case DLInformationTransferMRDC_r16_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *DLInformationTransferMRDC_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dLInformationTransferMRDCR16Constraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(dLInformationTransferMRDC_r16CriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case DLInformationTransferMRDC_r16_CriticalExtensions_C1:
			ie.CriticalExtensions.C1 = &struct {
				Choice                        int
				DlInformationTransferMRDC_r16 *DLInformationTransferMRDC_r16_IEs
				Spare3                        *struct{}
				Spare2                        *struct{}
				Spare1                        *struct{}
			}{}
			choiceDec := d.NewChoiceDecoder(dLInformationTransferMRDCR16CriticalExtensionsC1Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.CriticalExtensions.C1).Choice = int(index)
			switch index {
			case DLInformationTransferMRDC_r16_CriticalExtensions_C1_DlInformationTransferMRDC_r16:
				(*ie.CriticalExtensions.C1).DlInformationTransferMRDC_r16 = new(DLInformationTransferMRDC_r16_IEs)
				if err := (*ie.CriticalExtensions.C1).DlInformationTransferMRDC_r16.Decode(d); err != nil {
					return err
				}
			case DLInformationTransferMRDC_r16_CriticalExtensions_C1_Spare3:
				(*ie.CriticalExtensions.C1).Spare3 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case DLInformationTransferMRDC_r16_CriticalExtensions_C1_Spare2:
				(*ie.CriticalExtensions.C1).Spare2 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case DLInformationTransferMRDC_r16_CriticalExtensions_C1_Spare1:
				(*ie.CriticalExtensions.C1).Spare1 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			}
		case DLInformationTransferMRDC_r16_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
