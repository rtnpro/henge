// +build !ignore_autogenerated

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	route_api "github.com/openshift/origin/pkg/route/api"
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func init() {
	if err := api.Scheme.AddGeneratedConversionFuncs(
		Convert_v1_Route_To_api_Route,
		Convert_api_Route_To_v1_Route,
		Convert_v1_RouteIngress_To_api_RouteIngress,
		Convert_api_RouteIngress_To_v1_RouteIngress,
		Convert_v1_RouteIngressCondition_To_api_RouteIngressCondition,
		Convert_api_RouteIngressCondition_To_v1_RouteIngressCondition,
		Convert_v1_RouteList_To_api_RouteList,
		Convert_api_RouteList_To_v1_RouteList,
		Convert_v1_RoutePort_To_api_RoutePort,
		Convert_api_RoutePort_To_v1_RoutePort,
		Convert_v1_RouteSpec_To_api_RouteSpec,
		Convert_api_RouteSpec_To_v1_RouteSpec,
		Convert_v1_RouteStatus_To_api_RouteStatus,
		Convert_api_RouteStatus_To_v1_RouteStatus,
		Convert_v1_RouteTargetReference_To_api_RouteTargetReference,
		Convert_api_RouteTargetReference_To_v1_RouteTargetReference,
		Convert_v1_RouterShard_To_api_RouterShard,
		Convert_api_RouterShard_To_v1_RouterShard,
		Convert_v1_TLSConfig_To_api_TLSConfig,
		Convert_api_TLSConfig_To_v1_TLSConfig,
	); err != nil {
		// if one of the conversion functions is malformed, detect it immediately.
		panic(err)
	}
}

func autoConvert_v1_Route_To_api_Route(in *Route, out *route_api.Route, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_v1_RouteSpec_To_api_RouteSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_RouteStatus_To_api_RouteStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_Route_To_api_Route(in *Route, out *route_api.Route, s conversion.Scope) error {
	return autoConvert_v1_Route_To_api_Route(in, out, s)
}

func autoConvert_api_Route_To_v1_Route(in *route_api.Route, out *Route, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_api_RouteSpec_To_v1_RouteSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_api_RouteStatus_To_v1_RouteStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_Route_To_v1_Route(in *route_api.Route, out *Route, s conversion.Scope) error {
	return autoConvert_api_Route_To_v1_Route(in, out, s)
}

func autoConvert_v1_RouteIngress_To_api_RouteIngress(in *RouteIngress, out *route_api.RouteIngress, s conversion.Scope) error {
	out.Host = in.Host
	out.RouterName = in.RouterName
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]route_api.RouteIngressCondition, len(*in))
		for i := range *in {
			if err := Convert_v1_RouteIngressCondition_To_api_RouteIngressCondition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	return nil
}

func Convert_v1_RouteIngress_To_api_RouteIngress(in *RouteIngress, out *route_api.RouteIngress, s conversion.Scope) error {
	return autoConvert_v1_RouteIngress_To_api_RouteIngress(in, out, s)
}

func autoConvert_api_RouteIngress_To_v1_RouteIngress(in *route_api.RouteIngress, out *RouteIngress, s conversion.Scope) error {
	out.Host = in.Host
	out.RouterName = in.RouterName
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]RouteIngressCondition, len(*in))
		for i := range *in {
			if err := Convert_api_RouteIngressCondition_To_v1_RouteIngressCondition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	return nil
}

func Convert_api_RouteIngress_To_v1_RouteIngress(in *route_api.RouteIngress, out *RouteIngress, s conversion.Scope) error {
	return autoConvert_api_RouteIngress_To_v1_RouteIngress(in, out, s)
}

func autoConvert_v1_RouteIngressCondition_To_api_RouteIngressCondition(in *RouteIngressCondition, out *route_api.RouteIngressCondition, s conversion.Scope) error {
	out.Type = route_api.RouteIngressConditionType(in.Type)
	out.Status = api.ConditionStatus(in.Status)
	out.Reason = in.Reason
	out.Message = in.Message
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = new(unversioned.Time)
		if err := api.Convert_unversioned_Time_To_unversioned_Time(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.LastTransitionTime = nil
	}
	return nil
}

func Convert_v1_RouteIngressCondition_To_api_RouteIngressCondition(in *RouteIngressCondition, out *route_api.RouteIngressCondition, s conversion.Scope) error {
	return autoConvert_v1_RouteIngressCondition_To_api_RouteIngressCondition(in, out, s)
}

func autoConvert_api_RouteIngressCondition_To_v1_RouteIngressCondition(in *route_api.RouteIngressCondition, out *RouteIngressCondition, s conversion.Scope) error {
	out.Type = RouteIngressConditionType(in.Type)
	out.Status = api_v1.ConditionStatus(in.Status)
	out.Reason = in.Reason
	out.Message = in.Message
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = new(unversioned.Time)
		if err := api.Convert_unversioned_Time_To_unversioned_Time(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.LastTransitionTime = nil
	}
	return nil
}

