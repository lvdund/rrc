package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type GapConfig_r17 struct {
	MeasGapId_r17              MeasGapId_r17                           `madatory`
	GapType_r17                GapConfig_r17_gapType_r17               `madatory`
	GapOffset_r17              int64                                   `lb:0,ub:159,madatory`
	Mgl_r17                    GapConfig_r17_mgl_r17                   `madatory`
	Mgrp_r17                   GapConfig_r17_mgrp_r17                  `madatory`
	Mgta_r17                   GapConfig_r17_mgta_r17                  `madatory`
	RefServCellIndicator_r17   *GapConfig_r17_refServCellIndicator_r17 `optional`
	RefFR2_ServCellAsyncCA_r17 *ServCellIndex                          `optional`
	PreConfigInd_r17           *GapConfig_r17_preConfigInd_r17         `optional`
	NcsgInd_r17                *GapConfig_r17_ncsgInd_r17              `optional`
	GapAssociationPRS_r17      *GapConfig_r17_gapAssociationPRS_r17    `optional`
	GapSharing_r17             *MeasGapSharingScheme                   `optional`
	GapPriority_r17            *GapPriority_r17                        `optional`
}

func (ie *GapConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.RefServCellIndicator_r17 != nil, ie.RefFR2_ServCellAsyncCA_r17 != nil, ie.PreConfigInd_r17 != nil, ie.NcsgInd_r17 != nil, ie.GapAssociationPRS_r17 != nil, ie.GapSharing_r17 != nil, ie.GapPriority_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.MeasGapId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MeasGapId_r17", err)
	}
	if err = ie.GapType_r17.Encode(w); err != nil {
		return utils.WrapError("Encode GapType_r17", err)
	}
	if err = w.WriteInteger(ie.GapOffset_r17, &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
		return utils.WrapError("WriteInteger GapOffset_r17", err)
	}
	if err = ie.Mgl_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Mgl_r17", err)
	}
	if err = ie.Mgrp_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Mgrp_r17", err)
	}
	if err = ie.Mgta_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Mgta_r17", err)
	}
	if ie.RefServCellIndicator_r17 != nil {
		if err = ie.RefServCellIndicator_r17.Encode(w); err != nil {
			return utils.WrapError("Encode RefServCellIndicator_r17", err)
		}
	}
	if ie.RefFR2_ServCellAsyncCA_r17 != nil {
		if err = ie.RefFR2_ServCellAsyncCA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode RefFR2_ServCellAsyncCA_r17", err)
		}
	}
	if ie.PreConfigInd_r17 != nil {
		if err = ie.PreConfigInd_r17.Encode(w); err != nil {
			return utils.WrapError("Encode PreConfigInd_r17", err)
		}
	}
	if ie.NcsgInd_r17 != nil {
		if err = ie.NcsgInd_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NcsgInd_r17", err)
		}
	}
	if ie.GapAssociationPRS_r17 != nil {
		if err = ie.GapAssociationPRS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode GapAssociationPRS_r17", err)
		}
	}
	if ie.GapSharing_r17 != nil {
		if err = ie.GapSharing_r17.Encode(w); err != nil {
			return utils.WrapError("Encode GapSharing_r17", err)
		}
	}
	if ie.GapPriority_r17 != nil {
		if err = ie.GapPriority_r17.Encode(w); err != nil {
			return utils.WrapError("Encode GapPriority_r17", err)
		}
	}
	return nil
}

func (ie *GapConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var RefServCellIndicator_r17Present bool
	if RefServCellIndicator_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var RefFR2_ServCellAsyncCA_r17Present bool
	if RefFR2_ServCellAsyncCA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PreConfigInd_r17Present bool
	if PreConfigInd_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NcsgInd_r17Present bool
	if NcsgInd_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var GapAssociationPRS_r17Present bool
	if GapAssociationPRS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var GapSharing_r17Present bool
	if GapSharing_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var GapPriority_r17Present bool
	if GapPriority_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.MeasGapId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MeasGapId_r17", err)
	}
	if err = ie.GapType_r17.Decode(r); err != nil {
		return utils.WrapError("Decode GapType_r17", err)
	}
	var tmp_int_GapOffset_r17 int64
	if tmp_int_GapOffset_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
		return utils.WrapError("ReadInteger GapOffset_r17", err)
	}
	ie.GapOffset_r17 = tmp_int_GapOffset_r17
	if err = ie.Mgl_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Mgl_r17", err)
	}
	if err = ie.Mgrp_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Mgrp_r17", err)
	}
	if err = ie.Mgta_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Mgta_r17", err)
	}
	if RefServCellIndicator_r17Present {
		ie.RefServCellIndicator_r17 = new(GapConfig_r17_refServCellIndicator_r17)
		if err = ie.RefServCellIndicator_r17.Decode(r); err != nil {
			return utils.WrapError("Decode RefServCellIndicator_r17", err)
		}
	}
	if RefFR2_ServCellAsyncCA_r17Present {
		ie.RefFR2_ServCellAsyncCA_r17 = new(ServCellIndex)
		if err = ie.RefFR2_ServCellAsyncCA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode RefFR2_ServCellAsyncCA_r17", err)
		}
	}
	if PreConfigInd_r17Present {
		ie.PreConfigInd_r17 = new(GapConfig_r17_preConfigInd_r17)
		if err = ie.PreConfigInd_r17.Decode(r); err != nil {
			return utils.WrapError("Decode PreConfigInd_r17", err)
		}
	}
	if NcsgInd_r17Present {
		ie.NcsgInd_r17 = new(GapConfig_r17_ncsgInd_r17)
		if err = ie.NcsgInd_r17.Decode(r); err != nil {
			return utils.WrapError("Decode NcsgInd_r17", err)
		}
	}
	if GapAssociationPRS_r17Present {
		ie.GapAssociationPRS_r17 = new(GapConfig_r17_gapAssociationPRS_r17)
		if err = ie.GapAssociationPRS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode GapAssociationPRS_r17", err)
		}
	}
	if GapSharing_r17Present {
		ie.GapSharing_r17 = new(MeasGapSharingScheme)
		if err = ie.GapSharing_r17.Decode(r); err != nil {
			return utils.WrapError("Decode GapSharing_r17", err)
		}
	}
	if GapPriority_r17Present {
		ie.GapPriority_r17 = new(GapPriority_r17)
		if err = ie.GapPriority_r17.Decode(r); err != nil {
			return utils.WrapError("Decode GapPriority_r17", err)
		}
	}
	return nil
}
