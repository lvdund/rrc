package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SCCH_MessageType_MessageClassExtension_C2_Choice_nothing uint64 = iota
	SCCH_MessageType_MessageClassExtension_C2_Choice_NotificationMessageSidelink_r17
	SCCH_MessageType_MessageClassExtension_C2_Choice_UeAssistanceInformationSidelink_r17
	SCCH_MessageType_MessageClassExtension_C2_Choice_Spare6
	SCCH_MessageType_MessageClassExtension_C2_Choice_Spare5
	SCCH_MessageType_MessageClassExtension_C2_Choice_Spare4
	SCCH_MessageType_MessageClassExtension_C2_Choice_Spare3
	SCCH_MessageType_MessageClassExtension_C2_Choice_Spare2
	SCCH_MessageType_MessageClassExtension_C2_Choice_Spare1
)

type SCCH_MessageType_MessageClassExtension_C2 struct {
	Choice                              uint64
	NotificationMessageSidelink_r17     *NotificationMessageSidelink_r17
	UeAssistanceInformationSidelink_r17 *UEAssistanceInformationSidelink_r17
	Spare6                              uper.NULL `madatory`
	Spare5                              uper.NULL `madatory`
	Spare4                              uper.NULL `madatory`
	Spare3                              uper.NULL `madatory`
	Spare2                              uper.NULL `madatory`
	Spare1                              uper.NULL `madatory`
}

func (ie *SCCH_MessageType_MessageClassExtension_C2) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SCCH_MessageType_MessageClassExtension_C2_Choice_NotificationMessageSidelink_r17:
		if err = ie.NotificationMessageSidelink_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode NotificationMessageSidelink_r17", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_UeAssistanceInformationSidelink_r17:
		if err = ie.UeAssistanceInformationSidelink_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode UeAssistanceInformationSidelink_r17", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_Spare6:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare6", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_Spare5:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare5", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_Spare4:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare4", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_Spare3:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare3", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_Spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare2", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_Spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SCCH_MessageType_MessageClassExtension_C2) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SCCH_MessageType_MessageClassExtension_C2_Choice_NotificationMessageSidelink_r17:
		ie.NotificationMessageSidelink_r17 = new(NotificationMessageSidelink_r17)
		if err = ie.NotificationMessageSidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode NotificationMessageSidelink_r17", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_UeAssistanceInformationSidelink_r17:
		ie.UeAssistanceInformationSidelink_r17 = new(UEAssistanceInformationSidelink_r17)
		if err = ie.UeAssistanceInformationSidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode UeAssistanceInformationSidelink_r17", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_Spare6:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare6", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_Spare5:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare5", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_Spare4:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare4", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_Spare3:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare3", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_Spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare2", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_Spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
