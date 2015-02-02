//  http://manual.iij.jp/gp/gpapi/11454973.html

package protocol

type GetSnatRuleListArg struct {
	GlServiceCode string `gpapi:"GlServiceCode,required"` // 対象FW+LBオプションのglサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象FW+LBオプションを含むgpサービスコード
}

type GetSnatRuleListResponse struct {
	CommonResponse
	SnatRuleList []struct {
		ToSourceIpAddress string // 変換先IPアドレス
		SourceIpAddress   string // 変換元IPアドレス
	}
	GlServiceCode string // 対象FW+LBオプションのglサービスコード
}
