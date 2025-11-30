package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_CapabilityAddXDD_Mode_v1530 struct {
	Eutra_ParametersXDD_Diff EUTRA_ParametersXDD_Diff `madatory`
}

func (ie *UE_NR_CapabilityAddXDD_Mode_v1530) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Eutra_ParametersXDD_Diff.Encode(w); err != nil {
		return utils.WrapError("Encode Eutra_ParametersXDD_Diff", err)
	}
	return nil
}

func (ie *UE_NR_CapabilityAddXDD_Mode_v1530) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Eutra_ParametersXDD_Diff.Decode(r); err != nil {
		return utils.WrapError("Decode Eutra_ParametersXDD_Diff", err)
	}
	return nil
}
