单一全局互斥锁(Sched.Lock) 和集中状态存储的存在，导致所有 Goroutine 相关操作，比如创建、重新调度等，都要上锁；
Goroutine 传递问题：M 经常在 M 之间传递“可运行”的 Goroutine，这导致调度延迟增大，也增加了额外的性能损耗；
每个 M 都做内存缓存，导致内存占用过高，数据局部性较差；
由于系统调用（syscall）而形成的频繁的工作线程阻塞和解除阻塞，导致额外的性能损耗。