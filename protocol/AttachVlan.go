//  http://manual.iij.jp/gp/gpapi/11454992.html

package protocol

type AttachVlanArg struct {
	GcServiceCode string `gpapi:"GcServiceCode,required"` // 対象仮想サーバのgcサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象仮想サーバを含むgpサービスコード
	GxServiceCode string `gpapi:"GxServiceCode,required"` // 接続するVLANオプションのgxサービスコード
}

type AttachVlanResponse struct {
	Current struct {
		Status string // 対象仮想サーバの現在の状態 (仮想サーバの状態遷移)
	}
	GxServiceCode string // 接続するVLANオプションのgxサービスコード
	CommonResponse
	GcServiceCode string // 対象仮想サーバのgcサービスコード
	Previous      struct {
		Status string // 対象仮想サーバのAPIリクエスト前の状態 (仮想サーバの状態遷移)
	}
}
