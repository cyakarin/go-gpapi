//  http://manual.iij.jp/gp/gpapi/11468122.html

package protocol

type GetGpServiceCodeListArg struct {
}

type GetGpServiceCodeListResponse struct {
	CommonResponse
	GpServiceCodeList []string // 操作可能なgpサービスコードのリスト
}
