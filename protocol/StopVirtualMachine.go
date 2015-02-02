//  http://manual.iij.jp/gp/gpapi/9544752.html

package protocol

type StopVirtualMachineArg struct {
	Force         string `gpapi:"Force,optional"`         // 強制的に仮想サーバを停止します  値: Yes, No  デフォルト値: No
	GcServiceCode string `gpapi:"GcServiceCode,required"` // 対象仮想サーバのgcサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象仮想サーバを含むgpサービスコード
}

type StopVirtualMachineResponse struct {
	Current struct {
		Status string // 対象仮想サーバの現在の状態 (仮想サーバの状態遷移)
	}
	CommonResponse
	GcServiceCode string // 対象仮想サーバのgcサービスコード
	Previous      struct {
		Status string // 対象仮想サーバのAPIリクエスト前の状態 (仮想サーバの状態遷移)
	}
}
