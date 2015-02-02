//  http://manual.iij.jp/gp/gpapi/11454905.html

package protocol

type AddVirtualMachinesArg struct {
	ContractNum            string `gpapi:"ContractNum,required"`            // 作成する仮想サーバ数
	DiskOption_1_DiskSpace string `gpapi:"DiskOption.1.DiskSpace,optional"` // 同時に申し込む追加ディスクオプションの容量 (1つめ)  値: 100, 300, 500, HS300
	DiskOption_2_DiskSpace string `gpapi:"DiskOption.2.DiskSpace,optional"` // 同時に申し込む追加ディスクオプションの容量 (2つめ)  値: 100, 300, 500, HS300
	GlobalAddressOptionNum string `gpapi:"GlobalAddressOptionNum,optional"` // 同時に申し込む追加グローバルアドレスオプションの個数
	GpServiceCode          string `gpapi:"GpServiceCode,required"`          // 仮想サーバを作成するgpサービスコード
	LLocationNum           string `gpapi:"LLocationNum,optional"`           // ロケーションLに作成する仮想サーバ数
	OS                     string `gpapi:"OS,required"`                     // 仮想サーバのOS (仮想サーバの性能表記)
	RLocationNum           string `gpapi:"RLocationNum,optional"`           // ロケーションRに作成する仮想サーバ数
	VirtualMachineType     string `gpapi:"VirtualMachineType,required"`     // 仮想サーバの品目 (仮想サーバの性能表記)
}

type AddVirtualMachinesResponse struct {
	GpServiceCode     string   // 対象仮想サーバを含むgpサービスコード
	GcServiceCodeList []string // 新規に作成された仮想サーバのgcサービスコードのリスト
	CommonResponse
}
