// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DL-DCCH-MessageType (line 58).

var dLDCCHMessageTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "c1"},
		{Name: "messageClassExtension"},
	},
}

const (
	DL_DCCH_MessageType_C1                    = 0
	DL_DCCH_MessageType_MessageClassExtension = 1
)

var dLDCCHMessageTypeC1Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rrcReconfiguration"},
		{Name: "rrcResume"},
		{Name: "rrcRelease"},
		{Name: "rrcReestablishment"},
		{Name: "securityModeCommand"},
		{Name: "dlInformationTransfer"},
		{Name: "ueCapabilityEnquiry"},
		{Name: "counterCheck"},
		{Name: "mobilityFromNRCommand"},
		{Name: "dlDedicatedMessageSegment-r16"},
		{Name: "ueInformationRequest-r16"},
		{Name: "dlInformationTransferMRDC-r16"},
		{Name: "loggedMeasurementConfiguration-r16"},
		{Name: "spare3"},
		{Name: "spare2"},
		{Name: "spare1"},
	},
}

const (
	DL_DCCH_MessageType_C1_RrcReconfiguration                 = 0
	DL_DCCH_MessageType_C1_RrcResume                          = 1
	DL_DCCH_MessageType_C1_RrcRelease                         = 2
	DL_DCCH_MessageType_C1_RrcReestablishment                 = 3
	DL_DCCH_MessageType_C1_SecurityModeCommand                = 4
	DL_DCCH_MessageType_C1_DlInformationTransfer              = 5
	DL_DCCH_MessageType_C1_UeCapabilityEnquiry                = 6
	DL_DCCH_MessageType_C1_CounterCheck                       = 7
	DL_DCCH_MessageType_C1_MobilityFromNRCommand              = 8
	DL_DCCH_MessageType_C1_DlDedicatedMessageSegment_r16      = 9
	DL_DCCH_MessageType_C1_UeInformationRequest_r16           = 10
	DL_DCCH_MessageType_C1_DlInformationTransferMRDC_r16      = 11
	DL_DCCH_MessageType_C1_LoggedMeasurementConfiguration_r16 = 12
	DL_DCCH_MessageType_C1_Spare3                             = 13
	DL_DCCH_MessageType_C1_Spare2                             = 14
	DL_DCCH_MessageType_C1_Spare1                             = 15
)

type DL_DCCH_MessageType struct {
	Choice int
	C1     *struct {
		Choice                             int
		RrcReconfiguration                 *RRCReconfiguration
		RrcResume                          *RRCResume
		RrcRelease                         *RRCRelease
		RrcReestablishment                 *RRCReestablishment
		SecurityModeCommand                *SecurityModeCommand
		DlInformationTransfer              *DLInformationTransfer
		UeCapabilityEnquiry                *UECapabilityEnquiry
		CounterCheck                       *CounterCheck
		MobilityFromNRCommand              *MobilityFromNRCommand
		DlDedicatedMessageSegment_r16      *DLDedicatedMessageSegment_r16
		UeInformationRequest_r16           *UEInformationRequest_r16
		DlInformationTransferMRDC_r16      *DLInformationTransferMRDC_r16
		LoggedMeasurementConfiguration_r16 *LoggedMeasurementConfiguration_r16
		Spare3                             *struct{}
		Spare2                             *struct{}
		Spare1                             *struct{}
	}
	MessageClassExtension *struct{}
}

