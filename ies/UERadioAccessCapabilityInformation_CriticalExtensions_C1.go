package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_nothing uint64 = iota
	UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_UeRadioAccessCapabilityInformation
	UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare7
	UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare6
	UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare5
	UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare4
	UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare3
	UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare2
	UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare1
)

type UERadioAccessCapabilityInformation_CriticalExtensions_C1 struct {
	Choice                             uint64
	UeRadioAccessCapabilityInformation *UERadioAccessCapabilityInformation_IEs
	Spare7                             uper.NULL `madatory`
	Spare6                             uper.NULL `madatory`
	Spare5                             uper.NULL `madatory`
	Spare4                             uper.NULL `madatory`
	Spare3                             uper.NULL `madatory`
	Spare2                             uper.NULL `madatory`
	Spare1                             uper.NULL `madatory`
}

func (ie *UERadioAccessCapabilityInformation_CriticalExtensions_C1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_UeRadioAccessCapabilityInformation:
		if err = ie.UeRadioAccessCapabilityInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode UeRadioAccessCapabilityInformation", err)
		}
	case UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare7:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare7", err)
		}
	case UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare6:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare6", err)
		}
	case UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare5:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare5", err)
		}
	case UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare4:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare4", err)
		}
	case UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare3:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare3", err)
		}
	case UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare2", err)
		}
	case UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UERadioAccessCapabilityInformation_CriticalExtensions_C1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_UeRadioAccessCapabilityInformation:
		ie.UeRadioAccessCapabilityInformation = new(UERadioAccessCapabilityInformation_IEs)
		if err = ie.UeRadioAccessCapabilityInformation.Decode(r); err != nil {
			return utils.WrapError("Decode UeRadioAccessCapabilityInformation", err)
		}
	case UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare7:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare7", err)
		}
	case UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare6:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare6", err)
		}
	case UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare5:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare5", err)
		}
	case UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare4:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare4", err)
		}
	case UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare3:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare3", err)
		}
	case UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare2", err)
		}
	case UERadioAccessCapabilityInformation_CriticalExtensions_C1_Choice_Spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
