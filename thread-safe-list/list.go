package list

import (
	"sync"
)

type ThreadSafeList struct {
	List     []interface{}
	Lock     sync.Mutex
}

func (list *ThreadSafeList) Pop() (pop interface{}, get bool) {

	(*list).Lock.Lock()

	if len((*list).List) == 0 {
		(*list).Lock.Unlock()
		return nil, false
	}

	pop = (*list).List[0]
	get = true

	(*list).List = (*list).List[1:]

	(*list).Lock.Unlock()

	return
}

func (list *ThreadSafeList) Push(v interface{}) bool {

	(*list).Lock.Lock()

	(*list).List = append([]interface{}{v}, (*list).List...)

	(*list).Lock.Unlock()

	return true
}

func (list *ThreadSafeList) Size() int {
	return len((*list).List)
}
