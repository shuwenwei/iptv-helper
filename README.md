## iptv-helper
### Quickstart
修改iptv.toml
```toml
[user]
#用户名
Username = "username" 
#密码
Password = "password"
[app]
#每个goroutine观看的时间(分钟)
Tasktime = 20
#同时观看数量(goroutine数), 总播放时间 = Tasknum * Tasktime
Tasknum = 10
```