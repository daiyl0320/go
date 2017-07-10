package framework

import "fmt"

func Register(obj Runnable) {
	Store.Lock()
	defer Store.Unlock()
	if _, ok := Store.mStore[obj.Id()]; ok {
		fmt.Printf("Object [%v] has already been registered, omit this request, please make sure object id is unique.\n", obj.Id())
		return
	}
	Store.mStore[obj.Id()] = obj
}

func Run() {
	Store.Lock()
	defer Store.Unlock()
	for _, v := range Store.mStore {
		fmt.Printf("Object Id: [%v]\nDetail: [%v]\n", v.Id(), v.Description())
		v.Run()
	}
}
