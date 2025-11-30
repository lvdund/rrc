package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NR_DL_PRS_Resource_r17 struct {
	Nr_DL_PRS_ResourceID_r17         NR_DL_PRS_ResourceID_r17                                `madatory`
	Dl_PRS_SequenceID_r17            int64                                                   `lb:0,ub:4095,madatory`
	Dl_PRS_CombSizeN_AndReOffset_r17 NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17 `lb:0,ub:1,madatory`
	Dl_PRS_ResourceSlotOffset_r17    int64                                                   `lb:0,ub:maxNrofPRS_ResourceOffsetValue_1_r17,madatory,ext`
	Dl_PRS_ResourceSymbolOffset_r17  int64                                                   `lb:0,ub:12,madatory,ext`
	Dl_PRS_QCL_Info_r17              *DL_PRS_QCL_Info_r17                                    `optional,ext`
}

func (ie *NR_DL_PRS_Resource_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Nr_DL_PRS_ResourceID_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Nr_DL_PRS_ResourceID_r17", err)
	}
	if err = w.WriteInteger(ie.Dl_PRS_SequenceID_r17, &aper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
		return utils.WrapError("WriteInteger Dl_PRS_SequenceID_r17", err)
	}
	if err = ie.Dl_PRS_CombSizeN_AndReOffset_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Dl_PRS_CombSizeN_AndReOffset_r17", err)
	}
	return nil
}

func (ie *NR_DL_PRS_Resource_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Nr_DL_PRS_ResourceID_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Nr_DL_PRS_ResourceID_r17", err)
	}
	var tmp_int_Dl_PRS_SequenceID_r17 int64
	if tmp_int_Dl_PRS_SequenceID_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
		return utils.WrapError("ReadInteger Dl_PRS_SequenceID_r17", err)
	}
	ie.Dl_PRS_SequenceID_r17 = tmp_int_Dl_PRS_SequenceID_r17
	if err = ie.Dl_PRS_CombSizeN_AndReOffset_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Dl_PRS_CombSizeN_AndReOffset_r17", err)
	}
	return nil
}
