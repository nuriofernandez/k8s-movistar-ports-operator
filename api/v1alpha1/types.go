package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	GroupVersion  = schema.GroupVersion{Group: "nurio.me", Version: "v1alpha1"}
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}
	AddToScheme   = SchemeBuilder.AddToScheme
)

type MovistarPortSpec struct {
	ExternalPort int32  `json:"externalPort"`
	InternalPort int32  `json:"internalPort"`
	Protocol     string `json:"protocol"`
	Host         string `json:"host"`
}

type MovistarPort struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MovistarPortSpec `json:"spec,omitempty"`
}

type MovistarPortList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MovistarPort `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MovistarPort{}, &MovistarPortList{})
}

func (in *MovistarPort) DeepCopyObject() runtime.Object     { c := in.DeepCopy(); return c }
func (in *MovistarPort) DeepCopy() *MovistarPort            { out := *in; return &out }
func (in *MovistarPortList) DeepCopyObject() runtime.Object { c := in.DeepCopy(); return c }
func (in *MovistarPortList) DeepCopy() *MovistarPortList    { out := *in; return &out }
