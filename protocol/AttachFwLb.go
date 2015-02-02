//  http://manual.iij.jp/gp/gpapi/12442746.html

package protocol

type AttachFwLbArg struct {
	GcServiceCode string `gpapi:"GcServiceCode,required"` // 対象仮想サーバのgcサービスコード
	GlServiceCode string `gpapi:"GlServiceCode,required"` // 接続するFW+LBオプションのglサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象仮想サーバ・FW+LBオプションを含むgpサービスコード
}

type AttachFwLbResponse struct {
	Current struct {
		Status string // 対象仮想サーバの現在の状態 (仮想サーバの状態遷移)
	}
	CommonResponse
	GlServiceCode string // 接続するFW+LBオプションのglサービスコード
	GcServiceCode string // 対象仮想サーバのgcサービスコード
	Previous      struct {
		Status string // 対象仮想サーバのAPIリクエスト前の状態(仮想サーバの状態遷移)
	}
}
