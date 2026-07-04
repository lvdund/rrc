// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UL-DCCH-MessageType (line 166).

var uLDCCHMessageTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "c1"},
		{Name: "messageClassExtension"},
	},
}

const (
	UL_DCCH_MessageType_C1                    = 0
	UL_DCCH_MessageType_MessageClassExtension = 1
)

var uLDCCHMessageTypeC1Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "measurementReport"},
		{Name: "rrcReconfigurationComplete"},
		{Name: "rrcSetupComplete"},
		{Name: "rrcReestablishmentComplete"},
		{Name: "rrcResumeComplete"},
		{Name: "securityModeComplete"},
		{Name: "securityModeFailure"},
		{Name: "ulInformationTransfer"},
		{Name: "locationMeasurementIndication"},
		{Name: "ueCapabilityInformation"},
		{Name: "counterCheckResponse"},
		{Name: "ueAssistanceInformation"},
		{Name: "failureInformation"},
		{Name: "ulInformationTransferMRDC"},
		{Name: "scgFailureInformation"},
		{Name: "scgFailureInformationEUTRA"},
	},
}

const (
	UL_DCCH_MessageType_C1_MeasurementReport             = 0
	UL_DCCH_MessageType_C1_RrcReconfigurationComplete    = 1
	UL_DCCH_MessageType_C1_RrcSetupComplete              = 2
	UL_DCCH_MessageType_C1_RrcReestablishmentComplete    = 3
	UL_DCCH_MessageType_C1_RrcResumeComplete             = 4
	UL_DCCH_MessageType_C1_SecurityModeComplete          = 5
	UL_DCCH_MessageType_C1_SecurityModeFailure           = 6
	UL_DCCH_MessageType_C1_UlInformationTransfer         = 7
	UL_DCCH_MessageType_C1_LocationMeasurementIndication = 8
	UL_DCCH_MessageType_C1_UeCapabilityInformation       = 9
	UL_DCCH_MessageType_C1_CounterCheckResponse          = 10
	UL_DCCH_MessageType_C1_UeAssistanceInformation       = 11
	UL_DCCH_MessageType_C1_FailureInformation            = 12
	UL_DCCH_MessageType_C1_UlInformationTransferMRDC     = 13
	UL_DCCH_MessageType_C1_ScgFailureInformation         = 14
	UL_DCCH_MessageType_C1_ScgFailureInformationEUTRA    = 15
)

var uLDCCHMessageTypeMessageClassExtensionConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "c2"},
		{Name: "messageClassExtensionFuture-r16"},
	},
}

const (
	UL_DCCH_MessageType_MessageClassExtension_C2                              = 0
	UL_DCCH_MessageType_MessageClassExtension_MessageClassExtensionFuture_r16 = 1
)

var uLDCCHMessageTypeMessageClassExtensionC2Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ulDedicatedMessageSegment-r16"},
		{Name: "dedicatedSIBRequest-r16"},
		{Name: "mcgFailureInformation-r16"},
		{Name: "ueInformationResponse-r16"},
		{Name: "sidelinkUEInformationNR-r16"},
		{Name: "ulInformationTransferIRAT-r16"},
		{Name: "iabOtherInformation-r16"},
		{Name: "mbsInterestIndication-r17"},
		{Name: "uePositioningAssistanceInfo-r17"},
		{Name: "measurementReportAppLayer-r17"},
		{Name: "indirectPathFailureInformation-r18"},
		{Name: "spare5"},
		{Name: "spare4"},
		{Name: "spare3"},
		{Name: "spare2"},
		{Name: "spare1"},
	},
}

