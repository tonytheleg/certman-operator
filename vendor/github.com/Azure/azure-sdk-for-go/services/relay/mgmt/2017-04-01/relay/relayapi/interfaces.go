package relayapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/relay/mgmt/2017-04-01/relay"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result relay.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result relay.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*relay.OperationsClient)(nil)

// NamespacesClientAPI contains the set of methods on the NamespacesClient type.
type NamespacesClientAPI interface {
	CheckNameAvailabilityMethod(ctx context.Context, parameters relay.CheckNameAvailability) (result relay.CheckNameAvailabilityResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, parameters relay.Namespace) (result relay.NamespacesCreateOrUpdateFuture, err error)
	CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string, parameters relay.AuthorizationRule) (result relay.AuthorizationRule, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string) (result relay.NamespacesDeleteFuture, err error)
	DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string) (result relay.Namespace, err error)
	GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string) (result relay.AuthorizationRule, err error)
	List(ctx context.Context) (result relay.NamespaceListResultPage, err error)
	ListComplete(ctx context.Context) (result relay.NamespaceListResultIterator, err error)
	ListAuthorizationRules(ctx context.Context, resourceGroupName string, namespaceName string) (result relay.AuthorizationRuleListResultPage, err error)
	ListAuthorizationRulesComplete(ctx context.Context, resourceGroupName string, namespaceName string) (result relay.AuthorizationRuleListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result relay.NamespaceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result relay.NamespaceListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string) (result relay.AccessKeys, err error)
	RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string, parameters relay.RegenerateAccessKeyParameters) (result relay.AccessKeys, err error)
	Update(ctx context.Context, resourceGroupName string, namespaceName string, parameters relay.UpdateParameters) (result relay.Namespace, err error)
}

var _ NamespacesClientAPI = (*relay.NamespacesClient)(nil)

// HybridConnectionsClientAPI contains the set of methods on the HybridConnectionsClient type.
type HybridConnectionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, parameters relay.HybridConnection) (result relay.HybridConnection, err error)
	CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, authorizationRuleName string, parameters relay.AuthorizationRule) (result relay.AuthorizationRule, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string) (result autorest.Response, err error)
	DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, authorizationRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string) (result relay.HybridConnection, err error)
	GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, authorizationRuleName string) (result relay.AuthorizationRule, err error)
	ListAuthorizationRules(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string) (result relay.AuthorizationRuleListResultPage, err error)
	ListAuthorizationRulesComplete(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string) (result relay.AuthorizationRuleListResultIterator, err error)
	ListByNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (result relay.HybridConnectionListResultPage, err error)
	ListByNamespaceComplete(ctx context.Context, resourceGroupName string, namespaceName string) (result relay.HybridConnectionListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, authorizationRuleName string) (result relay.AccessKeys, err error)
	RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, authorizationRuleName string, parameters relay.RegenerateAccessKeyParameters) (result relay.AccessKeys, err error)
}

var _ HybridConnectionsClientAPI = (*relay.HybridConnectionsClient)(nil)

// WCFRelaysClientAPI contains the set of methods on the WCFRelaysClient type.
type WCFRelaysClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, parameters relay.WcfRelay) (result relay.WcfRelay, err error)
	CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, authorizationRuleName string, parameters relay.AuthorizationRule) (result relay.AuthorizationRule, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string, relayName string) (result autorest.Response, err error)
	DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, authorizationRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string, relayName string) (result relay.WcfRelay, err error)
	GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, authorizationRuleName string) (result relay.AuthorizationRule, err error)
	ListAuthorizationRules(ctx context.Context, resourceGroupName string, namespaceName string, relayName string) (result relay.AuthorizationRuleListResultPage, err error)
	ListAuthorizationRulesComplete(ctx context.Context, resourceGroupName string, namespaceName string, relayName string) (result relay.AuthorizationRuleListResultIterator, err error)
	ListByNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (result relay.WcfRelaysListResultPage, err error)
	ListByNamespaceComplete(ctx context.Context, resourceGroupName string, namespaceName string) (result relay.WcfRelaysListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, authorizationRuleName string) (result relay.AccessKeys, err error)
	RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, authorizationRuleName string, parameters relay.RegenerateAccessKeyParameters) (result relay.AccessKeys, err error)
}

var _ WCFRelaysClientAPI = (*relay.WCFRelaysClient)(nil)