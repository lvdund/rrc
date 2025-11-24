package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RMTC_Config_r16_tci_StateInfo_r17 struct {
	Tci_StateId_r17    TCI_StateId    `madatory`
	Ref_ServCellId_r17 *ServCellIndex `optional`
}

func (ie *RMTC_Config_r16_tci_StateInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ref_ServCellId_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Tci_StateId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Tci_StateId_r17", err)
	}
	if ie.Ref_ServCellId_r17 != nil {
		if err = ie.Ref_ServCellId_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ref_ServCellId_r17", err)
		}
	}
	return nil
}

func (ie *RMTC_Config_r16_tci_StateInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var Ref_ServCellId_r17Present bool
	if Ref_ServCellId_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Tci_StateId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Tci_StateId_r17", err)
	}
	if Ref_ServCellId_r17Present {
		ie.Ref_ServCellId_r17 = new(ServCellIndex)
		if err = ie.Ref_ServCellId_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ref_ServCellId_r17", err)
		}
	}
	return nil
}
