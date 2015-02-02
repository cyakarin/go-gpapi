//  http://manual.iij.jp/gp/gpapi/12442739.html

package protocol

type AddLbPoolArg struct {
	GlServiceCode string   `gpapi:"GlServiceCode,required"` // 対象FW+LBオプションのglサービスコード
	GpServiceCode string   `gpapi:"GpServiceCode,required"` // 対象FW+LBオプションを含むgpサービスコード
	NodeIpAddress []string `gpapi:"NodeIpAddress,optional"` // ノードのIPアドレス
	NodePort      []string `gpapi:"NodePort,optional"`      // ノードのポート
	Pool          string   `gpapi:"Pool,required"`          // 追加するプールの名称
}

type AddLbPoolResponse struct {
	CommonResponse
	NodeList []struct {
		Status    string // ノードの状態 (有効・ドレイン・無効)  値: Active, Draining, Disable
		IpAddress string // ノードのIPアドレス
		Port      string // ノードのポート
	}
	GlServiceCode string // 対象FW+LBオプションのglサービスコード
	Pool          string // プール名称
}
