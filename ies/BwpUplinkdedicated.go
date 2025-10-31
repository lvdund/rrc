package ies

import "rrc/utils"

// BWP-UplinkDedicated ::= SEQUENCE
// Extensible
type BwpUplinkdedicated struct {
	PucchConfig                                        *utils.Setuprelease[PucchConfig]
	PuschConfig                                        *utils.Setuprelease[PuschConfig]
	Configuredgrantconfig                              *utils.Setuprelease[Configuredgrantconfig]
	SrsConfig                                          *utils.Setuprelease[SrsConfig]
	Beamfailurerecoveryconfig                          *utils.Setuprelease[Beamfailurerecoveryconfig]
	SlPucchConfigR16                                   *utils.Setuprelease[PucchConfig]                    `ext`
	CpExtensionc2R16                                   *utils.INTEGER                                      `lb:0,ub:28,ext`
	CpExtensionc3R16                                   *utils.INTEGER                                      `lb:0,ub:28,ext`
	UseinterlacepucchPuschR16                          *BwpUplinkdedicatedUseinterlacepucchPuschR16        `ext`
	PucchConfigurationlistR16                          *utils.Setuprelease[PucchConfigurationlistR16]      `ext`
	LbtFailurerecoveryconfigR16                        *utils.Setuprelease[LbtFailurerecoveryconfigR16]    `ext`
	ConfiguredgrantconfigtoaddmodlistR16               *ConfiguredgrantconfigtoaddmodlistR16               `ext`
	ConfiguredgrantconfigtoreleaselistR16              *ConfiguredgrantconfigtoreleaselistR16              `ext`
	Configuredgrantconfigtype2deactivationstatelistR16 *Configuredgrantconfigtype2deactivationstatelistR16 `ext`
	UlTciStatelistR17                                  *BwpUplinkdedicatedUlTciStatelistR17                `ext`
	UlPowercontrolR17                                  *UplinkPowercontrolidR17                            `ext`
	PucchConfigurationlistmulticast1R17                *utils.Setuprelease[PucchConfigurationlistR16]      `ext`
	PucchConfigurationlistmulticast2R17                *utils.Setuprelease[PucchConfigurationlistR16]      `ext`
	PucchConfigmulticast1R17                           *utils.Setuprelease[PucchConfig]                    `ext`
	PucchConfigmulticast2R17                           *utils.Setuprelease[PucchConfig]                    `ext`
	PathlossreferencerstoaddmodlistR17                 *[]PathlossreferencersR17                           `lb:1,ub:maxNrofPathlossReferenceRSsR17,ext`
	PathlossreferencerstoreleaselistR17                *[]PathlossreferencersIdR17                         `lb:1,ub:maxNrofPathlossReferenceRSsR17,ext`
}
