//  http://manual.iij.jp/gp/gpapi/12442885.html

package protocol

type SetLbPoolArg struct {
	GlServiceCode string   `gpapi:"GlServiceCode,required"` // 対象FW+LBオプションのglサービスコード
	GpServiceCode string   `gpapi:"GpServiceCode,required"` // 対象FW+LBオプションを含むgpサービスコード
	NodeIpAddress []string `gpapi:"NodeIpAddress,optional"` // ノードのIPアドレス
	NodePort      []string `gpapi:"NodePort,optional"`      // ノードのポート
	Pool          string   `gpapi:"Pool,required"`          // 設定を変更するプールの名称
}

type SetLbPoolResponse struct {
	CommonResponse
	NodeList []struct {
		Status    string // ノードの状態 (有効・ドレイン・無効)  値: Active, Draining, Disable
		IpAddress string // ノードのIPアドレス
		Port      string // ノードのポート
	}
	GlServiceCode string // 対象FW+LBオプションのglサービスコード
	Pool          string // プール名称
}
