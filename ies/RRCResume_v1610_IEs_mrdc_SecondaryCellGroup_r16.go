package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16_Choice_nothing uint64 = iota
	RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16_Choice_Nr_SCG_r16
	RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16_Choice_Eutra_SCG_r16
)

type RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16 struct {
	Choice        uint64
	Nr_SCG_r16    []byte `madatory`
	Eutra_SCG_r16 []byte `madatory`
}

func (ie *RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16_Choice_Nr_SCG_r16:
		if err = w.WriteOctetString(ie.Nr_SCG_r16, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			err = utils.WrapError("Encode Nr_SCG_r16", err)
		}
	case RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16_Choice_Eutra_SCG_r16:
		if err = w.WriteOctetString(ie.Eutra_SCG_r16, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			err = utils.WrapError("Encode Eutra_SCG_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16_Choice_Nr_SCG_r16:
		var tmp_os_Nr_SCG_r16 []byte
		if tmp_os_Nr_SCG_r16, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Nr_SCG_r16", err)
		}
		ie.Nr_SCG_r16 = tmp_os_Nr_SCG_r16
	case RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16_Choice_Eutra_SCG_r16:
		var tmp_os_Eutra_SCG_r16 []byte
		if tmp_os_Eutra_SCG_r16, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Eutra_SCG_r16", err)
		}
		ie.Eutra_SCG_r16 = tmp_os_Eutra_SCG_r16
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
