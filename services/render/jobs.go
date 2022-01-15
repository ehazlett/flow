package render

import (
	"context"

	api "git.underland.io/ehazlett/finca/api/services/render/v1"
	ptypes "github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (s *service) ListJobs(ctx context.Context, r *api.ListJobsRequest) (*api.ListJobsResponse, error) {
	jobs, err := s.ds.GetJobs(ctx)
	if err != nil {
		return nil, err
	}

	return &api.ListJobsResponse{
		Jobs: jobs,
	}, nil
}

func (s *service) GetJob(ctx context.Context, id string) (*api.Job, error) {
	job, err := s.ds.GetJob(ctx, id)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (s *service) DeleteJob(ctx context.Context, r *api.DeleteJobRequest) (*ptypes.Empty, error) {
	job, err := s.ds.GetJob(ctx, r.ID)
	if err != nil {
		return empty, errors.Wrapf(err, "error getting job %s from datastore", r.ID)
	}

	if err := s.ds.DeleteJob(ctx, r.ID); err != nil {
		return empty, errors.Wrapf(err, "error deleting job %s from datastore", r.ID)
	}

	js, err := s.natsClient.JetStream()
	if err != nil {
		return empty, err
	}

	// TODO: check nats for the job and delete
	if job.SequenceID != 0 {
		if err := js.DeleteMsg(s.config.NATSJobSubject, job.SequenceID); err != nil {
			return empty, err
		}
	}

	logrus.Infof("deleted job %s", r.ID)
	return empty, nil
}