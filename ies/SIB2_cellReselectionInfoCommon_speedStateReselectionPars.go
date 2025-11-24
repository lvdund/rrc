package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB2_cellReselectionInfoCommon_speedStateReselectionPars struct {
	MobilityStateParameters MobilityStateParameters                                           `madatory`
	Q_HystSF                SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF `madatory`
}

func (ie *SIB2_cellReselectionInfoCommon_speedStateReselectionPars) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.MobilityStateParameters.Encode(w); err != nil {
		return utils.WrapError("Encode MobilityStateParameters", err)
	}
	if err = ie.Q_HystSF.Encode(w); err != nil {
		return utils.WrapError("Encode Q_HystSF", err)
	}
	return nil
}

func (ie *SIB2_cellReselectionInfoCommon_speedStateReselectionPars) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.MobilityStateParameters.Decode(r); err != nil {
		return utils.WrapError("Decode MobilityStateParameters", err)
	}
	if err = ie.Q_HystSF.Decode(r); err != nil {
		return utils.WrapError("Decode Q_HystSF", err)
	}
	return nil
}
