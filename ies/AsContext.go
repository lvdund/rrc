package ies

import "rrc/utils"

// AS-Context ::= SEQUENCE
// Extensible
type AsContext struct {
	Reestablishmentinfo             *Reestablishmentinfo
	Configrestrictinfo              *Configrestrictinfoscg
	RanNotificationareainfo         *RanNotificationareainfo     `ext`
	Ueassistanceinformation         *utils.OCTETSTRING           `ext`
	Selectedbandcombinationsn       *Bandcombinationinfosn       `ext`
	ConfigrestrictinfodapsR16       *ConfigrestrictinfodapsR16   `ext`
	SidelinkueinformationnrR16      *utils.OCTETSTRING           `ext`
	SidelinkueinformationeutraR16   *utils.OCTETSTRING           `ext`
	UeassistanceinformationeutraR16 *utils.OCTETSTRING           `ext`
	UeassistanceinformationscgR16   *utils.OCTETSTRING           `ext`
	NeedforgapsinfonrR16            *NeedforgapsinfonrR16        `ext`
	ConfigrestrictinfodapsV1640     *ConfigrestrictinfodapsV1640 `ext`
	NeedforgapncsgInfonrR17         *NeedforgapncsgInfonrR17     `ext`
	NeedforgapncsgInfoeutraR17      *NeedforgapncsgInfoeutraR17  `ext`
	MbsinterestindicationR17        *utils.OCTETSTRING           `ext`
}
