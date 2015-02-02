//  http://manual.iij.jp/gp/gpapi/12442716.html

package protocol

type AddLbNodeArg struct {
	GlServiceCode string `gpapi:"GlServiceCode,required"` // 対象FW+LBオプションのglサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象FW+LBオプションを含むgpサービスコード
	NodeIpAddress string `gpapi:"NodeIpAddress,required"` // 追加するノードのIPアドレス
	NodePort      string `gpapi:"NodePort,required"`      // 追加するノードのポート
	Pool          string `gpapi:"Pool,required"`          // 追加対象のプールの名称
}

type AddLbNodeResponse struct {
	CommonResponse
	NodeList []struct {
		Status    string // ノードの状態 (有効・ドレイン・無効)  値: Active, Draining, Disable
		IpAddress string // ノードのIPアドレス
		Port      string // ノードのポート
	}
	GlServiceCode string // 対象FW+LBオプションのglサービスコード
	Pool          string // 追加対象のプール名称
}
