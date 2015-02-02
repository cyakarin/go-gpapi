//  http://manual.iij.jp/gp/gpapi/11454987.html

package protocol

type ChangeFwLbOptionTypeArg struct {
	GlServiceCode string `gpapi:"GlServiceCode,required"` // 対象のFW+LBオプションのglサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象のFW+LBオプションを含むgpサービスコード
	Type          string `gpapi:"Type,required"`          // 変更するFW+LBオプションの品目 (ロードバランサの性能表記)
}

type ChangeFwLbOptionTypeResponse struct {
	CommonResponse
	Type          string // 変更後のFW+LBオプションの品目 (ロードバランサの性能表記)
	GlServiceCode string // 対象のFW+LBオプションのglサービスコード
}
