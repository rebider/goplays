// Code generated by tool/gen.go.
// DO NOT EDIT!

package pb

import (
	"errors"
)

//打包消息
func Packet(msg interface{}) (uint32, []byte, error) {
	switch msg.(type) {
	case *SJtpayOrder:
		msg.(*SJtpayOrder).Code = 3003
		b, err := msg.(*SJtpayOrder).Marshal()
		return 3003, b, err
	case *SWxpayQuery:
		msg.(*SWxpayQuery).Code = 3004
		b, err := msg.(*SWxpayQuery).Marshal()
		return 3004, b, err
	case *SWxLogin:
		msg.(*SWxLogin).Code = 1004
		b, err := msg.(*SWxLogin).Marshal()
		return 1004, b, err
	case *SHuiYinDealerList:
		msg.(*SHuiYinDealerList).Code = 6014
		b, err := msg.(*SHuiYinDealerList).Marshal()
		return 6014, b, err
	case *SLoginOut:
		msg.(*SLoginOut).Code = 1006
		b, err := msg.(*SLoginOut).Marshal()
		return 1006, b, err
	case *STourist:
		msg.(*STourist).Code = 1010
		b, err := msg.(*STourist).Marshal()
		return 1010, b, err
	case *SUserData:
		msg.(*SUserData).Code = 1022
		b, err := msg.(*SUserData).Marshal()
		return 1022, b, err
	case *SHuiYinRecords:
		msg.(*SHuiYinRecords).Code = 6007
		b, err := msg.(*SHuiYinRecords).Marshal()
		return 6007, b, err
	case *SHuiYinGames:
		msg.(*SHuiYinGames).Code = 6008
		b, err := msg.(*SHuiYinGames).Marshal()
		return 6008, b, err
	case *SHuiYinEnterRoom:
		msg.(*SHuiYinEnterRoom).Code = 6009
		b, err := msg.(*SHuiYinEnterRoom).Marshal()
		return 6009, b, err
	case *SHuiYinLeave:
		msg.(*SHuiYinLeave).Code = 6011
		b, err := msg.(*SHuiYinLeave).Marshal()
		return 6011, b, err
	case *SWxpayOrder:
		msg.(*SWxpayOrder).Code = 3002
		b, err := msg.(*SWxpayOrder).Marshal()
		return 3002, b, err
	case *SRegist:
		msg.(*SRegist).Code = 1002
		b, err := msg.(*SRegist).Marshal()
		return 1002, b, err
	case *SGetCurrency:
		msg.(*SGetCurrency).Code = 1024
		b, err := msg.(*SGetCurrency).Marshal()
		return 1024, b, err
	case *SPushCurrency:
		msg.(*SPushCurrency).Code = 1028
		b, err := msg.(*SPushCurrency).Marshal()
		return 1028, b, err
	case *SHuiYinRoomRoles:
		msg.(*SHuiYinRoomRoles).Code = 6010
		b, err := msg.(*SHuiYinRoomRoles).Marshal()
		return 6010, b, err
	case *SHuiYinDeskState:
		msg.(*SHuiYinDeskState).Code = 6020
		b, err := msg.(*SHuiYinDeskState).Marshal()
		return 6020, b, err
	case *SGetLastWins:
		msg.(*SGetLastWins).Code = 6023
		b, err := msg.(*SGetLastWins).Marshal()
		return 6023, b, err
	case *SApplePay:
		msg.(*SApplePay).Code = 3006
		b, err := msg.(*SApplePay).Marshal()
		return 3006, b, err
	case *SBroadcast:
		msg.(*SBroadcast).Code = 2006
		b, err := msg.(*SBroadcast).Marshal()
		return 2006, b, err
	case *SLogin:
		msg.(*SLogin).Code = 1000
		b, err := msg.(*SLogin).Marshal()
		return 1000, b, err
	case *SHuiYinDealer:
		msg.(*SHuiYinDealer).Code = 6013
		b, err := msg.(*SHuiYinDealer).Marshal()
		return 6013, b, err
	case *SHuiYinRoomList:
		msg.(*SHuiYinRoomList).Code = 6018
		b, err := msg.(*SHuiYinRoomList).Marshal()
		return 6018, b, err
	case *SGetTrend:
		msg.(*SGetTrend).Code = 6022
		b, err := msg.(*SGetTrend).Marshal()
		return 6022, b, err
	case *SChatText:
		msg.(*SChatText).Code = 2003
		b, err := msg.(*SChatText).Marshal()
		return 2003, b, err
	case *SNotice:
		msg.(*SNotice).Code = 2008
		b, err := msg.(*SNotice).Marshal()
		return 2008, b, err
	case *SPing:
		msg.(*SPing).Code = 1050
		b, err := msg.(*SPing).Marshal()
		return 1050, b, err
	case *SPk10Record:
		msg.(*SPk10Record).Code = 6015
		b, err := msg.(*SPk10Record).Marshal()
		return 6015, b, err
	case *SHuiYinGameover:
		msg.(*SHuiYinGameover).Code = 6017
		b, err := msg.(*SHuiYinGameover).Marshal()
		return 6017, b, err
	case *SHuiYinSit:
		msg.(*SHuiYinSit).Code = 6019
		b, err := msg.(*SHuiYinSit).Marshal()
		return 6019, b, err
	case *SHuiYinPushBeDealer:
		msg.(*SHuiYinPushBeDealer).Code = 6027
		b, err := msg.(*SHuiYinPushBeDealer).Marshal()
		return 6027, b, err
	case *SBuy:
		msg.(*SBuy).Code = 3000
		b, err := msg.(*SBuy).Marshal()
		return 3000, b, err
	case *SChatVoice:
		msg.(*SChatVoice).Code = 2004
		b, err := msg.(*SChatVoice).Marshal()
		return 2004, b, err
	case *SHuiYinPushDealer:
		msg.(*SHuiYinPushDealer).Code = 6021
		b, err := msg.(*SHuiYinPushDealer).Marshal()
		return 6021, b, err
	case *SHuiYinProfit:
		msg.(*SHuiYinProfit).Code = 6025
		b, err := msg.(*SHuiYinProfit).Marshal()
		return 6025, b, err
	case *SShop:
		msg.(*SShop).Code = 3010
		b, err := msg.(*SShop).Marshal()
		return 3010, b, err
	case *SResetPwd:
		msg.(*SResetPwd).Code = 1008
		b, err := msg.(*SResetPwd).Marshal()
		return 1008, b, err
	case *SHuiYinRoomBet:
		msg.(*SHuiYinRoomBet).Code = 6012
		b, err := msg.(*SHuiYinRoomBet).Marshal()
		return 6012, b, err
	case *SHuiYinCamein:
		msg.(*SHuiYinCamein).Code = 6016
		b, err := msg.(*SHuiYinCamein).Marshal()
		return 6016, b, err
	case *SGetOpenResult:
		msg.(*SGetOpenResult).Code = 6024
		b, err := msg.(*SGetOpenResult).Marshal()
		return 6024, b, err
	case *SHuiYinDeskBetInfo:
		msg.(*SHuiYinDeskBetInfo).Code = 6026
		b, err := msg.(*SHuiYinDeskBetInfo).Marshal()
		return 6026, b, err
	default:
		return 0, []byte{}, errors.New("unknown message")
	}
}