package d1exe

import (
	"bytes"

	"github.com/decomp/exp/bin"
	"github.com/pkg/errors"
	"github.com/sanctuary/formats/level/miniset"
)

// ref: v1.09b.
const (
	// l1 minisets.
	L1StairsUp1Addr        = 0x479EC8 // STAIRSUP
	L1StairsUp2Addr        = 0x479EEC // L5STAIRSUP
	L1StairsDownAddr       = 0x479F10 // STAIRSDOWN
	L1CandlestickAddr      = 0x479F2C // LAMPS
	L1StairsDownPoisonAddr = 0x479F38 // PWATERIN
	// l2 minisets.
	L2VertArch1Addr    = 0x4849E4 // VARCH1
	L2VertArch2Addr    = 0x4849F8 // VARCH2
	L2VertArch3Addr    = 0x484A0C // VARCH3
	L2VertArch4Addr    = 0x484A20 // VARCH4
	L2VertArch5Addr    = 0x484A34 // VARCH5
	L2VertArch6Addr    = 0x484A48 // VARCH6
	L2VertArch7Addr    = 0x484A5C // VARCH7
	L2VertArch8Addr    = 0x484A70 // VARCH8
	L2VertArch9Addr    = 0x484A84 // VARCH9
	L2VertArch10Addr   = 0x484A98 // VARCH10
	L2VertArch11Addr   = 0x484AAC // VARCH11
	L2VertArch12Addr   = 0x484AC0 // VARCH12
	L2VertArch13Addr   = 0x484AD4 // VARCH13
	L2VertArch14Addr   = 0x484AE8 // VARCH14
	L2VertArch15Addr   = 0x484AFC // VARCH15
	L2VertArch16Addr   = 0x484B10 // VARCH16
	L2VertArch17Addr   = 0x484B24 // VARCH17
	L2VertArch18Addr   = 0x484B34 // VARCH18
	L2VertArch19Addr   = 0x484B44 // VARCH19
	L2VertArch20Addr   = 0x484B54 // VARCH20
	L2VertArch21Addr   = 0x484B64 // VARCH21
	L2VertArch22Addr   = 0x484B74 // VARCH22
	L2VertArch23Addr   = 0x484B84 // VARCH23
	L2VertArch24Addr   = 0x484B94 // VARCH24
	L2VertArch25Addr   = 0x484BA4 // VARCH25
	L2VertArch26Addr   = 0x484BB8 // VARCH26
	L2VertArch27Addr   = 0x484BCC // VARCH27
	L2VertArch28Addr   = 0x484BE0 // VARCH28
	L2VertArch29Addr   = 0x484BF4 // VARCH29
	L2VertArch30Addr   = 0x484C08 // VARCH30
	L2VertArch31Addr   = 0x484C1C // VARCH31
	L2VertArch32Addr   = 0x484C30 // VARCH32
	L2VertArch33Addr   = 0x484C44 // VARCH33
	L2VertArch34Addr   = 0x484C58 // VARCH34
	L2VertArch35Addr   = 0x484C6C // VARCH35
	L2VertArch36Addr   = 0x484C80 // VARCH36
	L2VertArch37Addr   = 0x484C94 // VARCH37
	L2VertArch38Addr   = 0x484CA8 // VARCH38
	L2VertArch39Addr   = 0x484CBC // VARCH39
	L2VertArch40Addr   = 0x484CD0 // VARCH40
	L2HorizArch1Addr   = 0x484CE4 // HARCH1
	L2HorizArch2Addr   = 0x484CF4 // HARCH2
	L2HorizArch3Addr   = 0x484D04 // HARCH3
	L2HorizArch4Addr   = 0x484D14 // HARCH4
	L2HorizArch5Addr   = 0x484D24 // HARCH5
	L2HorizArch6Addr   = 0x484D34 // HARCH6
	L2HorizArch7Addr   = 0x484D44 // HARCH7
	L2HorizArch8Addr   = 0x484D54 // HARCH8
	L2HorizArch9Addr   = 0x484D64 // HARCH9
	L2HorizArch10Addr  = 0x484D74 // HARCH10
	L2HorizArch11Addr  = 0x484D84 // HARCH11
	L2HorizArch12Addr  = 0x484D94 // HARCH12
	L2HorizArch13Addr  = 0x484DA4 // HARCH13
	L2HorizArch14Addr  = 0x484DB4 // HARCH14
	L2HorizArch15Addr  = 0x484DC4 // HARCH15
	L2HorizArch16Addr  = 0x484DD4 // HARCH16
	L2HorizArch17Addr  = 0x484DE4 // HARCH17
	L2HorizArch18Addr  = 0x484DF4 // HARCH18
	L2HorizArch19Addr  = 0x484E04 // HARCH19
	L2HorizArch20Addr  = 0x484E14 // HARCH20
	L2HorizArch21Addr  = 0x484E24 // HARCH21
	L2HorizArch22Addr  = 0x484E34 // HARCH22
	L2HorizArch23Addr  = 0x484E44 // HARCH23
	L2HorizArch24Addr  = 0x484E54 // HARCH24
	L2HorizArch25Addr  = 0x484E64 // HARCH25
	L2HorizArch26Addr  = 0x484E74 // HARCH26
	L2HorizArch27Addr  = 0x484E84 // HARCH27
	L2HorizArch28Addr  = 0x484E94 // HARCH28
	L2HorizArch29Addr  = 0x484EA4 // HARCH29
	L2HorizArch30Addr  = 0x484EB4 // HARCH30
	L2HorizArch31Addr  = 0x484EC4 // HARCH31
	L2HorizArch32Addr  = 0x484ED4 // HARCH32
	L2HorizArch33Addr  = 0x484EE4 // HARCH33
	L2HorizArch34Addr  = 0x484EF4 // HARCH34
	L2HorizArch35Addr  = 0x484F04 // HARCH35
	L2HorizArch36Addr  = 0x484F14 // HARCH36
	L2HorizArch37Addr  = 0x484F24 // HARCH37
	L2HorizArch38Addr  = 0x484F34 // HARCH38
	L2HorizArch39Addr  = 0x484F44 // HARCH39
	L2HorizArch40Addr  = 0x484F54 // HARCH40
	L2StairsUpAddr     = 0x484F64 // USTAIRS
	L2StairsDownAddr   = 0x484F88 // DSTAIRS
	L2StairsToTownAddr = 0x484FAC // WARPSTAIRS
	// TODO: find better name.
	L2CrushedColumnAddr = 0x484FD0 // CRUSHCOL
	L2Big1Addr          = 0x484FE4 // BIG1
	L2Big2Addr          = 0x484FF0 // BIG2
	L2Big3Addr          = 0x484FFC // BIG3
	L2Big4Addr          = 0x485004 // BIG4
	L2Big5Addr          = 0x48500C // BIG5
	L2Big6Addr          = 0x485018 // BIG6
	L2Big7Addr          = 0x485020 // BIG7
	L2Big8Addr          = 0x485028 // BIG8
	L2Big9Addr          = 0x485034 // BIG9
	L2Big10Addr         = 0x485040 // BIG10
	L2Ruins1Addr        = 0x48504C // RUINS1
	L2Ruins2Addr        = 0x485050 // RUINS2
	L2Ruins3Addr        = 0x485054 // RUINS3
	L2Ruins4Addr        = 0x485058 // RUINS4
	L2Ruins5Addr        = 0x48505C // RUINS5
	L2Ruins6Addr        = 0x485060 // RUINS6
	L2Ruins7Addr        = 0x485064 // RUINS7
	L2Pancreas1Addr     = 0x485068 // PANCREAS1
	L2Pancreas2Addr     = 0x485088 // PANCREAS2
	L2CenterDoor1Addr   = 0x4850A8 // CTRDOOR1
	L2CenterDoor2Addr   = 0x4850BC // CTRDOOR2
	L2CenterDoor3Addr   = 0x4850D0 // CTRDOOR3
	L2CenterDoor4Addr   = 0x4850E4 // CTRDOOR4
	L2CenterDoor5Addr   = 0x4850F8 // CTRDOOR5
	L2CenterDoor6Addr   = 0x48510C // CTRDOOR6
	L2CenterDoor7Addr   = 0x485120 // CTRDOOR7
	L2CenterDoor8Addr   = 0x485134 // CTRDOOR8
)

