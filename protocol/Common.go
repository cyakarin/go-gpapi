//  http://manual.iij.jp/gp/gpapi/9564008.html

package protocol

type CommonArg struct {
	APIVersion       string `gpapi:"APIVersion,required"`       // APIバージョン  参照: API一覧
	AccessKeyId      string `gpapi:"AccessKeyId,required"`      // 使用するAccessKey  参照: AccessKey
	Action           string `gpapi:"Action,required"`           // 各APIで規定されるAction  参照: API一覧
	Expire           string `gpapi:"Expire,required"`           // Signatureの有効期間APIリクエストの発行から24時間未満の値が指定可能  書式: YYYY-MM-DDThh:mm:ssZ
	Signature        string `gpapi:"Signature,required"`        // リクエストの署名
	SignatureMethod  string `gpapi:"SignatureMethod,required"`  // Singature生成に用いるハッシュアルゴリズム  値: HmacSHA256, HmacSHA1
	SignatureVersion string `gpapi:"SignatureVersion,required"` // Signature生成ロジックのバージョン  値: 2
}

type CommonResponse struct {
	ErrorType    string //
	RequestId    string // APIリクエスト毎に割り当てられるユニークなID
	ErrorMessage string //
}
