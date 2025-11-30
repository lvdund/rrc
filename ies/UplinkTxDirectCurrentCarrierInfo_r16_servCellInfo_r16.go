package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UplinkTxDirectCurrentCarrierInfo_r16_servCellInfo_r16_Choice_nothing uint64 = iota
	UplinkTxDirectCurrentCarrierInfo_r16_servCellInfo_r16_Choice_Bwp_Id_r16
	UplinkTxDirectCurrentCarrierInfo_r16_servCellInfo_r16_Choice_DeactivatedCarrier_r16
)

type UplinkTxDirectCurrentCarrierInfo_r16_servCellInfo_r16 struct {
	Choice                 uint64
	Bwp_Id_r16             *BWP_Id
	DeactivatedCarrier_r16 *UplinkTxDirectCurrentCarrierInfo_r16_servCellInfo_r16_deactivatedCarrier_r16
}

func (ie *UplinkTxDirectCurrentCarrierInfo_r16_servCellInfo_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UplinkTxDirectCurrentCarrierInfo_r16_servCellInfo_r16_Choice_Bwp_Id_r16:
		if err = ie.Bwp_Id_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Bwp_Id_r16", err)
		}
	case UplinkTxDirectCurrentCarrierInfo_r16_servCellInfo_r16_Choice_DeactivatedCarrier_r16:
		if err = ie.DeactivatedCarrier_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode DeactivatedCarrier_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UplinkTxDirectCurrentCarrierInfo_r16_servCellInfo_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UplinkTxDirectCurrentCarrierInfo_r16_servCellInfo_r16_Choice_Bwp_Id_r16:
		ie.Bwp_Id_r16 = new(BWP_Id)
		if err = ie.Bwp_Id_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Bwp_Id_r16", err)
		}
	case UplinkTxDirectCurrentCarrierInfo_r16_servCellInfo_r16_Choice_DeactivatedCarrier_r16:
		ie.DeactivatedCarrier_r16 = new(UplinkTxDirectCurrentCarrierInfo_r16_servCellInfo_r16_deactivatedCarrier_r16)
		if err = ie.DeactivatedCarrier_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DeactivatedCarrier_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
