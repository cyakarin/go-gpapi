//  http://manual.iij.jp/gp/gpapi/11454884.html

package protocol

type InitializeVirtualMachineArg struct {
	GcServiceCode string `gpapi:"GcServiceCode,required"` // 対象仮想サーバのgcサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象仮想サーバを含むgpサービスコード
	VarDeviceName string `gpapi:"VarDeviceName,optional"` // 追加ディスクのマウント先  値: /dev/sdb, /dev/xvdb
}

type InitializeVirtualMachineResponse struct {
	Current struct {
		Status string // 対象仮想サーバの現在の状態 (仮想サーバの状態遷移)
	}
	CommonResponse
	GcServiceCode string // 対象仮想サーバのgcサービスコード
	Previous      struct {
		Status string // 対象仮想サーバのAPIリクエスト前の状態 (仮想サーバの状態遷移)
	}
}
