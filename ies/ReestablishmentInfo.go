package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ReestablishmentInfo struct {
	SourcePhysCellId          PhysCellId            `madatory`
	TargetCellShortMAC_I      ShortMAC_I            `madatory`
	AdditionalReestabInfoList *ReestabNCellInfoList `optional`
}

func (ie *ReestablishmentInfo) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.AdditionalReestabInfoList != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.SourcePhysCellId.Encode(w); err != nil {
		return utils.WrapError("Encode SourcePhysCellId", err)
	}
	if err = ie.TargetCellShortMAC_I.Encode(w); err != nil {
		return utils.WrapError("Encode TargetCellShortMAC_I", err)
	}
	if ie.AdditionalReestabInfoList != nil {
		if err = ie.AdditionalReestabInfoList.Encode(w); err != nil {
			return utils.WrapError("Encode AdditionalReestabInfoList", err)
		}
	}
	return nil
}

func (ie *ReestablishmentInfo) Decode(r *aper.AperReader) error {
	var err error
	var AdditionalReestabInfoListPresent bool
	if AdditionalReestabInfoListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.SourcePhysCellId.Decode(r); err != nil {
		return utils.WrapError("Decode SourcePhysCellId", err)
	}
	if err = ie.TargetCellShortMAC_I.Decode(r); err != nil {
		return utils.WrapError("Decode TargetCellShortMAC_I", err)
	}
	if AdditionalReestabInfoListPresent {
		ie.AdditionalReestabInfoList = new(ReestabNCellInfoList)
		if err = ie.AdditionalReestabInfoList.Decode(r); err != nil {
			return utils.WrapError("Decode AdditionalReestabInfoList", err)
		}
	}
	return nil
}
