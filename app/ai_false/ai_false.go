/*
暂时只有服务器监控
*/
package ai_false

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/fsUtils"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/zero"
	"math"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() { // 插件主体
	zero.Default().OnFullMatchGroup([]string{"检查身体", "自检", "启动自检", "系统状态"}, ZeroBot.AdminPermission).FirstPriority().SetBlock(true).
		Handle(func(ctx *ZeroBot.Ctx) {
			ctx.SendChain(message.Text(
				"* 系统架构: ",systemPlatform(), "\n",
				"* CPU占用率: ", cpuPercent(), "%\n",
				"* RAM占用率: ", memPercent(), "\n",
				//"* 硬盘活动率: ", diskPercent(), "%\n" ,
				"* 硬盘使用率: ", diskUsage(),"\n",
				"* 系统运行时间：",systemRunningTime(),"\n",
				"* 小雫的各项资源使用：\n",
				selfDump(),
			),
			)
		})
	zero.Default().OnFullMatch("!获取群信息", ZeroBot.OnlyGroup).Handle(func(ctx *ZeroBot.Ctx) {
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

	used := uint64(0)
	all := uint64(0)

	if stat, err := disk.Usage(fsUtils.Getwd()); err == nil {
		used = stat.Used / 1024 / 1024
		all = stat.Total / 1024 / 1024
	}

	usage = fmt.Sprintf("%dM/%dM",used, all)
	return
}

func systemRunningTime()(t string) {
	_time , _ := host.Uptime()
	t = fmt.Sprintf("%vh", _time / 3600)
	return
}

func systemPlatform()(t string){
	info,_ := host.Info()
	t = fmt.Sprintf("%s %s", info.KernelArch, info.OS)
	return
}

func selfDump()(t string)  {
	memStats := runtime.MemStats{}
	runtime.ReadMemStats(&memStats)
	t = fmt.Sprintf("\t * 使用内存：%dKB\n", (memStats.Sys)/1024)
	t += fmt.Sprintf("\t * GC次数：%d\n", memStats.NumGC)
	t += fmt.Sprintf("\t * 累计GC暂停服务时间：%dns\n", memStats.PauseTotalNs)
	return
}
