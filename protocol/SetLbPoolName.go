//  http://manual.iij.jp/gp/gpapi/12442888.html

package protocol

type SetLbPoolNameArg struct {
	CurrentPoolName string `gpapi:"CurrentPoolName,required"` // 設定を変更するプールの現在の名称
	GlServiceCode   string `gpapi:"GlServiceCode,required"`   // 対象FW+LBオプションのglサービスコード
	GpServiceCode   string `gpapi:"GpServiceCode,required"`   // 対象FW+LBオプションを含むgpサービスコード
	NewPoolName     string `gpapi:"NewPoolName,required"`     // 設定を変更するプールの新しい名称
}

type SetLbPoolNameResponse struct {
	CommonResponse
	NodeList []struct {
		Status    string // ノードの状態 (有効・ドレイン・無効)  値: Active, Draining, Disable
		IpAddress string // ノードのIPアドレス
		Port      string // ノードのポート
	}
	GlServiceCode string // 対象FW+LBオプションのglサービスコード
	Pool          string // 変更したプール名称
}
