/*
 * Copyright (c) 2019 Huawei Technologies Co., Ltd.
 * MeshAccelerating is licensed under the Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *     http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
 * PURPOSE.
 * See the Mulan PSL v2 for more details.
 * Author: LemmyHuang
 * Create: 2021-10-09
 */

package maps

// #cgo CFLAGS: -I../../bpf/include
// #include "filter.h"
import "C"
import (
	"github.com/cilium/ebpf"
	"openeuler.io/mesh/pkg/bpf"
)

// CFilter = C.filter_t
type CFilter struct {
	Entry	C.filter_t
}

func (cf *CFilter) Lookup(key *GoMapKey) error {
	return bpf.Obj.SockConn.CgroupSockObjects.CgroupSockMaps.Filter.
		Lookup(key, &cf.Entry)
}

func (cf *CFilter) Update(key *GoMapKey) error {
	return bpf.Obj.SockConn.CgroupSockObjects.CgroupSockMaps.Filter.
		Update(key, &cf.Entry, ebpf.UpdateAny)
}

func (cf *CFilter) Delete(key *GoMapKey) error {
	return bpf.Obj.SockConn.CgroupSockObjects.CgroupSockMaps.Filter.
		Delete(key)
}

type GoFilter struct {

}

func (cf *CFilter) ToGolang() *GoFilter {
	return nil
}

func (gf *GoFilter) ToClang() *CFilter {
	return nil
}

// CFilterChain = C.filter_chain_t
type CFilterChain struct {
	Entry	C.filter_chain_t
}

func (cfc *CFilterChain) Lookup(key *GoMapKey) error {
	return bpf.Obj.SockConn.CgroupSockObjects.CgroupSockMaps.FilterChain.
		Lookup(key, &cfc.Entry)
}

func (cfc *CFilterChain) Update(key *GoMapKey) error {
	return bpf.Obj.SockConn.CgroupSockObjects.CgroupSockMaps.FilterChain.
		Update(key, &cfc.Entry, ebpf.UpdateAny)
}

func (cfc *CFilterChain) Delete(key *GoMapKey) error {
	return bpf.Obj.SockConn.CgroupSockObjects.CgroupSockMaps.FilterChain.
		Delete(key)
}

type GoFilterChain struct {

}

func (cfc *CFilterChain) ToGolang() *GoFilterChain {
	return nil
}

func (gfc *GoFilterChain) ToClang() *CFilterChain {
	return nil
}
