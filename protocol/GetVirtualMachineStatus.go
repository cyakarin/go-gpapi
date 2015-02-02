//  http://manual.iij.jp/gp/gpapi/11454867.html

package protocol

type GetVirtualMachineStatusArg struct {
	GcServiceCode string `gpapi:"GcServiceCode,required"` // 対象仮想サーバのgcサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象仮想サーバを含むgpサービスコード
}

type GetVirtualMachineStatusResponse struct {
	Status string // 対象仮想サーバの現在の状態 (仮想サーバの状態遷移)
	CommonResponse
	GcServiceCode string // 対象仮想サーバのgcサービスコード
}
