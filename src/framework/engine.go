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

func RunAll() {
	Store.Lock()
	defer Store.Unlock()
	for _, v := range Store.mStore {
		fmt.Println("===================================================")
		fmt.Printf("Object Id: [%v]\nDetail: [%v]\n", v.Id(), v.Description())
		v.Run()
	}
}

func RunSingle(id string) {
	Store.Lock()
	defer Store.Unlock()
	if v, ok := Store.mStore[id]; ok {
		fmt.Println("===================================================")
		fmt.Printf("Object Id: [%v]\nDetail: [%v]\n", v.Id(), v.Description())
		v.Run()
	} else {
		fmt.Printf("Can't find object with Id [%s]\n", id)
	}
}
