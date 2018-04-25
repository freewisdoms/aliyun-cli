package green

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

// ImageAsyncScan invokes the green.ImageAsyncScan API synchronously
// api document: https://help.aliyun.com/api/green/imageasyncscan.html
func (client *Client) ImageAsyncScan(request *ImageAsyncScanRequest) (response *ImageAsyncScanResponse, err error) {
	response = CreateImageAsyncScanResponse()
	err = client.DoAction(request, response)
	return
}

// ImageAsyncScanWithChan invokes the green.ImageAsyncScan API asynchronously
// api document: https://help.aliyun.com/api/green/imageasyncscan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ImageAsyncScanWithChan(request *ImageAsyncScanRequest) (<-chan *ImageAsyncScanResponse, <-chan error) {
	responseChan := make(chan *ImageAsyncScanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImageAsyncScan(request)
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

// ImageAsyncScanWithCallback invokes the green.ImageAsyncScan API asynchronously
// api document: https://help.aliyun.com/api/green/imageasyncscan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ImageAsyncScanWithCallback(request *ImageAsyncScanRequest, callback func(response *ImageAsyncScanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImageAsyncScanResponse
		var err error
		defer close(result)
		response, err = client.ImageAsyncScan(request)
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

// ImageAsyncScanRequest is the request struct for api ImageAsyncScan
type ImageAsyncScanRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// ImageAsyncScanResponse is the response struct for api ImageAsyncScan
type ImageAsyncScanResponse struct {
	*responses.BaseResponse
}

// CreateImageAsyncScanRequest creates a request to invoke ImageAsyncScan API
func CreateImageAsyncScanRequest() (request *ImageAsyncScanRequest) {
	request = &ImageAsyncScanRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-25", "ImageAsyncScan", "/green/image/asyncscan", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateImageAsyncScanResponse creates a response to parse from ImageAsyncScan response
func CreateImageAsyncScanResponse() (response *ImageAsyncScanResponse) {
	response = &ImageAsyncScanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}