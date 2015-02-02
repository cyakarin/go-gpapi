//  http://manual.iij.jp/gp/gpapi/11454997.html

package protocol

type AddVlanOptionArg struct {
	GpServiceCode string `gpapi:"GpServiceCode,required"` // VLANオプションを申し込むgpサービスコード
}

type AddVlanOptionResponse struct {
	GpServiceCode string // VLANオプションを申し込むgpサービスコード
	GxServiceCode string // 新規に追加されたVLANオプションのgxサービスコード
	CommonResponse
}
