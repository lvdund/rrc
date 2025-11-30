package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NR_DL_PRS_PDC_Info_r17 struct {
	Nr_DL_PRS_PDC_ResourceSet_r17 *NR_DL_PRS_PDC_ResourceSet_r17 `optional`
}

func (ie *NR_DL_PRS_PDC_Info_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Nr_DL_PRS_PDC_ResourceSet_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Nr_DL_PRS_PDC_ResourceSet_r17 != nil {
		if err = ie.Nr_DL_PRS_PDC_ResourceSet_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Nr_DL_PRS_PDC_ResourceSet_r17", err)
		}
	}
	return nil
}

func (ie *NR_DL_PRS_PDC_Info_r17) Decode(r *aper.AperReader) error {
	var err error
	var Nr_DL_PRS_PDC_ResourceSet_r17Present bool
	if Nr_DL_PRS_PDC_ResourceSet_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Nr_DL_PRS_PDC_ResourceSet_r17Present {
		ie.Nr_DL_PRS_PDC_ResourceSet_r17 = new(NR_DL_PRS_PDC_ResourceSet_r17)
		if err = ie.Nr_DL_PRS_PDC_ResourceSet_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Nr_DL_PRS_PDC_ResourceSet_r17", err)
		}
	}
	return nil
}
