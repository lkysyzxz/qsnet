package util

import "sync"

type QMap struct {
	sync.RWMutex
	mp map[interface{}]interface{}
}
type RangeFunc func(key interface{},value interface{})
func (this *QMap)Range(f RangeFunc){
	this.Lock()
	defer this.Unlock()
	if len(this.mp)==0{
		return
	}
	for k,v := range this.mp{
		f(k,v)
	}
}

func (this *QMap)NonBlockRange(f RangeFunc){
	if len(this.mp)==0{
		return
	}
	for k,v := range this.mp{
		f(k,v)
	}
}

func (this *QMap)Clear(){
	this.Lock()
	defer this.Unlock()
	if len(this.mp)==0{
		return
	}
	for k,_ := range this.mp{
		delete(this.mp, k)
	}
}

func (this *QMap)GetAnyValue()interface{}{
	this.RLock()
	defer this.RUnlock()
	for _,v := range this.mp{
		return v
	}
	return nil
}

func (this *QMap)GetAnyKey()interface{}{
	this.RLock()
	defer this.RUnlock()
	for k,_ := range this.mp{
		return k
	}
	return nil
}

func (this *QMap) Get(k interface{}) interface{} {
	this.RLock()
	defer this.RUnlock()
	if val, ok := this.mp[k]; ok {
		return val
	} else {
		return nil
	}
}

func (this *QMap) Remove(k interface{}) {
	this.Lock()
	defer this.Unlock()
	delete(this.mp, k)
}

func (this *QMap) Add(k, v interface{}) bool {
	this.Lock()
	defer this.Unlock()
	val, ok := this.mp[k]
	if !ok || val != v {
		this.mp[k] = v
	} else {
		return false
	}
	return true
}

func (this *QMap)Length()int{
	this.RLock()
	defer this.RUnlock()
	return len(this.mp)
}

func (this *QMap) HasKey(k interface{}) bool {
	this.RLock()
	defer this.RUnlock()
	_, ok := this.mp[k]
	return ok
}

func NewQMap() *QMap {
	return &QMap{
		mp:make(map[interface{}]interface{}),
	}
}
