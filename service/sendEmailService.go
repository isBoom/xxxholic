package service

import (
	"fmt"
	"os"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dm"
	"xxxholic/util"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)
type SendEmail struct {
	Email string
	Msg string
}
func (s *SendEmail) Send() error{
	client, err := dm.NewClientWithAccessKey("cn-hangzhou", os.Getenv("OSS_Email_AccessKeyId"), os.Getenv("OSS_Email_AccessKeySecret"))
	request := dm.CreateSingleSendMailRequest()
	request.Scheme = "https"

	request.AccountName = "tian@xxxholic.top"
	request.AddressType = requests.NewInteger(1)
	request.ReplyToAddress = requests.NewBoolean(true)
	request.ToAddress = s.Email
	request.Subject = "验证码"
	request.FromAlias="xxxholic"
	request.HtmlBody = s.Msg
	response, err := client.SingleSendMail(request)
	if err != nil {
		return err
	}
	util.Log().Println(fmt.Sprintf("response is %#v\n", response))
	return nil
}
