package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TCI_UL_State_r17 struct {
	Tci_UL_State_Id_r17        TCI_UL_State_Id_r17                  `madatory`
	ServingCellId_r17          *ServCellIndex                       `optional`
	Bwp_Id_r17                 *BWP_Id                              `optional`
	ReferenceSignal_r17        TCI_UL_State_r17_referenceSignal_r17 `madatory`
	AdditionalPCI_r17          *AdditionalPCIIndex_r17              `optional`
	Ul_powerControl_r17        *Uplink_powerControlId_r17           `optional`
	PathlossReferenceRS_Id_r17 *PathlossReferenceRS_Id_r17          `optional`
}

func (ie *TCI_UL_State_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ServingCellId_r17 != nil, ie.Bwp_Id_r17 != nil, ie.AdditionalPCI_r17 != nil, ie.Ul_powerControl_r17 != nil, ie.PathlossReferenceRS_Id_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Tci_UL_State_Id_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Tci_UL_State_Id_r17", err)
	}
	if ie.ServingCellId_r17 != nil {
		if err = ie.ServingCellId_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ServingCellId_r17", err)
		}
	}
	if ie.Bwp_Id_r17 != nil {
		if err = ie.Bwp_Id_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Bwp_Id_r17", err)
		}
	}
	if err = ie.ReferenceSignal_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ReferenceSignal_r17", err)
	}
	if ie.AdditionalPCI_r17 != nil {
		if err = ie.AdditionalPCI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode AdditionalPCI_r17", err)
		}
	}
	if ie.Ul_powerControl_r17 != nil {
		if err = ie.Ul_powerControl_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_powerControl_r17", err)
		}
	}
	if ie.PathlossReferenceRS_Id_r17 != nil {
		if err = ie.PathlossReferenceRS_Id_r17.Encode(w); err != nil {
			return utils.WrapError("Encode PathlossReferenceRS_Id_r17", err)
		}
	}
	return nil
}

func (ie *TCI_UL_State_r17) Decode(r *uper.UperReader) error {
	var err error
	var ServingCellId_r17Present bool
	if ServingCellId_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Bwp_Id_r17Present bool
	if Bwp_Id_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AdditionalPCI_r17Present bool
	if AdditionalPCI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_powerControl_r17Present bool
	if Ul_powerControl_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PathlossReferenceRS_Id_r17Present bool
	if PathlossReferenceRS_Id_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Tci_UL_State_Id_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Tci_UL_State_Id_r17", err)
	}
	if ServingCellId_r17Present {
		ie.ServingCellId_r17 = new(ServCellIndex)
		if err = ie.ServingCellId_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ServingCellId_r17", err)
		}
	}
	if Bwp_Id_r17Present {
		ie.Bwp_Id_r17 = new(BWP_Id)
		if err = ie.Bwp_Id_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Bwp_Id_r17", err)
		}
	}
	if err = ie.ReferenceSignal_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ReferenceSignal_r17", err)
	}
	if AdditionalPCI_r17Present {
		ie.AdditionalPCI_r17 = new(AdditionalPCIIndex_r17)
		if err = ie.AdditionalPCI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode AdditionalPCI_r17", err)
		}
	}
	if Ul_powerControl_r17Present {
		ie.Ul_powerControl_r17 = new(Uplink_powerControlId_r17)
		if err = ie.Ul_powerControl_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_powerControl_r17", err)
		}
	}
	if PathlossReferenceRS_Id_r17Present {
		ie.PathlossReferenceRS_Id_r17 = new(PathlossReferenceRS_Id_r17)
		if err = ie.PathlossReferenceRS_Id_r17.Decode(r); err != nil {
			return utils.WrapError("Decode PathlossReferenceRS_Id_r17", err)
		}
	}
	return nil
}