const (
	UL_DCCH_MessageType_MessageClassExtension_C2_UlDedicatedMessageSegment_r16      = 0
	UL_DCCH_MessageType_MessageClassExtension_C2_DedicatedSIBRequest_r16            = 1
	UL_DCCH_MessageType_MessageClassExtension_C2_McgFailureInformation_r16          = 2
	UL_DCCH_MessageType_MessageClassExtension_C2_UeInformationResponse_r16          = 3
	UL_DCCH_MessageType_MessageClassExtension_C2_SidelinkUEInformationNR_r16        = 4
	UL_DCCH_MessageType_MessageClassExtension_C2_UlInformationTransferIRAT_r16      = 5
	UL_DCCH_MessageType_MessageClassExtension_C2_IabOtherInformation_r16            = 6
	UL_DCCH_MessageType_MessageClassExtension_C2_MbsInterestIndication_r17          = 7
	UL_DCCH_MessageType_MessageClassExtension_C2_UePositioningAssistanceInfo_r17    = 8
	UL_DCCH_MessageType_MessageClassExtension_C2_MeasurementReportAppLayer_r17      = 9
	UL_DCCH_MessageType_MessageClassExtension_C2_IndirectPathFailureInformation_r18 = 10
	UL_DCCH_MessageType_MessageClassExtension_C2_Spare5                             = 11
	UL_DCCH_MessageType_MessageClassExtension_C2_Spare4                             = 12
	UL_DCCH_MessageType_MessageClassExtension_C2_Spare3                             = 13
	UL_DCCH_MessageType_MessageClassExtension_C2_Spare2                             = 14
	UL_DCCH_MessageType_MessageClassExtension_C2_Spare1                             = 15
)

type UL_DCCH_MessageType struct {
	Choice int
	C1     *struct {
		Choice                        int
		MeasurementReport             *MeasurementReport
		RrcReconfigurationComplete    *RRCReconfigurationComplete
		RrcSetupComplete              *RRCSetupComplete
		RrcReestablishmentComplete    *RRCReestablishmentComplete
		RrcResumeComplete             *RRCResumeComplete
		SecurityModeComplete          *SecurityModeComplete
		SecurityModeFailure           *SecurityModeFailure
		UlInformationTransfer         *ULInformationTransfer
		LocationMeasurementIndication *LocationMeasurementIndication
		UeCapabilityInformation       *UECapabilityInformation
		CounterCheckResponse          *CounterCheckResponse
		UeAssistanceInformation       *UEAssistanceInformation
		FailureInformation            *FailureInformation
		UlInformationTransferMRDC     *ULInformationTransferMRDC
		ScgFailureInformation         *SCGFailureInformation
		ScgFailureInformationEUTRA    *SCGFailureInformationEUTRA
	}
	MessageClassExtension *struct {
		Choice int
		C2     *struct {
			Choice                             int
			UlDedicatedMessageSegment_r16      *ULDedicatedMessageSegment_r16
			DedicatedSIBRequest_r16            *DedicatedSIBRequest_r16
			McgFailureInformation_r16          *MCGFailureInformation_r16
			UeInformationResponse_r16          *UEInformationResponse_r16
			SidelinkUEInformationNR_r16        *SidelinkUEInformationNR_r16
			UlInformationTransferIRAT_r16      *ULInformationTransferIRAT_r16
			IabOtherInformation_r16            *IABOtherInformation_r16
			MbsInterestIndication_r17          *MBSInterestIndication_r17
			UePositioningAssistanceInfo_r17    *UEPositioningAssistanceInfo_r17
			MeasurementReportAppLayer_r17      *MeasurementReportAppLayer_r17
			IndirectPathFailureInformation_r18 *IndirectPathFailureInformation_r18
			Spare5                             *struct{}
			Spare4                             *struct{}
			Spare3                             *struct{}
			Spare2                             *struct{}
			Spare1                             *struct{}
		}
		MessageClassExtensionFuture_r16 *struct{}
	}
}

