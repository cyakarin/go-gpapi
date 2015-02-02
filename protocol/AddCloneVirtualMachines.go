//  http://manual.iij.jp/gp/gpapi/11454899.html

package protocol

type AddCloneVirtualMachinesArg struct {
	ContractNum        int    `gpapi:"ContractNum,required"`        // 新しく作成する仮想サーバ数  値: 1～ (整数)
	GpServiceCode      string `gpapi:"GpServiceCode,required"`      // 対象仮想サーバを含むgpサービスコード
	LLocationNum       string `gpapi:"LLocationNum,optional"`       // ロケーションLに作成する仮想サーバ数
	RLocationNum       string `gpapi:"RLocationNum,optional"`       // ロケーションRに作成する仮想サーバ数
	SrcGcServiceCode   string `gpapi:"SrcGcServiceCode,required"`   // クローニング元仮想サーバのgcサービスコード
	SrcGpServiceCode   string `gpapi:"SrcGpServiceCode,optional"`   // クローニング元仮想サーバを含むgpサービスコード
	VirtualMachineType string `gpapi:"VirtualMachineType,required"` // 仮想サーバの品目 (仮想サーバの性能表記)
}

type AddCloneVirtualMachinesResponse struct {
	GpServiceCode     string   // 対象仮想サーバを含むgpサービスコード
	GcServiceCodeList []string // 新規に作成された仮想サーバのgcサービスコードのリスト
	SrcGcServiceCode  string   // クローニング元仮想サーバのgcサービスコード
	CommonResponse
}
