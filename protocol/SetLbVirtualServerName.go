//  http://manual.iij.jp/gp/gpapi/12442894.html

package protocol

type SetLbVirtualServerNameArg struct {
	CurrentVirtualServerName string `gpapi:"CurrentVirtualServerName,required"` // 変更するLB仮想サービスの現在の名称
	GlServiceCode            string `gpapi:"GlServiceCode,required"`            // 対象FW+LBオプションのglサービスコード
	GpServiceCode            string `gpapi:"GpServiceCode,required"`            // 対象FW+LBオプションを含むgpサービスコード
	NewVirtualServerName     string `gpapi:"NewVirtualServerName,required"`     // 変更するLB仮想サービスの新しい名称
}

type SetLbVirtualServerNameResponse struct {
	Protocol          string   // LB仮想サービスのプロトコル  (ロードバランサで使用するプロトコルの表記)
	TrafficNameList   []string // LB仮想サービスが使用するLBグローバルアドレスのリスト
	VirtualServerName string   // LB仮想サービスの変更後の名称
	Enabled           string   // LB仮想サービスの状態 (有効, 無効)  値: Yes, No
	CommonResponse
	GlServiceCode string // 対象FW+LBオプションのglサービスコード
	Port          string // LB仮想サービスのポート
	Pool          string // LB仮想サービスが使用するプール
}
