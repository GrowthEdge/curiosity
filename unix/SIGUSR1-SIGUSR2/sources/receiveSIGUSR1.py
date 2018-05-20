import signal
import os
from time import sleep


# 处理信号的函数
def handlerSIGUSR1(signalNumber, stack):
    print("shutup!!!")


# 把信号`USR1` 交给 `handlerSIGUSR1` 来处理
signal.signal(signal.SIGUSR1, handlerSIGUSR1)
print(os.getpid())

# 为了让这个程序一直执行搞一个死循环
while (True):
    print("hi man")
    sleep(1)
