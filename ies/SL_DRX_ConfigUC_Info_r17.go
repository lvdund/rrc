package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_DRX_ConfigUC_Info_r17 struct {
	Sl_DestinationIndex_r17 *SL_DestinationIndex_r16 `optional`
	Sl_DRX_ConfigUC_r17     *SL_DRX_ConfigUC_r17     `optional`
}

func (ie *SL_DRX_ConfigUC_Info_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_DestinationIndex_r17 != nil, ie.Sl_DRX_ConfigUC_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_DestinationIndex_r17 != nil {
		if err = ie.Sl_DestinationIndex_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_DestinationIndex_r17", err)
		}
	}
	if ie.Sl_DRX_ConfigUC_r17 != nil {
		if err = ie.Sl_DRX_ConfigUC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_DRX_ConfigUC_r17", err)
		}
	}
	return nil
}

func (ie *SL_DRX_ConfigUC_Info_r17) Decode(r *uper.UperReader) error {
	var err error
	var Sl_DestinationIndex_r17Present bool
	if Sl_DestinationIndex_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_DRX_ConfigUC_r17Present bool
	if Sl_DRX_ConfigUC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_DestinationIndex_r17Present {
		ie.Sl_DestinationIndex_r17 = new(SL_DestinationIndex_r16)
		if err = ie.Sl_DestinationIndex_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_DestinationIndex_r17", err)
		}
	}
	if Sl_DRX_ConfigUC_r17Present {
		ie.Sl_DRX_ConfigUC_r17 = new(SL_DRX_ConfigUC_r17)
		if err = ie.Sl_DRX_ConfigUC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_DRX_ConfigUC_r17", err)
		}
	}
	return nil
}
