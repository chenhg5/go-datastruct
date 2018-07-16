package list

import (
	"time"
	"fmt"
	"sync"
)

func main() {
	InitRedPacket()

	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {

		go func(user int) {
			wg.Add(1)
			time.Sleep(time.Second)
			money, ok := GrabPacket()
			fmt.Println("用户：", user, "，抢到：",money, "元", "，是否抢成功", ok)
			wg.Done()
		}(i)

		go func(user int) {
			wg.Add(1)
			time.Sleep(time.Second)
			PushPacket()
			fmt.Println("用户：", user, "，塞进去：3元")
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("剩下的列表：", (*RedPacketList).List)
}

var RedPacketList *ThreadSafeList
var GrabLock sync.Mutex

func InitRedPacket() {
	RedPacketList = &ThreadSafeList{
		[]interface{}{4, 5, 2, 4, 6, 2, 0, 3, 5, 7},
		0,
		10,
		GrabLock,
	}
}

func GrabPacket() (interface{}, bool) {
	return RedPacketList.Pop()
}

func PushPacket()  {
	RedPacketList.Push(3)
}