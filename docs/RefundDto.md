# RefundDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Refund ID | [optional] 
**Object** | **string** | 対象種類の表記 | [optional] [default to refund]
**ChargeId** | **string** | Charge ID | [optional] 
**LiveMode** | **bool** | 本番モードかどうか - false テストモード - true 本番モード  | [optional] 
**Amount** | **int32** | 返金金額。全額返金、及び amount を指定することで金額の部分返金を行うことができます。 | [optional] 
**Currency** | **string** | 通貨コード (ISO_4217) | [optional] 
**Reason** | **string** | 返金理由 | [optional] 
**Status** | **string** | 返金状態 | [optional] 
**RefundedTime** | **int64** | 返金を行う時間のUTCタイムスタンプ。 | [optional] 
**CreateTime** | **int64** | 返金新規時間のUTCタイムスタンプ。 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


