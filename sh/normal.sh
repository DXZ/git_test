##linux 端口检测
netstat -nat | grep -i "8102" | wc -l  #统计 含字符8102 的网络连接数量
watch -n 1 -d 'ss dst 10.10.7.179:9100' #每秒刷新查看左右连接 目标为10.10.7.179:9100的tcp连接
ss dst 10.10.7.179:9100|wc -l  统计目标为10.10.7.179:9100的tcp连接 总数
## 查看nginx upstream list


