//  http://manual.iij.jp/gp/gpapi/11454963.html

package protocol

type DescribeFwArg struct {
	GlServiceCode string `gpapi:"GlServiceCode,required"` // 対象FW+LBオプションのglサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象FW+LBオプションを含むgpサービスコード
}

type DescribeFwResponse struct {
	GlServiceCode     string // 対象FW+LBオプションのglサービスコード
	Redundant         string // 対象FW+LBオプションが冗長構成かどうか  値: Yes, No
	Label             string // 対象FW+LBオプションのラベル
	FilteringRuleList []struct {
		Out struct {
			IPv4List []struct {
				Source struct {
					IpAddress string // 送信元IPv4アドレス/ネットマスク
				}
				Destination struct {
					Protocol  string // 送信先プロトコル  値: UDP, TCP
					IpAddress string // 送信先IPv4アドレス/ネットマスク
					Port      string // 送信先ポート番号
				}
				Target string // 動作  値: ACCEPT, DROP, REJECT
			}
		}
		In struct {
			IPv4List []struct {
				Source struct {
					IpAddress string // 送信元IPv4アドレス/ネットマスク
				}
				Destination struct {
					Protocol  string // 送信先プロトコル  値: UDP, TCP
					IpAddress string // 送信先IPv4アドレス/ネットマスク
					Port      string // 送信先ポート番号
				}
				Target string // 動作  値: ACCEPT, DROP, REJECT
			}
			IPv6List []struct {
				Source struct {
					IpAddress string // 送信元IPv6アドレス/ネットマスク
				}
				Destination struct {
					Protocol  string // 送信先プロトコル  値: UDP, TCP
					IpAddress string // 送信先IPv6アドレス/ネットマスク
					Port      string // 送信先ポート番号
				}
				Target string // 動作  値: ACCEPT, DROP, REJECT
			}
		}
	}
	CommonResponse
	Type string // 対象FW+LBオプションの品目 (ロードバランサの性能表記)
}
