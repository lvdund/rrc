package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SCCH_MessageType_C1_Choice_nothing uint64 = iota
	SCCH_MessageType_C1_Choice_MeasurementReportSidelink
	SCCH_MessageType_C1_Choice_RrcReconfigurationSidelink
	SCCH_MessageType_C1_Choice_RrcReconfigurationCompleteSidelink
	SCCH_MessageType_C1_Choice_RrcReconfigurationFailureSidelink
	SCCH_MessageType_C1_Choice_UeCapabilityEnquirySidelink
	SCCH_MessageType_C1_Choice_UeCapabilityInformationSidelink
	SCCH_MessageType_C1_Choice_UuMessageTransferSidelink_r17
	SCCH_MessageType_C1_Choice_RemoteUEInformationSidelink_r17
)

type SCCH_MessageType_C1 struct {
	Choice                             uint64
	MeasurementReportSidelink          *MeasurementReportSidelink
	RrcReconfigurationSidelink         *RRCReconfigurationSidelink
	RrcReconfigurationCompleteSidelink *RRCReconfigurationCompleteSidelink
	RrcReconfigurationFailureSidelink  *RRCReconfigurationFailureSidelink
	UeCapabilityEnquirySidelink        *UECapabilityEnquirySidelink
	UeCapabilityInformationSidelink    *UECapabilityInformationSidelink
	UuMessageTransferSidelink_r17      *UuMessageTransferSidelink_r17
	RemoteUEInformationSidelink_r17    *RemoteUEInformationSidelink_r17
}

func (ie *SCCH_MessageType_C1) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SCCH_MessageType_C1_Choice_MeasurementReportSidelink:
		if err = ie.MeasurementReportSidelink.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasurementReportSidelink", err)
		}
	case SCCH_MessageType_C1_Choice_RrcReconfigurationSidelink:
		if err = ie.RrcReconfigurationSidelink.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcReconfigurationSidelink", err)
		}
	case SCCH_MessageType_C1_Choice_RrcReconfigurationCompleteSidelink:
		if err = ie.RrcReconfigurationCompleteSidelink.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcReconfigurationCompleteSidelink", err)
		}
	case SCCH_MessageType_C1_Choice_RrcReconfigurationFailureSidelink:
		if err = ie.RrcReconfigurationFailureSidelink.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcReconfigurationFailureSidelink", err)
		}
	case SCCH_MessageType_C1_Choice_UeCapabilityEnquirySidelink:
		if err = ie.UeCapabilityEnquirySidelink.Encode(w); err != nil {
			err = utils.WrapError("Encode UeCapabilityEnquirySidelink", err)
		}
	case SCCH_MessageType_C1_Choice_UeCapabilityInformationSidelink:
		if err = ie.UeCapabilityInformationSidelink.Encode(w); err != nil {
			err = utils.WrapError("Encode UeCapabilityInformationSidelink", err)
		}
	case SCCH_MessageType_C1_Choice_UuMessageTransferSidelink_r17:
		if err = ie.UuMessageTransferSidelink_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode UuMessageTransferSidelink_r17", err)
		}
	case SCCH_MessageType_C1_Choice_RemoteUEInformationSidelink_r17:
		if err = ie.RemoteUEInformationSidelink_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode RemoteUEInformationSidelink_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SCCH_MessageType_C1) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SCCH_MessageType_C1_Choice_MeasurementReportSidelink:
		ie.MeasurementReportSidelink = new(MeasurementReportSidelink)
		if err = ie.MeasurementReportSidelink.Decode(r); err != nil {
			return utils.WrapError("Decode MeasurementReportSidelink", err)
		}
	case SCCH_MessageType_C1_Choice_RrcReconfigurationSidelink:
		ie.RrcReconfigurationSidelink = new(RRCReconfigurationSidelink)
		if err = ie.RrcReconfigurationSidelink.Decode(r); err != nil {
			return utils.WrapError("Decode RrcReconfigurationSidelink", err)
		}
	case SCCH_MessageType_C1_Choice_RrcReconfigurationCompleteSidelink:
		ie.RrcReconfigurationCompleteSidelink = new(RRCReconfigurationCompleteSidelink)
		if err = ie.RrcReconfigurationCompleteSidelink.Decode(r); err != nil {
			return utils.WrapError("Decode RrcReconfigurationCompleteSidelink", err)
		}
	case SCCH_MessageType_C1_Choice_RrcReconfigurationFailureSidelink:
		ie.RrcReconfigurationFailureSidelink = new(RRCReconfigurationFailureSidelink)
		if err = ie.RrcReconfigurationFailureSidelink.Decode(r); err != nil {
			return utils.WrapError("Decode RrcReconfigurationFailureSidelink", err)
		}
	case SCCH_MessageType_C1_Choice_UeCapabilityEnquirySidelink:
		ie.UeCapabilityEnquirySidelink = new(UECapabilityEnquirySidelink)
		if err = ie.UeCapabilityEnquirySidelink.Decode(r); err != nil {
			return utils.WrapError("Decode UeCapabilityEnquirySidelink", err)
		}
	case SCCH_MessageType_C1_Choice_UeCapabilityInformationSidelink:
		ie.UeCapabilityInformationSidelink = new(UECapabilityInformationSidelink)
		if err = ie.UeCapabilityInformationSidelink.Decode(r); err != nil {
			return utils.WrapError("Decode UeCapabilityInformationSidelink", err)
		}
	case SCCH_MessageType_C1_Choice_UuMessageTransferSidelink_r17:
		ie.UuMessageTransferSidelink_r17 = new(UuMessageTransferSidelink_r17)
		if err = ie.UuMessageTransferSidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode UuMessageTransferSidelink_r17", err)
		}
	case SCCH_MessageType_C1_Choice_RemoteUEInformationSidelink_r17:
		ie.RemoteUEInformationSidelink_r17 = new(RemoteUEInformationSidelink_r17)
		if err = ie.RemoteUEInformationSidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode RemoteUEInformationSidelink_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
