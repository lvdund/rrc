package ies

// Uplink-powerControl-r17 ::= SEQUENCE
type UplinkPowercontrolR17 struct {
	UlPowercontrolidR17   UplinkPowercontrolidR17
	P0alphasetforpuschR17 *P0alphasetR17
	P0alphasetforpucchR17 *P0alphasetR17
	P0alphasetforsrsR17   *P0alphasetR17
}
