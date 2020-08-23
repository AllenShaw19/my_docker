# my_docker

## 基础知识

### linux proc 文件系统

1. linux 下的/proc 文件系统是由内核提供的，只包含系统运行时的信息（系统内存，mount设备信息，硬件配置等）
2. 只存在于内存中，不占用外存空间
3. 以文件系统的形式，为访问内核数据的操作提供接口

/proc目录下，有很多数字，这些都是为每个进程创建的空间，数字是进程的pid

/proc/N 
/proc/N/cmdline 
/proc/N/cwd/
proc/N/environ/
proc/N/exe/
proc/N/fd/
proc/N/maps/
proc/N/mem/
proc/N/root/
proc/N/stat/
proc/N/statm/
proc/N/status/
proc/self/