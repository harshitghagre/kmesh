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

type KmeshSendmsgBpfSockTuple struct {
	Ipv4 struct {
		Saddr uint32
		Daddr uint32
		Sport uint16
		Dport uint16
	}
	_ [24]byte
}

type KmeshSendmsgBuf struct{ Data [40]int8 }

type KmeshSendmsgKmeshConfig struct {
	BpfLogLevel      uint32
	NodeIp           [4]uint32
	PodGateway       [4]uint32
	AuthzOffload     uint32
	EnableMonitoring uint32
}

type KmeshSendmsgLogEvent struct {
	Ret uint32
	Msg [255]int8
	_   [1]byte
}

// LoadKmeshSendmsg returns the embedded CollectionSpec for KmeshSendmsg.
func LoadKmeshSendmsg() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_KmeshSendmsgBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load KmeshSendmsg: %w", err)
	}

	return spec, err
}

// LoadKmeshSendmsgObjects loads KmeshSendmsg and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*KmeshSendmsgObjects
//	*KmeshSendmsgPrograms
//	*KmeshSendmsgMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadKmeshSendmsgObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadKmeshSendmsg()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// KmeshSendmsgSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshSendmsgSpecs struct {
	KmeshSendmsgProgramSpecs
	KmeshSendmsgMapSpecs
}

// KmeshSendmsgSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshSendmsgProgramSpecs struct {
	SendmsgProg *ebpf.ProgramSpec `ebpf:"sendmsg_prog"`
}

// KmeshSendmsgMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshSendmsgMapSpecs struct {
	KmConfigmap *ebpf.MapSpec `ebpf:"km_configmap"`
	KmDstinfo   *ebpf.MapSpec `ebpf:"km_dstinfo"`
	KmEvents    *ebpf.MapSpec `ebpf:"km_events"`
	KmLogbuf    *ebpf.MapSpec `ebpf:"km_logbuf"`
	KmTmpbuf    *ebpf.MapSpec `ebpf:"km_tmpbuf"`
}

// KmeshSendmsgObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshSendmsgObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshSendmsgObjects struct {
	KmeshSendmsgPrograms
	KmeshSendmsgMaps
}

func (o *KmeshSendmsgObjects) Close() error {
	return _KmeshSendmsgClose(
		&o.KmeshSendmsgPrograms,
		&o.KmeshSendmsgMaps,
	)
}

// KmeshSendmsgMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshSendmsgObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshSendmsgMaps struct {
	KmConfigmap *ebpf.Map `ebpf:"km_configmap"`
	KmDstinfo   *ebpf.Map `ebpf:"km_dstinfo"`
	KmEvents    *ebpf.Map `ebpf:"km_events"`
	KmLogbuf    *ebpf.Map `ebpf:"km_logbuf"`
	KmTmpbuf    *ebpf.Map `ebpf:"km_tmpbuf"`
}

func (m *KmeshSendmsgMaps) Close() error {
	return _KmeshSendmsgClose(
		m.KmConfigmap,
		m.KmDstinfo,
		m.KmEvents,
		m.KmLogbuf,
		m.KmTmpbuf,
	)
}

// KmeshSendmsgPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshSendmsgObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshSendmsgPrograms struct {
	SendmsgProg *ebpf.Program `ebpf:"sendmsg_prog"`
}

func (p *KmeshSendmsgPrograms) Close() error {
	return _KmeshSendmsgClose(
		p.SendmsgProg,
	)
}

func _KmeshSendmsgClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed kmeshsendmsg_bpfeb.o
var _KmeshSendmsgBytes []byte
