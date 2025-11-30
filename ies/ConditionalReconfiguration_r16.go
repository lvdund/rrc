package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ConditionalReconfiguration_r16 struct {
	AttemptCondReconfig_r16      *ConditionalReconfiguration_r16_attemptCondReconfig_r16 `optional`
	CondReconfigToRemoveList_r16 *CondReconfigToRemoveList_r16                           `optional`
	CondReconfigToAddModList_r16 *CondReconfigToAddModList_r16                           `optional`
}

func (ie *ConditionalReconfiguration_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.AttemptCondReconfig_r16 != nil, ie.CondReconfigToRemoveList_r16 != nil, ie.CondReconfigToAddModList_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.AttemptCondReconfig_r16 != nil {
		if err = ie.AttemptCondReconfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AttemptCondReconfig_r16", err)
		}
	}
	if ie.CondReconfigToRemoveList_r16 != nil {
		if err = ie.CondReconfigToRemoveList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CondReconfigToRemoveList_r16", err)
		}
	}
	if ie.CondReconfigToAddModList_r16 != nil {
		if err = ie.CondReconfigToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CondReconfigToAddModList_r16", err)
		}
	}
	return nil
}

func (ie *ConditionalReconfiguration_r16) Decode(r *aper.AperReader) error {
	var err error
	var AttemptCondReconfig_r16Present bool
	if AttemptCondReconfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CondReconfigToRemoveList_r16Present bool
	if CondReconfigToRemoveList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CondReconfigToAddModList_r16Present bool
	if CondReconfigToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if AttemptCondReconfig_r16Present {
		ie.AttemptCondReconfig_r16 = new(ConditionalReconfiguration_r16_attemptCondReconfig_r16)
		if err = ie.AttemptCondReconfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode AttemptCondReconfig_r16", err)
		}
	}
	if CondReconfigToRemoveList_r16Present {
		ie.CondReconfigToRemoveList_r16 = new(CondReconfigToRemoveList_r16)
		if err = ie.CondReconfigToRemoveList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CondReconfigToRemoveList_r16", err)
		}
	}
	if CondReconfigToAddModList_r16Present {
		ie.CondReconfigToAddModList_r16 = new(CondReconfigToAddModList_r16)
		if err = ie.CondReconfigToAddModList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CondReconfigToAddModList_r16", err)
		}
	}
	return nil
}
