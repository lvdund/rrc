package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UL_DCCH_MessageType_C1_Choice_nothing uint64 = iota
	UL_DCCH_MessageType_C1_Choice_MeasurementReport
	UL_DCCH_MessageType_C1_Choice_RrcReconfigurationComplete
	UL_DCCH_MessageType_C1_Choice_RrcSetupComplete
	UL_DCCH_MessageType_C1_Choice_RrcReestablishmentComplete
	UL_DCCH_MessageType_C1_Choice_RrcResumeComplete
	UL_DCCH_MessageType_C1_Choice_SecurityModeComplete
	UL_DCCH_MessageType_C1_Choice_SecurityModeFailure
	UL_DCCH_MessageType_C1_Choice_UlInformationTransfer
	UL_DCCH_MessageType_C1_Choice_LocationMeasurementIndication
	UL_DCCH_MessageType_C1_Choice_UeCapabilityInformation
	UL_DCCH_MessageType_C1_Choice_CounterCheckResponse
	UL_DCCH_MessageType_C1_Choice_UeAssistanceInformation
	UL_DCCH_MessageType_C1_Choice_FailureInformation
	UL_DCCH_MessageType_C1_Choice_UlInformationTransferMRDC
	UL_DCCH_MessageType_C1_Choice_ScgFailureInformation
	UL_DCCH_MessageType_C1_Choice_ScgFailureInformationEUTRA
)

