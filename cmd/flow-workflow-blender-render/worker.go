// Copyright 2022 Evan Hazlett
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ehazlett/flow/cmd/flow-workflow-blender-render/processor"
	"github.com/ehazlett/flow/pkg/workflows"
	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

func workerAction(clix *cli.Context) error {
	id := clix.String("id")
	blenderPath := clix.String("blender-path")

	logrus.Infof("starting %s: id=%s", workflowName, id)

	cfg := &workflows.Config{
		ID:                    clix.String("id"),
		Type:                  workflowType,
		Address:               clix.String("address"),
		TLSCertificate:        clix.String("tls-certificate"),
		TLSKey:                clix.String("tls-key"),
		TLSInsecureSkipVerify: clix.Bool("tls-skip-verify"),
		MaxWorkflows:          clix.Uint64("max-workflows"),
		ServiceToken:          clix.String("service-token"),
		WorkflowTimeout:       clix.Duration("workflow-timeout"),
	}

	p, err := processor.NewProcessor(blenderPath, cfg)
	if err != nil {
		return err
	}

	h, err := workflows.NewWorkflowHandler(cfg, p)
	if err != nil {
		return err
	}

	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	doneCh := make(chan bool, 1)
	errCh := make(chan error, 1)

	var runErr error
	go func() {
		for {
			select {
			case err := <-errCh:
				runErr = err
				doneCh <- true
				return
			case sig := <-signals:
				switch sig {
				case syscall.SIGTERM, syscall.SIGINT:
					if err := h.Stop(); err != nil {
						logrus.Error(err)
					}
					doneCh <- true
					return
				default:
					logrus.Warnf("unhandled signal %s", sig)
				}
			}
		}
	}()

	go func() {
		ctx := context.Background()
		if err := h.Run(ctx); err != nil {
			errCh <- err
			return
		}
		doneCh <- true
	}()

	<-doneCh

	logrus.Info("shutting down")

	return runErr
}