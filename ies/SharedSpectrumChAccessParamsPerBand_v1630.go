package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SharedSpectrumChAccessParamsPerBand_v1630 struct {
	Dl_ReceptionIntraCellGuardband_r16 *SharedSpectrumChAccessParamsPerBand_v1630_dl_ReceptionIntraCellGuardband_r16 `optional`
	Dl_ReceptionLBT_subsetRB_r16       *SharedSpectrumChAccessParamsPerBand_v1630_dl_ReceptionLBT_subsetRB_r16       `optional`
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1630) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Dl_ReceptionIntraCellGuardband_r16 != nil, ie.Dl_ReceptionLBT_subsetRB_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Dl_ReceptionIntraCellGuardband_r16 != nil {
		if err = ie.Dl_ReceptionIntraCellGuardband_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Dl_ReceptionIntraCellGuardband_r16", err)
		}
	}
	if ie.Dl_ReceptionLBT_subsetRB_r16 != nil {
		if err = ie.Dl_ReceptionLBT_subsetRB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Dl_ReceptionLBT_subsetRB_r16", err)
		}
	}
	return nil
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1630) Decode(r *uper.UperReader) error {
	var err error
	var Dl_ReceptionIntraCellGuardband_r16Present bool
	if Dl_ReceptionIntraCellGuardband_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dl_ReceptionLBT_subsetRB_r16Present bool
	if Dl_ReceptionLBT_subsetRB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Dl_ReceptionIntraCellGuardband_r16Present {
		ie.Dl_ReceptionIntraCellGuardband_r16 = new(SharedSpectrumChAccessParamsPerBand_v1630_dl_ReceptionIntraCellGuardband_r16)
		if err = ie.Dl_ReceptionIntraCellGuardband_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Dl_ReceptionIntraCellGuardband_r16", err)
		}
	}
	if Dl_ReceptionLBT_subsetRB_r16Present {
		ie.Dl_ReceptionLBT_subsetRB_r16 = new(SharedSpectrumChAccessParamsPerBand_v1630_dl_ReceptionLBT_subsetRB_r16)
		if err = ie.Dl_ReceptionLBT_subsetRB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Dl_ReceptionLBT_subsetRB_r16", err)
		}
	}
	return nil
}
