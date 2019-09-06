# ChargeReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | 金額 | [optional] 
**Currency** | **string** | 通貨コード (ISO_4217) | [optional] [default to JPY]
**PaymentMethod** | [**PaymentMethodType**](PaymentMethodType.md) |  | [optional] 
**OrderNo** | **string** | お客様側のシステムオーダーNo（例：注文番号、決済IDなど） | [optional] 
**Description** | **string** | 決済に関する説明 | [optional] 
**Extra** | [**map[string]map[string]interface{}**](map[string]interface{}.md) | 決済に関する追加情報がある場合に利用します。 | [optional] 
**Metadata** | [**map[string]map[string]interface{}**](map[string]interface{}.md) | メタデータ | [optional] 
**ClientIp** | **string** | Client IP アドレス | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


