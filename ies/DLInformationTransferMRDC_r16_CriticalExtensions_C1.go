package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DLInformationTransferMRDC_r16_CriticalExtensions_C1_Choice_nothing uint64 = iota
	DLInformationTransferMRDC_r16_CriticalExtensions_C1_Choice_DlInformationTransferMRDC_r16
	DLInformationTransferMRDC_r16_CriticalExtensions_C1_Choice_Spare3
	DLInformationTransferMRDC_r16_CriticalExtensions_C1_Choice_Spare2
	DLInformationTransferMRDC_r16_CriticalExtensions_C1_Choice_Spare1
)

type DLInformationTransferMRDC_r16_CriticalExtensions_C1 struct {
	Choice                        uint64
	DlInformationTransferMRDC_r16 *DLInformationTransferMRDC_r16_IEs
	Spare3                        uper.NULL `madatory`
	Spare2                        uper.NULL `madatory`
	Spare1                        uper.NULL `madatory`
}

func (ie *DLInformationTransferMRDC_r16_CriticalExtensions_C1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DLInformationTransferMRDC_r16_CriticalExtensions_C1_Choice_DlInformationTransferMRDC_r16:
		if err = ie.DlInformationTransferMRDC_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode DlInformationTransferMRDC_r16", err)
		}
	case DLInformationTransferMRDC_r16_CriticalExtensions_C1_Choice_Spare3:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare3", err)
		}
	case DLInformationTransferMRDC_r16_CriticalExtensions_C1_Choice_Spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare2", err)
		}
	case DLInformationTransferMRDC_r16_CriticalExtensions_C1_Choice_Spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DLInformationTransferMRDC_r16_CriticalExtensions_C1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DLInformationTransferMRDC_r16_CriticalExtensions_C1_Choice_DlInformationTransferMRDC_r16:
		ie.DlInformationTransferMRDC_r16 = new(DLInformationTransferMRDC_r16_IEs)
		if err = ie.DlInformationTransferMRDC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DlInformationTransferMRDC_r16", err)
		}
	case DLInformationTransferMRDC_r16_CriticalExtensions_C1_Choice_Spare3:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare3", err)
		}
	case DLInformationTransferMRDC_r16_CriticalExtensions_C1_Choice_Spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare2", err)
		}
	case DLInformationTransferMRDC_r16_CriticalExtensions_C1_Choice_Spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
