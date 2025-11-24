package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_Grp_CarrierTypes_r16 struct {
	Fr1_NonSharedTDD_r16 *PUCCH_Grp_CarrierTypes_r16_fr1_NonSharedTDD_r16 `optional`
	Fr1_SharedTDD_r16    *PUCCH_Grp_CarrierTypes_r16_fr1_SharedTDD_r16    `optional`
	Fr1_NonSharedFDD_r16 *PUCCH_Grp_CarrierTypes_r16_fr1_NonSharedFDD_r16 `optional`
	Fr2_r16              *PUCCH_Grp_CarrierTypes_r16_fr2_r16              `optional`
}

func (ie *PUCCH_Grp_CarrierTypes_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Fr1_NonSharedTDD_r16 != nil, ie.Fr1_SharedTDD_r16 != nil, ie.Fr1_NonSharedFDD_r16 != nil, ie.Fr2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Fr1_NonSharedTDD_r16 != nil {
		if err = ie.Fr1_NonSharedTDD_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1_NonSharedTDD_r16", err)
		}
	}
	if ie.Fr1_SharedTDD_r16 != nil {
		if err = ie.Fr1_SharedTDD_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1_SharedTDD_r16", err)
		}
	}
	if ie.Fr1_NonSharedFDD_r16 != nil {
		if err = ie.Fr1_NonSharedFDD_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1_NonSharedFDD_r16", err)
		}
	}
	if ie.Fr2_r16 != nil {
		if err = ie.Fr2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Fr2_r16", err)
		}
	}
	return nil
}

func (ie *PUCCH_Grp_CarrierTypes_r16) Decode(r *uper.UperReader) error {
	var err error
	var Fr1_NonSharedTDD_r16Present bool
	if Fr1_NonSharedTDD_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr1_SharedTDD_r16Present bool
	if Fr1_SharedTDD_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr1_NonSharedFDD_r16Present bool
	if Fr1_NonSharedFDD_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr2_r16Present bool
	if Fr2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Fr1_NonSharedTDD_r16Present {
		ie.Fr1_NonSharedTDD_r16 = new(PUCCH_Grp_CarrierTypes_r16_fr1_NonSharedTDD_r16)
		if err = ie.Fr1_NonSharedTDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1_NonSharedTDD_r16", err)
		}
	}
	if Fr1_SharedTDD_r16Present {
		ie.Fr1_SharedTDD_r16 = new(PUCCH_Grp_CarrierTypes_r16_fr1_SharedTDD_r16)
		if err = ie.Fr1_SharedTDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1_SharedTDD_r16", err)
		}
	}
	if Fr1_NonSharedFDD_r16Present {
		ie.Fr1_NonSharedFDD_r16 = new(PUCCH_Grp_CarrierTypes_r16_fr1_NonSharedFDD_r16)
		if err = ie.Fr1_NonSharedFDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1_NonSharedFDD_r16", err)
		}
	}
	if Fr2_r16Present {
		ie.Fr2_r16 = new(PUCCH_Grp_CarrierTypes_r16_fr2_r16)
		if err = ie.Fr2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Fr2_r16", err)
		}
	}
	return nil
}
