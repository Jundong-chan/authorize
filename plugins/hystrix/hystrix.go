package pkghystrix

import (
	"github.com/afex/hystrix-go/hystrix"
)

var DefaultCommand = "defaultcommand"

// type HystrixServiceimpl struct {
// 	Command    string
// 	Timeout    int //断路器允许远程调用过程的时间，超时直接判定失败,单位是ms
// 	Maxrequest int //规定最大的并行请求数，超过的请求会直接判定为失败
// 	Trigger    int //触发断路器的最小并发请求数
// 	Sleep      int //断路器开启后休眠的时间，过了就转为半开状态
// 	Allowerr   int //允许出现请求失败的个数（包括超时），达到值则开启断路器
//
// }

// func (Hy HystrixServiceimpl) NewCommand() {
// 	hystrix.ConfigureCommand(Hy.Command, hystrix.CommandConfig{
// 		Timeout:                Hy.Timeout,
// 		MaxConcurrentRequests:  Hy.Maxrequest,
// 		RequestVolumeThreshold: Hy.Trigger,
// 		SleepWindow:            Hy.Sleep,
// 		ErrorPercentThreshold:  Hy.Allowerr,
// 	})
// }

func NewDefaultCommand() string {
	hystrix.ConfigureCommand(DefaultCommand, hystrix.CommandConfig{
		Timeout:                500,
		MaxConcurrentRequests:  500,
		RequestVolumeThreshold: 20,
		SleepWindow:            30,
		ErrorPercentThreshold:  5,
	})
	return DefaultCommand
}
