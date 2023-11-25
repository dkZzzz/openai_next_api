# simple_openai_next_api

## 来源
From "github.com/sashabaranov/go-openai"

对这个项目：
1. 进行了简化，只有Chat的发起以及返回
2. 更改了apikey，改为支持next_api


## 使用
```
go install github.com/dkZzzz/openai_next_api
```

## example
```
package main

import next_api "github.com/dkZzzz/openai_next_api"

func main() {
	questions := make([]string, 0)
	next_api_key := ""
	next_api.Chat(next_api_key, questions)
}

```

## 补充
如果想要用回openai的apikey，直接在config.go 里的 openaiAPIURLv1 更改接口url即可，key用回openai的sk-开头的key

