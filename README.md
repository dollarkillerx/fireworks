# fireworks
fireworks container based cicd

### MVP

- Gitlab WebHook
- base
- Backend Web

### conf demo

backend: `configs/config.yaml`

``` 
ServerAddr: http://127.0.0.1:8751
ListenAddr: 0.0.0.0:8751
Debug: false
AgentToken: 5ddbae5a-3d43-4f63-a0b6-c3cdac657379
JWTToken: 2749aa90-bf93-4314-8d6f-ae24f7f7fe49
BasicAdministrator:
  Email: fireworks@fireworks.com
  Name: fireworks
  Password: fireworks

PostgresConfig:
  Host: 127.0.0.1
  Port: 5432
  User: root
  Password: root
  DBName: fireworks
```

agent: `configs/config.yaml`

``` 
AgentName: test_local
AgentIP: 127.0.0.1
Token: 5ddbae5a-3d43-4f63-a0b6-c3cdac657376
Workspace: /home/fireworks/workspace/test
Description: 这是测试服务器
BackendAddr: http://127.0.0.1:8751
```
