//  http://manual.iij.jp/gp/gpapi/11454873.html

package protocol

type ImportRootSshPublicKeyArg struct {
	GcServiceCode string `gpapi:"GcServiceCode,required"` // 対象仮想サーバのgcサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象仮想サーバを含むgpサービスコード
	PublicKey     string `gpapi:"PublicKey,required"`     // 追加するssh公開鍵
}

type ImportRootSshPublicKeyResponse struct {
	Current struct {
		Status string // 対象仮想サーバの現在の状態 (仮想サーバの状態遷移)
	}
	PublicKey string // 追加したssh公開鍵
	CommonResponse
	GcServiceCode string // 対象仮想サーバのgcサービスコード
	Previous      struct {
		Status string // 対象仮想サーバのAPIリクエスト前の状態 (仮想サーバの状態遷移)
	}
}
