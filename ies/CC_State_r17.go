package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CC_State_r17 struct {
	DlCarrier_r17 *CarrierState_r17 `optional`
	UlCarrier_r17 *CarrierState_r17 `optional`
}

func (ie *CC_State_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.DlCarrier_r17 != nil, ie.UlCarrier_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.DlCarrier_r17 != nil {
		if err = ie.DlCarrier_r17.Encode(w); err != nil {
			return utils.WrapError("Encode DlCarrier_r17", err)
		}
	}
	if ie.UlCarrier_r17 != nil {
		if err = ie.UlCarrier_r17.Encode(w); err != nil {
			return utils.WrapError("Encode UlCarrier_r17", err)
		}
	}
	return nil
}

func (ie *CC_State_r17) Decode(r *aper.AperReader) error {
	var err error
	var DlCarrier_r17Present bool
	if DlCarrier_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var UlCarrier_r17Present bool
	if UlCarrier_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if DlCarrier_r17Present {
		ie.DlCarrier_r17 = new(CarrierState_r17)
		if err = ie.DlCarrier_r17.Decode(r); err != nil {
			return utils.WrapError("Decode DlCarrier_r17", err)
		}
	}
	if UlCarrier_r17Present {
		ie.UlCarrier_r17 = new(CarrierState_r17)
		if err = ie.UlCarrier_r17.Decode(r); err != nil {
			return utils.WrapError("Decode UlCarrier_r17", err)
		}
	}
	return nil
}
