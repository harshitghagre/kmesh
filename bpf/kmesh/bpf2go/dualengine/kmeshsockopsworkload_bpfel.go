// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64

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
	KmDstinfo     *ebpf.MapSpec `ebpf:"km_dstinfo"`
	KmEndpoint    *ebpf.MapSpec `ebpf:"km_endpoint"`
	KmEvents      *ebpf.MapSpec `ebpf:"km_events"`
	KmFrontend    *ebpf.MapSpec `ebpf:"km_frontend"`
	KmLogbuf      *ebpf.MapSpec `ebpf:"km_logbuf"`
	KmManage      *ebpf.MapSpec `ebpf:"km_manage"`
	KmPerfInfo    *ebpf.MapSpec `ebpf:"km_perf_info"`
	KmPerfMap     *ebpf.MapSpec `ebpf:"km_perf_map"`
	KmService     *ebpf.MapSpec `ebpf:"km_service"`
	KmSocket      *ebpf.MapSpec `ebpf:"km_socket"`
	KmSockstorage *ebpf.MapSpec `ebpf:"km_sockstorage"`
	KmTcpinfo     *ebpf.MapSpec `ebpf:"km_tcpinfo"`
	KmTmpbuf      *ebpf.MapSpec `ebpf:"km_tmpbuf"`
	KmTuple       *ebpf.MapSpec `ebpf:"km_tuple"`
	KmWlpolicy    *ebpf.MapSpec `ebpf:"km_wlpolicy"`
	KmeshMap1600  *ebpf.MapSpec `ebpf:"kmesh_map1600"`
	KmeshMap192   *ebpf.MapSpec `ebpf:"kmesh_map192"`
	KmeshMap296   *ebpf.MapSpec `ebpf:"kmesh_map296"`
	KmeshMap64    *ebpf.MapSpec `ebpf:"kmesh_map64"`
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
	KmDstinfo     *ebpf.Map `ebpf:"km_dstinfo"`
	KmEndpoint    *ebpf.Map `ebpf:"km_endpoint"`
	KmEvents      *ebpf.Map `ebpf:"km_events"`
	KmFrontend    *ebpf.Map `ebpf:"km_frontend"`
	KmLogbuf      *ebpf.Map `ebpf:"km_logbuf"`
	KmManage      *ebpf.Map `ebpf:"km_manage"`
	KmPerfInfo    *ebpf.Map `ebpf:"km_perf_info"`
	KmPerfMap     *ebpf.Map `ebpf:"km_perf_map"`
	KmService     *ebpf.Map `ebpf:"km_service"`
	KmSocket      *ebpf.Map `ebpf:"km_socket"`
	KmSockstorage *ebpf.Map `ebpf:"km_sockstorage"`
	KmTcpinfo     *ebpf.Map `ebpf:"km_tcpinfo"`
	KmTmpbuf      *ebpf.Map `ebpf:"km_tmpbuf"`
	KmTuple       *ebpf.Map `ebpf:"km_tuple"`
	KmWlpolicy    *ebpf.Map `ebpf:"km_wlpolicy"`
	KmeshMap1600  *ebpf.Map `ebpf:"kmesh_map1600"`
	KmeshMap192   *ebpf.Map `ebpf:"kmesh_map192"`
	KmeshMap296   *ebpf.Map `ebpf:"kmesh_map296"`
	KmeshMap64    *ebpf.Map `ebpf:"kmesh_map64"`
}

func (m *KmeshSockopsWorkloadMaps) Close() error {
	return _KmeshSockopsWorkloadClose(
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
		m.KmSocket,
		m.KmSockstorage,
		m.KmTcpinfo,
		m.KmTmpbuf,
		m.KmTuple,
		m.KmWlpolicy,
		m.KmeshMap1600,
		m.KmeshMap192,
		m.KmeshMap296,
		m.KmeshMap64,
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
//go:embed kmeshsockopsworkload_bpfel.o
var _KmeshSockopsWorkloadBytes []byte
