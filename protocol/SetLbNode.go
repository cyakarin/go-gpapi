//  http://manual.iij.jp/gp/gpapi/12442852.html

package protocol

type SetLbNodeArg struct {
	CurrentNodeIpAddress string `gpapi:"CurrentNodeIpAddress,required"` // 変更するノードの現在のIPアドレス
	CurrentNodePort      string `gpapi:"CurrentNodePort,required"`      // 変更するノードの現在のポート
	GlServiceCode        string `gpapi:"GlServiceCode,required"`        // 対象FW+LBオプションのglサービスコード
	GpServiceCode        string `gpapi:"GpServiceCode,required"`        // 対象FW+LBオプションを含むgpサービスコード
	NewNodeIpAddress     string `gpapi:"NewNodeIpAddress,required"`     // 変更するノードの新しいIPアドレス
	NewNodePort          string `gpapi:"NewNodePort,required"`          // 変更するノードの新しいポート
	Pool                 string `gpapi:"Pool,required"`                 // 変更するノードを含むプールの名称
}

type SetLbNodeResponse struct {
	CommonResponse
	NodeList []struct {
		Status    string // ノードの状態 (有効・ドレイン・無効)  値: Active, Draining, Disable
		IpAddress string // ノードのIPアドレス
		Port      string // ノードのポート
	}
	GlServiceCode string // 対象FW+LBオプションのglサービスコード
	Pool          string // 追加対象のプール名称
}
