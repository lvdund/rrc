package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_nothing uint64 = iota
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_UlDedicatedMessageSegment_r16
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_DedicatedSIBRequest_r16
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_McgFailureInformation_r16
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_UeInformationResponse_r16
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_SidelinkUEInformationNR_r16
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_UlInformationTransferIRAT_r16
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_IabOtherInformation_r16
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_MbsInterestIndication_r17
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_UePositioningAssistanceInfo_r17
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_MeasurementReportAppLayer_r17
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_Spare6
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_Spare5
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_Spare4
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_Spare3
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_Spare2
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_Spare1
)

type UL_DCCH_MessageType_MessageClassExtension_C2 struct {
	Choice                          uint64
	UlDedicatedMessageSegment_r16   *ULDedicatedMessageSegment_r16
	DedicatedSIBRequest_r16         *DedicatedSIBRequest_r16
	McgFailureInformation_r16       *MCGFailureInformation_r16
	UeInformationResponse_r16       *UEInformationResponse_r16
	SidelinkUEInformationNR_r16     *SidelinkUEInformationNR_r16
	UlInformationTransferIRAT_r16   *ULInformationTransferIRAT_r16
	IabOtherInformation_r16         *IABOtherInformation_r16
	MbsInterestIndication_r17       *MBSInterestIndication_r17
	UePositioningAssistanceInfo_r17 *UEPositioningAssistanceInfo_r17
	MeasurementReportAppLayer_r17   *MeasurementReportAppLayer_r17
	Spare6                          uper.NULL `madatory`
	Spare5                          uper.NULL `madatory`
	Spare4                          uper.NULL `madatory`
	Spare3                          uper.NULL `madatory`
	Spare2                          uper.NULL `madatory`
	Spare1                          uper.NULL `madatory`
}

func (ie *UL_DCCH_MessageType_MessageClassExtension_C2) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 16, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_UlDedicatedMessageSegment_r16:
		if err = ie.UlDedicatedMessageSegment_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode UlDedicatedMessageSegment_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_DedicatedSIBRequest_r16:
		if err = ie.DedicatedSIBRequest_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode DedicatedSIBRequest_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_McgFailureInformation_r16:
		if err = ie.McgFailureInformation_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode McgFailureInformation_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_UeInformationResponse_r16:
		if err = ie.UeInformationResponse_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode UeInformationResponse_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_SidelinkUEInformationNR_r16:
		if err = ie.SidelinkUEInformationNR_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode SidelinkUEInformationNR_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_UlInformationTransferIRAT_r16:
		if err = ie.UlInformationTransferIRAT_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode UlInformationTransferIRAT_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_IabOtherInformation_r16:
		if err = ie.IabOtherInformation_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode IabOtherInformation_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_MbsInterestIndication_r17:
		if err = ie.MbsInterestIndication_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode MbsInterestIndication_r17", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_UePositioningAssistanceInfo_r17:
		if err = ie.UePositioningAssistanceInfo_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode UePositioningAssistanceInfo_r17", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_MeasurementReportAppLayer_r17:
		if err = ie.MeasurementReportAppLayer_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasurementReportAppLayer_r17", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_Spare6:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare6", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_Spare5:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare5", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_Spare4:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare4", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_Spare3:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare3", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_Spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare2", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_Spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UL_DCCH_MessageType_MessageClassExtension_C2) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(16, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_UlDedicatedMessageSegment_r16:
		ie.UlDedicatedMessageSegment_r16 = new(ULDedicatedMessageSegment_r16)
		if err = ie.UlDedicatedMessageSegment_r16.Decode(r); err != nil {
			return utils.WrapError("Decode UlDedicatedMessageSegment_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_DedicatedSIBRequest_r16:
		ie.DedicatedSIBRequest_r16 = new(DedicatedSIBRequest_r16)
		if err = ie.DedicatedSIBRequest_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DedicatedSIBRequest_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_McgFailureInformation_r16:
		ie.McgFailureInformation_r16 = new(MCGFailureInformation_r16)
		if err = ie.McgFailureInformation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode McgFailureInformation_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_UeInformationResponse_r16:
		ie.UeInformationResponse_r16 = new(UEInformationResponse_r16)
		if err = ie.UeInformationResponse_r16.Decode(r); err != nil {
			return utils.WrapError("Decode UeInformationResponse_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_SidelinkUEInformationNR_r16:
		ie.SidelinkUEInformationNR_r16 = new(SidelinkUEInformationNR_r16)
		if err = ie.SidelinkUEInformationNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SidelinkUEInformationNR_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_UlInformationTransferIRAT_r16:
		ie.UlInformationTransferIRAT_r16 = new(ULInformationTransferIRAT_r16)
		if err = ie.UlInformationTransferIRAT_r16.Decode(r); err != nil {
			return utils.WrapError("Decode UlInformationTransferIRAT_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_IabOtherInformation_r16:
		ie.IabOtherInformation_r16 = new(IABOtherInformation_r16)
		if err = ie.IabOtherInformation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IabOtherInformation_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_MbsInterestIndication_r17:
		ie.MbsInterestIndication_r17 = new(MBSInterestIndication_r17)
		if err = ie.MbsInterestIndication_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MbsInterestIndication_r17", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_UePositioningAssistanceInfo_r17:
		ie.UePositioningAssistanceInfo_r17 = new(UEPositioningAssistanceInfo_r17)
		if err = ie.UePositioningAssistanceInfo_r17.Decode(r); err != nil {
			return utils.WrapError("Decode UePositioningAssistanceInfo_r17", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_MeasurementReportAppLayer_r17:
		ie.MeasurementReportAppLayer_r17 = new(MeasurementReportAppLayer_r17)
		if err = ie.MeasurementReportAppLayer_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MeasurementReportAppLayer_r17", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_Spare6:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare6", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_Spare5:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare5", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_Spare4:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare4", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_Spare3:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare3", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_Spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare2", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_Spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
