client_2_nginx_head='ss -n state ESTABLISHED dport = :9100|head'
client_2_nginx_num=`ss -n state ESTABLISHED dport = :2304|grep -v 'Port'|wc -l`

nginx_2_client_head='ss -n state ESTABLISHED sport = :9100|head'
nginx_2_client_num=`ss -n state ESTABLISHED sport = :2304|grep -v 'Port'|wc -l`

nginx_2_upstream_head='ss -n state ESTABLISHED dport = :9001|head'
nginx_2_upstream_num=`ss -n state ESTABLISHED dport = :2303|grep -v 'Port'|wc -l`

upstream_2_nginx_head='ss -n state ESTABLISHED sport = :9001|head'
upstream_2_nginx_num=`ss -n state ESTABLISHED sport = :2303|grep -v 'Port'|wc -l`

echo "client_2_nginx_num:"$client_2_nginx_num
#echo "client_2_nginx_head:"
#eval $client_2_nginx_head

echo "nginx_2_clinet_num:"$nginx_2_client_num
#echo "nginx_2_clinet_head:"
#eval $nginx_2_client_head

echo "nginx_2_upstream_num:"$nginx_2_upstream_num
#echo "nginx_2_upstream_head:"
#eval $nginx_2_upstream_head

echo "upstream_2_nginx_num:"$upstream_2_nginx_num
#echo "upstream_2_nginx_head:"
#eval $upstream_2_nginx_head
