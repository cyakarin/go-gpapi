//  http://manual.iij.jp/gp/gpapi/11454901.html

package protocol

type AddDiskOptionArg struct {
	DiskSpace     string `gpapi:"DiskSpace,required"`     // 追加するディスクの容量 (種類)  値: 100, 300, 500, HS300
	GcServiceCode string `gpapi:"GcServiceCode,required"` // 対象仮想サーバのgcサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象仮想サーバを含むgpサービスコード
}

type AddDiskOptionResponse struct {
	CommonResponse
	GcServiceCode string // 対象仮想サーバのgcサービスコード
	DiskSpace     string // 追加されたディスクの容量(種類)
}
