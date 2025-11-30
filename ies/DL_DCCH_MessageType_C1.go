package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DL_DCCH_MessageType_C1_Choice_nothing uint64 = iota
	DL_DCCH_MessageType_C1_Choice_RrcReconfiguration
	DL_DCCH_MessageType_C1_Choice_RrcResume
	DL_DCCH_MessageType_C1_Choice_RrcRelease
	DL_DCCH_MessageType_C1_Choice_RrcReestablishment
	DL_DCCH_MessageType_C1_Choice_SecurityModeCommand
	DL_DCCH_MessageType_C1_Choice_DlInformationTransfer
	DL_DCCH_MessageType_C1_Choice_UeCapabilityEnquiry
	DL_DCCH_MessageType_C1_Choice_CounterCheck
	DL_DCCH_MessageType_C1_Choice_MobilityFromNRCommand
	DL_DCCH_MessageType_C1_Choice_DlDedicatedMessageSegment_r16
	DL_DCCH_MessageType_C1_Choice_UeInformationRequest_r16
	DL_DCCH_MessageType_C1_Choice_DlInformationTransferMRDC_r16
	DL_DCCH_MessageType_C1_Choice_LoggedMeasurementConfiguration_r16
	DL_DCCH_MessageType_C1_Choice_Spare3
	DL_DCCH_MessageType_C1_Choice_Spare2
	DL_DCCH_MessageType_C1_Choice_Spare1
)

type DL_DCCH_MessageType_C1 struct {
	Choice                             uint64
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
	Spare3                             aper.NULL `madatory`
	Spare2                             aper.NULL `madatory`
	Spare1                             aper.NULL `madatory`
}