func (ie *DL_DCCH_MessageType) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(dLDCCHMessageTypeConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_DCCH_MessageType_C1:
		choiceEnc := e.NewChoiceEncoder(dLDCCHMessageTypeC1Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.C1).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.C1).Choice {
		case DL_DCCH_MessageType_C1_RrcReconfiguration:
			if err := (*ie.C1).RrcReconfiguration.Encode(e); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_RrcResume:
			if err := (*ie.C1).RrcResume.Encode(e); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_RrcRelease:
			if err := (*ie.C1).RrcRelease.Encode(e); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_RrcReestablishment:
			if err := (*ie.C1).RrcReestablishment.Encode(e); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_SecurityModeCommand:
			if err := (*ie.C1).SecurityModeCommand.Encode(e); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_DlInformationTransfer:
			if err := (*ie.C1).DlInformationTransfer.Encode(e); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_UeCapabilityEnquiry:
			if err := (*ie.C1).UeCapabilityEnquiry.Encode(e); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_CounterCheck:
			if err := (*ie.C1).CounterCheck.Encode(e); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_MobilityFromNRCommand:
			if err := (*ie.C1).MobilityFromNRCommand.Encode(e); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_DlDedicatedMessageSegment_r16:
			if err := (*ie.C1).DlDedicatedMessageSegment_r16.Encode(e); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_UeInformationRequest_r16:
			if err := (*ie.C1).UeInformationRequest_r16.Encode(e); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_DlInformationTransferMRDC_r16:
			if err := (*ie.C1).DlInformationTransferMRDC_r16.Encode(e); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_LoggedMeasurementConfiguration_r16:
			if err := (*ie.C1).LoggedMeasurementConfiguration_r16.Encode(e); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_Spare3:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_Spare2:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_Spare1:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		}
	case DL_DCCH_MessageType_MessageClassExtension:
		// empty SEQUENCE alternative
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "DL-DCCH-MessageType",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *DL_DCCH_MessageType) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(dLDCCHMessageTypeConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "DL-DCCH-MessageType", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case DL_DCCH_MessageType_C1:
		ie.C1 = &struct {
			Choice                             int
			RrcReconfiguration                 *RRCReconfiguration
			RrcResume                          *RRCResume
			RrcRelease                         *RRCRelease
			RrcReestablishment                 *RRCReestablishment
			SecurityModeCommand                *SecurityModeCommand
			DlInformationTransfer              *DLInformationTransfer
			UeCapabilityEnquiry                *UECapabilityEnquiry
			CounterCheck                       *CounterCheck
			MobilityFromNRCommand              *MobilityFromNRCommand
			DlDedicatedMessageSegment_r16      *DLDedicatedMessageSegment_r16
			UeInformationRequest_r16           *UEInformationRequest_r16
			DlInformationTransferMRDC_r16      *DLInformationTransferMRDC_r16
			LoggedMeasurementConfiguration_r16 *LoggedMeasurementConfiguration_r16
			Spare3                             *struct{}
			Spare2                             *struct{}
			Spare1                             *struct{}
		}{}
		choiceDec := d.NewChoiceDecoder(dLDCCHMessageTypeC1Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.C1).Choice = int(index)
		switch index {
		case DL_DCCH_MessageType_C1_RrcReconfiguration:
			(*ie.C1).RrcReconfiguration = new(RRCReconfiguration)
			if err := (*ie.C1).RrcReconfiguration.Decode(d); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_RrcResume:
			(*ie.C1).RrcResume = new(RRCResume)
			if err := (*ie.C1).RrcResume.Decode(d); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_RrcRelease:
			(*ie.C1).RrcRelease = new(RRCRelease)
			if err := (*ie.C1).RrcRelease.Decode(d); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_RrcReestablishment:
			(*ie.C1).RrcReestablishment = new(RRCReestablishment)
			if err := (*ie.C1).RrcReestablishment.Decode(d); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_SecurityModeCommand:
			(*ie.C1).SecurityModeCommand = new(SecurityModeCommand)
			if err := (*ie.C1).SecurityModeCommand.Decode(d); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_DlInformationTransfer:
			(*ie.C1).DlInformationTransfer = new(DLInformationTransfer)
			if err := (*ie.C1).DlInformationTransfer.Decode(d); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_UeCapabilityEnquiry:
			(*ie.C1).UeCapabilityEnquiry = new(UECapabilityEnquiry)
			if err := (*ie.C1).UeCapabilityEnquiry.Decode(d); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_CounterCheck:
			(*ie.C1).CounterCheck = new(CounterCheck)
			if err := (*ie.C1).CounterCheck.Decode(d); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_MobilityFromNRCommand:
			(*ie.C1).MobilityFromNRCommand = new(MobilityFromNRCommand)
			if err := (*ie.C1).MobilityFromNRCommand.Decode(d); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_DlDedicatedMessageSegment_r16:
			(*ie.C1).DlDedicatedMessageSegment_r16 = new(DLDedicatedMessageSegment_r16)
			if err := (*ie.C1).DlDedicatedMessageSegment_r16.Decode(d); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_UeInformationRequest_r16:
			(*ie.C1).UeInformationRequest_r16 = new(UEInformationRequest_r16)
			if err := (*ie.C1).UeInformationRequest_r16.Decode(d); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_DlInformationTransferMRDC_r16:
			(*ie.C1).DlInformationTransferMRDC_r16 = new(DLInformationTransferMRDC_r16)
			if err := (*ie.C1).DlInformationTransferMRDC_r16.Decode(d); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_LoggedMeasurementConfiguration_r16:
			(*ie.C1).LoggedMeasurementConfiguration_r16 = new(LoggedMeasurementConfiguration_r16)
			if err := (*ie.C1).LoggedMeasurementConfiguration_r16.Decode(d); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_Spare3:
			(*ie.C1).Spare3 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_Spare2:
			(*ie.C1).Spare2 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case DL_DCCH_MessageType_C1_Spare1:
			(*ie.C1).Spare1 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		}
	case DL_DCCH_MessageType_MessageClassExtension:
		ie.MessageClassExtension = &struct{}{}
		// empty SEQUENCE alternative
	default:
		return &per.DecodeError{TypeName: "DL-DCCH-MessageType", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
