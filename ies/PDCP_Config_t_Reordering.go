package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCP_Config_t_Reordering_Enum_ms0     aper.Enumerated = 0
	PDCP_Config_t_Reordering_Enum_ms1     aper.Enumerated = 1
	PDCP_Config_t_Reordering_Enum_ms2     aper.Enumerated = 2
	PDCP_Config_t_Reordering_Enum_ms4     aper.Enumerated = 3
	PDCP_Config_t_Reordering_Enum_ms5     aper.Enumerated = 4
	PDCP_Config_t_Reordering_Enum_ms8     aper.Enumerated = 5
	PDCP_Config_t_Reordering_Enum_ms10    aper.Enumerated = 6
	PDCP_Config_t_Reordering_Enum_ms15    aper.Enumerated = 7
	PDCP_Config_t_Reordering_Enum_ms20    aper.Enumerated = 8
	PDCP_Config_t_Reordering_Enum_ms30    aper.Enumerated = 9
	PDCP_Config_t_Reordering_Enum_ms40    aper.Enumerated = 10
	PDCP_Config_t_Reordering_Enum_ms50    aper.Enumerated = 11
	PDCP_Config_t_Reordering_Enum_ms60    aper.Enumerated = 12
	PDCP_Config_t_Reordering_Enum_ms80    aper.Enumerated = 13
	PDCP_Config_t_Reordering_Enum_ms100   aper.Enumerated = 14
	PDCP_Config_t_Reordering_Enum_ms120   aper.Enumerated = 15
	PDCP_Config_t_Reordering_Enum_ms140   aper.Enumerated = 16
	PDCP_Config_t_Reordering_Enum_ms160   aper.Enumerated = 17
	PDCP_Config_t_Reordering_Enum_ms180   aper.Enumerated = 18
	PDCP_Config_t_Reordering_Enum_ms200   aper.Enumerated = 19
	PDCP_Config_t_Reordering_Enum_ms220   aper.Enumerated = 20
	PDCP_Config_t_Reordering_Enum_ms240   aper.Enumerated = 21
	PDCP_Config_t_Reordering_Enum_ms260   aper.Enumerated = 22
	PDCP_Config_t_Reordering_Enum_ms280   aper.Enumerated = 23
	PDCP_Config_t_Reordering_Enum_ms300   aper.Enumerated = 24
	PDCP_Config_t_Reordering_Enum_ms500   aper.Enumerated = 25
	PDCP_Config_t_Reordering_Enum_ms750   aper.Enumerated = 26
	PDCP_Config_t_Reordering_Enum_ms1000  aper.Enumerated = 27
	PDCP_Config_t_Reordering_Enum_ms1250  aper.Enumerated = 28
	PDCP_Config_t_Reordering_Enum_ms1500  aper.Enumerated = 29
	PDCP_Config_t_Reordering_Enum_ms1750  aper.Enumerated = 30
	PDCP_Config_t_Reordering_Enum_ms2000  aper.Enumerated = 31
	PDCP_Config_t_Reordering_Enum_ms2250  aper.Enumerated = 32
	PDCP_Config_t_Reordering_Enum_ms2500  aper.Enumerated = 33
	PDCP_Config_t_Reordering_Enum_ms2750  aper.Enumerated = 34
	PDCP_Config_t_Reordering_Enum_ms3000  aper.Enumerated = 35
	PDCP_Config_t_Reordering_Enum_spare28 aper.Enumerated = 36
	PDCP_Config_t_Reordering_Enum_spare27 aper.Enumerated = 37
	PDCP_Config_t_Reordering_Enum_spare26 aper.Enumerated = 38
	PDCP_Config_t_Reordering_Enum_spare25 aper.Enumerated = 39
	PDCP_Config_t_Reordering_Enum_spare24 aper.Enumerated = 40
	PDCP_Config_t_Reordering_Enum_spare23 aper.Enumerated = 41
	PDCP_Config_t_Reordering_Enum_spare22 aper.Enumerated = 42
	PDCP_Config_t_Reordering_Enum_spare21 aper.Enumerated = 43
	PDCP_Config_t_Reordering_Enum_spare20 aper.Enumerated = 44
	PDCP_Config_t_Reordering_Enum_spare19 aper.Enumerated = 45
	PDCP_Config_t_Reordering_Enum_spare18 aper.Enumerated = 46
	PDCP_Config_t_Reordering_Enum_spare17 aper.Enumerated = 47
	PDCP_Config_t_Reordering_Enum_spare16 aper.Enumerated = 48
	PDCP_Config_t_Reordering_Enum_spare15 aper.Enumerated = 49
	PDCP_Config_t_Reordering_Enum_spare14 aper.Enumerated = 50
	PDCP_Config_t_Reordering_Enum_spare13 aper.Enumerated = 51
	PDCP_Config_t_Reordering_Enum_spare12 aper.Enumerated = 52
	PDCP_Config_t_Reordering_Enum_spare11 aper.Enumerated = 53
	PDCP_Config_t_Reordering_Enum_spare10 aper.Enumerated = 54
	PDCP_Config_t_Reordering_Enum_spare09 aper.Enumerated = 55
	PDCP_Config_t_Reordering_Enum_spare08 aper.Enumerated = 56
	PDCP_Config_t_Reordering_Enum_spare07 aper.Enumerated = 57
	PDCP_Config_t_Reordering_Enum_spare06 aper.Enumerated = 58
	PDCP_Config_t_Reordering_Enum_spare05 aper.Enumerated = 59
	PDCP_Config_t_Reordering_Enum_spare04 aper.Enumerated = 60
	PDCP_Config_t_Reordering_Enum_spare03 aper.Enumerated = 61
	PDCP_Config_t_Reordering_Enum_spare02 aper.Enumerated = 62
	PDCP_Config_t_Reordering_Enum_spare01 aper.Enumerated = 63
)

type PDCP_Config_t_Reordering struct {
	Value aper.Enumerated `lb:0,ub:63,madatory`
}

func (ie *PDCP_Config_t_Reordering) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("Encode PDCP_Config_t_Reordering", err)
	}
	return nil
}

func (ie *PDCP_Config_t_Reordering) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("Decode PDCP_Config_t_Reordering", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
