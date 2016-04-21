/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015-2016 Intel Corporation

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
package controlproxy

import (
	"errors"
	"time"

	"golang.org/x/net/context"

	"github.com/intelsdi-x/snap/control/rpc"
	"github.com/intelsdi-x/snap/core"
	"github.com/intelsdi-x/snap/core/ctypes"
	"github.com/intelsdi-x/snap/core/serror"
	"github.com/intelsdi-x/snap/internal/common"
)

// Implements managesMetrics interface provided by scheduler and
// proxies those calls to the grpc client.
type ControlProxy struct {
	Client rpc.MetricManagerClient
}

func New(c rpc.MetricManagerClient) ControlProxy {
	return ControlProxy{Client: c}
}

func (c ControlProxy) ExpandWildcards(namespace []string) ([][]string, serror.SnapError) {
	req := &rpc.ExpandWildcardsRequest{
		Namespace: namespace,
	}
	reply, err := c.Client.ExpandWildcards(context.Background(), req)
	if err != nil {
		return nil, serror.New(err)
	}
	if reply.Error != nil {
		return nil, common.ToSnapError(reply.Error)
	}
	nss := ToNSS(reply.NSS)
	return nss, nil
}
func (c ControlProxy) PublishMetrics(contentType string, content []byte, pluginName string, pluginVersion int, config map[string]ctypes.ConfigValue, taskID string) []error {
	req := GetPubProcReq(contentType, content, pluginName, pluginVersion, config, taskID)
	reply, err := c.Client.PublishMetrics(context.Background(), req)
	errs := make([]error, 0)
	if err != nil {
		errs = append(errs, err)
		return errs
	}
	rerrs := ReplyErrorsToErrors(reply.Errors)
	errs = append(errs, rerrs...)
	return errs
}

func (c ControlProxy) ProcessMetrics(contentType string, content []byte, pluginName string, pluginVersion int, config map[string]ctypes.ConfigValue, taskID string) (string, []byte, []error) {
	req := GetPubProcReq(contentType, content, pluginName, pluginVersion, config, taskID)
	reply, err := c.Client.ProcessMetrics(context.Background(), req)
	errs := make([]error, 0)
	if err != nil {
		errs = append(errs, err)
		return "", nil, errs
	}
	rerrs := ReplyErrorsToErrors(reply.Errors)
	errs = append(errs, rerrs...)
	return reply.ContentType, reply.Content, errs
}

func (c ControlProxy) CollectMetrics(mts []core.Metric, deadline time.Time, taskID string) ([]core.Metric, []error) {
	req := &rpc.CollectMetricsRequest{
		Metrics: common.NewMetrics(mts),
		Deadline: &common.Time{
			Sec:  deadline.Unix(),
			Nsec: int64(deadline.Nanosecond()),
		},
		TaskID: taskID,
	}
	reply, err := c.Client.CollectMetrics(context.Background(), req)
	errs := make([]error, 0)
	if err != nil {
		errs = append(errs, err)
		return nil, errs
	}
	rerrs := ReplyErrorsToErrors(reply.Errors)
	if len(rerrs) > 0 {
		errs = append(errs, rerrs...)
		return nil, errs
	}
	metrics := common.ToCoreMetrics(reply.Metrics)
	return metrics, nil
}

func (c ControlProxy) GetPluginContentTypes(n string, t core.PluginType, v int) ([]string, []string, error) {
	req := &rpc.GetPluginContentTypesRequest{
		Name:       n,
		PluginType: getPluginType(t),
		Version:    int32(v),
	}
	reply, err := c.Client.GetPluginContentTypes(context.Background(), req)
	if err != nil {
		return nil, nil, err
	}
	if reply.Error != "" {
		return nil, nil, errors.New(reply.Error)
	}
	return reply.AcceptedTypes, reply.ReturnedTypes, nil
}

