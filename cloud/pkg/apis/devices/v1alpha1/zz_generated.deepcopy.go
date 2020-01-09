// +build !ignore_autogenerated

/*
Copyright 2020 The KubeEdge Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BluetoothOperations) DeepCopyInto(out *BluetoothOperations) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BluetoothOperations.
func (in *BluetoothOperations) DeepCopy() *BluetoothOperations {
	if in == nil {
		return nil
	}
	out := new(BluetoothOperations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BluetoothReadConverter) DeepCopyInto(out *BluetoothReadConverter) {
	*out = *in
	if in.OrderOfOperations != nil {
		in, out := &in.OrderOfOperations, &out.OrderOfOperations
		*out = make([]BluetoothOperations, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BluetoothReadConverter.
func (in *BluetoothReadConverter) DeepCopy() *BluetoothReadConverter {
	if in == nil {
		return nil
	}
	out := new(BluetoothReadConverter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Device) DeepCopyInto(out *Device) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Device.
func (in *Device) DeepCopy() *Device {
	if in == nil {
		return nil
	}
	out := new(Device)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Device) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceList) DeepCopyInto(out *DeviceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Device, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceList.
func (in *DeviceList) DeepCopy() *DeviceList {
	if in == nil {
		return nil
	}
	out := new(DeviceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceModel) DeepCopyInto(out *DeviceModel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceModel.
func (in *DeviceModel) DeepCopy() *DeviceModel {
	if in == nil {
		return nil
	}
	out := new(DeviceModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceModel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceModelList) DeepCopyInto(out *DeviceModelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeviceModel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceModelList.
func (in *DeviceModelList) DeepCopy() *DeviceModelList {
	if in == nil {
		return nil
	}
	out := new(DeviceModelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceModelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceModelSpec) DeepCopyInto(out *DeviceModelSpec) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]DeviceProperty, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PropertyVisitors != nil {
		in, out := &in.PropertyVisitors, &out.PropertyVisitors
		*out = make([]DevicePropertyVisitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceModelSpec.
func (in *DeviceModelSpec) DeepCopy() *DeviceModelSpec {
	if in == nil {
		return nil
	}
	out := new(DeviceModelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceProperty) DeepCopyInto(out *DeviceProperty) {
	*out = *in
	in.Type.DeepCopyInto(&out.Type)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceProperty.
func (in *DeviceProperty) DeepCopy() *DeviceProperty {
	if in == nil {
		return nil
	}
	out := new(DeviceProperty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePropertyVisitor) DeepCopyInto(out *DevicePropertyVisitor) {
	*out = *in
	in.VisitorConfig.DeepCopyInto(&out.VisitorConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePropertyVisitor.
func (in *DevicePropertyVisitor) DeepCopy() *DevicePropertyVisitor {
	if in == nil {
		return nil
	}
	out := new(DevicePropertyVisitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSpec) DeepCopyInto(out *DeviceSpec) {
	*out = *in
	if in.DeviceModelRef != nil {
		in, out := &in.DeviceModelRef, &out.DeviceModelRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	in.Protocol.DeepCopyInto(&out.Protocol)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(v1.NodeSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSpec.
func (in *DeviceSpec) DeepCopy() *DeviceSpec {
	if in == nil {
		return nil
	}
	out := new(DeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceStatus) DeepCopyInto(out *DeviceStatus) {
	*out = *in
	if in.Twins != nil {
		in, out := &in.Twins, &out.Twins
		*out = make([]Twin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceStatus.
func (in *DeviceStatus) DeepCopy() *DeviceStatus {
	if in == nil {
		return nil
	}
	out := new(DeviceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PropertyType) DeepCopyInto(out *PropertyType) {
	*out = *in
	if in.Int != nil {
		in, out := &in.Int, &out.Int
		*out = new(PropertyTypeInt64)
		**out = **in
	}
	if in.String != nil {
		in, out := &in.String, &out.String
		*out = new(PropertyTypeString)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PropertyType.
func (in *PropertyType) DeepCopy() *PropertyType {
	if in == nil {
		return nil
	}
	out := new(PropertyType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PropertyTypeInt64) DeepCopyInto(out *PropertyTypeInt64) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PropertyTypeInt64.
func (in *PropertyTypeInt64) DeepCopy() *PropertyTypeInt64 {
	if in == nil {
		return nil
	}
	out := new(PropertyTypeInt64)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PropertyTypeString) DeepCopyInto(out *PropertyTypeString) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PropertyTypeString.
func (in *PropertyTypeString) DeepCopy() *PropertyTypeString {
	if in == nil {
		return nil
	}
	out := new(PropertyTypeString)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtocolConfig) DeepCopyInto(out *ProtocolConfig) {
	*out = *in
	if in.OpcUA != nil {
		in, out := &in.OpcUA, &out.OpcUA
		*out = new(ProtocolConfigOpcUA)
		**out = **in
	}
	if in.Modbus != nil {
		in, out := &in.Modbus, &out.Modbus
		*out = new(ProtocolConfigModbus)
		(*in).DeepCopyInto(*out)
	}
	if in.Bluetooth != nil {
		in, out := &in.Bluetooth, &out.Bluetooth
		*out = new(ProtocolConfigBluetooth)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtocolConfig.
func (in *ProtocolConfig) DeepCopy() *ProtocolConfig {
	if in == nil {
		return nil
	}
	out := new(ProtocolConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtocolConfigBluetooth) DeepCopyInto(out *ProtocolConfigBluetooth) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtocolConfigBluetooth.
func (in *ProtocolConfigBluetooth) DeepCopy() *ProtocolConfigBluetooth {
	if in == nil {
		return nil
	}
	out := new(ProtocolConfigBluetooth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtocolConfigModbus) DeepCopyInto(out *ProtocolConfigModbus) {
	*out = *in
	if in.RTU != nil {
		in, out := &in.RTU, &out.RTU
		*out = new(ProtocolConfigModbusRTU)
		**out = **in
	}
	if in.TCP != nil {
		in, out := &in.TCP, &out.TCP
		*out = new(ProtocolConfigModbusTCP)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtocolConfigModbus.
func (in *ProtocolConfigModbus) DeepCopy() *ProtocolConfigModbus {
	if in == nil {
		return nil
	}
	out := new(ProtocolConfigModbus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtocolConfigModbusRTU) DeepCopyInto(out *ProtocolConfigModbusRTU) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtocolConfigModbusRTU.
func (in *ProtocolConfigModbusRTU) DeepCopy() *ProtocolConfigModbusRTU {
	if in == nil {
		return nil
	}
	out := new(ProtocolConfigModbusRTU)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtocolConfigModbusTCP) DeepCopyInto(out *ProtocolConfigModbusTCP) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtocolConfigModbusTCP.
func (in *ProtocolConfigModbusTCP) DeepCopy() *ProtocolConfigModbusTCP {
	if in == nil {
		return nil
	}
	out := new(ProtocolConfigModbusTCP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtocolConfigOpcUA) DeepCopyInto(out *ProtocolConfigOpcUA) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtocolConfigOpcUA.
func (in *ProtocolConfigOpcUA) DeepCopy() *ProtocolConfigOpcUA {
	if in == nil {
		return nil
	}
	out := new(ProtocolConfigOpcUA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Twin) DeepCopyInto(out *Twin) {
	*out = *in
	in.Desired.DeepCopyInto(&out.Desired)
	in.Reported.DeepCopyInto(&out.Reported)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Twin.
func (in *Twin) DeepCopy() *Twin {
	if in == nil {
		return nil
	}
	out := new(Twin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinProperty) DeepCopyInto(out *TwinProperty) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinProperty.
func (in *TwinProperty) DeepCopy() *TwinProperty {
	if in == nil {
		return nil
	}
	out := new(TwinProperty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VisitorConfig) DeepCopyInto(out *VisitorConfig) {
	*out = *in
	if in.OpcUA != nil {
		in, out := &in.OpcUA, &out.OpcUA
		*out = new(VisitorConfigOPCUA)
		**out = **in
	}
	if in.Modbus != nil {
		in, out := &in.Modbus, &out.Modbus
		*out = new(VisitorConfigModbus)
		**out = **in
	}
	if in.Bluetooth != nil {
		in, out := &in.Bluetooth, &out.Bluetooth
		*out = new(VisitorConfigBluetooth)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VisitorConfig.
func (in *VisitorConfig) DeepCopy() *VisitorConfig {
	if in == nil {
		return nil
	}
	out := new(VisitorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VisitorConfigBluetooth) DeepCopyInto(out *VisitorConfigBluetooth) {
	*out = *in
	if in.DataWriteToBluetooth != nil {
		in, out := &in.DataWriteToBluetooth, &out.DataWriteToBluetooth
		*out = make(map[string][]byte, len(*in))
		for key, val := range *in {
			var outVal []byte
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]byte, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	in.BluetoothDataConverter.DeepCopyInto(&out.BluetoothDataConverter)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VisitorConfigBluetooth.
func (in *VisitorConfigBluetooth) DeepCopy() *VisitorConfigBluetooth {
	if in == nil {
		return nil
	}
	out := new(VisitorConfigBluetooth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VisitorConfigModbus) DeepCopyInto(out *VisitorConfigModbus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VisitorConfigModbus.
func (in *VisitorConfigModbus) DeepCopy() *VisitorConfigModbus {
	if in == nil {
		return nil
	}
	out := new(VisitorConfigModbus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VisitorConfigOPCUA) DeepCopyInto(out *VisitorConfigOPCUA) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VisitorConfigOPCUA.
func (in *VisitorConfigOPCUA) DeepCopy() *VisitorConfigOPCUA {
	if in == nil {
		return nil
	}
	out := new(VisitorConfigOPCUA)
	in.DeepCopyInto(out)
	return out
}
