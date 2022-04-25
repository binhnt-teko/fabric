/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package chaincode

import (
	"github.com/hyperledger/fabric/common/metrics"
	cmap "github.com/orcaman/concurrent-map"
)

var (
	launchDuration = metrics.HistogramOpts{
		Namespace:    "chaincode",
		Name:         "launch_duration",
		Help:         "The time to launch a chaincode.",
		LabelNames:   []string{"chaincode", "success"},
		StatsdFormat: "%{#fqname}.%{chaincode}.%{success}",
	}
	launchFailures = metrics.CounterOpts{
		Namespace:    "chaincode",
		Name:         "launch_failures",
		Help:         "The number of chaincode launches that have failed.",
		LabelNames:   []string{"chaincode"},
		StatsdFormat: "%{#fqname}.%{chaincode}",
	}
	launchTimeouts = metrics.CounterOpts{
		Namespace:    "chaincode",
		Name:         "launch_timeouts",
		Help:         "The number of chaincode launches that have timed out.",
		LabelNames:   []string{"chaincode"},
		StatsdFormat: "%{#fqname}.%{chaincode}",
	}

	shimRequestsReceived = metrics.CounterOpts{
		Namespace:    "chaincode",
		Name:         "shim_requests_received",
		Help:         "The number of chaincode shim requests received.",
		LabelNames:   []string{"type", "channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{type}.%{channel}.%{chaincode}",
	}
	shimRequestsCompleted = metrics.CounterOpts{
		Namespace:    "chaincode",
		Name:         "shim_requests_completed",
		Help:         "The number of chaincode shim requests completed.",
		LabelNames:   []string{"type", "channel", "chaincode", "success"},
		StatsdFormat: "%{#fqname}.%{type}.%{channel}.%{chaincode}.%{success}",
	}
	shimRequestDuration = metrics.HistogramOpts{
		Namespace:    "chaincode",
		Name:         "shim_request_duration",
		Help:         "The time to complete chaincode shim requests.",
		LabelNames:   []string{"type", "channel", "chaincode", "success"},
		StatsdFormat: "%{#fqname}.%{type}.%{channel}.%{chaincode}.%{success}",
	}
	executeTimeouts = metrics.CounterOpts{
		Namespace:    "chaincode",
		Name:         "execute_timeouts",
		Help:         "The number of chaincode executions (Init or Invoke) that have timed out.",
		LabelNames:   []string{"chaincode"},
		StatsdFormat: "%{#fqname}.%{chaincode}",
	}
	chaincodeInvokeDuration = metrics.HistogramOpts{
		Namespace:    "chaincode",
		Name:         "chaincode_invoke_duration",
		Help:         "The time to complete chaincode shim requests.",
		LabelNames:   []string{"channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}",
	}
	chaincodeCheckInvocation = metrics.HistogramOpts{
		Namespace:    "chaincode",
		Name:         "chaincode_check_invocation_duration",
		Help:         "The time to complete chaincode shim requests.",
		LabelNames:   []string{"channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}",
	}
	chaincodeLaunch = metrics.HistogramOpts{
		Namespace:    "chaincode",
		Name:         "chaincode_launch_duration",
		Help:         "The time to complete chaincode shim requests.",
		LabelNames:   []string{"channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}",
	}
	chaincodeExecute = metrics.HistogramOpts{
		Namespace:    "chaincode",
		Name:         "chaincode_execute_duration",
		Help:         "The time to complete chaincode shim requests.",
		LabelNames:   []string{"channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}",
	}
	chaincodeProposal = metrics.HistogramOpts{
		Namespace:    "chaincode",
		Name:         "chaincode_chaincode_proposal",
		Help:         "The time to complete chaincode shim requests.",
		LabelNames:   []string{"channel"},
		StatsdFormat: "%{#fqname}.%{channel}",
	}
	chaincodeProposalPrepare = metrics.HistogramOpts{
		Namespace:    "chaincode",
		Name:         "chaincode_chaincode_proposal_prepare",
		Help:         "The time to complete chaincode shim requests.",
		LabelNames:   []string{"channel"},
		StatsdFormat: "%{#fqname}.%{channel}",
	}
	chaincodeProposalSend = metrics.HistogramOpts{
		Namespace:    "chaincode",
		Name:         "chaincode_chaincode_proposal_send",
		Help:         "The time to complete chaincode shim requests.",
		LabelNames:   []string{"channel"},
		StatsdFormat: "%{#fqname}.%{channel}",
	}

	chaincodeProposalTransactionCount = metrics.CounterOpts{
		Namespace:    "chaincode",
		Name:         "transaction_count",
		Help:         "The number of chaincode launches that have timed out.",
		LabelNames:   []string{"channel"},
		StatsdFormat: "%{#fqname}.%{channel}",
	}

	chaincodeProposalTransactionDuration = metrics.HistogramOpts{
		Namespace:    "chaincode",
		Name:         "transaction_duration",
		Help:         "The time to complete chaincode shim requests.",
		LabelNames:   []string{"channel"},
		StatsdFormat: "%{#fqname}.%{channel}",
	}
	chaincodeProposalTransactionGetContext = metrics.HistogramOpts{
		Namespace:    "chaincode",
		Name:         "transaction_get_context",
		Help:         "The time to complete chaincode shim requests.",
		LabelNames:   []string{"channel"},
		StatsdFormat: "%{#fqname}.%{channel}",
	}
	chaincodeProposalTransactionClose = metrics.HistogramOpts{
		Namespace:    "chaincode",
		Name:         "transaction_close",
		Help:         "The time to complete chaincode shim requests.",
		LabelNames:   []string{"channel"},
		StatsdFormat: "%{#fqname}.%{channel}",
	}
	chaincodeProposalTransactionBeforeSend = metrics.HistogramOpts{
		Namespace:    "chaincode",
		Name:         "transaction_before_send",
		Help:         "The time to complete chaincode shim requests.",
		LabelNames:   []string{"channel"},
		StatsdFormat: "%{#fqname}.%{channel}",
	}
	chaincodeProposalTransactionSendTime = metrics.HistogramOpts{
		Namespace:    "chaincode",
		Name:         "transaction_send_time",
		Help:         "The time to complete chaincode shim requests.",
		LabelNames:   []string{"channel"},
		StatsdFormat: "%{#fqname}.%{channel}",
	}
)

type HandlerMetrics struct {
	ShimRequestsReceived  metrics.Counter
	ShimRequestsCompleted metrics.Counter
	ShimRequestDuration   metrics.Histogram
	ExecuteTimeouts       metrics.Counter
	// Luat add metrics
	ChaincodeInvokeDuration  metrics.Histogram
	ChaincodeCheckInvocation metrics.Histogram
	ChaincodeLaunch          metrics.Histogram
	ChaincodeExecute         metrics.Histogram
	ChaincodeProposal        metrics.Histogram
	ChaincodeProposalPrepare metrics.Histogram
	ChaincodeProposalSend    metrics.Histogram

	ChaincodeProposalTransactionMap           cmap.ConcurrentMap
	ChaincodeProposalTransactionBeforeSendMap cmap.ConcurrentMap

	ChaincodeProposalTransactionCount      metrics.Counter
	ChaincodeProposalTransactionDuration   metrics.Histogram
	ChaincodeProposalTransactionGetContext metrics.Histogram
	ChaincodeProposalTransactionClose      metrics.Histogram
	ChaincodeProposalTransactionBeforeSend metrics.Histogram
	ChaincodeProposalTransactionSendTime   metrics.Histogram
}

func NewHandlerMetrics(p metrics.Provider) *HandlerMetrics {
	return &HandlerMetrics{
		ShimRequestsReceived:  p.NewCounter(shimRequestsReceived),
		ShimRequestsCompleted: p.NewCounter(shimRequestsCompleted),
		ShimRequestDuration:   p.NewHistogram(shimRequestDuration),
		ExecuteTimeouts:       p.NewCounter(executeTimeouts),

		// Luat add metrics
		ChaincodeInvokeDuration:  p.NewHistogram(chaincodeInvokeDuration),
		ChaincodeCheckInvocation: p.NewHistogram(chaincodeCheckInvocation),
		ChaincodeLaunch:          p.NewHistogram(chaincodeLaunch),
		ChaincodeExecute:         p.NewHistogram(chaincodeExecute),
		ChaincodeProposal:        p.NewHistogram(chaincodeProposal),
		ChaincodeProposalPrepare: p.NewHistogram(chaincodeProposalPrepare),
		ChaincodeProposalSend:    p.NewHistogram(chaincodeProposalSend),

		ChaincodeProposalTransactionMap:           cmap.New(),
		ChaincodeProposalTransactionBeforeSendMap: cmap.New(),
		ChaincodeProposalTransactionCount:         p.NewCounter(chaincodeProposalTransactionCount),
		ChaincodeProposalTransactionDuration:      p.NewHistogram(chaincodeProposalTransactionDuration),
		ChaincodeProposalTransactionGetContext:    p.NewHistogram(chaincodeProposalTransactionGetContext),
		ChaincodeProposalTransactionClose:         p.NewHistogram(chaincodeProposalTransactionClose),
		ChaincodeProposalTransactionBeforeSend:    p.NewHistogram(chaincodeProposalTransactionBeforeSend),
		ChaincodeProposalTransactionSendTime:      p.NewHistogram(chaincodeProposalTransactionSendTime),
	}
}

type LaunchMetrics struct {
	LaunchDuration metrics.Histogram
	LaunchFailures metrics.Counter
	LaunchTimeouts metrics.Counter
}

func NewLaunchMetrics(p metrics.Provider) *LaunchMetrics {
	return &LaunchMetrics{
		LaunchDuration: p.NewHistogram(launchDuration),
		LaunchFailures: p.NewCounter(launchFailures),
		LaunchTimeouts: p.NewCounter(launchTimeouts),
	}
}