func Convert_api_RouteIngressCondition_To_v1_RouteIngressCondition(in *route_api.RouteIngressCondition, out *RouteIngressCondition, s conversion.Scope) error {
	return autoConvert_api_RouteIngressCondition_To_v1_RouteIngressCondition(in, out, s)
}

func autoConvert_v1_RouteList_To_api_RouteList(in *RouteList, out *route_api.RouteList, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]route_api.Route, len(*in))
		for i := range *in {
			if err := Convert_v1_Route_To_api_Route(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_RouteList_To_api_RouteList(in *RouteList, out *route_api.RouteList, s conversion.Scope) error {
	return autoConvert_v1_RouteList_To_api_RouteList(in, out, s)
}

func autoConvert_api_RouteList_To_v1_RouteList(in *route_api.RouteList, out *RouteList, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Route, len(*in))
		for i := range *in {
			if err := Convert_api_Route_To_v1_Route(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_api_RouteList_To_v1_RouteList(in *route_api.RouteList, out *RouteList, s conversion.Scope) error {
	return autoConvert_api_RouteList_To_v1_RouteList(in, out, s)
}

func autoConvert_v1_RoutePort_To_api_RoutePort(in *RoutePort, out *route_api.RoutePort, s conversion.Scope) error {
	if err := api.Convert_intstr_IntOrString_To_intstr_IntOrString(&in.TargetPort, &out.TargetPort, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_RoutePort_To_api_RoutePort(in *RoutePort, out *route_api.RoutePort, s conversion.Scope) error {
	return autoConvert_v1_RoutePort_To_api_RoutePort(in, out, s)
}

func autoConvert_api_RoutePort_To_v1_RoutePort(in *route_api.RoutePort, out *RoutePort, s conversion.Scope) error {
	if err := api.Convert_intstr_IntOrString_To_intstr_IntOrString(&in.TargetPort, &out.TargetPort, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_RoutePort_To_v1_RoutePort(in *route_api.RoutePort, out *RoutePort, s conversion.Scope) error {
	return autoConvert_api_RoutePort_To_v1_RoutePort(in, out, s)
}

func autoConvert_v1_RouteSpec_To_api_RouteSpec(in *RouteSpec, out *route_api.RouteSpec, s conversion.Scope) error {
	out.Host = in.Host
	out.Path = in.Path
	if err := Convert_v1_RouteTargetReference_To_api_RouteTargetReference(&in.To, &out.To, s); err != nil {
		return err
	}
	if in.AlternateBackends != nil {
		in, out := &in.AlternateBackends, &out.AlternateBackends
		*out = make([]route_api.RouteTargetReference, len(*in))
		for i := range *in {
			if err := Convert_v1_RouteTargetReference_To_api_RouteTargetReference(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.AlternateBackends = nil
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(route_api.RoutePort)
		if err := Convert_v1_RoutePort_To_api_RoutePort(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Port = nil
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(route_api.TLSConfig)
		if err := Convert_v1_TLSConfig_To_api_TLSConfig(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.TLS = nil
	}
	return nil
}

func Convert_v1_RouteSpec_To_api_RouteSpec(in *RouteSpec, out *route_api.RouteSpec, s conversion.Scope) error {
	return autoConvert_v1_RouteSpec_To_api_RouteSpec(in, out, s)
}

func autoConvert_api_RouteSpec_To_v1_RouteSpec(in *route_api.RouteSpec, out *RouteSpec, s conversion.Scope) error {
	out.Host = in.Host
	out.Path = in.Path
	if err := Convert_api_RouteTargetReference_To_v1_RouteTargetReference(&in.To, &out.To, s); err != nil {
		return err
	}
	if in.AlternateBackends != nil {
		in, out := &in.AlternateBackends, &out.AlternateBackends
		*out = make([]RouteTargetReference, len(*in))
		for i := range *in {
			if err := Convert_api_RouteTargetReference_To_v1_RouteTargetReference(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.AlternateBackends = nil
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(RoutePort)
		if err := Convert_api_RoutePort_To_v1_RoutePort(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Port = nil
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		if err := Convert_api_TLSConfig_To_v1_TLSConfig(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.TLS = nil
	}
	return nil
}

func Convert_api_RouteSpec_To_v1_RouteSpec(in *route_api.RouteSpec, out *RouteSpec, s conversion.Scope) error {
	return autoConvert_api_RouteSpec_To_v1_RouteSpec(in, out, s)
}

func autoConvert_v1_RouteStatus_To_api_RouteStatus(in *RouteStatus, out *route_api.RouteStatus, s conversion.Scope) error {
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]route_api.RouteIngress, len(*in))
		for i := range *in {
			if err := Convert_v1_RouteIngress_To_api_RouteIngress(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Ingress = nil
	}
	return nil
}

func Convert_v1_RouteStatus_To_api_RouteStatus(in *RouteStatus, out *route_api.RouteStatus, s conversion.Scope) error {
	return autoConvert_v1_RouteStatus_To_api_RouteStatus(in, out, s)
}

func autoConvert_api_RouteStatus_To_v1_RouteStatus(in *route_api.RouteStatus, out *RouteStatus, s conversion.Scope) error {
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]RouteIngress, len(*in))
		for i := range *in {
			if err := Convert_api_RouteIngress_To_v1_RouteIngress(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Ingress = nil
	}
	return nil
}

func Convert_api_RouteStatus_To_v1_RouteStatus(in *route_api.RouteStatus, out *RouteStatus, s conversion.Scope) error {
	return autoConvert_api_RouteStatus_To_v1_RouteStatus(in, out, s)
}

func autoConvert_v1_RouteTargetReference_To_api_RouteTargetReference(in *RouteTargetReference, out *route_api.RouteTargetReference, s conversion.Scope) error {
	SetDefaults_RouteTargetReference(in)
	out.Kind = in.Kind
	out.Name = in.Name
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int32)
		**out = **in
	} else {
		out.Weight = nil
	}
	return nil
}

func Convert_v1_RouteTargetReference_To_api_RouteTargetReference(in *RouteTargetReference, out *route_api.RouteTargetReference, s conversion.Scope) error {
	return autoConvert_v1_RouteTargetReference_To_api_RouteTargetReference(in, out, s)
}

func autoConvert_api_RouteTargetReference_To_v1_RouteTargetReference(in *route_api.RouteTargetReference, out *RouteTargetReference, s conversion.Scope) error {
	out.Kind = in.Kind
	out.Name = in.Name
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int32)
		**out = **in
	} else {
		out.Weight = nil
	}
	return nil
}

func Convert_api_RouteTargetReference_To_v1_RouteTargetReference(in *route_api.RouteTargetReference, out *RouteTargetReference, s conversion.Scope) error {
	return autoConvert_api_RouteTargetReference_To_v1_RouteTargetReference(in, out, s)
}

func autoConvert_v1_RouterShard_To_api_RouterShard(in *RouterShard, out *route_api.RouterShard, s conversion.Scope) error {
	out.ShardName = in.ShardName
	out.DNSSuffix = in.DNSSuffix
	return nil
}

func Convert_v1_RouterShard_To_api_RouterShard(in *RouterShard, out *route_api.RouterShard, s conversion.Scope) error {
	return autoConvert_v1_RouterShard_To_api_RouterShard(in, out, s)
}

func autoConvert_api_RouterShard_To_v1_RouterShard(in *route_api.RouterShard, out *RouterShard, s conversion.Scope) error {
	out.ShardName = in.ShardName
	out.DNSSuffix = in.DNSSuffix
	return nil
}

func Convert_api_RouterShard_To_v1_RouterShard(in *route_api.RouterShard, out *RouterShard, s conversion.Scope) error {
	return autoConvert_api_RouterShard_To_v1_RouterShard(in, out, s)
}

func autoConvert_v1_TLSConfig_To_api_TLSConfig(in *TLSConfig, out *route_api.TLSConfig, s conversion.Scope) error {
	SetDefaults_TLSConfig(in)
	out.Termination = route_api.TLSTerminationType(in.Termination)
	out.Certificate = in.Certificate
	out.Key = in.Key
	out.CACertificate = in.CACertificate
	out.DestinationCACertificate = in.DestinationCACertificate
	out.InsecureEdgeTerminationPolicy = route_api.InsecureEdgeTerminationPolicyType(in.InsecureEdgeTerminationPolicy)
	return nil
}

func Convert_v1_TLSConfig_To_api_TLSConfig(in *TLSConfig, out *route_api.TLSConfig, s conversion.Scope) error {
	return autoConvert_v1_TLSConfig_To_api_TLSConfig(in, out, s)
}

func autoConvert_api_TLSConfig_To_v1_TLSConfig(in *route_api.TLSConfig, out *TLSConfig, s conversion.Scope) error {
	out.Termination = TLSTerminationType(in.Termination)
	out.Certificate = in.Certificate
	out.Key = in.Key
	out.CACertificate = in.CACertificate
	out.DestinationCACertificate = in.DestinationCACertificate
	out.InsecureEdgeTerminationPolicy = InsecureEdgeTerminationPolicyType(in.InsecureEdgeTerminationPolicy)
	return nil
}

func Convert_api_TLSConfig_To_v1_TLSConfig(in *route_api.TLSConfig, out *TLSConfig, s conversion.Scope) error {
	return autoConvert_api_TLSConfig_To_v1_TLSConfig(in, out, s)
}
