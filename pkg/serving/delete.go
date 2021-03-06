package serving

import (
	"fmt"

	"github.com/kubeflow/arena/pkg/apis/types"
	"github.com/kubeflow/arena/pkg/workflow"
	log "github.com/sirupsen/logrus"
)

func DeleteServingJob(namespace, name, version string, jobType types.ServingJobType) error {
	job, err := SearchServingJob(namespace, name, version, jobType)
	if err != nil {
		return err
	}
	nameWithVersion := fmt.Sprintf("%v-%v", job.Name(), job.Version())
	servingType := string(job.Type())
	err = workflow.DeleteJob(nameWithVersion, namespace, servingType)
	if err != nil {
		return err
	}
	log.Infof("The Serving job %s with version %s has been deleted successfully", job.Name(), job.Version())
	return nil
}
