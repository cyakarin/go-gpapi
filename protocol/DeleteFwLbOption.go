//  http://manual.iij.jp/gp/gpapi/11454990.html

package protocol

type DeleteFwLbOptionArg struct {
	GlServiceCode string `gpapi:"GlServiceCode,required"` // 対象のFW+LBオプションのglサービスコード
	StopDate      string `gpapi:"StopDate,required"`      // 解約予定日  値: YYYYMMDD
}

type DeleteFwLbOptionResponse struct {
	CommonResponse
	StopDate      string // 解約予定日
	GlServiceCode string // 対象のFW+LBオプションのglサービスコード
}
