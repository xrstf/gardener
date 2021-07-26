// +build !ignore_autogenerated

/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	config "github.com/gardener/gardener/pkg/scheduler/apis/config"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	configv1alpha1 "k8s.io/component-base/config/v1alpha1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*BackupBucketSchedulerConfiguration)(nil), (*config.BackupBucketSchedulerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_BackupBucketSchedulerConfiguration_To_config_BackupBucketSchedulerConfiguration(a.(*BackupBucketSchedulerConfiguration), b.(*config.BackupBucketSchedulerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.BackupBucketSchedulerConfiguration)(nil), (*BackupBucketSchedulerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_BackupBucketSchedulerConfiguration_To_v1alpha1_BackupBucketSchedulerConfiguration(a.(*config.BackupBucketSchedulerConfiguration), b.(*BackupBucketSchedulerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LeaderElectionConfiguration)(nil), (*config.LeaderElectionConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration(a.(*LeaderElectionConfiguration), b.(*config.LeaderElectionConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.LeaderElectionConfiguration)(nil), (*LeaderElectionConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(a.(*config.LeaderElectionConfiguration), b.(*LeaderElectionConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SchedulerConfiguration)(nil), (*config.SchedulerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_SchedulerConfiguration_To_config_SchedulerConfiguration(a.(*SchedulerConfiguration), b.(*config.SchedulerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.SchedulerConfiguration)(nil), (*SchedulerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_SchedulerConfiguration_To_v1alpha1_SchedulerConfiguration(a.(*config.SchedulerConfiguration), b.(*SchedulerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SchedulerControllerConfiguration)(nil), (*config.SchedulerControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_SchedulerControllerConfiguration_To_config_SchedulerControllerConfiguration(a.(*SchedulerControllerConfiguration), b.(*config.SchedulerControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.SchedulerControllerConfiguration)(nil), (*SchedulerControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_SchedulerControllerConfiguration_To_v1alpha1_SchedulerControllerConfiguration(a.(*config.SchedulerControllerConfiguration), b.(*SchedulerControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Server)(nil), (*config.Server)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Server_To_config_Server(a.(*Server), b.(*config.Server), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Server)(nil), (*Server)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Server_To_v1alpha1_Server(a.(*config.Server), b.(*Server), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServerConfiguration)(nil), (*config.ServerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration(a.(*ServerConfiguration), b.(*config.ServerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ServerConfiguration)(nil), (*ServerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration(a.(*config.ServerConfiguration), b.(*ServerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ShootSchedulerConfiguration)(nil), (*config.ShootSchedulerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ShootSchedulerConfiguration_To_config_ShootSchedulerConfiguration(a.(*ShootSchedulerConfiguration), b.(*config.ShootSchedulerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ShootSchedulerConfiguration)(nil), (*ShootSchedulerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ShootSchedulerConfiguration_To_v1alpha1_ShootSchedulerConfiguration(a.(*config.ShootSchedulerConfiguration), b.(*ShootSchedulerConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_BackupBucketSchedulerConfiguration_To_config_BackupBucketSchedulerConfiguration(in *BackupBucketSchedulerConfiguration, out *config.BackupBucketSchedulerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	return nil
}

// Convert_v1alpha1_BackupBucketSchedulerConfiguration_To_config_BackupBucketSchedulerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_BackupBucketSchedulerConfiguration_To_config_BackupBucketSchedulerConfiguration(in *BackupBucketSchedulerConfiguration, out *config.BackupBucketSchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_BackupBucketSchedulerConfiguration_To_config_BackupBucketSchedulerConfiguration(in, out, s)
}

func autoConvert_config_BackupBucketSchedulerConfiguration_To_v1alpha1_BackupBucketSchedulerConfiguration(in *config.BackupBucketSchedulerConfiguration, out *BackupBucketSchedulerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	return nil
}

