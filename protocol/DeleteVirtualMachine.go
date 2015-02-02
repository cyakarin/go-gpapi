//  http://manual.iij.jp/gp/gpapi/11454953.html

package protocol

type DeleteVirtualMachineArg struct {
	GcServiceCode string `gpapi:"GcServiceCode,required"` // 対象仮想サーバのgcサービスコード
	StopDate      string `gpapi:"StopDate,required"`      // 解約予定日  値: YYYYMMDD
}

type DeleteVirtualMachineResponse struct {
	CommonResponse
	StopDate      string // 解約予定日
	GcServiceCode string // 対象仮想サーバのgcサービスコード
}
