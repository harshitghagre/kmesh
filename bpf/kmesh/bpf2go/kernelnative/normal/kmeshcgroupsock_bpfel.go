// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64

package normal

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type KmeshCgroupSockBuf struct{ Data [40]int8 }

type KmeshCgroupSockClusterSockData struct{ ClusterId uint32 }

type KmeshCgroupSockLogEvent struct {
	Ret uint32
	Msg [255]int8
	_   [1]byte
}

type KmeshCgroupSockManagerKey struct {
	NetnsCookie uint64
	_           [8]byte
}

type KmeshCgroupSockSockStorageData struct {
	ConnectNs      uint64
	Direction      uint8
	ConnectSuccess uint8
	_              [6]byte
}

// LoadKmeshCgroupSock returns the embedded CollectionSpec for KmeshCgroupSock.
func LoadKmeshCgroupSock() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_KmeshCgroupSockBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load KmeshCgroupSock: %w", err)
	}

	return spec, err
}

// LoadKmeshCgroupSockObjects loads KmeshCgroupSock and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*KmeshCgroupSockObjects
//	*KmeshCgroupSockPrograms
//	*KmeshCgroupSockMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadKmeshCgroupSockObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadKmeshCgroupSock()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// KmeshCgroupSockSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockSpecs struct {
	KmeshCgroupSockProgramSpecs
	KmeshCgroupSockMapSpecs
}

// KmeshCgroupSockSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockProgramSpecs struct {
	CgroupConnect4Prog *ebpf.ProgramSpec `ebpf:"cgroup_connect4_prog"`
	ClusterManager     *ebpf.ProgramSpec `ebpf:"cluster_manager"`
	FilterChainManager *ebpf.ProgramSpec `ebpf:"filter_chain_manager"`
	FilterManager      *ebpf.ProgramSpec `ebpf:"filter_manager"`
}

// KmeshCgroupSockMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockMapSpecs struct {
	BpfLogLevel         *ebpf.MapSpec `ebpf:"bpf_log_level"`
	InnerMap            *ebpf.MapSpec `ebpf:"inner_map"`
	KmeshCluster        *ebpf.MapSpec `ebpf:"kmesh_cluster"`
	KmeshClusterStats   *ebpf.MapSpec `ebpf:"kmesh_cluster_stats"`
	KmeshEvents         *ebpf.MapSpec `ebpf:"kmesh_events"`
	KmeshListener       *ebpf.MapSpec `ebpf:"kmesh_listener"`
	KmeshManage         *ebpf.MapSpec `ebpf:"kmesh_manage"`
	KmeshTailCallCtx    *ebpf.MapSpec `ebpf:"kmesh_tail_call_ctx"`
	KmeshTailCallProg   *ebpf.MapSpec `ebpf:"kmesh_tail_call_prog"`
	MapOfClusterEps     *ebpf.MapSpec `ebpf:"map_of_cluster_eps"`
	MapOfClusterEpsData *ebpf.MapSpec `ebpf:"map_of_cluster_eps_data"`
	MapOfClusterSock    *ebpf.MapSpec `ebpf:"map_of_cluster_sock"`
	MapOfSockStorage    *ebpf.MapSpec `ebpf:"map_of_sock_storage"`
	OuterMap            *ebpf.MapSpec `ebpf:"outer_map"`
	TmpBuf              *ebpf.MapSpec `ebpf:"tmp_buf"`
	TmpLogBuf           *ebpf.MapSpec `ebpf:"tmp_log_buf"`
}

// KmeshCgroupSockObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockObjects struct {
	KmeshCgroupSockPrograms
	KmeshCgroupSockMaps
}

func (o *KmeshCgroupSockObjects) Close() error {
	return _KmeshCgroupSockClose(
		&o.KmeshCgroupSockPrograms,
		&o.KmeshCgroupSockMaps,
	)
}

// KmeshCgroupSockMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockMaps struct {
	BpfLogLevel         *ebpf.Map `ebpf:"bpf_log_level"`
	InnerMap            *ebpf.Map `ebpf:"inner_map"`
	KmeshCluster        *ebpf.Map `ebpf:"kmesh_cluster"`
	KmeshClusterStats   *ebpf.Map `ebpf:"kmesh_cluster_stats"`
	KmeshEvents         *ebpf.Map `ebpf:"kmesh_events"`
	KmeshListener       *ebpf.Map `ebpf:"kmesh_listener"`
	KmeshManage         *ebpf.Map `ebpf:"kmesh_manage"`
	KmeshTailCallCtx    *ebpf.Map `ebpf:"kmesh_tail_call_ctx"`
	KmeshTailCallProg   *ebpf.Map `ebpf:"kmesh_tail_call_prog"`
	MapOfClusterEps     *ebpf.Map `ebpf:"map_of_cluster_eps"`
	MapOfClusterEpsData *ebpf.Map `ebpf:"map_of_cluster_eps_data"`
	MapOfClusterSock    *ebpf.Map `ebpf:"map_of_cluster_sock"`
	MapOfSockStorage    *ebpf.Map `ebpf:"map_of_sock_storage"`
	OuterMap            *ebpf.Map `ebpf:"outer_map"`
	TmpBuf              *ebpf.Map `ebpf:"tmp_buf"`
	TmpLogBuf           *ebpf.Map `ebpf:"tmp_log_buf"`
}

func (m *KmeshCgroupSockMaps) Close() error {
	return _KmeshCgroupSockClose(
		m.BpfLogLevel,
		m.InnerMap,
		m.KmeshCluster,
		m.KmeshClusterStats,
		m.KmeshEvents,
		m.KmeshListener,
		m.KmeshManage,
		m.KmeshTailCallCtx,
		m.KmeshTailCallProg,
		m.MapOfClusterEps,
		m.MapOfClusterEpsData,
		m.MapOfClusterSock,
		m.MapOfSockStorage,
		m.OuterMap,
		m.TmpBuf,
		m.TmpLogBuf,
	)
}

// KmeshCgroupSockPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockPrograms struct {
	CgroupConnect4Prog *ebpf.Program `ebpf:"cgroup_connect4_prog"`
	ClusterManager     *ebpf.Program `ebpf:"cluster_manager"`
	FilterChainManager *ebpf.Program `ebpf:"filter_chain_manager"`
	FilterManager      *ebpf.Program `ebpf:"filter_manager"`
}

func (p *KmeshCgroupSockPrograms) Close() error {
	return _KmeshCgroupSockClose(
		p.CgroupConnect4Prog,
		p.ClusterManager,
		p.FilterChainManager,
		p.FilterManager,
	)
}

func _KmeshCgroupSockClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed kmeshcgroupsock_bpfel.o
var _KmeshCgroupSockBytes []byte
