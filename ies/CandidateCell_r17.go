package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CandidateCell_r17 struct {
	PhysCellId_r17           PhysCellId `madatory`
	CondExecutionCondSCG_r17 *[]byte    `optional`
}

func (ie *CandidateCell_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.CondExecutionCondSCG_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.PhysCellId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode PhysCellId_r17", err)
	}
	if ie.CondExecutionCondSCG_r17 != nil {
		if err = w.WriteOctetString(*ie.CondExecutionCondSCG_r17, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode CondExecutionCondSCG_r17", err)
		}
	}
	return nil
}

func (ie *CandidateCell_r17) Decode(r *aper.AperReader) error {
	var err error
	var CondExecutionCondSCG_r17Present bool
	if CondExecutionCondSCG_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.PhysCellId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode PhysCellId_r17", err)
	}
	if CondExecutionCondSCG_r17Present {
		var tmp_os_CondExecutionCondSCG_r17 []byte
		if tmp_os_CondExecutionCondSCG_r17, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode CondExecutionCondSCG_r17", err)
		}
		ie.CondExecutionCondSCG_r17 = &tmp_os_CondExecutionCondSCG_r17
	}
	return nil
}
