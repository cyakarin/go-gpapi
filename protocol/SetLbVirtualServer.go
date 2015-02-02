//  http://manual.iij.jp/gp/gpapi/12442891.html

package protocol

type SetLbVirtualServerArg struct {
	GlServiceCode     string   `gpapi:"GlServiceCode,required"`     // 対象FW+LBオプションのglサービスコード
	GpServiceCode     string   `gpapi:"GpServiceCode,required"`     // 対象FW+LBオプションを含むgpサービスコード
	Pool              string   `gpapi:"Pool,required"`              // LB仮想サービスが使用するプール
	Port              string   `gpapi:"Port,required"`              // LB仮想サービスが使用するポート
	Protocol          string   `gpapi:"Protocol,required"`          // LB仮想サービスが使用するプロトコル  (ロードバランサで使用するプロトコルの表記)
	TrafficIpName     []string `gpapi:"TrafficIpName,optional"`     // LB仮想サービスが使用するLBグローバルアドレス
	VirtualServerName string   `gpapi:"VirtualServerName,required"` // 変更するLB仮想サービスの名称
}

type SetLbVirtualServerResponse struct {
	Protocol          string   // LB仮想サービスのプロトコル  (ロードバランサで使用するプロトコルの表記)
	TrafficNameList   []string // LB仮想サービスが使用するLBグローバルアドレスのリスト
	VirtualServername string   // LB仮想サービスの名称
	Enabled           string   // LB仮想サービスの状態 (有効, 無効)  値: Yes, No
	CommonResponse
	GlServiceCode string // 対象FW+LBオプションのglサービスコード
	Port          string // LB仮想サービスのポート
	Pool          string // LB仮想サービスが使用するプール
}
