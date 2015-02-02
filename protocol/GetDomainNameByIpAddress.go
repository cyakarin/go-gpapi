//  http://manual.iij.jp/gp/gpapi/11454967.html

package protocol

type GetDomainNameByIpAddressArg struct {
	GcServiceCode     string `gpapi:"GcServiceCode,optional"`     // 対象仮想サーバのgcサービスコード
	GlServiceCode     string `gpapi:"GlServiceCode,optional"`     // 対象FW+LBオプションのglサービスコード
	GpServiceCode     string `gpapi:"GpServiceCode,required"`     // 対象を含むgpサービスコード
	IPv4GlobalAddress string `gpapi:"IPv4GlobalAddress,required"` // 対象のIPv4グローバルアドレス
}

type GetDomainNameByIpAddressResponse struct {
	CommonResponse
	IPv4GlobalAddress string // 対象のIPv4グローバルアドレス
	IPv4DomainName    string // 対象IPv4グローバルアドレスのドメイン名 (FQDN)
	GlServiceCode     string // 対象仮想サーバのgcサービスコード  ※2
	GcServiceCode     string // 対象仮想サーバのgcサービスコード  ※2
}
