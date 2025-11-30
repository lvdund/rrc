package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RA_PrioritizationForSlicing_r17 struct {
	Ra_PrioritizationSliceInfoList_r17 RA_PrioritizationSliceInfoList_r17 `madatory`
}

func (ie *RA_PrioritizationForSlicing_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Ra_PrioritizationSliceInfoList_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Ra_PrioritizationSliceInfoList_r17", err)
	}
	return nil
}

func (ie *RA_PrioritizationForSlicing_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Ra_PrioritizationSliceInfoList_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Ra_PrioritizationSliceInfoList_r17", err)
	}
	return nil
}
