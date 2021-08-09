/*
暂时只有服务器监控
*/
package ai_false

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"math"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"

	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() { // 插件主体
	zero.OnFullMatchGroup([]string{"检查身体", "自检", "启动自检", "系统状态"}, zero.AdminPermission).FirstPriority().SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text(
				"* CPU占用率: ", cpuPercent(), "%\n",
				"* RAM占用率: ", memPercent(), "\n",
				//"* 硬盘活动率: ", diskPercent(), "%\n" ,
				"* 硬盘使用率: ", diskUsage(),"\n",
				"* 系统运行时间：",systemRunningTime(),"\n",

			),
			)
		})
	zero.OnFullMatch("!获取群信息", zero.OnlyGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text(fmt.Sprintf("%v", ctx.GetGroupInfo(ctx.Event.GroupID, false))))

	})
}

func cpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return math.Round(percent[0])
}

func memPercent() string {
	memInfo, _ := mem.VirtualMemory()
	return fmt.Sprintf("%vM/%vM %v%%",memInfo.Used / 1024 / 1024,memInfo.Total / 1024 / 1024,math.Round(memInfo.UsedPercent))
}

func diskPercent() float64 {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return math.Round(diskInfo.UsedPercent)
}

func diskUsage() (usage string) {
	//wd, _ := os.Getwd()
	//fs := syscall.Statfs_t{}
	//err := syscall.Statfs(wd, &fs)
	//if err != nil {
	//	return
	//}
	//
	//all := (fs.Blocks * uint64(fs.Bsize) )/ 1024 / 1024
	//free := (fs.Bfree * uint64(fs.Bsize) )/ 1024 / 1024
	//used := all - free

	used := 0
	all := 0

	usage = fmt.Sprintf("%dM/%dM",used, all)
	return
}

func systemRunningTime()(t string) {
	_time , _ := host.Uptime()
	t = fmt.Sprintf("%vh", _time / 3600)
	return
}

