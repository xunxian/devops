> 
>
> 学无止境，苦海无涯
# 关闭SELinux
```shell
临时关闭：setenforce 0
临时开启：setenforce 1
永久关闭：vi /etc/selinux/config
修改SELINUX=disabled，保存退出。
查看状态：getenforce
```
# 修改时区
```shell
ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
```
# 关闭firewalld
```shell
systemctl stop firewalld && systemctl disable firewalld
```
# 修改主机名
> https://blog.csdn.net/xuheng8600/article/details/79983927
```shell
hostnamectl set-hostname cn-linux-jenkins
vi /etc/sysconfig/network --> xxx
vi /etc/hostname  --> xxx
vi /etc/hosts -->xxx
```
# 更换YUM源
```shell
yum install wget -y
mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup
wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.163.com/.help/CentOS7-Base-163.repo
yum clean all
yum makecache
```
# 安装基本组件
```shell
yum install -y net-tools wget lsof lrzsz zip gzip unzip telnet git ntpdate
```
# 创建个人管理员用户
```shell
// 添加用户，给予wheel组成员管理员权限
useradd -m -c “My Admin user” myname
passwd myname
usermod -aG wheel myname
// 彻底删除用户
userdel -rf fangs
// 将用户从组里移除
id fangsh
gpasswd -d fangsh wheel
```
# 禁止root远程登录
```shell
sed -i 's/#PermitRootLogin yes/PermitRootLogin no/g' /etc/ssh/sshd_config
systemctl restart sshd
备注：可以使用sudo -i的方式登录，也可以使用su - root登录
vi /etc/passwd，将root账号的bash修改成nologin之后，sudo -i和su - root都无法登录
```
# 安装Nginx
```shell
sudo yum install gcc-c++
sudo yum install pcre pcre-devel
sudo yum install zlib zlib-devel
sudo yum install openssl openssl-devel
wget下载nginx源码包后解压编译
./configure --prefix=/opt/nginx
make && make install
# nginx.conf user使用fangsh
配置开机自启脚本
vi nginx.service
[Unit]
User=fangsh
Description=nginx
After=network.target
[Service]
Type=forking
PIDFile=/opt/nginx/logs/nginx.pid
ExecStart=/opt/nginx/sbin/nginx
ExecReload=/opt/nginx/sbin/nginx -s reload
ExecStop=/opt/nginx/sbin/nginx -s stop
PrivateTmp=true
[Install]
WantedBy=multi-user.target
保存路径：/usr/lib/systemd/system
```
# 使用钉钉内网穿透
```
git clone https://github.com/open-dingtalk/pierced.git
cd pierced/linux
./ding -config=./ding.cfg -subdomain=maxfang 8000
http://maxfang.vaiwan.com
```

# CentOS7安装EPEL(Extra Packages for Enterprise Linux)
```
yum install epel-release
yum install -y atop
```