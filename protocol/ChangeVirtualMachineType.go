//  http://manual.iij.jp/gp/gpapi/11454907.html

package protocol

type ChangeVirtualMachineTypeArg struct {
	GcServiceCode      string `gpapi:"GcServiceCode,required"`      // 対象仮想サーバのgcサービスコード
	GpServiceCode      string `gpapi:"GpServiceCode,required"`      // 対象仮想サーバを含むgpサービスコード
	VirtualMachineType string `gpapi:"VirtualMachineType,required"` // 仮想サーバの品目 (仮想サーバの性能表記)
}

type ChangeVirtualMachineTypeResponse struct {
	CommonResponse
	VirtualMachineType string // 変更後の仮想サーバの品目
	GcServiceCode      string // 対象仮想サーバのgcサービスコード
}
