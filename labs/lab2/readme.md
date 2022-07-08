## 使用shell实现一个docker


### 准备工作
```
tar xf nginx_image.tgz           // ubuntu/nginx docker镜像释放的文件
mkdir -p diff workdir nginx      // 创建实验需要的目录
```


### 实现一个docker
```
unshare --uts --mount --pid --fork --net bash     // 进入新的namespace空间
hostname cloudnative 

mount -t overlay overlay -o lowerdir=nginx_image,upperdir=diff,workdir=workdir nginx               // 挂载目录，  后面会讲到细节

mkdir -p nginx/old
cd nginx
pivot_root . ./old          // 把当前目录设置为根目录

mount -t proc proc /proc    // 挂载一个新的proc目录
umount -l /old

exec bash
```

观察hostname、pid、网卡等
