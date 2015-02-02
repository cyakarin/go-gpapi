//  http://manual.iij.jp/gp/gpapi/11455004.html

package protocol

type GetContractStatusArg struct {
	DiskDeviceName string `gpapi:"DiskDeviceName,optional"` // 追加ディスクオプションのデバイス名
	GcServiceCode  string `gpapi:"GcServiceCode,optional"`  // 仮想サーバのgcサービスコード
	GlServiceCode  string `gpapi:"GlServiceCode,optional"`  // FW+LBオプションのglサービスコード
	GpServiceCode  string `gpapi:"GpServiceCode,required"`  // 対象を含むGP契約のgpサービスコード
	GxServiceCode  string `gpapi:"GxServiceCode,optional"`  // VLANオプションのgxサービスコード
	IPv4Address    string `gpapi:"IPv4Address,optional"`    // 追加グローバルアドレスオプションのIPv4アドレス
}

type GetContractStatusResponse struct {
	GxServiceCode string // VLANオプションのサービスコード
	IPv4Address   string // 追加グローバルアドレスオプションのIPv4アドレス
	Status        string // 契約状態 (契約状態)
	CommonResponse
	GlServiceCode  string // FW+LBオプションのサービスコード
	GcServiceCode  string // 仮想サーバのサービスコード
	DiskDeviceName string // 追加ディスクオプションのデバイス名
}
