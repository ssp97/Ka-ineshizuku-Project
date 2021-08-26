<div align="center">
  <h2>Ka Ineshizuku</h2>
  <h2>夏禾雫</h2>
  夏禾雫是基于ZeroBot框架的聊天机器人<br><br>

[![YAYA](https://img.shields.io/badge/OneBot-YaYa-green.svg?style=social&logo=appveyor)](https://github.com/Yiwen-Chan/OneBot-YaYa)
[![GOCQ](https://img.shields.io/badge/OneBot-MiraiGo-green.svg?style=social&logo=appveyor)](https://github.com/Mrs4s/go-cqhttp)
[![OICQ](https://img.shields.io/badge/OneBot-OICQ-green.svg?style=social&logo=appveyor)](https://github.com/takayama-lily/node-onebot)
[![MIRAI](https://img.shields.io/badge/OneBot-Mirai-green.svg?style=social&logo=appveyor)](https://github.com/yyuueexxiinngg/onebot-kotlin)

<!--[![Go Report Card](https://goreportcard.com/badge/github.com/Yiwen-Chan/ZeroBot-App?style=flat-square&logo=go)](https://goreportcard.com/report/github.com/github.com/Yiwen-Chan/ZeroBot-App)-->
[![Badge](https://img.shields.io/badge/onebot-v11-black?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAHAAAABwCAMAAADxPgR5AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAAxQTFRF////29vbr6+vAAAAk1hCcwAAAAR0Uk5T////AEAqqfQAAAKcSURBVHja7NrbctswDATQXfD//zlpO7FlmwAWIOnOtNaTM5JwDMa8E+PNFz7g3waJ24fviyDPgfhz8fHP39cBcBL9KoJbQUxjA2iYqHL3FAnvzhL4GtVNUcoSZe6eSHizBcK5LL7dBr2AUZlev1ARRHCljzRALIEog6H3U6bCIyqIZdAT0eBuJYaGiJaHSjmkYIZd+qSGWAQnIaz2OArVnX6vrItQvbhZJtVGB5qX9wKqCMkb9W7aexfCO/rwQRBzsDIsYx4AOz0nhAtWu7bqkEQBO0Pr+Ftjt5fFCUEbm0Sbgdu8WSgJ5NgH2iu46R/o1UcBXJsFusWF/QUaz3RwJMEgngfaGGdSxJkE/Yg4lOBryBiMwvAhZrVMUUvwqU7F05b5WLaUIN4M4hRocQQRnEedgsn7TZB3UCpRrIJwQfqvGwsg18EnI2uSVNC8t+0QmMXogvbPg/xk+Mnw/6kW/rraUlvqgmFreAA09xW5t0AFlHrQZ3CsgvZm0FbHNKyBmheBKIF2cCA8A600aHPmFtRB1XvMsJAiza7LpPog0UJwccKdzw8rdf8MyN2ePYF896LC5hTzdZqxb6VNXInaupARLDNBWgI8spq4T0Qb5H4vWfPmHo8OyB1ito+AysNNz0oglj1U955sjUN9d41LnrX2D/u7eRwxyOaOpfyevCWbTgDEoilsOnu7zsKhjRCsnD/QzhdkYLBLXjiK4f3UWmcx2M7PO21CKVTH84638NTplt6JIQH0ZwCNuiWAfvuLhdrcOYPVO9eW3A67l7hZtgaY9GZo9AFc6cryjoeFBIWeU+npnk/nLE0OxCHL1eQsc1IciehjpJv5mqCsjeopaH6r15/MrxNnVhu7tmcslay2gO2Z1QfcfX0JMACG41/u0RrI9QAAAABJRU5ErkJggg==)](https://github.com/howmanybots/onebot)
[![Badge](https://img.shields.io/badge/zerobot-v1.2.1-black?style=flat-square&logo=go)](https://github.com/wdvxdr1123/ZeroBot)
[![License](https://img.shields.io/github/license/Yiwen-Chan/OneBot-YaYa.svg?style=flat-square&logo=gnu)](https://raw.githubusercontent.com/FloatTech/ZeroBot-App/master/LICENSE)
<!--[![qq group](https://img.shields.io/badge/group-1048452984-red?style=flat-square&logo=tencent-qq)](https://jq.qq.com/?_wv=1027&k=QMb7x1mM)-->
</div>


### 功能
- 聊天 `app/chat 来自ZeroBot-Plugin`
    - [x] [BOT名字]
    - [x] [戳一戳BOT]
    - [x] 空调开
    - [x] 空调关
    - [x] 群温度
    - [x] 设置温度[正整数]
- ATRI `app/atri 来自ZeroBot-Plugin`
    - [x] 具体指令看代码
    - 注：本插件基于 [ATRI](https://github.com/Kyomotoi/ATRI) ，为 Golang 移植版
- 群管 `app/manager 来自ZeroBot-Plugin`
    - [x] 禁言[@xxx][分钟]
    - [x] 解除禁言[@xxx]
    - [x] 我要自闭 [分钟]
    - [x] 开启全员禁言
    - [x] 解除全员禁言
    - [x] 升为管理[@xxx]
    - [x] 取消管理[@xxx]
    - [x] 修改名片[@xxx][xxx]
    - [x] 修改头衔[@xxx][xxx]
    - [x] 申请头衔[xxx]
    - [x] 踢出群聊[@xxx]
    - [x] 退出群聊[群号]
    - [x] *入群欢迎
    - [x] *退群通知
    - [x] 在[月份]月[日期]日的[小时]点[分钟]分时(用[url])提醒大家[消息]
    - [x] 在[月份]月[每周or周几]的[小时]点[分钟]分时(用[url])提醒大家[消息]
    - [x] 取消在[月份]月[日期]日的[小时]点[分钟]分的提醒
    - [x] 取消在[月份]月[每周or周几]的[小时]点[分钟]分的提醒
    - [x] 翻牌
    - [ ] 同意入群请求
    - [ ] 同意好友请求
    - [ ] 撤回[@xxx] [xxx]
    - [ ] 警告[@xxx]
    - [x] run[xxx]
- 在线代码运行 `app/runcode 来自ZeroBot-Plugin`
    - [x] >runcode help
    - [x] >runcode [on/off]
    - [x] >runcode [language] [code block]
- 涩图 `app/setutime 来自ZeroBot-Plugin，增加启动缓存功能`
    - [x] 来份[涩图/二次元/风景/车万]
    - [x] 添加[涩图/二次元/风景/车万][P站图片ID]
    - [x] 删除[涩图/二次元/风景/车万][P站图片ID]
    - [x] >setu status
- lolicon `app/lolicon 来自ZeroBot-Plugin`
    - [x] 来份萝莉
- 搜图 `app/saucenao 来自ZeroBot-Plugin`
    - [x] 以图搜图|搜索图片|以图识图[图片]
    - [x] 搜图[P站图片ID]
    - [ ] 新增其他图源
- AIfalse `app/ai_false 来自ZeroBot-Plugin`
    - [x] 查询计算机当前活跃度 [身体检查]
    - [ ] 简易语音
    - [ ] 爬图合成 [@xxx]
- EEAsst `app/EEAsst 电子助手`  
    - [x] 查封装尺寸 [尺寸0603]
    - [x] 查电阻大小 [电阻01B]
    - [ ] 丝印库
- Gag `app/gag 对一些行为进行禁言`
    - [x] 想要静静 [我想静静]
- jieba `app/jieba 结巴分词测试`
    - [x] 结巴分词 [jieba分词 balabala]
- snare `app/snare 群黑历史等群友间互相伤害功能`
    - [x] 使用 [随机陷害]
    - [x] 加图 [陷害加图]
    - [x] 删图 [陷害删图]
    - [x] 伪造聊天 [!伪造@xxx balabala]
- study `app/study 对话学习功能`
    - [x] 精准匹配
    - [x] 模糊匹配
    - [x] 分词模糊匹配
    - [x] 敷衍行为
    - 注：目前大部分语料来自[星野夜蝶Offiial](https://github.com/Giftia/ChatDACS)
- thunder `app/thunder 群游戏手捧雷`  
    - [x] 游戏功能 [手捧雷]
    - [x] 加倍功能  
    - [x] 小学数学加法
- haveAFriend `app/haveAFriend 我有个朋友说...图片生成功能`
    - [x] 针对自己 
    - [x] 针对群友 `我有个朋友说今天是个好日子@xxx`
- gifApp `app/gifApp 一些gif图片合成玩法`
    - [x] 摸头 `摸头@xxx`
- TODO...
### 兼容性
- [x] Linux 
- [ ] Windows (有待测试)
```
    因为CPP程序依赖问题，暂时不能使用"go run"进行运行，也不允许在中文路径下运行。
```
### 依赖组件
- oneBot v11 的兼容实现，例如[go-cqhttp](https://github.com/Mrs4s/go-cqhttp)
- 选择PostgreSql作为数据库时，需要安装对应的数据库并进行配置(可能会带来更好的性能)

### 数据库支持
- SQLite
- PostgreSql

### 快速使用
- 组件依赖请参考需要[go-cqhttp](https://github.com/Mrs4s/go-cqhttp)
- 在Action中，选择windows/linux版本，下载压缩包
- 解压，复制config.toml.template为config.toml，并修改主要配置
- 双击exe文件启动(windows)

### 我要修改
- 我想要修改一些配置上没有的怎么办
```
1. 点击右上角fork一份到自己仓库
2. 修改文件(如果是改改触发的表达式的话推荐在线修改)
3. 开启action自动编译
4. 等
5. 编译完成下载运行
```


### 特别感谢
- [ZeroBot](https://github.com/wdvxdr1123/ZeroBot)
- [ATRI](https://github.com/Kyomotoi/ATRI)
- [ZeroBot-Plugin](https://github.com/FloatTech/ZeroBot-Plugin)
- [ChatDACS](https://github.com/Giftia/ChatDACS)
- [小夜语料](https://github.com/Giftia/Project_Xiaoye)