// Convert_config_BackupBucketSchedulerConfiguration_To_v1alpha1_BackupBucketSchedulerConfiguration is an autogenerated conversion function.
func Convert_config_BackupBucketSchedulerConfiguration_To_v1alpha1_BackupBucketSchedulerConfiguration(in *config.BackupBucketSchedulerConfiguration, out *BackupBucketSchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_config_BackupBucketSchedulerConfiguration_To_v1alpha1_BackupBucketSchedulerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration(in *LeaderElectionConfiguration, out *config.LeaderElectionConfiguration, s conversion.Scope) error {
	if err := configv1alpha1.Convert_v1alpha1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration(&in.LeaderElectionConfiguration, &out.LeaderElectionConfiguration, s); err != nil {
		return err
	}
	out.LockObjectNamespace = in.LockObjectNamespace
	out.LockObjectName = in.LockObjectName
	return nil
}

// Convert_v1alpha1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration(in *LeaderElectionConfiguration, out *config.LeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration(in, out, s)
}

func autoConvert_config_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in *config.LeaderElectionConfiguration, out *LeaderElectionConfiguration, s conversion.Scope) error {
	if err := configv1alpha1.Convert_config_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElectionConfiguration, &out.LeaderElectionConfiguration, s); err != nil {
		return err
	}
	out.LockObjectNamespace = in.LockObjectNamespace
	out.LockObjectName = in.LockObjectName
	return nil
}

// Convert_config_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration is an autogenerated conversion function.
func Convert_config_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in *config.LeaderElectionConfiguration, out *LeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_config_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in, out, s)
}

func autoConvert_v1alpha1_SchedulerConfiguration_To_config_SchedulerConfiguration(in *SchedulerConfiguration, out *config.SchedulerConfiguration, s conversion.Scope) error {
	if err := configv1alpha1.Convert_v1alpha1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	out.LogLevel = in.LogLevel
	out.LogFormat = in.LogFormat
	if err := Convert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration(&in.Server, &out.Server, s); err != nil {
		return err
	}
	out.HealthServer = (*config.ServerConfiguration)(unsafe.Pointer(in.HealthServer))
	out.MetricsServer = (*config.ServerConfiguration)(unsafe.Pointer(in.MetricsServer))
	if err := Convert_v1alpha1_SchedulerControllerConfiguration_To_config_SchedulerControllerConfiguration(&in.Schedulers, &out.Schedulers, s); err != nil {
		return err
	}
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

// Convert_v1alpha1_SchedulerConfiguration_To_config_SchedulerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_SchedulerConfiguration_To_config_SchedulerConfiguration(in *SchedulerConfiguration, out *config.SchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_SchedulerConfiguration_To_config_SchedulerConfiguration(in, out, s)
}

