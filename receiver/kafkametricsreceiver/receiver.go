// Copyright  The OpenTelemetry Authors
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

package kafkametricsreceiver

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

var newMetricsReceiver = func(
	ctx context.Context,
	config Config,
	params component.ReceiverCreateParams,
	consumer consumer.MetricsConsumer,
) (component.MetricsReceiver, error) {
	scraperControllerOptions := make([]scraperhelper.ScraperControllerOption, 0, len(config.Scrapers))
	return scraperhelper.NewScraperControllerReceiver(
		&config.ScraperControllerSettings,
		params.Logger,
		consumer,
		scraperControllerOptions...,
	)
}
