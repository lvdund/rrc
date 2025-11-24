package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup_Choice_nothing uint64 = iota
	MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup_Choice_Nr_SCG
	MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup_Choice_Eutra_SCG
)

type MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup struct {
	Choice    uint64
	Nr_SCG    []byte `madatory`
	Eutra_SCG []byte `madatory`
}

func (ie *MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup_Choice_Nr_SCG:
		if err = w.WriteOctetString(ie.Nr_SCG, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			err = utils.WrapError("Encode Nr_SCG", err)
		}
	case MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup_Choice_Eutra_SCG:
		if err = w.WriteOctetString(ie.Eutra_SCG, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			err = utils.WrapError("Encode Eutra_SCG", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup_Choice_Nr_SCG:
		var tmp_os_Nr_SCG []byte
		if tmp_os_Nr_SCG, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Nr_SCG", err)
		}
		ie.Nr_SCG = tmp_os_Nr_SCG
	case MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup_Choice_Eutra_SCG:
		var tmp_os_Eutra_SCG []byte
		if tmp_os_Eutra_SCG, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Eutra_SCG", err)
		}
		ie.Eutra_SCG = tmp_os_Eutra_SCG
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