// parseMinisets parses the miniset DUN files contained within the executable.
func (exe *Executable) parseMinisets() error {
	addrs := []uint32{
		// l1.
		L1StairsUp1Addr,
		L1StairsUp2Addr,
		L1StairsDownAddr,
		L1CandlestickAddr,
		L1StairsDownPoisonAddr,
		// l2.
		L2VertArch1Addr,
		L2VertArch2Addr,
		L2VertArch3Addr,
		L2VertArch4Addr,
		L2VertArch5Addr,
		L2VertArch6Addr,
		L2VertArch7Addr,
		L2VertArch8Addr,
		L2VertArch9Addr,
		L2VertArch10Addr,
		L2VertArch11Addr,
		L2VertArch12Addr,
		L2VertArch13Addr,
		L2VertArch14Addr,
		L2VertArch15Addr,
		L2VertArch16Addr,
		L2VertArch17Addr,
		L2VertArch18Addr,
		L2VertArch19Addr,
		L2VertArch20Addr,
		L2VertArch21Addr,
		L2VertArch22Addr,
		L2VertArch23Addr,
		L2VertArch24Addr,
		L2VertArch25Addr,
		L2VertArch26Addr,
		L2VertArch27Addr,
		L2VertArch28Addr,
		L2VertArch29Addr,
		L2VertArch30Addr,
		L2VertArch31Addr,
		L2VertArch32Addr,
		L2VertArch33Addr,
		L2VertArch34Addr,
		L2VertArch35Addr,
		L2VertArch36Addr,
		L2VertArch37Addr,
		L2VertArch38Addr,
		L2VertArch39Addr,
		L2VertArch40Addr,
		L2HorizArch1Addr,
		L2HorizArch2Addr,
		L2HorizArch3Addr,
		L2HorizArch4Addr,
		L2HorizArch5Addr,
		L2HorizArch6Addr,
		L2HorizArch7Addr,
		L2HorizArch8Addr,
		L2HorizArch9Addr,
		L2HorizArch10Addr,
		L2HorizArch11Addr,
		L2HorizArch12Addr,
		L2HorizArch13Addr,
		L2HorizArch14Addr,
		L2HorizArch15Addr,
		L2HorizArch16Addr,
		L2HorizArch17Addr,
		L2HorizArch18Addr,
		L2HorizArch19Addr,
		L2HorizArch20Addr,
		L2HorizArch21Addr,
		L2HorizArch22Addr,
		L2HorizArch23Addr,
		L2HorizArch24Addr,
		L2HorizArch25Addr,
		L2HorizArch26Addr,
		L2HorizArch27Addr,
		L2HorizArch28Addr,
		L2HorizArch29Addr,
		L2HorizArch30Addr,
		L2HorizArch31Addr,
		L2HorizArch32Addr,
		L2HorizArch33Addr,
		L2HorizArch34Addr,
		L2HorizArch35Addr,
		L2HorizArch36Addr,
		L2HorizArch37Addr,
		L2HorizArch38Addr,
		L2HorizArch39Addr,
		L2HorizArch40Addr,
		L2StairsUpAddr,
		L2StairsDownAddr,
		L2StairsToTownAddr,
		L2CrushedColumnAddr,
		L2Big1Addr,
		L2Big2Addr,
		L2Big3Addr,
		L2Big4Addr,
		L2Big5Addr,
		L2Big6Addr,
		L2Big7Addr,
		L2Big8Addr,
		L2Big9Addr,
		L2Big10Addr,
		L2Ruins1Addr,
		L2Ruins2Addr,
		L2Ruins3Addr,
		L2Ruins4Addr,
		L2Ruins5Addr,
		L2Ruins6Addr,
		L2Ruins7Addr,
		L2Pancreas1Addr,
		L2Pancreas2Addr,
		L2CenterDoor1Addr,
		L2CenterDoor2Addr,
		L2CenterDoor3Addr,
		L2CenterDoor4Addr,
		L2CenterDoor5Addr,
		L2CenterDoor6Addr,
		L2CenterDoor7Addr,
		L2CenterDoor8Addr,
	}
	for _, addr := range addrs {
		m, err := exe.parseMiniset(addr)
		if err != nil {
			return errors.WithStack(err)
		}
		exe.Minisets[addr] = m
	}
	return nil
}

// parseMinisets parses the miniset DUN file at the given address contained
// within the executable.
func (exe *Executable) parseMiniset(minisetAddr uint32) (*miniset.Miniset, error) {
	dbg.Printf("parsing miniset at 0x%08X", minisetAddr)
	data := exe.file.Data(bin.Address(minisetAddr))
	r := bytes.NewReader(data)
	m, err := miniset.Parse(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return m, nil
}
