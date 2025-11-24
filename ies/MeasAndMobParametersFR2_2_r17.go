package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersFR2_2_r17 struct {
	HandoverInterF_r17            *MeasAndMobParametersFR2_2_r17_handoverInterF_r17            `optional`
	HandoverLTE_EPC_r17           *MeasAndMobParametersFR2_2_r17_handoverLTE_EPC_r17           `optional`
	HandoverLTE_5GC_r17           *MeasAndMobParametersFR2_2_r17_handoverLTE_5GC_r17           `optional`
	IdleInactiveNR_MeasReport_r17 *MeasAndMobParametersFR2_2_r17_idleInactiveNR_MeasReport_r17 `optional`
}

func (ie *MeasAndMobParametersFR2_2_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.HandoverInterF_r17 != nil, ie.HandoverLTE_EPC_r17 != nil, ie.HandoverLTE_5GC_r17 != nil, ie.IdleInactiveNR_MeasReport_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.HandoverInterF_r17 != nil {
		if err = ie.HandoverInterF_r17.Encode(w); err != nil {
			return utils.WrapError("Encode HandoverInterF_r17", err)
		}
	}
	if ie.HandoverLTE_EPC_r17 != nil {
		if err = ie.HandoverLTE_EPC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode HandoverLTE_EPC_r17", err)
		}
	}
	if ie.HandoverLTE_5GC_r17 != nil {
		if err = ie.HandoverLTE_5GC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode HandoverLTE_5GC_r17", err)
		}
	}
	if ie.IdleInactiveNR_MeasReport_r17 != nil {
		if err = ie.IdleInactiveNR_MeasReport_r17.Encode(w); err != nil {
			return utils.WrapError("Encode IdleInactiveNR_MeasReport_r17", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersFR2_2_r17) Decode(r *uper.UperReader) error {
	var err error
	var HandoverInterF_r17Present bool
	if HandoverInterF_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var HandoverLTE_EPC_r17Present bool
	if HandoverLTE_EPC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var HandoverLTE_5GC_r17Present bool
	if HandoverLTE_5GC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var IdleInactiveNR_MeasReport_r17Present bool
	if IdleInactiveNR_MeasReport_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if HandoverInterF_r17Present {
		ie.HandoverInterF_r17 = new(MeasAndMobParametersFR2_2_r17_handoverInterF_r17)
		if err = ie.HandoverInterF_r17.Decode(r); err != nil {
			return utils.WrapError("Decode HandoverInterF_r17", err)
		}
	}
	if HandoverLTE_EPC_r17Present {
		ie.HandoverLTE_EPC_r17 = new(MeasAndMobParametersFR2_2_r17_handoverLTE_EPC_r17)
		if err = ie.HandoverLTE_EPC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode HandoverLTE_EPC_r17", err)
		}
	}
	if HandoverLTE_5GC_r17Present {
		ie.HandoverLTE_5GC_r17 = new(MeasAndMobParametersFR2_2_r17_handoverLTE_5GC_r17)
		if err = ie.HandoverLTE_5GC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode HandoverLTE_5GC_r17", err)
		}
	}
	if IdleInactiveNR_MeasReport_r17Present {
		ie.IdleInactiveNR_MeasReport_r17 = new(MeasAndMobParametersFR2_2_r17_idleInactiveNR_MeasReport_r17)
		if err = ie.IdleInactiveNR_MeasReport_r17.Decode(r); err != nil {
			return utils.WrapError("Decode IdleInactiveNR_MeasReport_r17", err)
		}
	}
	return nil
}
