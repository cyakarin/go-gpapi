//  http://manual.iij.jp/gp/gpapi/12442776.html

package protocol

type DeleteLbVirtualServerArg struct {
	GlServiceCode     string `gpapi:"GlServiceCode,required"`     // 対象FW+LBオプションのglサービスコード
	GpServiceCode     string `gpapi:"GpServiceCode,required"`     // 対象FW+LBオプションを含むgpサービスコード
	VirtualServerName string `gpapi:"VirtualServerName,required"` // 削除するLB仮想サービスの名称
}

type DeleteLbVirtualServerResponse struct {
	Protocol          string   // LB仮想サービスのプロトコル  (ロードバランサで使用するプロトコルの表記)
	TrafficNameList   []string // LB仮想サービスが使用するLBグローバルアドレスのリスト
	VirtualServerName string   // 削除するLB仮想サービスの名称
	Enabled           string   // LB仮想サービスの状態 (有効, 無効)  値: Yes, No
	CommonResponse
	GlServiceCode string // 対象FW+LBオプションのglサービスコード
	Port          string // LB仮想サービスのポート
	Pool          string // LB仮想サービスが使用するプール
}
