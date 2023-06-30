package apistruct

import sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"

type TencentCloudStorageCredentialReq struct {
	OperationID string `json:"operationID"`
}

type TencentCloudStorageCredentialRespData struct {
	*sts.CredentialResult
	Region string `json:"region"`
	Bucket string `json:"bucket"`
}

type TencentCloudStorageCredentialResp struct {
	CosData TencentCloudStorageCredentialRespData `json:"-"`

	Data map[string]interface{} `json:"data"`
}