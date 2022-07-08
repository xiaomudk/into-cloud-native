# 通过golang和C体验系统调用


### 编译

```
go build readfile.go

gcc readfile.c -o readfile_c
```

### 追踪用户态函数调用

**strace参数**
```
-f 跟踪由fork调用所产生的子进程. 
-e trace=set 
只跟踪指定的系统 调用.例如:-e trace=open,close,read,write表示只跟踪这四个系统调用.默认的为set=all. 
-t 在输出中的每一行前加上时间信息. 
-tt 在输出中的每一行前加上时间信息,微秒级. 
-ttt 微秒级输出,以秒了表示时间. 
-T 显示每一调用所耗的时间. 
```

运行golang
```
strace -f -e open,openat,read,write,close ./readfile
```

运行c
```
strace -f -e open,openat,read,write,close ./readfile_c
```

运行cat
```
strace -f -e open,openat,read,write,close cat test.txt
```

三条命令的共同点
```
openat(AT_FDCWD, "test.txt", O_RDONLY)  = 3
read(3, "into the cloud native", 131072) = 21
write(1, "into the cloud native", 21into the cloud native)   = 21
read(3, "", 131072)                     = 0
close(3)                                = 0
close(1)                                = 0
close(2)                                = 0
```

### 内核调用追踪

**trace-cmd**
```
record 记录追踪事件到trace.data
	-p 
		function_graph 跟踪函数调用关系图
		function 追踪函数调用
	-P <pid> 追踪进程
	-g <function> 指定追踪的函数名
	--max-graph-depth <num>  跟踪函数调用关系图深度
	-c -F  同时追踪子进程

report  查看追踪结果
```

**实际操作**

```
trace-cmd record -p function_graph -g '*sys_openat' 

另外一个窗口cat一个文件
trace-cmd report

trace-cmd record -p function_graph -g '*sys_read' 
```