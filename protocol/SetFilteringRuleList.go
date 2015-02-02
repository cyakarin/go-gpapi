//  http://manual.iij.jp/gp/gpapi/11454979.html

package protocol

type SetFilteringRuleListArg struct {
	DestinationIpAddress []string `gpapi:"DestinationIpAddress,optional"` // 送信先IPアドレス/ネットマスク
	DestinationPort      []string `gpapi:"DestinationPort,optional"`      // 送信先ポート番号
	DestinationProtocol  []string `gpapi:"DestinationProtocol,optional"`  // 送信先プロトコル  値: UDP, TCP
	Direction            string   `gpapi:"Direction,required"`            // 設定するフィルタリングルールの通信方向  値: IN, OUT
	GlServiceCode        string   `gpapi:"GlServiceCode,required"`        // 対象FW+LBオプションのglサービスコード
	GpServiceCode        string   `gpapi:"GpServiceCode,required"`        // 対象FW+LBオプションを含むgpサービスコード
	IpVersion            string   `gpapi:"IpVersion,required"`            // 設定するフィルタリングルールのIPバージョン  値: 4, 6
	SourceIpAddress      []string `gpapi:"SourceIpAddress,optional"`      // 送信元IPアドレス/ネットマスク
	Target               []string `gpapi:"Target,optional"`               // 動作  値: ACCEPT, DROP, REJECT
}

type SetFilteringRuleListResponse struct {
	IpVersion         string // 設定したフィルタリングルールのIPバージョン  値: 4, 6
	FilteringRuleList []struct {
		Source struct {
			IpAddress string // 送信元IPアドレス/ネットマスク
		}
		Destination struct {
			Protocol  string // 送信先プロトコル  値: UDP, TCP
			IpAddress string // 送信先IPアドレス/ネットマスク
			Port      string // 送信先ポート番号
		}
		Target string // 動作  値: ACCEPT, DROP, REJECT
	}
	Direction     string // 設定したフィルタリングルールの通信方向  値: IN, OUT
	GlServiceCode string // 対象FW+LBオプションのglサービスコード
	CommonResponse
}
