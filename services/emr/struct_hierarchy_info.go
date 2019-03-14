package emr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// HierarchyInfo is a nested struct in emr response
type HierarchyInfo struct {
	Id                   int                                `json:"Id" xml:"Id"`
	NodeBizType          string                             `json:"NodeBizType" xml:"NodeBizType"`
	NodeType             string                             `json:"NodeType" xml:"NodeType"`
	RelateId             string                             `json:"RelateId" xml:"RelateId"`
	Name                 string                             `json:"Name" xml:"Name"`
	ParentId             string                             `json:"ParentId" xml:"ParentId"`
	ResourceOwnerId      string                             `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	UtcCreateTimestamp   int                                `json:"UtcCreateTimestamp" xml:"UtcCreateTimestamp"`
	UtcModifiedTimestamp int                                `json:"UtcModifiedTimestamp" xml:"UtcModifiedTimestamp"`
	NodeStatus           int                                `json:"NodeStatus" xml:"NodeStatus"`
	ExecutionPlan        ExecutionPlan                      `json:"ExecutionPlan" xml:"ExecutionPlan"`
	Job                  JobInListJobExecutionPlanHierarchy `json:"Job" xml:"Job"`
}