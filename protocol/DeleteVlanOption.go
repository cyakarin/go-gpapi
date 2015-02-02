//  http://manual.iij.jp/gp/gpapi/11454999.html

package protocol

type DeleteVlanOptionArg struct {
	GxServiceCode string `gpapi:"GxServiceCode,required"` // 対象仮想サーバのgxサービスコード
	StopDate      string `gpapi:"StopDate,required"`      // 解約予定日  値: YYYYMMDD
}

type DeleteVlanOptionResponse struct {
	GxServiceCode string // 解約するVLANオプションのサービスコード
	CommonResponse
	StopDate string // 解約予定日
}