func (ie *DL_DCCH_MessageType_C1) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 16, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_DCCH_MessageType_C1_Choice_RrcReconfiguration:
		if err = ie.RrcReconfiguration.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcReconfiguration", err)
		}
	case DL_DCCH_MessageType_C1_Choice_RrcResume:
		if err = ie.RrcResume.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcResume", err)
		}
	case DL_DCCH_MessageType_C1_Choice_RrcRelease:
		if err = ie.RrcRelease.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcRelease", err)
		}
	case DL_DCCH_MessageType_C1_Choice_RrcReestablishment:
		if err = ie.RrcReestablishment.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcReestablishment", err)
		}
	case DL_DCCH_MessageType_C1_Choice_SecurityModeCommand:
		if err = ie.SecurityModeCommand.Encode(w); err != nil {
			err = utils.WrapError("Encode SecurityModeCommand", err)
		}
	case DL_DCCH_MessageType_C1_Choice_DlInformationTransfer:
		if err = ie.DlInformationTransfer.Encode(w); err != nil {
			err = utils.WrapError("Encode DlInformationTransfer", err)
		}
	case DL_DCCH_MessageType_C1_Choice_UeCapabilityEnquiry:
		if err = ie.UeCapabilityEnquiry.Encode(w); err != nil {
			err = utils.WrapError("Encode UeCapabilityEnquiry", err)
		}
	case DL_DCCH_MessageType_C1_Choice_CounterCheck:
		if err = ie.CounterCheck.Encode(w); err != nil {
			err = utils.WrapError("Encode CounterCheck", err)
		}
	case DL_DCCH_MessageType_C1_Choice_MobilityFromNRCommand:
		if err = ie.MobilityFromNRCommand.Encode(w); err != nil {
			err = utils.WrapError("Encode MobilityFromNRCommand", err)
		}
	case DL_DCCH_MessageType_C1_Choice_DlDedicatedMessageSegment_r16:
		if err = ie.DlDedicatedMessageSegment_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode DlDedicatedMessageSegment_r16", err)
		}
	case DL_DCCH_MessageType_C1_Choice_UeInformationRequest_r16:
		if err = ie.UeInformationRequest_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode UeInformationRequest_r16", err)
		}
	case DL_DCCH_MessageType_C1_Choice_DlInformationTransferMRDC_r16:
		if err = ie.DlInformationTransferMRDC_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode DlInformationTransferMRDC_r16", err)
		}
	case DL_DCCH_MessageType_C1_Choice_LoggedMeasurementConfiguration_r16:
		if err = ie.LoggedMeasurementConfiguration_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode LoggedMeasurementConfiguration_r16", err)
		}
	case DL_DCCH_MessageType_C1_Choice_Spare3:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare3", err)
		}
	case DL_DCCH_MessageType_C1_Choice_Spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare2", err)
		}
	case DL_DCCH_MessageType_C1_Choice_Spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DL_DCCH_MessageType_C1) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(16, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_DCCH_MessageType_C1_Choice_RrcReconfiguration:
		ie.RrcReconfiguration = new(RRCReconfiguration)
		if err = ie.RrcReconfiguration.Decode(r); err != nil {
			return utils.WrapError("Decode RrcReconfiguration", err)
		}
	case DL_DCCH_MessageType_C1_Choice_RrcResume:
		ie.RrcResume = new(RRCResume)
		if err = ie.RrcResume.Decode(r); err != nil {
			return utils.WrapError("Decode RrcResume", err)
		}
	case DL_DCCH_MessageType_C1_Choice_RrcRelease:
		ie.RrcRelease = new(RRCRelease)
		if err = ie.RrcRelease.Decode(r); err != nil {
			return utils.WrapError("Decode RrcRelease", err)
		}
	case DL_DCCH_MessageType_C1_Choice_RrcReestablishment:
		ie.RrcReestablishment = new(RRCReestablishment)
		if err = ie.RrcReestablishment.Decode(r); err != nil {
			return utils.WrapError("Decode RrcReestablishment", err)
		}
	case DL_DCCH_MessageType_C1_Choice_SecurityModeCommand:
		ie.SecurityModeCommand = new(SecurityModeCommand)
		if err = ie.SecurityModeCommand.Decode(r); err != nil {
			return utils.WrapError("Decode SecurityModeCommand", err)
		}
	case DL_DCCH_MessageType_C1_Choice_DlInformationTransfer:
		ie.DlInformationTransfer = new(DLInformationTransfer)
		if err = ie.DlInformationTransfer.Decode(r); err != nil {
			return utils.WrapError("Decode DlInformationTransfer", err)
		}
	case DL_DCCH_MessageType_C1_Choice_UeCapabilityEnquiry:
		ie.UeCapabilityEnquiry = new(UECapabilityEnquiry)
		if err = ie.UeCapabilityEnquiry.Decode(r); err != nil {
			return utils.WrapError("Decode UeCapabilityEnquiry", err)
		}
	case DL_DCCH_MessageType_C1_Choice_CounterCheck:
		ie.CounterCheck = new(CounterCheck)
		if err = ie.CounterCheck.Decode(r); err != nil {
			return utils.WrapError("Decode CounterCheck", err)
		}
	case DL_DCCH_MessageType_C1_Choice_MobilityFromNRCommand:
		ie.MobilityFromNRCommand = new(MobilityFromNRCommand)
		if err = ie.MobilityFromNRCommand.Decode(r); err != nil {
			return utils.WrapError("Decode MobilityFromNRCommand", err)
		}
	case DL_DCCH_MessageType_C1_Choice_DlDedicatedMessageSegment_r16:
		ie.DlDedicatedMessageSegment_r16 = new(DLDedicatedMessageSegment_r16)
		if err = ie.DlDedicatedMessageSegment_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DlDedicatedMessageSegment_r16", err)
		}
	case DL_DCCH_MessageType_C1_Choice_UeInformationRequest_r16:
		ie.UeInformationRequest_r16 = new(UEInformationRequest_r16)
		if err = ie.UeInformationRequest_r16.Decode(r); err != nil {
			return utils.WrapError("Decode UeInformationRequest_r16", err)
		}
	case DL_DCCH_MessageType_C1_Choice_DlInformationTransferMRDC_r16:
		ie.DlInformationTransferMRDC_r16 = new(DLInformationTransferMRDC_r16)
		if err = ie.DlInformationTransferMRDC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DlInformationTransferMRDC_r16", err)
		}
	case DL_DCCH_MessageType_C1_Choice_LoggedMeasurementConfiguration_r16:
		ie.LoggedMeasurementConfiguration_r16 = new(LoggedMeasurementConfiguration_r16)
		if err = ie.LoggedMeasurementConfiguration_r16.Decode(r); err != nil {
			return utils.WrapError("Decode LoggedMeasurementConfiguration_r16", err)
		}
	case DL_DCCH_MessageType_C1_Choice_Spare3:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare3", err)
		}
	case DL_DCCH_MessageType_C1_Choice_Spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare2", err)
		}
	case DL_DCCH_MessageType_C1_Choice_Spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
