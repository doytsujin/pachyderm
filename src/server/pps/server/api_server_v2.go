package server

//import (
//	"context"
//	"path"
//	"time"
//
//	"github.com/gogo/protobuf/types"
//	"github.com/pachyderm/pachyderm/src/client/pfs"
//	"github.com/pachyderm/pachyderm/src/client/pkg/errors"
//	"github.com/pachyderm/pachyderm/src/client/pps"
//	"github.com/pachyderm/pachyderm/src/server/worker/common"
//	"github.com/pachyderm/pachyderm/src/server/worker/datum"
//)
//
//var _ pps.APIServer = &apiServerV2{}
//
//type apiServerV2 struct {
//	*apiServer
//}
//
//func newAPIServerV2(embeddedServer *apiServer) *apiServerV2 {
//	return &apiServerV2{apiServer: embeddedServer}
//}
//
//func (a *apiServerV2) InspectDatum(ctx context.Context, request *pps.InspectDatumRequest) (response *pps.DatumInfo, retErr error) {
//	func() { a.Log(request, nil, nil, 0) }()
//	defer func(start time.Time) { a.Log(request, response, retErr, time.Since(start)) }(time.Now())
//	if err := a.collectDatums(ctx, request.Datum.Job, func(meta *datum.Meta, pfsState *pfs.File) error {
//		if common.DatumIDV2(meta.Inputs) == request.Datum.ID {
//			response = convertDatumMetaToInfo(meta)
//			response.PfsState = pfsState
//		}
//		return nil
//	}); err != nil {
//		return nil, err
//	}
//	return response, nil
//}
//
//func (a *apiServerV2) ListDatumStream(request *pps.ListDatumRequest, server pps.API_ListDatumStreamServer) (retErr error) {
//	func() { a.Log(request, nil, nil, 0) }()
//	defer func(start time.Time) { a.Log(request, nil, retErr, time.Since(start)) }(time.Now())
//	return a.collectDatums(server.Context(), request.Job, func(meta *datum.Meta, _ *pfs.File) error {
//		return server.Send(&pps.ListDatumStreamResponse{
//			DatumInfo: convertDatumMetaToInfo(meta),
//		})
//	})
//}
//
//func convertDatumMetaToInfo(meta *datum.Meta) *pps.DatumInfo {
//	di := &pps.DatumInfo{
//		Datum: &pps.Datum{
//			Job: &pps.Job{
//				ID: meta.JobID,
//			},
//			ID: common.DatumIDV2(meta.Inputs),
//		},
//		State: convertDatumState(meta.State),
//		Stats: meta.Stats,
//	}
//	for _, input := range meta.Inputs {
//		di.Data = append(di.Data, input.FileInfo)
//	}
//	return di
//}
//
//// (bryce) this is a bit wonky, but it is necessary based on the dependency graph.
//func convertDatumState(state datum.State) pps.DatumState {
//	switch state {
//	case datum.State_FAILED:
//		return pps.DatumState_FAILED
//	case datum.State_RECOVERED:
//		return pps.DatumState_RECOVERED
//	default:
//		return pps.DatumState_SUCCESS
//	}
//}
//
//func (a *apiServerV2) collectDatums(ctx context.Context, job *pps.Job, cb func(*datum.Meta, *pfs.File) error) error {
//	jobInfo, err := a.InspectJob(ctx, &pps.InspectJobRequest{
//		Job: &pps.Job{
//			ID: job.ID,
//		},
//	})
//	if err != nil {
//		return err
//	}
//	if jobInfo.StatsCommit == nil {
//		return errors.Errorf("job not finished")
//	}
//	pachClient := a.env.GetPachClient(ctx)
//	fsi := datum.NewFileSetIterator(pachClient, jobInfo.StatsCommit.Repo.Name, jobInfo.StatsCommit.ID)
//	return fsi.Iterate(func(meta *datum.Meta) error {
//		// TODO: Potentially refactor into datum package (at least the path).
//		pfsState := &pfs.File{
//			Commit: jobInfo.StatsCommit,
//			Path:   "/" + path.Join(datum.PFSPrefix, common.DatumIDV2(meta.Inputs)),
//		}
//		return cb(meta, pfsState)
//	})
//}
//
//var errV1NotImplemented = errors.Errorf("V1 method not implemented")
//
//func (a *apiServerV2) ListDatum(_ context.Context, _ *pps.ListDatumRequest) (*pps.ListDatumResponse, error) {
//	return nil, errV1NotImplemented
//}
//
//func (a *apiServerV2) CreatePipeline(ctx context.Context, request *pps.CreatePipelineRequest) (*types.Empty, error) {
//	if request.TFJob != nil {
//		return nil, errors.Errorf("TFJob not implemented")
//	}
//	if request.HashtreeSpec != nil {
//		return nil, errors.Errorf("HashtreeSpec not implemented")
//	}
//	if request.Egress != nil {
//		return nil, errors.Errorf("Egress not implemented")
//	}
//	if request.S3Out {
//		return nil, errors.Errorf("S3Out not implemented")
//	}
//	var err error
//	pps.VisitInput(request.Input, func(input *pps.Input) {
//		if input.Join != nil {
//			err = errors.Errorf("Join input not implemented")
//		}
//	})
//	if err != nil {
//		return nil, err
//	}
//	if request.CacheSize != "" {
//		return nil, errors.Errorf("CacheSize not implemented")
//	}
//	request.EnableStats = true
//	if request.MaxQueueSize != 0 {
//		return nil, errors.Errorf("MaxQueueSize not implemented")
//	}
//	if request.Service != nil {
//		return nil, errors.Errorf("Service not implemented")
//	}
//	if request.Spout != nil {
//		return nil, errors.Errorf("Spout not implemented")
//	}
//	return a.apiServer.CreatePipeline(ctx, request)
//}
//
//func (a *apiServerV2) GarbageCollect(_ context.Context, _ *pps.GarbageCollectRequest) (*pps.GarbageCollectResponse, error) {
//	return nil, errV1NotImplemented
//}
