# Earthquake wake up

此專案是接收 [地牛 Wake Up!](https://eew.earthquake.tw) 的連動通知

**僅供非營利學術研究使用**。

## Usage

複製 `config.yaml`

```shell
make config
```

此專案可以自行擴充通知的方式

需要在 [configs/config.yaml](configs/config.yaml) 中 新增該通知的參數

並新增 `interface` 的 `notify` 方法，即可使用。

```go
type Notifier interface {
Notify(message string)
}
```

詳細可以直接參考 [LINE 通知](./notify/line/line.go) 的實作

## Reference

- [地牛 Wake Up!](https://eew.earthquake.tw/)