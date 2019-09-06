# ChargeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Charge ID | [optional] 
**Object** | **string** | 対象種類の表記 | [optional] [default to charge]
**LiveMode** | **bool** | 本番モードかどうか - false テストモード - true 本番モード  | [optional] 
**Amount** | **int32** | 支払い金額 | [optional] 
**Currency** | **string** | 通貨コード (ISO_4217) | [optional] [default to JPY]
**PaymentMethod** | [**PaymentMethodType**](PaymentMethodType.md) |  | [optional] 
**OrderNo** | **string** | お客様システム側のオーダーNo、例えば注文番号、決済IDなど | [optional] 
**Description** | **string** | 支払い説明文 | [optional] 
**Extra** | [**map[string]map[string]interface{}**](map[string]interface{}.md) | 支払いエキストラデータ | [optional] 
**Metadata** | [**map[string]map[string]interface{}**](map[string]interface{}.md) | 支払いメタデータ | [optional] 
**ClientIp** | **string** | Client IP アドレス | [optional] 
**Paid** | **bool** | 支払い済みフラグ | [optional] 
**Refunded** | **bool** | 返金済みフラグ | [optional] 
**Refunds** | [**RefundExtDto**](RefundExtDto.md) |  | [optional] 
**Status** | **string** | 支払い状態 | [optional] 
**Credential** | **string** | Client SDK の認証情報 | [optional] 
**PaidTime** | **int64** | 支払い時間のUTCタイムスタンプ | [optional] 
**RefundTime** | **int64** | 返金時間のUTCタイムスタンプ | [optional] 
**ExpiryTime** | **int64** | 支払い請求有効時間のUTCタイムスタンプ | [optional] 
**SettleTime** | **int64** | 支払い締め時間のUTCタイムスタンプ | [optional] 
**CreateTime** | **int64** | 支払い新規時間のUTCタイムスタンプ | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


