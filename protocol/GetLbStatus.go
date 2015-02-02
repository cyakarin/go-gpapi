//  http://manual.iij.jp/gp/gpapi/12442805.html

package protocol

type GetLbStatusArg struct {
	GlServiceCode string `gpapi:"GlServiceCode,required"` // 対象FW+LBオプションのglサービスコード
	GpServiceCode string `gpapi:"GpServiceCode,required"` // 対象FW+LBオプションを含むgpサービスコード
}

type GetLbStatusResponse struct {
	CommonResponse
	HostList []struct {
		Status   string // ホストの状態
		Location string // ホストの存在するロケーション  値: L, R
	}
	GlServiceCode string // 対象FW+LBオプションのglサービスコード
}
