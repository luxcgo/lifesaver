# 本项目已被迁移，现在由 Olive 组织维护

# 详情请查看 https://github.com/go-olive

# lifesaver

[![GoDoc](https://img.shields.io/badge/GoDoc-Reference-blue?style=for-the-badge&logo=go)](https://pkg.go.dev/github.com/luxcgo/lifesaver?tab=doc)
[![GoReport](https://goreportcard.com/badge/github.com/luxcgo/lifesaver?style=for-the-badge)](https://goreportcard.com/report/github.com/luxcgo/lifesaver)
[![Sourcegraph](https://img.shields.io/badge/view%20on-Sourcegraph-brightgreen.svg?style=for-the-badge&logo=sourcegraph)](https://sourcegraph.com/github.com/luxcgo/lifesaver)

## Save Lives!!

Lives are delicate and fleeting creatures, waiting to be captured by us. ❤

> 全自动录播、投稿工具
>
> 支持抖音直播、虎牙直播、B站直播、油管直播
>
> 支持B站投稿

## Usage

1. 安装 **[FFmpeg](https://ffmpeg.org/)**
2. 安装 **[streamlink](https://streamlink.github.io/)**（若不录制 YouTube 直播无需安装）
3. 安装 **[biliup-rs](https://github.com/ForgQi/biliup-rs)**（若不上传至哔哩哔哩无需安装）
4. 安装 [**lifesaver**](https://github.com/biliup/lifesaver)
    * 可直接在 [**releases**](https://github.com/biliup/lifesaver/releases) 中下载相应平台的执行文件
    * 或者本地构建`go install github.com/luxcgo/lifesaver/cmd/lifesaver@latest`
5. 命令行中运行
    * 直接下载可执行文件`/path/to/lifesaver -c /path/to/config.toml`
    * 本地构建`lifesaver -c /path/to/config.toml`

## Config.toml

template file to reference [config.toml](tmpl/config.toml)

```toml
[UploadConfig]
# 是否上传到 bilibili
Enable = false
# biliup-rs 可执行文件的路径
ExecPath = "biliup"
# biliup-rs 配置文件路径，为空的话走默认配置
Filepath = ""

[PlatformConfig]
# 若有录制抖音直播，可在无痕模式非登录状态下找下面的 cookie 填入即可
DouyinCookie = "__ac_nonce=06245c89100e7ab2dd536; __ac_signature=_02B4Z6wo00f01LjBMSAAAIDBwA.aJ.c4z1C44TWAAEx696;"

[[Shows]]
# 平台名，目前支持：
# "bilibili"
# "douyin"
# "huya"
# "youtube"
Platform = "bilibili"
# 房间号，支持字符串类型的房间号
RoomID = "21852"
# 主播名称
StreamerName = "老番茄"
```

## RoadMap

* 支持 go 原生对视频流的抓取
* 支持 go 原生对 bilibili 的投稿
* 增加 docker image
* 增加 mock test
* 增加 YouTube 投稿
* 增加对更多直播平台的支持
* 增加对程序运行状况的监控
* 增加网页端

## Credits

* [bililive-go](https://github.com/hr3lxphr6j/bililive-go)
* [biliup-rs](https://github.com/ForgQi/biliup-rs)
* [ffmpeg](https://ffmpeg.org/)
* [streamlink](https://streamlink.github.io/)
