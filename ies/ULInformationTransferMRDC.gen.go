// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ULInformationTransferMRDC (line 3673).

var uLInformationTransferMRDCConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var uLInformationTransferMRDCCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "c1"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	ULInformationTransferMRDC_CriticalExtensions_C1                       = 0
	ULInformationTransferMRDC_CriticalExtensions_CriticalExtensionsFuture = 1
)

var uLInformationTransferMRDCCriticalExtensionsC1Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ulInformationTransferMRDC"},
		{Name: "spare3"},
		{Name: "spare2"},
		{Name: "spare1"},
	},
}

const (
	ULInformationTransferMRDC_CriticalExtensions_C1_UlInformationTransferMRDC = 0
	ULInformationTransferMRDC_CriticalExtensions_C1_Spare3                    = 1
	ULInformationTransferMRDC_CriticalExtensions_C1_Spare2                    = 2
	ULInformationTransferMRDC_CriticalExtensions_C1_Spare1                    = 3
)

type ULInformationTransferMRDC struct {
	CriticalExtensions struct {
		Choice int
		C1     *struct {
			Choice                    int
			UlInformationTransferMRDC *ULInformationTransferMRDC_IEs
			Spare3                    *struct{}
			Spare2                    *struct{}
			Spare1                    *struct{}
		}
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *ULInformationTransferMRDC) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLInformationTransferMRDCConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(uLInformationTransferMRDCCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case ULInformationTransferMRDC_CriticalExtensions_C1:
			choiceEnc := e.NewChoiceEncoder(uLInformationTransferMRDCCriticalExtensionsC1Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.CriticalExtensions.C1).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.CriticalExtensions.C1).Choice {
			case ULInformationTransferMRDC_CriticalExtensions_C1_UlInformationTransferMRDC:
				if err := (*ie.CriticalExtensions.C1).UlInformationTransferMRDC.Encode(e); err != nil {
					return err
				}
			case ULInformationTransferMRDC_CriticalExtensions_C1_Spare3:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case ULInformationTransferMRDC_CriticalExtensions_C1_Spare2:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case ULInformationTransferMRDC_CriticalExtensions_C1_Spare1:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			}
		case ULInformationTransferMRDC_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *ULInformationTransferMRDC) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLInformationTransferMRDCConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(uLInformationTransferMRDCCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case ULInformationTransferMRDC_CriticalExtensions_C1:
			ie.CriticalExtensions.C1 = &struct {
				Choice                    int
				UlInformationTransferMRDC *ULInformationTransferMRDC_IEs
				Spare3                    *struct{}
				Spare2                    *struct{}
				Spare1                    *struct{}
			}{}
			choiceDec := d.NewChoiceDecoder(uLInformationTransferMRDCCriticalExtensionsC1Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.CriticalExtensions.C1).Choice = int(index)
			switch index {
			case ULInformationTransferMRDC_CriticalExtensions_C1_UlInformationTransferMRDC:
				(*ie.CriticalExtensions.C1).UlInformationTransferMRDC = new(ULInformationTransferMRDC_IEs)
				if err := (*ie.CriticalExtensions.C1).UlInformationTransferMRDC.Decode(d); err != nil {
					return err
				}
			case ULInformationTransferMRDC_CriticalExtensions_C1_Spare3:
				(*ie.CriticalExtensions.C1).Spare3 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case ULInformationTransferMRDC_CriticalExtensions_C1_Spare2:
				(*ie.CriticalExtensions.C1).Spare2 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case ULInformationTransferMRDC_CriticalExtensions_C1_Spare1:
				(*ie.CriticalExtensions.C1).Spare1 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			}
		case ULInformationTransferMRDC_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
