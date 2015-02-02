//  http://manual.iij.jp/gp/gpapi/11454958.html

package protocol

type AddVirtualIpAddressArg struct {
	GlServiceCode string `gpapi:"GlServiceCode,required"` // 対象FW+LBオプションのglサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象FW+LBオプションを含むgpサービスコード
	Name          string `gpapi:"Name,required"`          // 追加するグローバルアドレスにつける名称
}

type AddVirtualIpAddressResponse struct {
	CommonResponse
	GlServiceCode string // 対象FW+LBオプションのglサービスコード
	IPv4          struct {
		TrafficAddress string // 追加されたIPv4アドレス
		TrafficIpName  string // 追加されたIPv4アドレスの名称
	}
	IPv6 struct {
		TrafficAddress string // 追加されたIPv6アドレス
		TrafficIpName  string // 追加されたIPv6アドレスの名称
	}
}
