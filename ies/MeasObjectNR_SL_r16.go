package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasObjectNR_SL_r16 struct {
	Tx_PoolMeasToRemoveList_r16 *Tx_PoolMeasList_r16 `optional`
	Tx_PoolMeasToAddModList_r16 *Tx_PoolMeasList_r16 `optional`
}

func (ie *MeasObjectNR_SL_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Tx_PoolMeasToRemoveList_r16 != nil, ie.Tx_PoolMeasToAddModList_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Tx_PoolMeasToRemoveList_r16 != nil {
		if err = ie.Tx_PoolMeasToRemoveList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Tx_PoolMeasToRemoveList_r16", err)
		}
	}
	if ie.Tx_PoolMeasToAddModList_r16 != nil {
		if err = ie.Tx_PoolMeasToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Tx_PoolMeasToAddModList_r16", err)
		}
	}
	return nil
}

func (ie *MeasObjectNR_SL_r16) Decode(r *uper.UperReader) error {
	var err error
	var Tx_PoolMeasToRemoveList_r16Present bool
	if Tx_PoolMeasToRemoveList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Tx_PoolMeasToAddModList_r16Present bool
	if Tx_PoolMeasToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Tx_PoolMeasToRemoveList_r16Present {
		ie.Tx_PoolMeasToRemoveList_r16 = new(Tx_PoolMeasList_r16)
		if err = ie.Tx_PoolMeasToRemoveList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Tx_PoolMeasToRemoveList_r16", err)
		}
	}
	if Tx_PoolMeasToAddModList_r16Present {
		ie.Tx_PoolMeasToAddModList_r16 = new(Tx_PoolMeasList_r16)
		if err = ie.Tx_PoolMeasToAddModList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Tx_PoolMeasToAddModList_r16", err)
		}
	}
	return nil
}
