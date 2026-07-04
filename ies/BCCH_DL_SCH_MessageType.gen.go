// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BCCH-DL-SCH-MessageType (line 26).

var bCCHDLSCHMessageTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "c1"},
		{Name: "messageClassExtension"},
	},
}

const (
	BCCH_DL_SCH_MessageType_C1                    = 0
	BCCH_DL_SCH_MessageType_MessageClassExtension = 1
)

var bCCHDLSCHMessageTypeC1Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "systemInformation"},
		{Name: "systemInformationBlockType1"},
	},
}

const (
	BCCH_DL_SCH_MessageType_C1_SystemInformation           = 0
	BCCH_DL_SCH_MessageType_C1_SystemInformationBlockType1 = 1
)

type BCCH_DL_SCH_MessageType struct {
	Choice int
	C1     *struct {
		Choice                      int
		SystemInformation           *SystemInformation
		SystemInformationBlockType1 *SIB1
	}
	MessageClassExtension *struct{}
}

func (ie *BCCH_DL_SCH_MessageType) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(bCCHDLSCHMessageTypeConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case BCCH_DL_SCH_MessageType_C1:
		choiceEnc := e.NewChoiceEncoder(bCCHDLSCHMessageTypeC1Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.C1).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.C1).Choice {
		case BCCH_DL_SCH_MessageType_C1_SystemInformation:
			if err := (*ie.C1).SystemInformation.Encode(e); err != nil {
				return err
			}
		case BCCH_DL_SCH_MessageType_C1_SystemInformationBlockType1:
			if err := (*ie.C1).SystemInformationBlockType1.Encode(e); err != nil {
				return err
			}
		}
	case BCCH_DL_SCH_MessageType_MessageClassExtension:
		// empty SEQUENCE alternative
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "BCCH-DL-SCH-MessageType",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *BCCH_DL_SCH_MessageType) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(bCCHDLSCHMessageTypeConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "BCCH-DL-SCH-MessageType", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case BCCH_DL_SCH_MessageType_C1:
		ie.C1 = &struct {
			Choice                      int
			SystemInformation           *SystemInformation
			SystemInformationBlockType1 *SIB1
		}{}
		choiceDec := d.NewChoiceDecoder(bCCHDLSCHMessageTypeC1Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.C1).Choice = int(index)
		switch index {
		case BCCH_DL_SCH_MessageType_C1_SystemInformation:
			(*ie.C1).SystemInformation = new(SystemInformation)
			if err := (*ie.C1).SystemInformation.Decode(d); err != nil {
				return err
			}
		case BCCH_DL_SCH_MessageType_C1_SystemInformationBlockType1:
			(*ie.C1).SystemInformationBlockType1 = new(SIB1)
			if err := (*ie.C1).SystemInformationBlockType1.Decode(d); err != nil {
				return err
			}
		}
	case BCCH_DL_SCH_MessageType_MessageClassExtension:
		ie.MessageClassExtension = &struct{}{}
		// empty SEQUENCE alternative
	default:
		return &per.DecodeError{TypeName: "BCCH-DL-SCH-MessageType", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
