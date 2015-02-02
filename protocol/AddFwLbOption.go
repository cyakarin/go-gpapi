//  http://manual.iij.jp/gp/gpapi/11454985.html

package protocol

type AddFwLbOptionArg struct {
	GpServiceCode string `gpapi:"GpServiceCode,required"` // FW/LBオプションを追加するgpサービスコード
	Redundant     string `gpapi:"Redundant,required"`     // 対象FW/LBオプションが冗長構成かどうか  値: Yes, No
	Type          string `gpapi:"Type,required"`          // 対象FW/LBオプションの品目 (ロードバランサの性能表記)
}

type AddFwLbOptionResponse struct {
	GpServiceCode string // FW/LBオプションを追加するgpサービスコード
	CommonResponse
	GlServiceCode string // 追加されたFW/LBオプションのglサービスコード
}
