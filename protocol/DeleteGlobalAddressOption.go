//  http://manual.iij.jp/gp/gpapi/11454927.html

package protocol

type DeleteGlobalAddressOptionArg struct {
	GcServiceCode string `gpapi:"GcServiceCode,required"` // 対象仮想サーバのgcサービスコード
	IPv4Address   string `gpapi:"IPv4Address,required"`   // 解約対象のIPv4アドレス
	StopDate      string `gpapi:"StopDate,required"`      // 解約予定日  値: YYYYMMDD
}

type DeleteGlobalAddressOptionResponse struct {
	IPv4Address string // 解約対象のIPv4アドレス
	CommonResponse
	StopDate      string // 解約予定日
	GcServiceCode string // 対象仮想サーバのgcサービスコード
}
