package temporal

import (
	ssv1 "github.com/why-xn/go-temporal-skeleton/pkg/api/service/student/v1"
	swv1 "github.com/why-xn/go-temporal-skeleton/pkg/api/workflow/student/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

func StartWorker(c client.Client) {
	w := worker.New(c, WorkerQueue, worker.Options{})

	// This worker hosts both Workflow and Activity functions.
	w.RegisterWorkflow(swv1.StudentWorkflow().GetStudentList)
	w.RegisterWorkflow(swv1.StudentWorkflow().GetStudent)

	w.RegisterActivity(ssv1.StudentService().GetStudentList)
	w.RegisterActivity(ssv1.StudentService().GetStudent)

	// Start listening to the Task Queue.
	err := w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Alert Worker", err)
	}
}
