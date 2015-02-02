//  http://manual.iij.jp/gp/gpapi/11454969.html

package protocol

type GetFilteringRuleListArg struct {
	Direction     string `gpapi:"Direction,required"`     // 取得するフィルタリングルールの通信方向  値: IN, OUT
	GlServiceCode string `gpapi:"GlServiceCode,required"` // 対象FW+LBオプションのglサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象FW+LBオプションを含むgpサービスコード
	IpVersion     string `gpapi:"IpVersion,required"`     // 取得するフィルタリングルールのIPバージョン  値: 4, 6
}

type GetFilteringRuleListResponse struct {
	IpVersion         string // 取得したフィルタリングルールのIPバージョン  値: 4, 6
	FilteringRuleList []struct {
		Source struct {
			IpAddress string // 送信元IPアドレス/ネットマスク
		}
		Destination struct {
			Protocol  string // 送信先プロトコル  値: UDP, TCP
			IpAddress struct {
				string // 送信先IPアドレス/ネットマスク
			}
			Port string // 送信先ポート番号
		}
		Target string // 動作  値: ACCEPT, DROP, REJECT
	}
	Direction     string // 取得したフィルタリングルールの通信方向  値: IN, OUT
	GlServiceCode string // 対象FW+LBオプションのglサービスコード
	CommonResponse
}
