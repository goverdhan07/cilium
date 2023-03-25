package vpc

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateTrafficMirrorFilter invokes the vpc.CreateTrafficMirrorFilter API synchronously
func (client *Client) CreateTrafficMirrorFilter(request *CreateTrafficMirrorFilterRequest) (response *CreateTrafficMirrorFilterResponse, err error) {
	response = CreateCreateTrafficMirrorFilterResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTrafficMirrorFilterWithChan invokes the vpc.CreateTrafficMirrorFilter API asynchronously
func (client *Client) CreateTrafficMirrorFilterWithChan(request *CreateTrafficMirrorFilterRequest) (<-chan *CreateTrafficMirrorFilterResponse, <-chan error) {
	responseChan := make(chan *CreateTrafficMirrorFilterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTrafficMirrorFilter(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateTrafficMirrorFilterWithCallback invokes the vpc.CreateTrafficMirrorFilter API asynchronously
func (client *Client) CreateTrafficMirrorFilterWithCallback(request *CreateTrafficMirrorFilterRequest, callback func(response *CreateTrafficMirrorFilterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTrafficMirrorFilterResponse
		var err error
		defer close(result)
		response, err = client.CreateTrafficMirrorFilter(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateTrafficMirrorFilterRequest is the request struct for api CreateTrafficMirrorFilter
type CreateTrafficMirrorFilterRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                requests.Integer                         `position:"Query" name:"ResourceOwnerId"`
	ClientToken                    string                                   `position:"Query" name:"ClientToken"`
	IngressRules                   *[]CreateTrafficMirrorFilterIngressRules `position:"Query" name:"IngressRules"  type:"Repeated"`
	TrafficMirrorFilterName        string                                   `position:"Query" name:"TrafficMirrorFilterName"`
	ResourceGroupId                string                                   `position:"Query" name:"ResourceGroupId"`
	EgressRules                    *[]CreateTrafficMirrorFilterEgressRules  `position:"Query" name:"EgressRules"  type:"Repeated"`
	DryRun                         requests.Boolean                         `position:"Query" name:"DryRun"`
	ResourceOwnerAccount           string                                   `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                   string                                   `position:"Query" name:"OwnerAccount"`
	OwnerId                        requests.Integer                         `position:"Query" name:"OwnerId"`
	TrafficMirrorFilterDescription string                                   `position:"Query" name:"TrafficMirrorFilterDescription"`
}

// CreateTrafficMirrorFilterIngressRules is a repeated param struct in CreateTrafficMirrorFilterRequest
type CreateTrafficMirrorFilterIngressRules struct {
	Action               string `name:"Action"`
	SourceCidrBlock      string `name:"SourceCidrBlock"`
	Protocol             string `name:"Protocol"`
	DestinationPortRange string `name:"DestinationPortRange"`
	Priority             string `name:"Priority"`
	DestinationCidrBlock string `name:"DestinationCidrBlock"`
	SourcePortRange      string `name:"SourcePortRange"`
}

// CreateTrafficMirrorFilterEgressRules is a repeated param struct in CreateTrafficMirrorFilterRequest
type CreateTrafficMirrorFilterEgressRules struct {
	Action               string `name:"Action"`
	SourceCidrBlock      string `name:"SourceCidrBlock"`
	Protocol             string `name:"Protocol"`
	DestinationPortRange string `name:"DestinationPortRange"`
	Priority             string `name:"Priority"`
	DestinationCidrBlock string `name:"DestinationCidrBlock"`
	SourcePortRange      string `name:"SourcePortRange"`
}

// CreateTrafficMirrorFilterResponse is the response struct for api CreateTrafficMirrorFilter
type CreateTrafficMirrorFilterResponse struct {
	*responses.BaseResponse
	TrafficMirrorFilterId string `json:"TrafficMirrorFilterId" xml:"TrafficMirrorFilterId"`
	RequestId             string `json:"RequestId" xml:"RequestId"`
	ResourceGroupId       string `json:"ResourceGroupId" xml:"ResourceGroupId"`
}

// CreateCreateTrafficMirrorFilterRequest creates a request to invoke CreateTrafficMirrorFilter API
func CreateCreateTrafficMirrorFilterRequest() (request *CreateTrafficMirrorFilterRequest) {
	request = &CreateTrafficMirrorFilterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateTrafficMirrorFilter", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateTrafficMirrorFilterResponse creates a response to parse from CreateTrafficMirrorFilter response
func CreateCreateTrafficMirrorFilterResponse() (response *CreateTrafficMirrorFilterResponse) {
	response = &CreateTrafficMirrorFilterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
