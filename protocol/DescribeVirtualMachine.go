//  http://manual.iij.jp/gp/gpapi/11445331.html

package protocol

type DescribeVirtualMachineArg struct {
	GcServiceCode string `gpapi:"GcServiceCode,required"` // 対象仮想サーバのgcサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象仮想サーバを含むgpサービスコード
}

type DescribeVirtualMachineResponse struct {
	Status         string // 仮想サーバの状態 (仮想サーバの状態遷移)
	PrivateAddress struct {
		IPv4Address      string // IPv4アドレス
		NetworkInterface string // プライベートアドレスが割り当てられているインターフェース
	}
	VirtualMachineType string // 仮想サーバの品目 (仮想サーバの性能表記)
	MobileWebOption    string // モバイルWebオプション有無  値: Yes, No
	GlobalAddress      struct {
		IPv4Address      string // IPv4アドレス
		IPv4DomainName   string // IPv4アドレス DNS逆引き名
		IPv6Enabled      string // IPv6有効・無効  値: Yes, No
		IPv6Address      string // IPv6アドレス
		NetworkInterface string // グローバルアドレスが割り当てられているインターフェース
	}
	Vlan             string // VLANオプション (仮想サーバの性能表記)
	Label            string // 仮想サーバのラベル
	SecureMailOption string // セキュアメールオプション有無  値: Yes, No
	HostType         string // 仮想サーバのタイプ (仮想サーバの性能表記)
	GcServiceCode    string // 対象仮想サーバのgcサービスコード
	DiskOptionList   []struct {
		DeviceName string // 追加ディスクが接続されているデバイス名
		Size       string // ディスク容量 (仮想サーバの性能表記)
	}
	GlobalAddressOptionList []struct {
		IPv4 struct {
			NetworkInterface string // 追加グローバルアドレスが割り当てられているインターフェース
			IpAddress        string // IPv4アドレス
			DomainName       string // IPv4アドレス DNS逆引き名
		}
		IPv6 struct {
			IpAddress        string // IPv6アドレス
			Enabled          string // IPv6有効・無効  値: Yes, No
			NetworkInterface string // 追加グローバルアドレスが割り当てられているインターフェース
		}
	}
	Location string // 仮想サーバのロケーション  値: L, R  参考: B-1.2 基本サーバ構成（全プラン共通）ロケーション
	Memory   string // 仮想サーバのメモリ容量 (仮想サーバの性能表記)
	CommonResponse
	Disk string // 仮想サーバのディスクの容量 (仮想サーバの性能表記)
	OS   string // 仮想サーバのOS (仮想サーバの性能表記)
	CPU  string // 仮想サーバのCPU (仮想サーバの性能表記)
}
