//  http://manual.iij.jp/gp/gpapi/11454995.html

package protocol

type DetachVlanArg struct {
	GcServiceCode string `gpapi:"GcServiceCode,required"` // 対象仮想サーバのgcサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象仮想サーバを含むgpサービスコード
	GxServiceCode string `gpapi:"GxServiceCode,required"` // 切断するVLANオプションのgxサービスコード
}

type DetachVlanResponse struct {
	Current struct {
		Status string // 対象仮想サーバの現在の状態 (仮想サーバの状態遷移)
	}
	GxServiceCode string // 切断するVLANオプションのgxサービスコード
	CommonResponse
	GcServiceCode string // 対象仮想サーバのgcサービスコード
	Previous      struct {
		Status string // 対象仮想サーバのAPIリクエスト前の状態 (仮想サーバの状態遷移)
	}
}
