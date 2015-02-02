//  http://manual.iij.jp/gp/gpapi/11454971.html

package protocol

type GetLbHostListArg struct {
	GlServiceCode string `gpapi:"GlServiceCode,required"` // 対象FW+LBオプションのglサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象FW+LBオプションを含むgpサービスコード
}

type GetLbHostListResponse struct {
	HostListPrivateIpAddress struct {
		IPv4Address string // プライベートIPv4アドレス
	}
	GlServiceCode string // 対象FW+LBオプションのglサービスコード
	Redundant     string // 対象FW+LBオプションが冗長構成かどうか  値: Yes, No
	Label         string // 対象FW+LBオプションのラベル
	HostList      []struct {
		SoftwareVersion string // ロードバランサのソフトウェアバージョン  参考:C-4.3 LBのバージョンアップ
		GlobalIpAddress struct {
			IPv4Address string // グローバルIPv4アドレス
			IPv6Address string // グローバルIPv6アドレス
		}
		Location string // ホストのロケーション  値: L, R  参考:B-1.2 基本サーバ構成（全プラン共通）
	}
	CommonResponse
	Type string // 対象FW+LBオプションの品目 (ロードバランサの性能表記)
}
