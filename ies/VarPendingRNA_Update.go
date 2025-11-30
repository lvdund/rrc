package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type VarPendingRNA_Update struct {
	PendingRNA_Update *bool `optional`
}

func (ie *VarPendingRNA_Update) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.PendingRNA_Update != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.PendingRNA_Update != nil {
		if err = w.WriteBoolean(*ie.PendingRNA_Update); err != nil {
			return utils.WrapError("Encode PendingRNA_Update", err)
		}
	}
	return nil
}

func (ie *VarPendingRNA_Update) Decode(r *aper.AperReader) error {
	var err error
	var PendingRNA_UpdatePresent bool
	if PendingRNA_UpdatePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if PendingRNA_UpdatePresent {
		var tmp_bool_PendingRNA_Update bool
		if tmp_bool_PendingRNA_Update, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode PendingRNA_Update", err)
		}
		ie.PendingRNA_Update = &tmp_bool_PendingRNA_Update
	}
	return nil
}
