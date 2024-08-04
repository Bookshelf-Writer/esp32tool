package address

/* This file is automatically generated */

type SysConType byte

const (
	SYSCON_SYSCLK_CONF_REG    SysConType = 2 // R/W; Configures system clock frequency
	SYSCON_XTAL_TICK_CONF_REG SysConType = 3 // R/W; Configures the divider value of REF_TICK
	SYSCON_PLL_TICK_CONF_REG  SysConType = 4 // R/W; Configures the divider value of REF_TICK
	SYSCON_CK8M_TICK_CONF_REG SysConType = 5 // R/W; Configures the divider value of REF_TICK
	SYSCON_APLL_TICK_CONF_REG SysConType = 6 // R/W; Configures the divider value of REF_TICK
	SYSCON_DATE_REG           SysConType = 7 // R/W; Chip revision register
)

const (
	AdrBeginSYSCON_SYSCLK_CONF_REG    AddressType = 0x0000
	AdrBeginSYSCON_XTAL_TICK_CONF_REG AddressType = 0x0004
	AdrBeginSYSCON_PLL_TICK_CONF_REG  AddressType = 0x0008
	AdrBeginSYSCON_CK8M_TICK_CONF_REG AddressType = 0x000C
	AdrBeginSYSCON_APLL_TICK_CONF_REG AddressType = 0x003C
	AdrBeginSYSCON_DATE_REG           AddressType = 0x007C
)

const (
	AdrEndSYSCON_SYSCLK_CONF_REG    AddressType = 0x0004
	AdrEndSYSCON_XTAL_TICK_CONF_REG AddressType = 0x0008
	AdrEndSYSCON_PLL_TICK_CONF_REG  AddressType = 0x000C
	AdrEndSYSCON_CK8M_TICK_CONF_REG AddressType = 0x003C
	AdrEndSYSCON_APLL_TICK_CONF_REG AddressType = 0x007C
	AdrEndSYSCON_DATE_REG           AddressType = 0x007C + 5
)

var (
	AdrSYSCON_SYSCLK_CONF_REG    = AddressCellObj{AdrBeginSYSCON_SYSCLK_CONF_REG, AdrEndSYSCON_SYSCLK_CONF_REG}       // R/W; Configures system clock frequency
	AdrSYSCON_XTAL_TICK_CONF_REG = AddressCellObj{AdrBeginSYSCON_XTAL_TICK_CONF_REG, AdrEndSYSCON_XTAL_TICK_CONF_REG} // R/W; Configures the divider value of REF_TICK
	AdrSYSCON_PLL_TICK_CONF_REG  = AddressCellObj{AdrBeginSYSCON_PLL_TICK_CONF_REG, AdrEndSYSCON_PLL_TICK_CONF_REG}   // R/W; Configures the divider value of REF_TICK
	AdrSYSCON_CK8M_TICK_CONF_REG = AddressCellObj{AdrBeginSYSCON_CK8M_TICK_CONF_REG, AdrEndSYSCON_CK8M_TICK_CONF_REG} // R/W; Configures the divider value of REF_TICK
	AdrSYSCON_APLL_TICK_CONF_REG = AddressCellObj{AdrBeginSYSCON_APLL_TICK_CONF_REG, AdrEndSYSCON_APLL_TICK_CONF_REG} // R/W; Configures the divider value of REF_TICK
	AdrSYSCON_DATE_REG           = AddressCellObj{AdrBeginSYSCON_DATE_REG, AdrEndSYSCON_DATE_REG}                     // R/W; Chip revision register
)

var SysConMap = map[SysConType]*AddressCellObj{
	SYSCON_SYSCLK_CONF_REG:    &AdrSYSCON_SYSCLK_CONF_REG,
	SYSCON_XTAL_TICK_CONF_REG: &AdrSYSCON_XTAL_TICK_CONF_REG,
	SYSCON_PLL_TICK_CONF_REG:  &AdrSYSCON_PLL_TICK_CONF_REG,
	SYSCON_CK8M_TICK_CONF_REG: &AdrSYSCON_CK8M_TICK_CONF_REG,
	SYSCON_APLL_TICK_CONF_REG: &AdrSYSCON_APLL_TICK_CONF_REG,
	SYSCON_DATE_REG:           &AdrSYSCON_DATE_REG,
}
