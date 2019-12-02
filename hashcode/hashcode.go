package hashcode

import (
	"fmt"
	"hash/fnv"
	"io"
	"math/rand"
	"reflect"
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	Debug = true
)

var (
	ErrorNotSUPPORT = fmt.Errorf("can not support")

	intSize = unsafe.Sizeof(1)
	ptrSize = unsafe.Sizeof((*int)(nil))
)

type hashEngine struct {
	putFunc func(w io.Writer, v interface{})
}

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	// Slow-path.
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		f()
		atomic.StoreUint32(&o.done, 1)
	}
}

func (o *Once) IsDone() bool {
	if atomic.LoadUint32(&o.done) == 1 {
		return true
	}
	return false
}

type interfaceInfo struct {
	isHash bool
	/*-- kind of key type --*/
	kind reflect.Kind
	/*-- index of field if it is a field of struct --*/
	index []int
	/*-- field informations of struct --*/
	fields []*interfaceInfo
	/*-- element information of array --*/
	elementInfo *interfaceInfo
	size        int
}

func HashInterface(k interface{}, isRead bool) (uint32, error) {
	var code uint32
	var err error
	h := fnv.New32a()
	h.Reset()
	switch v := k.(type) {
	case bool:
		h.Write((*((*[1]byte)(unsafe.Pointer(&v))))[:])
		code = h.Sum32()
	case int:
		h.Write((*((*[]byte)(unsafe.Pointer(&v))))[:])
		code = h.Sum32()
	case int8:
		h.Write((*((*[1]byte)(unsafe.Pointer(&v))))[:])
		code = h.Sum32()
	case int16:
		h.Write((*((*[2]byte)(unsafe.Pointer(&v))))[:])
		code = h.Sum32()
	case int32:
		h.Write((*((*[4]byte)(unsafe.Pointer(&v))))[:])
		code = h.Sum32()
	case int64:
		h.Write((*((*[8]byte)(unsafe.Pointer(&v))))[:])
		code = h.Sum32()
	case uint:
		h.Write((*((*[]byte)(unsafe.Pointer(&v))))[:])
		code = h.Sum32()
	case uint8:
		h.Write((*((*[1]byte)(unsafe.Pointer(&v))))[:])
		code = h.Sum32()
	case uint16:
		h.Write((*((*[2]byte)(unsafe.Pointer(&v))))[:])
		code = h.Sum32()
	case uint32:
		h.Write((*((*[4]byte)(unsafe.Pointer(&v))))[:])
		code = h.Sum32()
	case uint64:
		h.Write((*((*[8]byte)(unsafe.Pointer(&v))))[:])
		code = h.Sum32()
	case uintptr:
		h.Write((*((*[]byte)(unsafe.Pointer(&v))))[:])
		code = h.Sum32()
	case float32:
		//Nan != Nan, so use a rand number to generate hash code
		if v != v {
			v = rand.Float32()
		}
		h.Write((*((*[4]byte)(unsafe.Pointer(&v))))[:])
		code = h.Sum32()
	case float64:
		//Nan != Nan, so use a rand number to generate hash code
		if v != v {
			v = rand.Float64()
		}
		h.Write((*((*[8]byte)(unsafe.Pointer(&v))))[:])
		code = h.Sum32()
	case complex64:
		h.Write((*((*[8]byte)(unsafe.Pointer(&v))))[:])
		code = h.Sum32()
	case complex128:
		h.Write((*((*[128]byte)(unsafe.Pointer(&v))))[:])
		code = h.Sum32()
	case string:
		h.Write([]byte(v))
		code = h.Sum32()
	default:
		//if is not simple type

		code = h.Sum32()
	}
	return code, err
}

func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}

	rv := reflect.ValueOf(v)
	k := rv.Type().Kind()
	switch k {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}
