// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package controllers

// DataValues is the data values type generated by the AzureFileCSIConfig CR
type DataValues struct {
	AzureFileCSI *DataValuesAzureFileCSI `yaml:"azureFileCSIDriver,omitempty"`
}

// DataValuesAzureFileCSI is the data values section of AzureFileCSI Secret
type DataValuesAzureFileCSI struct {
	Namespace          string `yaml:"namespace"`
	HTTPProxy          string `yaml:"http_proxy"`
	HTTPSProxy         string `yaml:"https_proxy"`
	NoProxy            string `yaml:"no_proxy"`
	DeploymentReplicas int32  `yaml:"deployment_replicas"`
}