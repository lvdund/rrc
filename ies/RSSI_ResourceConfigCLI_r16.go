package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RSSI_ResourceConfigCLI_r16 struct {
	Rssi_ResourceId_r16           RSSI_ResourceId_r16           `madatory`
	Rssi_SCS_r16                  SubcarrierSpacing             `madatory`
	StartPRB_r16                  int64                         `lb:0,ub:2169,madatory`
	NrofPRBs_r16                  int64                         `lb:4,ub:maxNrofPhysicalResourceBlocksPlus1,madatory`
	StartPosition_r16             int64                         `lb:0,ub:13,madatory`
	NrofSymbols_r16               int64                         `lb:1,ub:14,madatory`
	Rssi_PeriodicityAndOffset_r16 RSSI_PeriodicityAndOffset_r16 `madatory`
	RefServCellIndex_r16          *ServCellIndex                `optional`
}

func (ie *RSSI_ResourceConfigCLI_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.RefServCellIndex_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Rssi_ResourceId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Rssi_ResourceId_r16", err)
	}
	if err = ie.Rssi_SCS_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Rssi_SCS_r16", err)
	}
	if err = w.WriteInteger(ie.StartPRB_r16, &uper.Constraint{Lb: 0, Ub: 2169}, false); err != nil {
		return utils.WrapError("WriteInteger StartPRB_r16", err)
	}
	if err = w.WriteInteger(ie.NrofPRBs_r16, &uper.Constraint{Lb: 4, Ub: maxNrofPhysicalResourceBlocksPlus1}, false); err != nil {
		return utils.WrapError("WriteInteger NrofPRBs_r16", err)
	}
	if err = w.WriteInteger(ie.StartPosition_r16, &uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("WriteInteger StartPosition_r16", err)
	}
	if err = w.WriteInteger(ie.NrofSymbols_r16, &uper.Constraint{Lb: 1, Ub: 14}, false); err != nil {
		return utils.WrapError("WriteInteger NrofSymbols_r16", err)
	}
	if err = ie.Rssi_PeriodicityAndOffset_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Rssi_PeriodicityAndOffset_r16", err)
	}
	if ie.RefServCellIndex_r16 != nil {
		if err = ie.RefServCellIndex_r16.Encode(w); err != nil {
			return utils.WrapError("Encode RefServCellIndex_r16", err)
		}
	}
	return nil
}

func (ie *RSSI_ResourceConfigCLI_r16) Decode(r *uper.UperReader) error {
	var err error
	var RefServCellIndex_r16Present bool
	if RefServCellIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Rssi_ResourceId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Rssi_ResourceId_r16", err)
	}
	if err = ie.Rssi_SCS_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Rssi_SCS_r16", err)
	}
	var tmp_int_StartPRB_r16 int64
	if tmp_int_StartPRB_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2169}, false); err != nil {
		return utils.WrapError("ReadInteger StartPRB_r16", err)
	}
	ie.StartPRB_r16 = tmp_int_StartPRB_r16
	var tmp_int_NrofPRBs_r16 int64
	if tmp_int_NrofPRBs_r16, err = r.ReadInteger(&uper.Constraint{Lb: 4, Ub: maxNrofPhysicalResourceBlocksPlus1}, false); err != nil {
		return utils.WrapError("ReadInteger NrofPRBs_r16", err)
	}
	ie.NrofPRBs_r16 = tmp_int_NrofPRBs_r16
	var tmp_int_StartPosition_r16 int64
	if tmp_int_StartPosition_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("ReadInteger StartPosition_r16", err)
	}
	ie.StartPosition_r16 = tmp_int_StartPosition_r16
	var tmp_int_NrofSymbols_r16 int64
	if tmp_int_NrofSymbols_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 14}, false); err != nil {
		return utils.WrapError("ReadInteger NrofSymbols_r16", err)
	}
	ie.NrofSymbols_r16 = tmp_int_NrofSymbols_r16
	if err = ie.Rssi_PeriodicityAndOffset_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Rssi_PeriodicityAndOffset_r16", err)
	}
	if RefServCellIndex_r16Present {
		ie.RefServCellIndex_r16 = new(ServCellIndex)
		if err = ie.RefServCellIndex_r16.Decode(r); err != nil {
			return utils.WrapError("Decode RefServCellIndex_r16", err)
		}
	}
	return nil
}
