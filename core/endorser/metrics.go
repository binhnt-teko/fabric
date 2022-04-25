/*
Copyright State Street Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package endorser

import "github.com/hyperledger/fabric/common/metrics"

var (
	proposalDurationHistogramOpts = metrics.HistogramOpts{
		Namespace:    "endorser",
		Name:         "proposal_duration",
		Help:         "The time to complete a proposal.",
		LabelNames:   []string{"channel", "chaincode", "success"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}.%{success}",
	}

	receivedProposalsCounterOpts = metrics.CounterOpts{
		Namespace: "endorser",
		Name:      "proposals_received",
		Help:      "The number of proposals received.",
	}

	successfulProposalsCounterOpts = metrics.CounterOpts{
		Namespace: "endorser",
		Name:      "successful_proposals",
		Help:      "The number of successful proposals.",
	}

	proposalValidationFailureCounterOpts = metrics.CounterOpts{
		Namespace: "endorser",
		Name:      "proposal_validation_failures",
		Help:      "The number of proposals that have failed initial validation.",
	}

	proposalChannelACLFailureOpts = metrics.CounterOpts{
		Namespace:    "endorser",
		Name:         "proposal_acl_failures",
		Help:         "The number of proposals that failed ACL checks.",
		LabelNames:   []string{"channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}",
	}

	initFailureCounterOpts = metrics.CounterOpts{
		Namespace:    "endorser",
		Name:         "chaincode_instantiation_failures",
		Help:         "The number of chaincode instantiations or upgrade that have failed.",
		LabelNames:   []string{"channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}",
	}

	endorsementFailureCounterOpts = metrics.CounterOpts{
		Namespace:    "endorser",
		Name:         "endorsement_failures",
		Help:         "The number of failed endorsements.",
		LabelNames:   []string{"channel", "chaincode", "chaincodeerror"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}.%{chaincodeerror}",
	}

	duplicateTxsFailureCounterOpts = metrics.CounterOpts{
		Namespace:    "endorser",
		Name:         "duplicate_transaction_failures",
		Help:         "The number of failed proposals due to duplicate transaction ID.",
		LabelNames:   []string{"channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}",
	}

	simulationFailureCounterOpts = metrics.CounterOpts{
		Namespace:    "endorser",
		Name:         "proposal_simulation_failures",
		Help:         "The number of failed proposal simulations",
		LabelNames:   []string{"channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}",
	}

	// Luat metrics for process proposal successfully or error func: ProcessProposalSuccessfullyOrError
	processProposalSuccessfullyOrErrorHistogramOpts = metrics.HistogramOpts{
		Namespace:    "endorser",
		Name:         "process_proposal",
		Help:         "The time to complete a process proposal.",
		LabelNames:   []string{"channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}",
	}

	processProposalAcquireTxSimulator = metrics.HistogramOpts{
		Namespace:    "endorser",
		Name:         "process_proposal_acquire_tx_simulator",
		Help:         "The time to complete a process proposal.",
		LabelNames:   []string{"channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}",
	}

	processProposalEndoresementInfo = metrics.HistogramOpts{
		Namespace:    "endorser",
		Name:         "process_proposal_endoresement_info",
		Help:         "The time to complete a process proposal.",
		LabelNames:   []string{"channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}",
	}

	processProposalSimulateProposal = metrics.HistogramOpts{
		Namespace:    "endorser",
		Name:         "process_proposal_simulate_proposal",
		Help:         "The time to complete a process proposal.",
		LabelNames:   []string{"channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}",
	}

	processProposalResponsePayload = metrics.HistogramOpts{
		Namespace:    "endorser",
		Name:         "process_proposal_response_payload",
		Help:         "The time to complete a process proposal.",
		LabelNames:   []string{"channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}",
	}

	processProposalSimulateProposalCallChaincode = metrics.HistogramOpts{
		Namespace:    "endorser",
		Name:         "process_proposal_call_chaincode",
		Help:         "The time to complete a process proposal.",
		LabelNames:   []string{"channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}",
	}

	processProposalSimulateProposalCheckResult = metrics.HistogramOpts{
		Namespace:    "endorser",
		Name:         "process_proposal_check_result",
		Help:         "The time to complete a process proposal.",
		LabelNames:   []string{"channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}",
	}

	processProposalSimulateProposalReturnResult = metrics.HistogramOpts{
		Namespace:    "endorser",
		Name:         "process_proposal_return_result",
		Help:         "The time to complete a process proposal.",
		LabelNames:   []string{"channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}",
	}

	processProposalExecuteChaincode = metrics.HistogramOpts{
		Namespace:    "endorser",
		Name:         "process_proposal_execute_chaincode",
		Help:         "The time to complete a process proposal.",
		LabelNames:   []string{"channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}",
	}
)

type Metrics struct {
	ProposalDuration         metrics.Histogram
	ProposalsReceived        metrics.Counter
	SuccessfulProposals      metrics.Counter
	ProposalValidationFailed metrics.Counter
	ProposalACLCheckFailed   metrics.Counter
	InitFailed               metrics.Counter
	EndorsementsFailed       metrics.Counter
	DuplicateTxsFailure      metrics.Counter
	SimulationFailure        metrics.Counter

	// Luat : add metrics
	ProcessProposal                   metrics.Histogram
	ProcessProposalAcquireTxSimulator metrics.Histogram
	ProcessProposalEndoresementInfo   metrics.Histogram
	ProcessProposalSimulateProposal   metrics.Histogram
	ProcessProposalResponsePayload    metrics.Histogram

	ProcessProposalSimulateProposalCallChaincode metrics.Histogram
	ProcessProposalSimulateProposalCheckResult   metrics.Histogram
	ProcessProposalSimulateProposalReturnResult  metrics.Histogram
	ProcessProposalExecuteChaincode              metrics.Histogram
}

func NewMetrics(p metrics.Provider) *Metrics {
	return &Metrics{
		ProposalDuration:         p.NewHistogram(proposalDurationHistogramOpts),
		ProposalsReceived:        p.NewCounter(receivedProposalsCounterOpts),
		SuccessfulProposals:      p.NewCounter(successfulProposalsCounterOpts),
		ProposalValidationFailed: p.NewCounter(proposalValidationFailureCounterOpts),
		ProposalACLCheckFailed:   p.NewCounter(proposalChannelACLFailureOpts),
		InitFailed:               p.NewCounter(initFailureCounterOpts),
		EndorsementsFailed:       p.NewCounter(endorsementFailureCounterOpts),
		DuplicateTxsFailure:      p.NewCounter(duplicateTxsFailureCounterOpts),
		SimulationFailure:        p.NewCounter(simulationFailureCounterOpts),

		// Luat : add metrics
		ProcessProposal:                   p.NewHistogram(processProposalSuccessfullyOrErrorHistogramOpts),
		ProcessProposalAcquireTxSimulator: p.NewHistogram(processProposalAcquireTxSimulator),
		ProcessProposalEndoresementInfo:   p.NewHistogram(processProposalEndoresementInfo),
		ProcessProposalSimulateProposal:   p.NewHistogram(processProposalSimulateProposal),
		ProcessProposalResponsePayload:    p.NewHistogram(processProposalResponsePayload),

		ProcessProposalSimulateProposalCallChaincode: p.NewHistogram(processProposalSimulateProposalCallChaincode),
		ProcessProposalSimulateProposalCheckResult:   p.NewHistogram(processProposalSimulateProposalCheckResult),
		ProcessProposalSimulateProposalReturnResult:  p.NewHistogram(processProposalSimulateProposalReturnResult),
		ProcessProposalExecuteChaincode:              p.NewHistogram(processProposalExecuteChaincode),
	}
}
