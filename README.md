
# 概

交互式播报天气情况。目前以实现天气信息获取和播报信息合成，还需要解决语音唤醒和语义分析问题。

# 解

* speaker
    实现文字到语音的转换合成
* weather
    获取天气情况数据

# 注
该程序对天气情况的获取和语音合成均使用百度的RESTFul－API，其中设计到的AK和SEC不能公布，有兴趣可以自己申请，然后写到`config.json`中，格式请见`main.go`中的`WFConfig`结构。

# 参
[Baidu语音合成API](http://yuyin.baidu.com/docs/tts/136)
[Baidu Token申请API](http://yuyin.baidu.com/docs/tts/135#获取%20Access%20Token)

