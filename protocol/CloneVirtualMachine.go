//  http://manual.iij.jp/gp/gpapi/11454852.html

package protocol

type CloneVirtualMachineArg struct {
	DestGcServiceCode string `gpapi:"DestGcServiceCode,required"` // クローニング先仮想サーバのgcサービスコード
	DestGpServiceCode string `gpapi:"DestGpServiceCode,optional"` // クローニング先仮想サーバを含むgpサービスコード
	GpServiceCode     string `gpapi:"GpServiceCode,optional"`     // 対象仮想サーバを含むgpサービスコード
	SrcGcServiceCode  string `gpapi:"SrcGcServiceCode,required"`  // クローニング元仮想サーバのgcサービスコード
	SrcGpServiceCode  string `gpapi:"SrcGpServiceCode,optional"`  // クローニング元仮想サーバを含むgpサービスコード
}

type CloneVirtualMachineResponse struct {
	CommonResponse
	Source struct {
		Current struct {
			Status string // クローニング元仮想サーバの現在の状態 (仮想サーバの状態遷移)
		}
		GcServiceCode string // クローニング元仮想サーバのgcサービスコード
		Previous      struct {
			Status string // クローニング元仮想サーバのAPIリクエスト前の状態 (仮想サーバの状態遷移)
		}
	}
	Destination struct {
		Current struct {
			Status string // クローニング先仮想サーバの現在の状態 (仮想サーバの状態遷移)
		}
		GcServiceCode string // クローニング先仮想サーバのgcサービスコード
		Previous      struct {
			Status string // クローニング先仮想サーバのAPIリクエスト前の状態 (仮想サーバの状態遷移)
		}
	}
}
