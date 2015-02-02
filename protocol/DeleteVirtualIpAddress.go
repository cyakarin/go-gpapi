//  http://manual.iij.jp/gp/gpapi/11454960.html

package protocol

type DeleteVirtualIpAddressArg struct {
	GlServiceCode string `gpapi:"GlServiceCode,required"` // 対象FW+LBオプションのglサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象FW+LBオプションを含むgpサービスコード
	Name          string `gpapi:"Name,required"`          // 削除するグローバルアドレスの名称
}

type DeleteVirtualIpAddressResponse struct {
	CommonResponse
	GlServiceCode string // 対象FW+LBオプションのglサービスコード
	IPv4          struct {
		TrafficAddress string // 削除されたIPv4アドレス
		TrafficIpName  string // 削除されたIPv4アドレスの名称
	}
	IPv6 struct {
		TrafficAddress string // 削除されたIPv6アドレス
		TrafficIpName  string // 削除されたIPv6アドレスの名称
	}
}
