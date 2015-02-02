//  http://manual.iij.jp/gp/gpapi/11454897.html

package protocol

type SetLabelArg struct {
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象を含むgpサービスコード
	Label         string `gpapi:"Label,required"`         // 設定するラベル
	ServiceCode   string `gpapi:"ServiceCode,required"`   // ラベルを設定するサービスコード  値: gp########, gc########, gl########のいずれか  参考: サービスコード
}

type SetLabelResponse struct {
	GpServiceCode string // 対象のgpサービスコード ※1
	CommonResponse
	GlServiceCode string // 対象のglサービスコード ※1
	GcServiceCode string // 対象のgcサービスコード ※1
	Label         string // 設定されたラベル
}
