## goroutine

#### go func() 流程

1. go func()
2. 入局部队列
3. 局部队列满的情况下入全局队列
4. 线程对应p调度器获取局部队列的协程
5. 若线程对应p调度器获取局部队列为空，就从全局队列中获取协程
6. 若全局队列也为空，就从其他队列中获取
7. 调用一个协程
8. 执行一个func() 函数
9. 当 M 执行某一个 G 时候如果发生了 syscall 或则其余阻塞操作，M 会阻塞，如果当前有一些 G 在执行，runtime 会把这个线程 M 从 P 中摘除 (detach)，然后再创建一个新的操作系统的线程 (如果有空闲的线程可用就复用空闲线程) 来服务于这个 P；