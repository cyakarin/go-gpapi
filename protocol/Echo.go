//  http://manual.iij.jp/gp/gpapi/11446078.html

package protocol

type EchoArg struct {
	Param string `gpapi:"Param,required"` // テスト用パラメータ
}

type EchoResponse struct {
	CommonResponse
	Param string // リクエストパラメータParamの内容
}
