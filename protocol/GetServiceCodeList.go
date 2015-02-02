//  http://manual.iij.jp/gp/gpapi/11455006.html

package protocol

type GetServiceCodeListArg struct {
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象GP契約のgpサービスコード
}

type GetServiceCodeListResponse struct {
	GnbServiceCodeList []string // GP契約に含まれるNAS/Bオプションのサービスコードのリスト
	GxServiceCodeList  []string // GP契約に含まれるVLANオプションサービスコードのリスト
	GvsServiceCodeList []string // GP契約に含まれるVPN Type-Sオプションサービスコードのリスト
	CommonResponse
	GlServiceCodeList  []string // GP契約に含まれるFW+LBオプションサービスコードのリスト
	GcServiceCodeList  []string // GP契約に含まれる仮想サーバのサービスコードのリスト
	GvmServiceCodeList []string // GP契約に含まれるVPN Type-Mオプションサービスコードのリスト
	GomServiceCodeList []string // GP契約に含まれるシステム運用管理オプションサービスコードのリスト
	CustomerCode       string   // undocumented
	GpServiceCode      string   // 対象GP契約のgpサービスコード
}
