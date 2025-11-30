package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type VarMeasReport struct {
	MeasId                        MeasId                   `madatory`
	CellsTriggeredList            *CellsTriggeredList      `optional`
	NumberOfReportsSent           int64                    `madatory`
	Cli_TriggeredList_r16         *CLI_TriggeredList_r16   `optional`
	Tx_PoolMeasToAddModListNR_r16 *Tx_PoolMeasList_r16     `optional`
	RelaysTriggeredList_r17       *RelaysTriggeredList_r17 `optional`
}

func (ie *VarMeasReport) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.CellsTriggeredList != nil, ie.Cli_TriggeredList_r16 != nil, ie.Tx_PoolMeasToAddModListNR_r16 != nil, ie.RelaysTriggeredList_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.MeasId.Encode(w); err != nil {
		return utils.WrapError("Encode MeasId", err)
	}
	if ie.CellsTriggeredList != nil {
		if err = ie.CellsTriggeredList.Encode(w); err != nil {
			return utils.WrapError("Encode CellsTriggeredList", err)
		}
	}
	if err = w.WriteInteger(ie.NumberOfReportsSent, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteInteger NumberOfReportsSent", err)
	}
	if ie.Cli_TriggeredList_r16 != nil {
		if err = ie.Cli_TriggeredList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Cli_TriggeredList_r16", err)
		}
	}
	if ie.Tx_PoolMeasToAddModListNR_r16 != nil {
		if err = ie.Tx_PoolMeasToAddModListNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Tx_PoolMeasToAddModListNR_r16", err)
		}
	}
	if ie.RelaysTriggeredList_r17 != nil {
		if err = ie.RelaysTriggeredList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode RelaysTriggeredList_r17", err)
		}
	}
	return nil
}

func (ie *VarMeasReport) Decode(r *aper.AperReader) error {
	var err error
	var CellsTriggeredListPresent bool
	if CellsTriggeredListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Cli_TriggeredList_r16Present bool
	if Cli_TriggeredList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Tx_PoolMeasToAddModListNR_r16Present bool
	if Tx_PoolMeasToAddModListNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var RelaysTriggeredList_r17Present bool
	if RelaysTriggeredList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.MeasId.Decode(r); err != nil {
		return utils.WrapError("Decode MeasId", err)
	}
	if CellsTriggeredListPresent {
		ie.CellsTriggeredList = new(CellsTriggeredList)
		if err = ie.CellsTriggeredList.Decode(r); err != nil {
			return utils.WrapError("Decode CellsTriggeredList", err)
		}
	}
	var tmp_int_NumberOfReportsSent int64
	if tmp_int_NumberOfReportsSent, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadInteger NumberOfReportsSent", err)
	}
	ie.NumberOfReportsSent = tmp_int_NumberOfReportsSent
	if Cli_TriggeredList_r16Present {
		ie.Cli_TriggeredList_r16 = new(CLI_TriggeredList_r16)
		if err = ie.Cli_TriggeredList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Cli_TriggeredList_r16", err)
		}
	}
	if Tx_PoolMeasToAddModListNR_r16Present {
		ie.Tx_PoolMeasToAddModListNR_r16 = new(Tx_PoolMeasList_r16)
		if err = ie.Tx_PoolMeasToAddModListNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Tx_PoolMeasToAddModListNR_r16", err)
		}
	}
	if RelaysTriggeredList_r17Present {
		ie.RelaysTriggeredList_r17 = new(RelaysTriggeredList_r17)
		if err = ie.RelaysTriggeredList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode RelaysTriggeredList_r17", err)
		}
	}
	return nil
}
