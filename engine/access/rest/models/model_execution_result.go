/*
 * Access API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

type ExecutionResult struct {
	Id               string  `json:"id"`
	BlockId          string  `json:"block_id"`
	Events           []Event `json:"events"`
	Chunks           []Chunk `json:"chunks,omitempty"`
	PreviousResultId string  `json:"previous_result_id"`
	Links            *Links  `json:"_links,omitempty"`
}
