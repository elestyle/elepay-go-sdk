/*
 * elepay API リファレンス
 *
 * elepay APIはRESTをベースに構成された決済APIです。支払い処理、返金処理など、決済に関わる運用における様々なことができます。 
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package io.elepay.api

// 支払い情報リスト
type ChargesResponse struct {
	// 件数
	Total int32 `json:"total,omitempty"`
	// 支払い詳細内容
	Charges []ChargeDto `json:"charges,omitempty"`
	Error ElepayRestError `json:"error,omitempty"`
}
