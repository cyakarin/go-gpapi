//  http://manual.iij.jp/gp/gpapi/11455001.html

package protocol

type GetContractInformationArg struct {
	GcServiceCode string `gpapi:"GcServiceCode,optional"` // 対象仮想サーバのgcサービスコード
	GlServiceCode string `gpapi:"GlServiceCode,optional"` // 対象FW+LBオプションのglサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象GP契約のgpサービスコード
	GxServiceCode string `gpapi:"GxServiceCode,optional"` // 対象VLANオプションのgxサービスコード
}

type GetContractInformationResponse struct {
	Status         string // 契約状態 (契約状態)
	StartDate      string // 契約日
	VLANOptionList []struct {
		GxServiceCode string // サービスコード
		StartDate     string // 契約日
		StopDate      string // 解約日
		Status        string // 契約状態 (契約状態)
	}
	NASBOptionList []struct {
		Status         string // 契約状態 (契約状態)
		GnbServiceCode string // サービスコード
		StopDate       string // 解約日
		StartDate      string // 契約日
		Label          string // ラベル
	}
	FWLBOptionList []struct {
		Status        string // 契約状態 (契約状態)
		StartDate     string // 契約日
		StopDate      string // 解約日
		GlServiceCode string // サービスコード
		Label         string // ラベル
	}
	VPNTypeMOptionList []struct {
		Status         string // 契約状態 (契約状態)
		StartDate      string // 契約日
		GvmServiceCode string // サービスコード
		StopDate       string // 解約日
	}
	CustomerCode       string // 対象GP契約のカスタマーコード
	StopDate           string // 解約日
	VPNTypeSOptionList []struct {
		Status         string // 契約状態 (契約状態)
		StartDate      string // 契約日
		GvsServiceCode string // サービスコード
		StopDate       string // 解約日
	}
	GpServiceCodeLabel string // 対象GP契約のラベル
	CommonResponse
	VirtualMachineList []struct {
		Status               string // 契約状態 (契約状態)
		StartDate            string // 契約日
		StopDate             string // 解約日
		Label                string // ラベル
		SecureMailOptionList []struct {
			Status    string // 契約状態 (契約状態)
			StartDate string // 契約日
			StopDate  string // 解約日
		}
		DiskOptionList []struct {
			Status     string // 契約状態 (契約状態)
			DeviceName string // 追加ディスクオプションのデバイス名
			StopDate   string // 解約日
			StartDate  string // 契約日
		}
		GlobalAddressOptionList []struct {
			Ipv4Address string // 追加グローバルオプションのIPv4アドレス
			StartDate   string // 契約日
			StopDate    string // 解約日
			Status      string // 契約状態 (契約状態)
		}
		MobileWebOptionList []struct {
			Status    string // 契約状態 (契約状態)
			StartDate string // 契約日
			StopDate  string // 解約日
		}
		GcServiceCode string // サービスコード
	}
	SMMOptionList []struct {
		Status         string // 契約状態 (契約状態)
		GomServiceCode string // サービスコード
		StopDate       string // 解約日
		StartDate      string // 契約日
	}
	GpServiceCode string // 対象GP契約のgpサービスコード
}
