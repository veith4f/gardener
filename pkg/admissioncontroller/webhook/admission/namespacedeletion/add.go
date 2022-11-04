// Copyright (c) 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package namespacedeletion

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
)

const (
	// HandlerName is the name of this admission webhook handler.
	HandlerName = "namespace_validator"
	// WebhookPath is the HTTP handler path for this admission webhook handler.
	WebhookPath = "/webhooks/validate-namespace-deletion"
)

// AddToManager adds Handler to the given manager.
func (h *Handler) AddToManager(ctx context.Context, mgr manager.Manager) error {
	// Initialize caches here to ensure the readyz informer check will only succeed once informers required for this
	// handler have synced so that http requests can be served quicker with pre-syncronized caches.
	if _, err := mgr.GetCache().GetInformer(ctx, &corev1.Namespace{}); err != nil {
		return err
	}
	if _, err := mgr.GetCache().GetInformer(ctx, &gardencorev1beta1.Project{}); err != nil {
		return err
	}

	webhook := admission.
		WithCustomValidator(&corev1.Namespace{}, h).
		WithRecoverPanic(true)

	mgr.GetWebhookServer().Register(WebhookPath, webhook)
	return nil
}
