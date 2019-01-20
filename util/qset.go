package util

import "sync"

type QSet struct{
	mp map[interface{}]struct{}
	sync.RWMutex
}

func(self *QSet)Add(elem interface{})bool{
	self.Lock()
	defer self.Unlock()
	if _,ok := self.mp[elem];ok{
		return false
	}
	self.mp[elem]= struct{}{}
	return true
}

func(self *QSet)Contains(elem interface{})bool{
	self.RLock()
	defer self.RUnlock()
	_,found := self.mp[elem]
	return found
}

func(self *QSet)Remove(elem interface{}){
	self.Lock()
	defer self.Unlock()
	delete(self.mp,elem)
}

func(self *QSet)Clear(){
	self.Lock()
	defer self.Unlock()
	self.mp = make(map[interface{}]struct{})
}

func(self *QSet)Range(f func(interface{})){
	self.RLock()
	defer self.RUnlock()
	for elem := range self.mp {
		f(elem)
	}
}

func(self *QSet)Length()int{
	self.RLock()
	defer self.RUnlock()
	return len(self.mp)
}

func NewSet()QSet{
	return QSet{
		mp:make(map[interface{}]struct{}),
	}
}
