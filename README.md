# ban_ip


```
firewall-cmd --permanent --zone=public --new-ipset=blacklist --type=hash:ip
firewall-cmd --permanent --zone=public --add-rich-rule='rule source ipset=blacklist drop'
firewall-cmd --permanent --zone=public --ipset=blacklist --add-entry=222.222.222.222
firewall-cmd --reload
```
## 配置文件说明
```
{
  "nginx_file_path": "/Users/doovemax/TMP/api.log",  # nginx日志文件位置
  "ip_to_black_count":50, # 每分钟单个URI访问的次数
  # 日志文件格式 $http_x_forwarded_for  $status $body_bytes_sent 必须有，时间格式是time_local
  "log_format":"$http_x_forwarded_for $remote_addr - $remote_user [$time_local] \"$request_method $request_uri $server_protocol\" $status $body_bytes_sent \"$http_referer\" \"$http_user_agent\" \"$http_host\" $upstream_addr $upstream_response_time $request_time",
  # IP白名单
  "white_ip_list":[
    "127.0.0.1"
  ]
}
```