type UL_DCCH_MessageType_C1 struct {
	Choice                        uint64
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

func (ie *UL_DCCH_MessageType_C1) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 16, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_DCCH_MessageType_C1_Choice_MeasurementReport:
		if err = ie.MeasurementReport.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasurementReport", err)
		}
	case UL_DCCH_MessageType_C1_Choice_RrcReconfigurationComplete:
		if err = ie.RrcReconfigurationComplete.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcReconfigurationComplete", err)
		}
	case UL_DCCH_MessageType_C1_Choice_RrcSetupComplete:
		if err = ie.RrcSetupComplete.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcSetupComplete", err)
		}
	case UL_DCCH_MessageType_C1_Choice_RrcReestablishmentComplete:
		if err = ie.RrcReestablishmentComplete.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcReestablishmentComplete", err)
		}
	case UL_DCCH_MessageType_C1_Choice_RrcResumeComplete:
		if err = ie.RrcResumeComplete.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcResumeComplete", err)
		}
	case UL_DCCH_MessageType_C1_Choice_SecurityModeComplete:
		if err = ie.SecurityModeComplete.Encode(w); err != nil {
			err = utils.WrapError("Encode SecurityModeComplete", err)
		}
	case UL_DCCH_MessageType_C1_Choice_SecurityModeFailure:
		if err = ie.SecurityModeFailure.Encode(w); err != nil {
			err = utils.WrapError("Encode SecurityModeFailure", err)
		}
	case UL_DCCH_MessageType_C1_Choice_UlInformationTransfer:
		if err = ie.UlInformationTransfer.Encode(w); err != nil {
			err = utils.WrapError("Encode UlInformationTransfer", err)
		}
	case UL_DCCH_MessageType_C1_Choice_LocationMeasurementIndication:
		if err = ie.LocationMeasurementIndication.Encode(w); err != nil {
			err = utils.WrapError("Encode LocationMeasurementIndication", err)
		}
	case UL_DCCH_MessageType_C1_Choice_UeCapabilityInformation:
		if err = ie.UeCapabilityInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode UeCapabilityInformation", err)
		}
	case UL_DCCH_MessageType_C1_Choice_CounterCheckResponse:
		if err = ie.CounterCheckResponse.Encode(w); err != nil {
			err = utils.WrapError("Encode CounterCheckResponse", err)
		}
	case UL_DCCH_MessageType_C1_Choice_UeAssistanceInformation:
		if err = ie.UeAssistanceInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode UeAssistanceInformation", err)
		}
	case UL_DCCH_MessageType_C1_Choice_FailureInformation:
		if err = ie.FailureInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode FailureInformation", err)
		}
	case UL_DCCH_MessageType_C1_Choice_UlInformationTransferMRDC:
		if err = ie.UlInformationTransferMRDC.Encode(w); err != nil {
			err = utils.WrapError("Encode UlInformationTransferMRDC", err)
		}
	case UL_DCCH_MessageType_C1_Choice_ScgFailureInformation:
		if err = ie.ScgFailureInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode ScgFailureInformation", err)
		}
	case UL_DCCH_MessageType_C1_Choice_ScgFailureInformationEUTRA:
		if err = ie.ScgFailureInformationEUTRA.Encode(w); err != nil {
			err = utils.WrapError("Encode ScgFailureInformationEUTRA", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UL_DCCH_MessageType_C1) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(16, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_DCCH_MessageType_C1_Choice_MeasurementReport:
		ie.MeasurementReport = new(MeasurementReport)
		if err = ie.MeasurementReport.Decode(r); err != nil {
			return utils.WrapError("Decode MeasurementReport", err)
		}
	case UL_DCCH_MessageType_C1_Choice_RrcReconfigurationComplete:
		ie.RrcReconfigurationComplete = new(RRCReconfigurationComplete)
		if err = ie.RrcReconfigurationComplete.Decode(r); err != nil {
			return utils.WrapError("Decode RrcReconfigurationComplete", err)
		}
	case UL_DCCH_MessageType_C1_Choice_RrcSetupComplete:
		ie.RrcSetupComplete = new(RRCSetupComplete)
		if err = ie.RrcSetupComplete.Decode(r); err != nil {
			return utils.WrapError("Decode RrcSetupComplete", err)
		}
	case UL_DCCH_MessageType_C1_Choice_RrcReestablishmentComplete:
		ie.RrcReestablishmentComplete = new(RRCReestablishmentComplete)
		if err = ie.RrcReestablishmentComplete.Decode(r); err != nil {
			return utils.WrapError("Decode RrcReestablishmentComplete", err)
		}
	case UL_DCCH_MessageType_C1_Choice_RrcResumeComplete:
		ie.RrcResumeComplete = new(RRCResumeComplete)
		if err = ie.RrcResumeComplete.Decode(r); err != nil {
			return utils.WrapError("Decode RrcResumeComplete", err)
		}
	case UL_DCCH_MessageType_C1_Choice_SecurityModeComplete:
		ie.SecurityModeComplete = new(SecurityModeComplete)
		if err = ie.SecurityModeComplete.Decode(r); err != nil {
			return utils.WrapError("Decode SecurityModeComplete", err)
		}
	case UL_DCCH_MessageType_C1_Choice_SecurityModeFailure:
		ie.SecurityModeFailure = new(SecurityModeFailure)
		if err = ie.SecurityModeFailure.Decode(r); err != nil {
			return utils.WrapError("Decode SecurityModeFailure", err)
		}
	case UL_DCCH_MessageType_C1_Choice_UlInformationTransfer:
		ie.UlInformationTransfer = new(ULInformationTransfer)
		if err = ie.UlInformationTransfer.Decode(r); err != nil {
			return utils.WrapError("Decode UlInformationTransfer", err)
		}
	case UL_DCCH_MessageType_C1_Choice_LocationMeasurementIndication:
		ie.LocationMeasurementIndication = new(LocationMeasurementIndication)
		if err = ie.LocationMeasurementIndication.Decode(r); err != nil {
			return utils.WrapError("Decode LocationMeasurementIndication", err)
		}
	case UL_DCCH_MessageType_C1_Choice_UeCapabilityInformation:
		ie.UeCapabilityInformation = new(UECapabilityInformation)
		if err = ie.UeCapabilityInformation.Decode(r); err != nil {
			return utils.WrapError("Decode UeCapabilityInformation", err)
		}
	case UL_DCCH_MessageType_C1_Choice_CounterCheckResponse:
		ie.CounterCheckResponse = new(CounterCheckResponse)
		if err = ie.CounterCheckResponse.Decode(r); err != nil {
			return utils.WrapError("Decode CounterCheckResponse", err)
		}
	case UL_DCCH_MessageType_C1_Choice_UeAssistanceInformation:
		ie.UeAssistanceInformation = new(UEAssistanceInformation)
		if err = ie.UeAssistanceInformation.Decode(r); err != nil {
			return utils.WrapError("Decode UeAssistanceInformation", err)
		}
	case UL_DCCH_MessageType_C1_Choice_FailureInformation:
		ie.FailureInformation = new(FailureInformation)
		if err = ie.FailureInformation.Decode(r); err != nil {
			return utils.WrapError("Decode FailureInformation", err)
		}
	case UL_DCCH_MessageType_C1_Choice_UlInformationTransferMRDC:
		ie.UlInformationTransferMRDC = new(ULInformationTransferMRDC)
		if err = ie.UlInformationTransferMRDC.Decode(r); err != nil {
			return utils.WrapError("Decode UlInformationTransferMRDC", err)
		}
	case UL_DCCH_MessageType_C1_Choice_ScgFailureInformation:
		ie.ScgFailureInformation = new(SCGFailureInformation)
		if err = ie.ScgFailureInformation.Decode(r); err != nil {
			return utils.WrapError("Decode ScgFailureInformation", err)
		}
	case UL_DCCH_MessageType_C1_Choice_ScgFailureInformationEUTRA:
		ie.ScgFailureInformationEUTRA = new(SCGFailureInformationEUTRA)
		if err = ie.ScgFailureInformationEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode ScgFailureInformationEUTRA", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