func autoConvert_config_SchedulerConfiguration_To_v1alpha1_SchedulerConfiguration(in *config.SchedulerConfiguration, out *SchedulerConfiguration, s conversion.Scope) error {
	if err := configv1alpha1.Convert_config_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	if err := Convert_config_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	out.LogLevel = in.LogLevel
	out.LogFormat = in.LogFormat
	if err := Convert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration(&in.Server, &out.Server, s); err != nil {
		return err
	}
	out.HealthServer = (*ServerConfiguration)(unsafe.Pointer(in.HealthServer))
	out.MetricsServer = (*ServerConfiguration)(unsafe.Pointer(in.MetricsServer))
	if err := Convert_config_SchedulerControllerConfiguration_To_v1alpha1_SchedulerControllerConfiguration(&in.Schedulers, &out.Schedulers, s); err != nil {
		return err
	}
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

// Convert_config_SchedulerConfiguration_To_v1alpha1_SchedulerConfiguration is an autogenerated conversion function.
func Convert_config_SchedulerConfiguration_To_v1alpha1_SchedulerConfiguration(in *config.SchedulerConfiguration, out *SchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_config_SchedulerConfiguration_To_v1alpha1_SchedulerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_SchedulerControllerConfiguration_To_config_SchedulerControllerConfiguration(in *SchedulerControllerConfiguration, out *config.SchedulerControllerConfiguration, s conversion.Scope) error {
	out.BackupBucket = (*config.BackupBucketSchedulerConfiguration)(unsafe.Pointer(in.BackupBucket))
	out.Shoot = (*config.ShootSchedulerConfiguration)(unsafe.Pointer(in.Shoot))
	return nil
}

// Convert_v1alpha1_SchedulerControllerConfiguration_To_config_SchedulerControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_SchedulerControllerConfiguration_To_config_SchedulerControllerConfiguration(in *SchedulerControllerConfiguration, out *config.SchedulerControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_SchedulerControllerConfiguration_To_config_SchedulerControllerConfiguration(in, out, s)
}

func autoConvert_config_SchedulerControllerConfiguration_To_v1alpha1_SchedulerControllerConfiguration(in *config.SchedulerControllerConfiguration, out *SchedulerControllerConfiguration, s conversion.Scope) error {
	out.BackupBucket = (*BackupBucketSchedulerConfiguration)(unsafe.Pointer(in.BackupBucket))
	out.Shoot = (*ShootSchedulerConfiguration)(unsafe.Pointer(in.Shoot))
	return nil
}

// Convert_config_SchedulerControllerConfiguration_To_v1alpha1_SchedulerControllerConfiguration is an autogenerated conversion function.
func Convert_config_SchedulerControllerConfiguration_To_v1alpha1_SchedulerControllerConfiguration(in *config.SchedulerControllerConfiguration, out *SchedulerControllerConfiguration, s conversion.Scope) error {
	return autoConvert_config_SchedulerControllerConfiguration_To_v1alpha1_SchedulerControllerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_Server_To_config_Server(in *Server, out *config.Server, s conversion.Scope) error {
	out.BindAddress = in.BindAddress
	out.Port = in.Port
	return nil
}

// Convert_v1alpha1_Server_To_config_Server is an autogenerated conversion function.
func Convert_v1alpha1_Server_To_config_Server(in *Server, out *config.Server, s conversion.Scope) error {
	return autoConvert_v1alpha1_Server_To_config_Server(in, out, s)
}

func autoConvert_config_Server_To_v1alpha1_Server(in *config.Server, out *Server, s conversion.Scope) error {
	out.BindAddress = in.BindAddress
	out.Port = in.Port
	return nil
}

// Convert_config_Server_To_v1alpha1_Server is an autogenerated conversion function.
func Convert_config_Server_To_v1alpha1_Server(in *config.Server, out *Server, s conversion.Scope) error {
	return autoConvert_config_Server_To_v1alpha1_Server(in, out, s)
}

func autoConvert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration(in *ServerConfiguration, out *config.ServerConfiguration, s conversion.Scope) error {
	if err := Convert_v1alpha1_Server_To_config_Server(&in.HTTP, &out.HTTP, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration(in *ServerConfiguration, out *config.ServerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration(in, out, s)
}

func autoConvert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration(in *config.ServerConfiguration, out *ServerConfiguration, s conversion.Scope) error {
	if err := Convert_config_Server_To_v1alpha1_Server(&in.HTTP, &out.HTTP, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration is an autogenerated conversion function.
func Convert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration(in *config.ServerConfiguration, out *ServerConfiguration, s conversion.Scope) error {
	return autoConvert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ShootSchedulerConfiguration_To_config_ShootSchedulerConfiguration(in *ShootSchedulerConfiguration, out *config.ShootSchedulerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	out.Strategy = config.CandidateDeterminationStrategy(in.Strategy)
	return nil
}

// Convert_v1alpha1_ShootSchedulerConfiguration_To_config_ShootSchedulerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ShootSchedulerConfiguration_To_config_ShootSchedulerConfiguration(in *ShootSchedulerConfiguration, out *config.ShootSchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ShootSchedulerConfiguration_To_config_ShootSchedulerConfiguration(in, out, s)
}

func autoConvert_config_ShootSchedulerConfiguration_To_v1alpha1_ShootSchedulerConfiguration(in *config.ShootSchedulerConfiguration, out *ShootSchedulerConfiguration, s conversion.Scope) error {
	out.ConcurrentSyncs = in.ConcurrentSyncs
	out.Strategy = CandidateDeterminationStrategy(in.Strategy)
	return nil
}

// Convert_config_ShootSchedulerConfiguration_To_v1alpha1_ShootSchedulerConfiguration is an autogenerated conversion function.
func Convert_config_ShootSchedulerConfiguration_To_v1alpha1_ShootSchedulerConfiguration(in *config.ShootSchedulerConfiguration, out *ShootSchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_config_ShootSchedulerConfiguration_To_v1alpha1_ShootSchedulerConfiguration(in, out, s)
}
