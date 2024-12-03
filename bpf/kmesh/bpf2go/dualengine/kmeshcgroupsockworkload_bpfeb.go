// Code generated by bpf2go; DO NOT EDIT.
//go:build mips || mips64 || ppc64 || s390x

package dualengine

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type KmeshCgroupSockWorkloadBpfSockTuple struct {
	Ipv4 struct {
		Saddr uint32
		Daddr uint32
		Sport uint16
		Dport uint16
	}
	_ [24]byte
}

type KmeshCgroupSockWorkloadBuf struct{ Data [40]int8 }

type KmeshCgroupSockWorkloadKmeshConfig struct {
	BpfLogLevel      uint32
	NodeIp           [4]uint32
	PodGateway       [4]uint32
	AuthzOffload     uint32
	EnableMonitoring uint32
}

type KmeshCgroupSockWorkloadLogEvent struct {
	Ret uint32
	Msg [255]int8
	_   [1]byte
}

type KmeshCgroupSockWorkloadManagerKey struct {
	NetnsCookie uint64
	_           [8]byte
}

type KmeshCgroupSockWorkloadOperationUsageData struct {
	StartTime     uint64
	EndTime       uint64
	PidTgid       uint64
	OperationType uint32
	_             [4]byte
}

type KmeshCgroupSockWorkloadOperationUsageKey struct {
	SocketCookie  uint64
	OperationType uint32
	_             [4]byte
}

type KmeshCgroupSockWorkloadSockStorageData struct {
	ConnectNs      uint64
	Direction      uint8
	ConnectSuccess uint8
	_              [6]byte
}

// LoadKmeshCgroupSockWorkload returns the embedded CollectionSpec for KmeshCgroupSockWorkload.
func LoadKmeshCgroupSockWorkload() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_KmeshCgroupSockWorkloadBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load KmeshCgroupSockWorkload: %w", err)
	}

	return spec, err
}

// LoadKmeshCgroupSockWorkloadObjects loads KmeshCgroupSockWorkload and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*KmeshCgroupSockWorkloadObjects
//	*KmeshCgroupSockWorkloadPrograms
//	*KmeshCgroupSockWorkloadMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadKmeshCgroupSockWorkloadObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadKmeshCgroupSockWorkload()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// KmeshCgroupSockWorkloadSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockWorkloadSpecs struct {
	KmeshCgroupSockWorkloadProgramSpecs
	KmeshCgroupSockWorkloadMapSpecs
}

// KmeshCgroupSockWorkloadSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockWorkloadProgramSpecs struct {
	CgroupConnect4Prog *ebpf.ProgramSpec `ebpf:"cgroup_connect4_prog"`
	CgroupConnect6Prog *ebpf.ProgramSpec `ebpf:"cgroup_connect6_prog"`
}

// KmeshCgroupSockWorkloadMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockWorkloadMapSpecs struct {
	KmAuth         *ebpf.MapSpec `ebpf:"km_auth"`
	KmBackend      *ebpf.MapSpec `ebpf:"km_backend"`
	KmConfigmap    *ebpf.MapSpec `ebpf:"km_configmap"`
	KmDstinfo      *ebpf.MapSpec `ebpf:"km_dstinfo"`
	KmEndpoint     *ebpf.MapSpec `ebpf:"km_endpoint"`
	KmEvents       *ebpf.MapSpec `ebpf:"km_events"`
	KmFrontend     *ebpf.MapSpec `ebpf:"km_frontend"`
	KmLogbuf       *ebpf.MapSpec `ebpf:"km_logbuf"`
	KmManage       *ebpf.MapSpec `ebpf:"km_manage"`
	KmPerfInfo     *ebpf.MapSpec `ebpf:"km_perf_info"`
	KmPerfMap      *ebpf.MapSpec `ebpf:"km_perf_map"`
	KmService      *ebpf.MapSpec `ebpf:"km_service"`
	KmSockStorage  *ebpf.MapSpec `ebpf:"km_sockStorage"`
	KmTailcallprog *ebpf.MapSpec `ebpf:"km_tailcallprog"`
	KmTcpinfo      *ebpf.MapSpec `ebpf:"km_tcpinfo"`
	KmTmpbuf       *ebpf.MapSpec `ebpf:"km_tmpbuf"`
	KmTuple        *ebpf.MapSpec `ebpf:"km_tuple"`
	KmWlpolicy     *ebpf.MapSpec `ebpf:"km_wlpolicy"`
	KmXdptailcall  *ebpf.MapSpec `ebpf:"km_xdptailcall"`
	KmeshMap1600   *ebpf.MapSpec `ebpf:"kmesh_map1600"`
	KmeshMap192    *ebpf.MapSpec `ebpf:"kmesh_map192"`
	KmeshMap296    *ebpf.MapSpec `ebpf:"kmesh_map296"`
	KmeshMap64     *ebpf.MapSpec `ebpf:"kmesh_map64"`
}

// KmeshCgroupSockWorkloadObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockWorkloadObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockWorkloadObjects struct {
	KmeshCgroupSockWorkloadPrograms
	KmeshCgroupSockWorkloadMaps
}

func (o *KmeshCgroupSockWorkloadObjects) Close() error {
	return _KmeshCgroupSockWorkloadClose(
		&o.KmeshCgroupSockWorkloadPrograms,
		&o.KmeshCgroupSockWorkloadMaps,
	)
}

// KmeshCgroupSockWorkloadMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockWorkloadObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockWorkloadMaps struct {
	KmAuth         *ebpf.Map `ebpf:"km_auth"`
	KmBackend      *ebpf.Map `ebpf:"km_backend"`
	KmConfigmap    *ebpf.Map `ebpf:"km_configmap"`
	KmDstinfo      *ebpf.Map `ebpf:"km_dstinfo"`
	KmEndpoint     *ebpf.Map `ebpf:"km_endpoint"`
	KmEvents       *ebpf.Map `ebpf:"km_events"`
	KmFrontend     *ebpf.Map `ebpf:"km_frontend"`
	KmLogbuf       *ebpf.Map `ebpf:"km_logbuf"`
	KmManage       *ebpf.Map `ebpf:"km_manage"`
	KmPerfInfo     *ebpf.Map `ebpf:"km_perf_info"`
	KmPerfMap      *ebpf.Map `ebpf:"km_perf_map"`
	KmService      *ebpf.Map `ebpf:"km_service"`
	KmSockStorage  *ebpf.Map `ebpf:"km_sockStorage"`
	KmTailcallprog *ebpf.Map `ebpf:"km_tailcallprog"`
	KmTcpinfo      *ebpf.Map `ebpf:"km_tcpinfo"`
	KmTmpbuf       *ebpf.Map `ebpf:"km_tmpbuf"`
	KmTuple        *ebpf.Map `ebpf:"km_tuple"`
	KmWlpolicy     *ebpf.Map `ebpf:"km_wlpolicy"`
	KmXdptailcall  *ebpf.Map `ebpf:"km_xdptailcall"`
	KmeshMap1600   *ebpf.Map `ebpf:"kmesh_map1600"`
	KmeshMap192    *ebpf.Map `ebpf:"kmesh_map192"`
	KmeshMap296    *ebpf.Map `ebpf:"kmesh_map296"`
	KmeshMap64     *ebpf.Map `ebpf:"kmesh_map64"`
}

func (m *KmeshCgroupSockWorkloadMaps) Close() error {
	return _KmeshCgroupSockWorkloadClose(
		m.KmAuth,
		m.KmBackend,
		m.KmConfigmap,
		m.KmDstinfo,
		m.KmEndpoint,
		m.KmEvents,
		m.KmFrontend,
		m.KmLogbuf,
		m.KmManage,
		m.KmPerfInfo,
		m.KmPerfMap,
		m.KmService,
		m.KmSockStorage,
		m.KmTailcallprog,
		m.KmTcpinfo,
		m.KmTmpbuf,
		m.KmTuple,
		m.KmWlpolicy,
		m.KmXdptailcall,
		m.KmeshMap1600,
		m.KmeshMap192,
		m.KmeshMap296,
		m.KmeshMap64,
	)
}

// KmeshCgroupSockWorkloadPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockWorkloadObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockWorkloadPrograms struct {
	CgroupConnect4Prog *ebpf.Program `ebpf:"cgroup_connect4_prog"`
	CgroupConnect6Prog *ebpf.Program `ebpf:"cgroup_connect6_prog"`
}

func (p *KmeshCgroupSockWorkloadPrograms) Close() error {
	return _KmeshCgroupSockWorkloadClose(
		p.CgroupConnect4Prog,
		p.CgroupConnect6Prog,
	)
}

func _KmeshCgroupSockWorkloadClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed kmeshcgroupsockworkload_bpfeb.o
var _KmeshCgroupSockWorkloadBytes []byte
