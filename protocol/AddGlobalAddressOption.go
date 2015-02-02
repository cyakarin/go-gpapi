//  http://manual.iij.jp/gp/gpapi/11454903.html

package protocol

type AddGlobalAddressOptionArg struct {
	GcServiceCode string `gpapi:"GcServiceCode,required"` // 対象仮想サーバを含むgpサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象仮想サーバを含むgpサービスコード
	Num           int    `gpapi:"Num,required"`           // 追加するグローバルアドレスの数  値: 1～3 (整数)  ※IPv4, IPv6アドレスのセットで1つと数えます
}

type AddGlobalAddressOptionResponse struct {
	CommonResponse
	Num           string // 追加されたグローバルアドレスの数
	GcServiceCode string // 対象仮想サーバのgcサービスコード
}
