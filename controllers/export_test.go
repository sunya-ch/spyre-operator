/*
 * +-------------------------------------------------------------------+
 * | Copyright (c) 2025, 2026 IBM Corp.                                |
 * | SPDX-License-Identifier: Apache-2.0                               |
 * +-------------------------------------------------------------------+
 */

package controllers

import (
	"context"

	spyrev1alpha1 "github.com/ibm-aiu/spyre-operator/api/v1alpha1"
	"github.com/ibm-aiu/spyre-operator/internal/state"
	"go.uber.org/zap/zapcore"
	"k8s.io/apimachinery/pkg/types"
)

var ProcessOverallStatus = processOverallStatus
var NodeUpdateNeedsReconcile = nodeUpdateNeedsReconcile

func (rec *SpyreClusterPolicyReconciler) SetStateController(stateController *state.StateController) {
	rec.stateController = stateController
}

func (rec *SpyreClusterPolicyReconciler) ApplyLogLevel(logLevel zapcore.Level) {
	rec.stateController.SetLogLevel(logLevel)
}

func (rec *SpyreClusterPolicyReconciler) UpdateCRState(ctx context.Context, namespacedName types.NamespacedName, state spyrev1alpha1.State, message string) error {
	return rec.updateCRState(ctx, namespacedName, state, message)
}
