//  http://manual.iij.jp/gp/gpapi/11454983.html

package protocol

type SetVirtualIpAddressNameArg struct {
	CurrentName   string `gpapi:"CurrentName,required"`   // 変更対象のグローバルアドレスの名称
	GlServiceCode string `gpapi:"GlServiceCode,required"` // 対象FW+LBオプションのglサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象FW+LBオプションを含むgpサービスコード
	NewName       string `gpapi:"NewName,required"`       // 新しいグローバルアドレスの名称
}

type SetVirtualIpAddressNameResponse struct {
	Current struct {
		IPv4 struct {
			TrafficIpAddress string // IPv4アドレス
			TrafficIpName    string // IPv4グローバルアドレス名称
		}
		IPv6 struct {
			TrafficIpAddress string // IPv6アドレス
			TrafficIpName    string // IPv6グローバルアドレス名称
		}
	}
	CommonResponse
	GlServiceCode string // 対象FW+LBオプションのglサービスコード
	Previous      struct {
		IPv4 struct {
			TrafficIpAddress string // IPv4アドレス
			TrafficIpName    string // IPv4グローバルアドレス名称
		}
		IPv6 struct {
			TrafficIpAddress string // IPv6アドレス
			TrafficIpName    string // IPv6グローバルアドレス名称
		}
	}
}