func (c ControlProxy) ValidateDeps(mts []core.Metric, plugins []core.SubscribedPlugin) []serror.SnapError {
	req := &rpc.ValidateDepsRequest{
		Metrics: common.NewMetrics(mts),
		Plugins: common.ToSubPluginsMsg(plugins),
	}
	reply, err := c.Client.ValidateDeps(context.Background(), req)
	if err != nil {
		return []serror.SnapError{serror.New(err)}
	}
	serrs := common.ConvertSnapErrors(reply.Errors)
	return serrs
}

func (c ControlProxy) SubscribeDeps(taskID string, mts []core.Metric, plugins []core.Plugin) []serror.SnapError {
	req := depsRequest(taskID, mts, plugins)
	reply, err := c.Client.SubscribeDeps(context.Background(), req)
	if err != nil {
		return []serror.SnapError{serror.New(err)}
	}
	serrs := common.ConvertSnapErrors(reply.Errors)
	return serrs
}

func (c ControlProxy) UnsubscribeDeps(taskID string, mts []core.Metric, plugins []core.Plugin) []serror.SnapError {
	req := depsRequest(taskID, mts, plugins)
	reply, err := c.Client.UnsubscribeDeps(context.Background(), req)
	if err != nil {
		return []serror.SnapError{serror.New(err)}
	}
	serrs := common.ConvertSnapErrors(reply.Errors)
	return serrs
}

func (c ControlProxy) MatchQueryToNamespaces(namespace []string) ([][]string, serror.SnapError) {
	req := &rpc.ExpandWildcardsRequest{
		Namespace: namespace,
	}
	reply, err := c.Client.MatchQueryToNamespaces(context.Background(), req)
	if err != nil {
		return nil, serror.New(err)
	}
	if reply.Error != nil {
		return nil, common.ToSnapError(reply.Error)
	}
	nss := ToNSS(reply.NSS)
	return nss, nil
}

///---------Util-------------------------------------------------------------------------
// TODO(CDR): move the functions used by outside callers only to a util package and make functions used only
// in this file non-exported.
func getPluginType(t core.PluginType) int32 {
	val := int32(-1)
	switch t {
	case core.CollectorPluginType:
		val = 0
	case core.ProcessorPluginType:
		val = 1
	case core.PublisherPluginType:
		val = 2
	}
	return val
}

func depsRequest(taskID string, mts []core.Metric, plugins []core.Plugin) *rpc.SubscribeDepsRequest {
	req := &rpc.SubscribeDepsRequest{
		Metrics: common.NewMetrics(mts),
		Plugins: common.ToCorePluginsMsg(plugins),
		TaskId:  taskID,
	}
	return req
}

func ConvertNSS(nss [][]string) []*rpc.ArrString {
	res := make([]*rpc.ArrString, len(nss))
	for i := range nss {
		var tmp rpc.ArrString
		tmp.S = nss[i]
		res[i] = &tmp
	}
	return res
}

func ToNSS(arr []*rpc.ArrString) [][]string {
	nss := make([][]string, len(arr))
	for i, v := range arr {
		nss[i] = v.S
	}
	return nss
}

func ErrorsToStrings(in []error) []string {
	if len(in) == 0 {
		return []string{}
	}
	erro := make([]string, len(in))
	for i, e := range in {
		erro[i] = e.Error()
	}
	return erro
}

func ReplyErrorsToErrors(errs []string) []error {
	if len(errs) == 0 {
		return []error{}
	}
	var erro []error
	for _, e := range errs {
		erro = append(erro, errors.New(e))
	}
	return erro
}

func GetPubProcReq(contentType string, content []byte, pluginName string, pluginVersion int, config map[string]ctypes.ConfigValue, taskID string) *rpc.PubProcMetricsRequest {
	newConfig := common.ToConfigMap(config)
	request := &rpc.PubProcMetricsRequest{
		ContentType:   contentType,
		Content:       content,
		PluginName:    pluginName,
		PluginVersion: int64(pluginVersion),
		Config:        newConfig,
		TaskId:        taskID,
	}
	return request
}
