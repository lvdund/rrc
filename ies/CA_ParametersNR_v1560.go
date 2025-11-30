package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1560 struct {
	DiffNumerologyWithinPUCCH_GroupLargerSCS *CA_ParametersNR_v1560_diffNumerologyWithinPUCCH_GroupLargerSCS `optional`
}

func (ie *CA_ParametersNR_v1560) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.DiffNumerologyWithinPUCCH_GroupLargerSCS != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.DiffNumerologyWithinPUCCH_GroupLargerSCS != nil {
		if err = ie.DiffNumerologyWithinPUCCH_GroupLargerSCS.Encode(w); err != nil {
			return utils.WrapError("Encode DiffNumerologyWithinPUCCH_GroupLargerSCS", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1560) Decode(r *aper.AperReader) error {
	var err error
	var DiffNumerologyWithinPUCCH_GroupLargerSCSPresent bool
	if DiffNumerologyWithinPUCCH_GroupLargerSCSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if DiffNumerologyWithinPUCCH_GroupLargerSCSPresent {
		ie.DiffNumerologyWithinPUCCH_GroupLargerSCS = new(CA_ParametersNR_v1560_diffNumerologyWithinPUCCH_GroupLargerSCS)
		if err = ie.DiffNumerologyWithinPUCCH_GroupLargerSCS.Decode(r); err != nil {
			return utils.WrapError("Decode DiffNumerologyWithinPUCCH_GroupLargerSCS", err)
		}
	}
	return nil
}
