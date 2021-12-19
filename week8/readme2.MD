#### 第二个作业
写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。
10 20 50 100 200 1k 5k 
#### 测试概述
测试10w的key,在不同value下，写入前后内存的大小
#### 测试过程
#### 测试前
redis_version:2.4.5
redis_git_sha1:00000000
redis_git_dirty:0
arch_bits:64
multiplexing_api:winsock2
process_id:13880
uptime_in_seconds:120051
uptime_in_days:1
lru_clock:414896
used_cpu_sys:1871.67
used_cpu_user:816.41
used_cpu_sys_children:0.00
used_cpu_user_children:0.00
connected_clients:1
connected_slaves:0
client_longest_output_list:0
client_biggest_input_buf:0
blocked_clients:0
used_memory:1189920
used_memory_human:1.13M
used_memory_rss:1189920
used_memory_peak:5234630848
used_memory_peak_human:4.88G
mem_fragmentation_ratio:1.00
mem_allocator:libc
loading:0
aof_enabled:0
changes_since_last_save:-2173446
bgsave_in_progress:0
last_save_time:1639926360
bgrewriteaof_in_progress:0
total_connections_received:3548
total_commands_processed:60515000
expired_keys:0
evicted_keys:0
keyspace_hits:95509348
keyspace_misses:5000405
pubsub_channels:0
pubsub_patterns:0
latest_fork_usec:0
vm_enabled:0
role:master
db1:keys=2,expires=2
.\redis-benchmark.exe -c 30 -n 100000 -d 10 -q -r -t set

#### 步骤