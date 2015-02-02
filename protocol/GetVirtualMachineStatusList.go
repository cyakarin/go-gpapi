//  http://manual.iij.jp/gp/gpapi/11454870.html

package protocol

type GetVirtualMachineStatusListArg struct {
	GcServiceCode []string `gpapi:"GcServiceCode,optional"` // 対象仮想サーバのgcサービスコード
	GpServiceCode string   `gpapi:"GpServiceCode,required"` // 対象仮想サーバを含むgpサービスコード
}

type GetVirtualMachineStatusListResponse struct {
	GpServiceCode string // 対象仮想サーバを含むgpサービスコード
	CommonResponse
	VirtualMachineStatusList []struct {
		Status        string // 対象仮想サーバの現在の状態 (仮想サーバの状態遷移)
		GcServiceCode string // 対象仮想サーバのgcサービスコード
	}
}
