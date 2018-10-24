
# https://blog.csdn.net/qq_34021712/article/details/72567786
# yum -y install make gcc gcc-c++ kernel-devel m4 ncurses-devel openssl-devel
# yum -y install ncurses-devel
# wget http://erlang.org/download/otp_src_21.0.tar.gz
# tar xvfz otp_src_21.0.tar.gz
# cd ./otp_src_21.0
# ./configure 
# make install
# ln -s /usr/local/bin/erl /usr/bin/erl

# https://blog.csdn.net/liang_henry/article/details/79584843
sudo vi /etc/yum.repos.d/rabbitmq-erlang.repo
# https://github.com/rabbitmq/erlang-rpm
'''
[rabbitmq-erlang]
name=rabbitmq-erlang
baseurl=https://dl.bintray.com/rabbitmq/rpm/erlang/21/el/7
gpgcheck=1
gpgkey=https://dl.bintray.com/rabbitmq/Keys/rabbitmq-release-signing-key.asc
repo_gpgcheck=0
enabled=1
'''

 yum install -y erlang


yum install -y socat
rpm -Uvh https://github.com/rabbitmq/rabbitmq-server/releases/download/v3.7.7/rabbitmq-server-3.7.7-1.el7.noarch.rpm