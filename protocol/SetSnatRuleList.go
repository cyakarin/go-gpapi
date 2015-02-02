//  http://manual.iij.jp/gp/gpapi/11454981.html

package protocol

type SetSnatRuleListArg struct {
	GlServiceCode     string   `gpapi:"GlServiceCode,required"`     // 対象FW+LBオプションのglサービスコード
	GpServiceCode     string   `gpapi:"GpServiceCode,required"`     // 対象FW+LBオプションを含むgpサービスコード
	SourceIpAddress   []string `gpapi:"SourceIpAddress,optional"`   // 変換元IPアドレス
	ToSourceIpAddress []string `gpapi:"ToSourceIpAddress,optional"` // 変換先IPアドレス
}

type SetSnatRuleListResponse struct {
	CommonResponse
	SnatRuleList []struct {
		ToSourceIpAddress string // 変換先IPアドレス
		SourceIpAddress   string // 変換元IPアドレス
	}
	GlServiceCode string // 対象FW+LBオプションのglサービスコード
}
