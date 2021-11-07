### CentOS7 YUM方式安装Mysql

> Mysql官网：https://www.mysql.com/downloads/

```
1、【DOWNLOADS】-【MySQL Community (GPL) Downloads »】-【MySQL Yum Repository】-【Red Hat Enterprise Linux 7 / Oracle Linux 7 (Architecture Independent), RPM Package】-【Donwoad】
2、上传mysql80-community-release-el7-4.noarch.rpm到Linux服务器
3、rpm -ivh mysql80-community-release-el7-4.noarch.rpm
4、vi /etc/yum.repos.d/mysql-community.repo
5、[mysql57-community]
   ……
   enabled=1
   ……
   
   [mysql80-community]
   ……
   enabled=0
   ……
6、yum search mysql | grep mysql-community-server
7、yum info mysql-community-server.x86_64
8、yum install mysql-community-server.x86_64
  
```