func (ie *UL_DCCH_MessageType) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(uLDCCHMessageTypeConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_DCCH_MessageType_C1:
		choiceEnc := e.NewChoiceEncoder(uLDCCHMessageTypeC1Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.C1).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.C1).Choice {
		case UL_DCCH_MessageType_C1_MeasurementReport:
			if err := (*ie.C1).MeasurementReport.Encode(e); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_RrcReconfigurationComplete:
			if err := (*ie.C1).RrcReconfigurationComplete.Encode(e); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_RrcSetupComplete:
			if err := (*ie.C1).RrcSetupComplete.Encode(e); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_RrcReestablishmentComplete:
			if err := (*ie.C1).RrcReestablishmentComplete.Encode(e); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_RrcResumeComplete:
			if err := (*ie.C1).RrcResumeComplete.Encode(e); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_SecurityModeComplete:
			if err := (*ie.C1).SecurityModeComplete.Encode(e); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_SecurityModeFailure:
			if err := (*ie.C1).SecurityModeFailure.Encode(e); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_UlInformationTransfer:
			if err := (*ie.C1).UlInformationTransfer.Encode(e); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_LocationMeasurementIndication:
			if err := (*ie.C1).LocationMeasurementIndication.Encode(e); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_UeCapabilityInformation:
			if err := (*ie.C1).UeCapabilityInformation.Encode(e); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_CounterCheckResponse:
			if err := (*ie.C1).CounterCheckResponse.Encode(e); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_UeAssistanceInformation:
			if err := (*ie.C1).UeAssistanceInformation.Encode(e); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_FailureInformation:
			if err := (*ie.C1).FailureInformation.Encode(e); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_UlInformationTransferMRDC:
			if err := (*ie.C1).UlInformationTransferMRDC.Encode(e); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_ScgFailureInformation:
			if err := (*ie.C1).ScgFailureInformation.Encode(e); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_ScgFailureInformationEUTRA:
			if err := (*ie.C1).ScgFailureInformationEUTRA.Encode(e); err != nil {
				return err
			}
		}
	case UL_DCCH_MessageType_MessageClassExtension:
		choiceEnc := e.NewChoiceEncoder(uLDCCHMessageTypeMessageClassExtensionConstraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.MessageClassExtension).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.MessageClassExtension).Choice {
		case UL_DCCH_MessageType_MessageClassExtension_C2:
			choiceEnc := e.NewChoiceEncoder(uLDCCHMessageTypeMessageClassExtensionC2Constraints)
			if err := choiceEnc.EncodeChoice(int64((*(*ie.MessageClassExtension).C2).Choice), false, nil); err != nil {
				return err
			}
			switch (*(*ie.MessageClassExtension).C2).Choice {
			case UL_DCCH_MessageType_MessageClassExtension_C2_UlDedicatedMessageSegment_r16:
				if err := (*(*ie.MessageClassExtension).C2).UlDedicatedMessageSegment_r16.Encode(e); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_DedicatedSIBRequest_r16:
				if err := (*(*ie.MessageClassExtension).C2).DedicatedSIBRequest_r16.Encode(e); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_McgFailureInformation_r16:
				if err := (*(*ie.MessageClassExtension).C2).McgFailureInformation_r16.Encode(e); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_UeInformationResponse_r16:
				if err := (*(*ie.MessageClassExtension).C2).UeInformationResponse_r16.Encode(e); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_SidelinkUEInformationNR_r16:
				if err := (*(*ie.MessageClassExtension).C2).SidelinkUEInformationNR_r16.Encode(e); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_UlInformationTransferIRAT_r16:
				if err := (*(*ie.MessageClassExtension).C2).UlInformationTransferIRAT_r16.Encode(e); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_IabOtherInformation_r16:
				if err := (*(*ie.MessageClassExtension).C2).IabOtherInformation_r16.Encode(e); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_MbsInterestIndication_r17:
				if err := (*(*ie.MessageClassExtension).C2).MbsInterestIndication_r17.Encode(e); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_UePositioningAssistanceInfo_r17:
				if err := (*(*ie.MessageClassExtension).C2).UePositioningAssistanceInfo_r17.Encode(e); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_MeasurementReportAppLayer_r17:
				if err := (*(*ie.MessageClassExtension).C2).MeasurementReportAppLayer_r17.Encode(e); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_IndirectPathFailureInformation_r18:
				if err := (*(*ie.MessageClassExtension).C2).IndirectPathFailureInformation_r18.Encode(e); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_Spare5:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_Spare4:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_Spare3:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_Spare2:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_Spare1:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			}
		case UL_DCCH_MessageType_MessageClassExtension_MessageClassExtensionFuture_r16:
			// empty SEQUENCE alternative
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "UL-DCCH-MessageType",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *UL_DCCH_MessageType) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(uLDCCHMessageTypeConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "UL-DCCH-MessageType", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case UL_DCCH_MessageType_C1:
		ie.C1 = &struct {
			Choice                        int
			MeasurementReport             *MeasurementReport
			RrcReconfigurationComplete    *RRCReconfigurationComplete
			RrcSetupComplete              *RRCSetupComplete
			RrcReestablishmentComplete    *RRCReestablishmentComplete
			RrcResumeComplete             *RRCResumeComplete
			SecurityModeComplete          *SecurityModeComplete
			SecurityModeFailure           *SecurityModeFailure
			UlInformationTransfer         *ULInformationTransfer
			LocationMeasurementIndication *LocationMeasurementIndication
			UeCapabilityInformation       *UECapabilityInformation
			CounterCheckResponse          *CounterCheckResponse
			UeAssistanceInformation       *UEAssistanceInformation
			FailureInformation            *FailureInformation
			UlInformationTransferMRDC     *ULInformationTransferMRDC
			ScgFailureInformation         *SCGFailureInformation
			ScgFailureInformationEUTRA    *SCGFailureInformationEUTRA
		}{}
		choiceDec := d.NewChoiceDecoder(uLDCCHMessageTypeC1Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.C1).Choice = int(index)
		switch index {
		case UL_DCCH_MessageType_C1_MeasurementReport:
			(*ie.C1).MeasurementReport = new(MeasurementReport)
			if err := (*ie.C1).MeasurementReport.Decode(d); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_RrcReconfigurationComplete:
			(*ie.C1).RrcReconfigurationComplete = new(RRCReconfigurationComplete)
			if err := (*ie.C1).RrcReconfigurationComplete.Decode(d); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_RrcSetupComplete:
			(*ie.C1).RrcSetupComplete = new(RRCSetupComplete)
			if err := (*ie.C1).RrcSetupComplete.Decode(d); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_RrcReestablishmentComplete:
			(*ie.C1).RrcReestablishmentComplete = new(RRCReestablishmentComplete)
			if err := (*ie.C1).RrcReestablishmentComplete.Decode(d); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_RrcResumeComplete:
			(*ie.C1).RrcResumeComplete = new(RRCResumeComplete)
			if err := (*ie.C1).RrcResumeComplete.Decode(d); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_SecurityModeComplete:
			(*ie.C1).SecurityModeComplete = new(SecurityModeComplete)
			if err := (*ie.C1).SecurityModeComplete.Decode(d); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_SecurityModeFailure:
			(*ie.C1).SecurityModeFailure = new(SecurityModeFailure)
			if err := (*ie.C1).SecurityModeFailure.Decode(d); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_UlInformationTransfer:
			(*ie.C1).UlInformationTransfer = new(ULInformationTransfer)
			if err := (*ie.C1).UlInformationTransfer.Decode(d); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_LocationMeasurementIndication:
			(*ie.C1).LocationMeasurementIndication = new(LocationMeasurementIndication)
			if err := (*ie.C1).LocationMeasurementIndication.Decode(d); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_UeCapabilityInformation:
			(*ie.C1).UeCapabilityInformation = new(UECapabilityInformation)
			if err := (*ie.C1).UeCapabilityInformation.Decode(d); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_CounterCheckResponse:
			(*ie.C1).CounterCheckResponse = new(CounterCheckResponse)
			if err := (*ie.C1).CounterCheckResponse.Decode(d); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_UeAssistanceInformation:
			(*ie.C1).UeAssistanceInformation = new(UEAssistanceInformation)
			if err := (*ie.C1).UeAssistanceInformation.Decode(d); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_FailureInformation:
			(*ie.C1).FailureInformation = new(FailureInformation)
			if err := (*ie.C1).FailureInformation.Decode(d); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_UlInformationTransferMRDC:
			(*ie.C1).UlInformationTransferMRDC = new(ULInformationTransferMRDC)
			if err := (*ie.C1).UlInformationTransferMRDC.Decode(d); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_ScgFailureInformation:
			(*ie.C1).ScgFailureInformation = new(SCGFailureInformation)
			if err := (*ie.C1).ScgFailureInformation.Decode(d); err != nil {
				return err
			}
		case UL_DCCH_MessageType_C1_ScgFailureInformationEUTRA:
			(*ie.C1).ScgFailureInformationEUTRA = new(SCGFailureInformationEUTRA)
			if err := (*ie.C1).ScgFailureInformationEUTRA.Decode(d); err != nil {
				return err
			}
		}
	case UL_DCCH_MessageType_MessageClassExtension:
		ie.MessageClassExtension = &struct {
			Choice int
			C2     *struct {
				Choice                             int
				UlDedicatedMessageSegment_r16      *ULDedicatedMessageSegment_r16
				DedicatedSIBRequest_r16            *DedicatedSIBRequest_r16
				McgFailureInformation_r16          *MCGFailureInformation_r16
				UeInformationResponse_r16          *UEInformationResponse_r16
				SidelinkUEInformationNR_r16        *SidelinkUEInformationNR_r16
				UlInformationTransferIRAT_r16      *ULInformationTransferIRAT_r16
				IabOtherInformation_r16            *IABOtherInformation_r16
				MbsInterestIndication_r17          *MBSInterestIndication_r17
				UePositioningAssistanceInfo_r17    *UEPositioningAssistanceInfo_r17
				MeasurementReportAppLayer_r17      *MeasurementReportAppLayer_r17
				IndirectPathFailureInformation_r18 *IndirectPathFailureInformation_r18
				Spare5                             *struct{}
				Spare4                             *struct{}
				Spare3                             *struct{}
				Spare2                             *struct{}
				Spare1                             *struct{}
			}
			MessageClassExtensionFuture_r16 *struct{}
		}{}
		choiceDec := d.NewChoiceDecoder(uLDCCHMessageTypeMessageClassExtensionConstraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.MessageClassExtension).Choice = int(index)
		switch index {
		case UL_DCCH_MessageType_MessageClassExtension_C2:
			(*ie.MessageClassExtension).C2 = &struct {
				Choice                             int
				UlDedicatedMessageSegment_r16      *ULDedicatedMessageSegment_r16
				DedicatedSIBRequest_r16            *DedicatedSIBRequest_r16
				McgFailureInformation_r16          *MCGFailureInformation_r16
				UeInformationResponse_r16          *UEInformationResponse_r16
				SidelinkUEInformationNR_r16        *SidelinkUEInformationNR_r16
				UlInformationTransferIRAT_r16      *ULInformationTransferIRAT_r16
				IabOtherInformation_r16            *IABOtherInformation_r16
				MbsInterestIndication_r17          *MBSInterestIndication_r17
				UePositioningAssistanceInfo_r17    *UEPositioningAssistanceInfo_r17
				MeasurementReportAppLayer_r17      *MeasurementReportAppLayer_r17
				IndirectPathFailureInformation_r18 *IndirectPathFailureInformation_r18
				Spare5                             *struct{}
				Spare4                             *struct{}
				Spare3                             *struct{}
				Spare2                             *struct{}
				Spare1                             *struct{}
			}{}
			choiceDec := d.NewChoiceDecoder(uLDCCHMessageTypeMessageClassExtensionC2Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*(*ie.MessageClassExtension).C2).Choice = int(index)
			switch index {
			case UL_DCCH_MessageType_MessageClassExtension_C2_UlDedicatedMessageSegment_r16:
				(*(*ie.MessageClassExtension).C2).UlDedicatedMessageSegment_r16 = new(ULDedicatedMessageSegment_r16)
				if err := (*(*ie.MessageClassExtension).C2).UlDedicatedMessageSegment_r16.Decode(d); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_DedicatedSIBRequest_r16:
				(*(*ie.MessageClassExtension).C2).DedicatedSIBRequest_r16 = new(DedicatedSIBRequest_r16)
				if err := (*(*ie.MessageClassExtension).C2).DedicatedSIBRequest_r16.Decode(d); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_McgFailureInformation_r16:
				(*(*ie.MessageClassExtension).C2).McgFailureInformation_r16 = new(MCGFailureInformation_r16)
				if err := (*(*ie.MessageClassExtension).C2).McgFailureInformation_r16.Decode(d); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_UeInformationResponse_r16:
				(*(*ie.MessageClassExtension).C2).UeInformationResponse_r16 = new(UEInformationResponse_r16)
				if err := (*(*ie.MessageClassExtension).C2).UeInformationResponse_r16.Decode(d); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_SidelinkUEInformationNR_r16:
				(*(*ie.MessageClassExtension).C2).SidelinkUEInformationNR_r16 = new(SidelinkUEInformationNR_r16)
				if err := (*(*ie.MessageClassExtension).C2).SidelinkUEInformationNR_r16.Decode(d); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_UlInformationTransferIRAT_r16:
				(*(*ie.MessageClassExtension).C2).UlInformationTransferIRAT_r16 = new(ULInformationTransferIRAT_r16)
				if err := (*(*ie.MessageClassExtension).C2).UlInformationTransferIRAT_r16.Decode(d); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_IabOtherInformation_r16:
				(*(*ie.MessageClassExtension).C2).IabOtherInformation_r16 = new(IABOtherInformation_r16)
				if err := (*(*ie.MessageClassExtension).C2).IabOtherInformation_r16.Decode(d); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_MbsInterestIndication_r17:
				(*(*ie.MessageClassExtension).C2).MbsInterestIndication_r17 = new(MBSInterestIndication_r17)
				if err := (*(*ie.MessageClassExtension).C2).MbsInterestIndication_r17.Decode(d); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_UePositioningAssistanceInfo_r17:
				(*(*ie.MessageClassExtension).C2).UePositioningAssistanceInfo_r17 = new(UEPositioningAssistanceInfo_r17)
				if err := (*(*ie.MessageClassExtension).C2).UePositioningAssistanceInfo_r17.Decode(d); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_MeasurementReportAppLayer_r17:
				(*(*ie.MessageClassExtension).C2).MeasurementReportAppLayer_r17 = new(MeasurementReportAppLayer_r17)
				if err := (*(*ie.MessageClassExtension).C2).MeasurementReportAppLayer_r17.Decode(d); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_IndirectPathFailureInformation_r18:
				(*(*ie.MessageClassExtension).C2).IndirectPathFailureInformation_r18 = new(IndirectPathFailureInformation_r18)
				if err := (*(*ie.MessageClassExtension).C2).IndirectPathFailureInformation_r18.Decode(d); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_Spare5:
				(*(*ie.MessageClassExtension).C2).Spare5 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_Spare4:
				(*(*ie.MessageClassExtension).C2).Spare4 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_Spare3:
				(*(*ie.MessageClassExtension).C2).Spare3 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_Spare2:
				(*(*ie.MessageClassExtension).C2).Spare2 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case UL_DCCH_MessageType_MessageClassExtension_C2_Spare1:
				(*(*ie.MessageClassExtension).C2).Spare1 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			}
		case UL_DCCH_MessageType_MessageClassExtension_MessageClassExtensionFuture_r16:
			(*ie.MessageClassExtension).MessageClassExtensionFuture_r16 = &struct{}{}
			// empty SEQUENCE alternative
		}
	default:
		return &per.DecodeError{TypeName: "UL-DCCH-MessageType", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
