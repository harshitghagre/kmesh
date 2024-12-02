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

type KmeshSockopsWorkloadBpfSockTuple struct {
	Ipv4 struct {
		Saddr uint32
		Daddr uint32
		Sport uint16
		Dport uint16
	}
	_ [24]byte
}

type KmeshSockopsWorkloadBuf struct{ Data [40]int8 }

type KmeshSockopsWorkloadKmeshConfig struct {
	BpfLogLevel      uint32
	NodeIp           [4]uint32
	PodGateway       [4]uint32
	AuthzOffload     uint32
	EnableMonitoring uint32
}

type KmeshSockopsWorkloadLogEvent struct {
	Ret uint32
	Msg [255]int8
	_   [1]byte
}

type KmeshSockopsWorkloadManagerKey struct {
	NetnsCookie uint64
	_           [8]byte
}

type KmeshSockopsWorkloadOperationUsageData struct {
	StartTime     uint64
	EndTime       uint64
	PidTgid       uint64
	OperationType uint32
	_             [4]byte
}

type KmeshSockopsWorkloadOperationUsageKey struct {
	SocketCookie  uint64
	OperationType uint32
	_             [4]byte
}

type KmeshSockopsWorkloadSockStorageData struct {
	ConnectNs      uint64
	Direction      uint8
	ConnectSuccess uint8
	_              [6]byte
}

// LoadKmeshSockopsWorkload returns the embedded CollectionSpec for KmeshSockopsWorkload.
func LoadKmeshSockopsWorkload() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_KmeshSockopsWorkloadBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load KmeshSockopsWorkload: %w", err)
	}

	return spec, err
}

// LoadKmeshSockopsWorkloadObjects loads KmeshSockopsWorkload and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*KmeshSockopsWorkloadObjects
//	*KmeshSockopsWorkloadPrograms
//	*KmeshSockopsWorkloadMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadKmeshSockopsWorkloadObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadKmeshSockopsWorkload()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// KmeshSockopsWorkloadSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshSockopsWorkloadSpecs struct {
	KmeshSockopsWorkloadProgramSpecs
	KmeshSockopsWorkloadMapSpecs
}

// KmeshSockopsWorkloadSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshSockopsWorkloadProgramSpecs struct {
	SockopsProg *ebpf.ProgramSpec `ebpf:"sockops_prog"`
}

// KmeshSockopsWorkloadMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshSockopsWorkloadMapSpecs struct {
	KmAuth        *ebpf.MapSpec `ebpf:"km_auth"`
	KmBackend     *ebpf.MapSpec `ebpf:"km_backend"`
	KmConfigmap   *ebpf.MapSpec `ebpf:"km_configmap"`
	KmDstInfo     *ebpf.MapSpec `ebpf:"km_dstInfo"`
	KmEndpoint    *ebpf.MapSpec `ebpf:"km_endpoint"`
	KmEvents      *ebpf.MapSpec `ebpf:"km_events"`
	KmFrontend    *ebpf.MapSpec `ebpf:"km_frontend"`
	KmLogbuf      *ebpf.MapSpec `ebpf:"km_logbuf"`
	KmManage      *ebpf.MapSpec `ebpf:"km_manage"`
	KmService     *ebpf.MapSpec `ebpf:"km_service"`
	KmSockStorage *ebpf.MapSpec `ebpf:"km_sockStorage"`
	KmSocket      *ebpf.MapSpec `ebpf:"km_socket"`
	KmTcpInfo     *ebpf.MapSpec `ebpf:"km_tcpInfo"`
	KmTmpbuf      *ebpf.MapSpec `ebpf:"km_tmpbuf"`
	KmTuple       *ebpf.MapSpec `ebpf:"km_tuple"`
	KmWlpolicy    *ebpf.MapSpec `ebpf:"km_wlpolicy"`
	KmeshMap1600  *ebpf.MapSpec `ebpf:"kmesh_map1600"`
	KmeshMap192   *ebpf.MapSpec `ebpf:"kmesh_map192"`
	KmeshMap296   *ebpf.MapSpec `ebpf:"kmesh_map296"`
	KmeshMap64    *ebpf.MapSpec `ebpf:"kmesh_map64"`
	KmeshPerfInfo *ebpf.MapSpec `ebpf:"kmesh_perf_info"`
	KmeshPerfMap  *ebpf.MapSpec `ebpf:"kmesh_perf_map"`
}

// KmeshSockopsWorkloadObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshSockopsWorkloadObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshSockopsWorkloadObjects struct {
	KmeshSockopsWorkloadPrograms
	KmeshSockopsWorkloadMaps
}

func (o *KmeshSockopsWorkloadObjects) Close() error {
	return _KmeshSockopsWorkloadClose(
		&o.KmeshSockopsWorkloadPrograms,
		&o.KmeshSockopsWorkloadMaps,
	)
}

// KmeshSockopsWorkloadMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshSockopsWorkloadObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshSockopsWorkloadMaps struct {
	KmAuth        *ebpf.Map `ebpf:"km_auth"`
	KmBackend     *ebpf.Map `ebpf:"km_backend"`
	KmConfigmap   *ebpf.Map `ebpf:"km_configmap"`
	KmDstInfo     *ebpf.Map `ebpf:"km_dstInfo"`
	KmEndpoint    *ebpf.Map `ebpf:"km_endpoint"`
	KmEvents      *ebpf.Map `ebpf:"km_events"`
	KmFrontend    *ebpf.Map `ebpf:"km_frontend"`
	KmLogbuf      *ebpf.Map `ebpf:"km_logbuf"`
	KmManage      *ebpf.Map `ebpf:"km_manage"`
	KmService     *ebpf.Map `ebpf:"km_service"`
	KmSockStorage *ebpf.Map `ebpf:"km_sockStorage"`
	KmSocket      *ebpf.Map `ebpf:"km_socket"`
	KmTcpInfo     *ebpf.Map `ebpf:"km_tcpInfo"`
	KmTmpbuf      *ebpf.Map `ebpf:"km_tmpbuf"`
	KmTuple       *ebpf.Map `ebpf:"km_tuple"`
	KmWlpolicy    *ebpf.Map `ebpf:"km_wlpolicy"`
	KmeshMap1600  *ebpf.Map `ebpf:"kmesh_map1600"`
	KmeshMap192   *ebpf.Map `ebpf:"kmesh_map192"`
	KmeshMap296   *ebpf.Map `ebpf:"kmesh_map296"`
	KmeshMap64    *ebpf.Map `ebpf:"kmesh_map64"`
	KmeshPerfInfo *ebpf.Map `ebpf:"kmesh_perf_info"`
	KmeshPerfMap  *ebpf.Map `ebpf:"kmesh_perf_map"`
}

func (m *KmeshSockopsWorkloadMaps) Close() error {
	return _KmeshSockopsWorkloadClose(
		m.KmAuth,
		m.KmBackend,
		m.KmConfigmap,
		m.KmDstInfo,
		m.KmEndpoint,
		m.KmEvents,
		m.KmFrontend,
		m.KmLogbuf,
		m.KmManage,
		m.KmService,
		m.KmSockStorage,
		m.KmSocket,
		m.KmTcpInfo,
		m.KmTmpbuf,
		m.KmTuple,
		m.KmWlpolicy,
		m.KmeshMap1600,
		m.KmeshMap192,
		m.KmeshMap296,
		m.KmeshMap64,
		m.KmeshPerfInfo,
		m.KmeshPerfMap,
	)
}

// KmeshSockopsWorkloadPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshSockopsWorkloadObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshSockopsWorkloadPrograms struct {
	SockopsProg *ebpf.Program `ebpf:"sockops_prog"`
}

func (p *KmeshSockopsWorkloadPrograms) Close() error {
	return _KmeshSockopsWorkloadClose(
		p.SockopsProg,
	)
}

func _KmeshSockopsWorkloadClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed kmeshsockopsworkload_bpfeb.o
var _KmeshSockopsWorkloadBytes []byte
