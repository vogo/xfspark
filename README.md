# xfspark 讯飞星火认知大模型 SDK

## Web SDK

安装本地命令行:
```bash
go install github.com/vogo/xfspark/cmd/sparkchat@latest

export APP_ID=xxx
export API_KEY=xxx
export API_SECRET=xxx

sparkchat
# You: AI会替代人们的工作吗？
# AI: 当然会了
```

系统库依赖调用范例:
```bash
s := chat.NewServer(appID, apiKey, apiSecret)
session, sessionErr := s.GetSession("<unique_user_id>")
answer, sendErr := session.Send(question)
fmt.Println("AI: ", answer)
```

