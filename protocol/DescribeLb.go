//  http://manual.iij.jp/gp/gpapi/11454965.html

package protocol

type DescribeLbArg struct {
	GlServiceCode string `gpapi:"GlServiceCode,required"` // 対象FW+LBオプションのglサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象FW+LBオプションを含むgpサービスコード
}

type DescribeLbResponse struct {
	TrafficIpList []struct {
		Primary string // FW+LBオプションに初期状態で割立てられたLBグローバルアドレスかどうか  値: Yes, No
		IPv4    struct {
			TrafficIpAddress string // LBグローバルアドレスのIPv4アドレス
			TrafficIpName    string // LBグローバルアドレスの名称
			DomainName       string // LBグローバルアドレスのDNS逆引き名
		}
		IPv6 struct {
			TrafficIpAddress string // LBグローバルアドレスのIPv6アドレス
			TrafficIpName    string // LBグローバルアドレスの名称
			DomainName       string // LBグローバルアドレスのDNS逆引き名
		}
	}
	PoolList []struct {
		NodeList []struct {
			Status    string // ノードの状態 (有効・ドレイン・無効)  値: Active, Draining, Disable
			IpAddress string // ノードのIPアドレス
			Port      string // ノードのポート
		}
		Name string // プール名称
	}
	GlServiceCode     string // 対象FW+LBオプションのglサービスコード
	Redundant         string // 対象FW+LBオプションが冗長構成かどうか  値: Yes, No
	Label             string // 対象FW+LBオプションのラベル
	VirtualServerList []struct {
		Protocol          string   // LB仮想サービスのプロトコル  (ロードバランサで使用するプロトコルの表記)
		Name              string   // LB仮想サービスの名称
		Enabled           string   // LB仮想サービスの状態 (有効, 無効)  値: Yes, No
		TrafficIpNameList []string // LB仮想サービスが利用するLBグローバルアドレスのリスト
		Port              string   // LB仮想サービスのポート
		Pool              string   // LB仮想サービスが利用するプール
	}
	CommonResponse
	Type string // 対象FW+LBオプションの品目 (ロードバランサの性能表記)
}
