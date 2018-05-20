SIGUSR1 SIGUSR2
================

* 问题1：怎么能快速实验下这两个信号？
* 问题2：SIGUSR1、SIGUSR2 有什么惯例的使用方法吗？
* 问题3：做一个信号接收的最小DEMO演示

听说这两个信号是用来给“用户编程用的”，具体要做啥是用户来决定，从这个命名来看为了以后有`SIGUSER3` `SIGUSER4` ... 留了方案，虽然这一点认知没什么卵用。

听说linux有64个信号啊，NND太TMD多了，真的是好尴尬！


### 问题1：怎么能快速实验下这两个信号？



我觉得应该先找个一`python`的案例，直觉上觉得`python`从表面上更容易理解一点：

`receiveSIGUSR1.py`

```python 
import signal
import os
from time import sleep

def handlerSIGUSR1(signalNumber, stack):
    print("shutup!!!")


signal.signal(signal.SIGUSR1, handlerSIGUSR1)
print(os.getpid())

while (True):
    print("hi man")
    sleep(1)



```

代码出来内心还是安稳了很多....

奔跑起来吧 `python3 receiveSIGUSR1.py`

然后新开一个终端输入 `kill -USR1 {上边脚本打印出来的进程ID}`

还是很好玩哈哈。

