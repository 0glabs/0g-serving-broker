// Code generated by gen; DO NOT EDIT.

package model

import (
	"fmt"
	"strings"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
)

func ValidateUpdateAccount(oldVal, newVal Account) error {
	fields := []string{}
	if newVal.ID != nil && !apiequality.Semantic.DeepEqual(newVal.ID, oldVal.ID){
		fields = append(fields, "id")
	}
	
	if !apiequality.Semantic.DeepEqual(newVal.User, oldVal.User){
		fields = append(fields, "user")
	}
	if !apiequality.Semantic.DeepEqual(newVal.Provider, oldVal.Provider){
		fields = append(fields, "provider")
	}

	if len(fields) > 0 {
		return fmt.Errorf("update field: [%s] not allowed", strings.Join(fields, ","))
	}
	return nil
}

func ValidateUpdateRequest(oldVal, newVal Request) error {
	fields := []string{}
	if newVal.ID != nil && !apiequality.Semantic.DeepEqual(newVal.ID, oldVal.ID){
		fields = append(fields, "id")
	}
	
	if !apiequality.Semantic.DeepEqual(newVal.CreatedAt, oldVal.CreatedAt){
		fields = append(fields, "createdAt")
	}
	if !apiequality.Semantic.DeepEqual(newVal.UserAddress, oldVal.UserAddress){
		fields = append(fields, "userAddress")
	}
	if !apiequality.Semantic.DeepEqual(newVal.Nonce, oldVal.Nonce){
		fields = append(fields, "nonce")
	}
	if !apiequality.Semantic.DeepEqual(newVal.ServiceName, oldVal.ServiceName){
		fields = append(fields, "serviceName")
	}
	if !apiequality.Semantic.DeepEqual(newVal.InputCount, oldVal.InputCount){
		fields = append(fields, "inputCount")
	}
	if !apiequality.Semantic.DeepEqual(newVal.PreviousOutputCount, oldVal.PreviousOutputCount){
		fields = append(fields, "previousOutputCount")
	}
	if !apiequality.Semantic.DeepEqual(newVal.Signature, oldVal.Signature){
		fields = append(fields, "signature")
	}

	if len(fields) > 0 {
		return fmt.Errorf("update field: [%s] not allowed", strings.Join(fields, ","))
	}
	return nil
}

func ValidateUpdateService(oldVal, newVal Service) error {
	fields := []string{}
	if newVal.ID != nil && !apiequality.Semantic.DeepEqual(newVal.ID, oldVal.ID){
		fields = append(fields, "id")
	}
	
	if !apiequality.Semantic.DeepEqual(newVal.Name, oldVal.Name){
		fields = append(fields, "name")
	}

	if len(fields) > 0 {
		return fmt.Errorf("update field: [%s] not allowed", strings.Join(fields, ","))
	}
	return nil
}
