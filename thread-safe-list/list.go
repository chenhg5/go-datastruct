package list

import (
	"sync"
)

type ThreadSafeList struct {
	List     []interface{}
	CurIndex uint32
	Len      uint32
	Lock     sync.Mutex
}

func (list *ThreadSafeList) Pop() (pop interface{}, get bool) {

	(*list).Lock.Lock()

	if (*list).CurIndex == (*list).Len {
		(*list).Lock.Unlock()
		return 0, false
	}

	pop = (*list).List[(*list).CurIndex]
	get = true

	(*list).CurIndex++

	(*list).Lock.Unlock()

	return
}

func (list *ThreadSafeList) Push(v interface{}) bool {
	(*list).Lock.Lock()

	if (*list).CurIndex < 0 {
		panic("bug list index error")
	}

	if (*list).CurIndex == 0 {
		(*list).List = append([]interface{}{v}, (*list).List...)
		(*list).Len++

		(*list).Lock.Unlock()
		return true
	}

	(*list).CurIndex--
	(*list).List[(*list).CurIndex] = v

	(*list).Lock.Unlock()
	return true
}

func (list *ThreadSafeList) Size() int {
	return int((*list).Len - (*list).CurIndex)
}
