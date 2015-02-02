//  http://manual.iij.jp/gp/gpapi/11454910.html

package protocol

type DeleteDiskOptionArg struct {
	DiskDeviceName string `gpapi:"DiskDeviceName,required"` // 解約対象の仮想ディスクのデバイス名
	GcServiceCode  string `gpapi:"GcServiceCode,required"`  // 対象仮想サーバのgcサービスコード
	StopDate       string `gpapi:"StopDate,required"`       // 解約予定日  値: YYYYMMDD
}

type DeleteDiskOptionResponse struct {
	CommonResponse
	StopDate       string // 解約予定日
	GcServiceCode  string // 対象仮想サーバのgcサービスコード
	DiskDeviceName string // 解約対象の仮想ディスクのデバイス名
}
